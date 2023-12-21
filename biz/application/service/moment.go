package service

import (
	"context"
	"net/url"

	"github.com/google/wire"
	"github.com/jinzhu/copier"
	"github.com/samber/lo"
	"github.com/xh-polaris/gopkg/errors"
	genbasic "github.com/xh-polaris/service-idl-gen-go/kitex_gen/basic"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowchat/content"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/platform/sts"

	"github.com/xh-polaris/meowchat-core-api/biz/adaptor"
	"github.com/xh-polaris/meowchat-core-api/biz/application/dto/basic"
	"github.com/xh-polaris/meowchat-core-api/biz/application/dto/meowchat/core_api"
	"github.com/xh-polaris/meowchat-core-api/biz/domain/service"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/config"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/consts"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/meowchat_content"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/meowchat_user"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/platform_sts"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/util"
)

type IMomentService interface {
	DeleteMoment(ctx context.Context, req *core_api.DeleteMomentReq) (*core_api.DeleteMomentResp, error)
	GetMomentDetail(ctx context.Context, req *core_api.GetMomentDetailReq) (*core_api.GetMomentDetailResp, error)
	GetMomentPreviews(ctx context.Context, req *core_api.GetMomentPreviewsReq) (*core_api.GetMomentPreviewsResp, error)
	NewMoment(ctx context.Context, req *core_api.NewMomentReq) (*core_api.NewMomentResp, error)
}

type MomentService struct {
	Config              *config.Config
	MomentDomainService service.IMomentDomainService
	MeowchatContent     meowchat_content.IMeowchatContent
	MeowchatUser        meowchat_user.IMeowchatUser
	PlatformSts         platform_sts.IPlatformSts
}

var MomentServiceSet = wire.NewSet(
	wire.Struct(new(MomentService), "*"),
	wire.Bind(new(IMomentService), new(*MomentService)),
)

func (s *MomentService) DeleteMoment(ctx context.Context, req *core_api.DeleteMomentReq) (*core_api.DeleteMomentResp, error) {
	user := adaptor.ExtractUserMeta(ctx)
	if user.GetUserId() == "" {
		return nil, consts.ErrNotAuthentication
	}
	resp := new(core_api.DeleteMomentResp)
	_, err := s.MeowchatContent.DeleteMoment(ctx, &content.DeleteMomentReq{
		MomentId: req.MomentId,
	})
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *MomentService) GetMomentDetail(ctx context.Context, req *core_api.GetMomentDetailReq) (*core_api.GetMomentDetailResp, error) {
	resp := new(core_api.GetMomentDetailResp)
	userMeta := adaptor.ExtractUserMeta(ctx)
	data, err := s.MeowchatContent.RetrieveMoment(ctx, &content.RetrieveMomentReq{MomentId: req.MomentId})
	if err != nil {
		return nil, err
	}

	resp.Moment = new(core_api.Moment)
	err = copier.Copy(resp.Moment, data.Moment)
	if err != nil {
		return nil, err
	}

	util.ParallelRun([]func(){
		func() {
			if data.Moment.GetCatId() != "" {
				_ = s.MomentDomainService.LoadCats(ctx, resp.Moment, []string{data.Moment.GetCatId()})
			}
		},
		func() {
			_ = s.MomentDomainService.LoadAuthor(ctx, resp.Moment, data.Moment.UserId)
		},
		func() {
			_ = s.MomentDomainService.LoadLikeCount(ctx, resp.Moment)
		},
		func() {
			_ = s.MomentDomainService.LoadCommentCount(ctx, resp.Moment)
		},
		func() {
			_ = s.MomentDomainService.LoadIsCurrentUserLiked(ctx, resp.Moment, userMeta.UserId)
		},
	})

	return resp, nil
}

func (s *MomentService) GetMomentPreviews(ctx context.Context, req *core_api.GetMomentPreviewsReq) (*core_api.GetMomentPreviewsResp, error) {
	resp := new(core_api.GetMomentPreviewsResp)
	var data *content.ListMomentResp

	if req.PaginationOption == nil {
		req.PaginationOption = &basic.PaginationOptions{}
	}
	if req.PaginationOption.Limit == nil {
		req.PaginationOption.Limit = lo.ToPtr[int64](10)
	}
	request := &content.ListMomentReq{
		FilterOptions: &content.MomentFilterOptions{
			OnlyUserId:      req.OnlyUserId,
			OnlyCommunityId: req.CommunityId,
		},
		PaginationOptions: &genbasic.PaginationOptions{
			Limit:     req.PaginationOption.Limit,
			Backward:  req.PaginationOption.Backward,
			LastToken: req.PaginationOption.LastToken,
		},
	}
	if req.GetKeyword() != "" {
		request.SearchOptions = &content.SearchOptions{
			Type: &content.SearchOptions_AllFieldsKey{AllFieldsKey: req.GetKeyword()},
		}
	}
	if req.PaginationOption.LastToken == nil {
		request.PaginationOptions.Offset = lo.EmptyableToPtr(req.PaginationOption.GetLimit() * req.PaginationOption.GetPage())
	}
	data, err := s.MeowchatContent.ListMoment(ctx, request)
	if err != nil {
		return nil, err
	}

	resp.Total = data.Total
	resp.Token = data.Token
	resp.Moments = make([]*core_api.Moment, 0)
	err = copier.Copy(&resp.Moments, data.Moments)
	if err != nil {
		return nil, err
	}

	// 并发获取额外信息
	util.ParallelRun(lo.Map(data.Moments, func(moment *content.Moment, i int) func() {
		return func() {
			// 并发获取用户信息、点赞数、评论数
			util.ParallelRun([]func(){
				func() {
					_ = s.MomentDomainService.LoadAuthor(ctx, resp.Moments[i], moment.UserId)
				},
				func() {
					_ = s.MomentDomainService.LoadLikeCount(ctx, resp.Moments[i])
				},
				func() {
					_ = s.MomentDomainService.LoadCommentCount(ctx, resp.Moments[i])
				},
				func() {
					if moment.GetCatId() != "" {
						_ = s.MomentDomainService.LoadCats(ctx, resp.Moments[i], []string{moment.GetCatId()})
					}
				},
			})
		}
	}))
	return resp, nil
}

func (s *MomentService) NewMoment(ctx context.Context, req *core_api.NewMomentReq) (*core_api.NewMomentResp, error) {
	user := adaptor.ExtractUserMeta(ctx)
	if user.GetUserId() == "" {
		return nil, consts.ErrNotAuthentication
	}
	resp := new(core_api.NewMomentResp)
	m := new(content.Moment)

	if req.GetText()+req.GetTitle() != "" {
		r, err := s.PlatformSts.TextCheck(ctx, &sts.TextCheckReq{
			Text:  req.GetText() + req.GetTitle(),
			User:  user,
			Scene: 2,
			Title: req.Title,
		})
		if err != nil {
			return nil, err
		}
		if r.Pass == false {
			return nil, errors.NewBizError(10001, "TextCheck don't pass")
		}
	}

	urls := make([]string, len(req.Photos))
	for i := 0; i < len(req.Photos); i++ {
		var u *url.URL
		u, err := url.Parse(req.Photos[i])
		if err != nil {
			return nil, err
		}
		u.Host = s.Config.CdnHost
		req.Photos[i] = u.String()
		urls[i] = req.Photos[i]
	}
	res, err := s.PlatformSts.PhotoCheck(ctx, &sts.PhotoCheckReq{
		User: user,
		Url:  urls,
	})
	if err != nil {
		return nil, err
	}
	if res.Pass == false {
		return nil, errors.NewBizError(10002, "PhotoCheck don't pass")
	}

	err = copier.Copy(m, req)
	if err != nil {
		return nil, err
	}

	m.UserId = user.GetUserId()

	if req.GetId() == "" {
		var data *content.CreateMomentResp
		data, err = s.MeowchatContent.CreateMoment(ctx, &content.CreateMomentReq{Moment: m})
		resp.MomentId = data.MomentId
		if data.GetGetFish() == true {
			_, err = s.MeowchatContent.AddUserFish(ctx, &content.AddUserFishReq{
				UserId: user.UserId,
				Fish:   s.Config.Fish.Content[data.GetFishTimes-1],
			})
			if err == nil {
				resp.GetFishNum = s.Config.Fish.Content[data.GetFishTimes-1]
			}
		}
		resp.GetFish = data.GetFish
		resp.GetFishTimes = data.GetFishTimes
	} else {
		_, err = s.MeowchatContent.UpdateMoment(ctx, &content.UpdateMomentReq{Moment: m})
		resp.MomentId = *req.Id
		resp.GetFish = false
		resp.GetFishTimes = 0
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

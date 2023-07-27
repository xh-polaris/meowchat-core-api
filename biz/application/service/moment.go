package service

import (
	"context"
	"github.com/google/wire"
	"github.com/jinzhu/copier"
	"net/url"

	"github.com/xh-polaris/meowchat-core-api/biz/application/dto/meowchat/core_api"
	user1 "github.com/xh-polaris/meowchat-core-api/biz/application/dto/meowchat/user"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/config"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/meowchat_moment"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/meowchat_user"

	"github.com/xh-polaris/meowchat-moment-rpc/pb"
	pb2 "github.com/xh-polaris/meowchat-user-rpc/pb"
)

type IMomentService interface {
	DeleteMoment(ctx context.Context, req *core_api.DeleteMomentReq) (*core_api.DeleteMomentResp, error)
	GetMomentDetail(ctx context.Context, req *core_api.GetMomentDetailReq) (*core_api.GetMomentDetailResp, error)
	GetMomentPreviews(ctx context.Context, req *core_api.GetMomentPreviewsReq) (*core_api.GetMomentPreviewsResp, error)
	NewMoment(ctx context.Context, req *core_api.NewMomentReq) (*core_api.NewMomentResp, error)
	SearchMoment(ctx context.Context, req *core_api.SearchMomentReq) (*core_api.SearchMomentResp, error)
}

type MomentService struct {
	Config *config.Config
	Moment meowchat_moment.IMeowchatMoment
	User   meowchat_user.IMeowchatUser
}

var MomentServiceSet = wire.NewSet(
	wire.Struct(new(MomentService), "*"),
	wire.Bind(new(IMomentService), new(*MomentService)),
)

var PageSize int64 = 10

func (s *MomentService) DeleteMoment(ctx context.Context, req *core_api.DeleteMomentReq) (*core_api.DeleteMomentResp, error) {
	resp := new(core_api.DeleteMomentResp)
	_, err := s.Moment.DeleteMoment(ctx, &pb.DeleteMomentReq{
		MomentId: req.MomentId,
	})
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *MomentService) GetMomentDetail(ctx context.Context, req *core_api.GetMomentDetailReq) (*core_api.GetMomentDetailResp, error) {
	resp := new(core_api.GetMomentDetailResp)
	data, err := s.Moment.RetrieveMoment(ctx, &pb.RetrieveMomentReq{MomentId: req.MomentId})
	if err != nil {
		return nil, err
	}

	err = copier.Copy(&resp.Moment, data.Moment)
	if err != nil {
		return nil, err
	}

	user, err := s.User.GetUser(ctx, &pb2.GetUserReq{UserId: data.Moment.UserId})
	if err == nil {
		resp.Moment.UserId = &user1.UserPreview{
			Id:        user.User.Id,
			Nickname:  user.User.Nickname,
			AvatarUrl: user.User.AvatarUrl,
		}
	}
	return resp, nil
}

func (s *MomentService) GetMomentPreviews(ctx context.Context, req *core_api.GetMomentPreviewsReq) (*core_api.GetMomentPreviewsResp, error) {
	resp := new(core_api.GetMomentPreviewsResp)
	var data *pb.ListMomentResp

	if req.PaginationOption.Limit == nil {
		req.PaginationOption.Limit = &PageSize
	}
	request := &pb.ListMomentReq{
		FilterOptions: &pb.FilterOptions{
			OnlyUserId:      req.OnlyUserId,
			OnlyCommunityId: req.CommunityId,
		},
		PaginationOptions: &pb.PaginationOptions{
			Offset:    new(int64),
			Limit:     req.PaginationOption.Limit,
			Backward:  req.PaginationOption.Backward,
			LastToken: req.PaginationOption.LastToken,
		},
	}
	*request.PaginationOptions.Offset = req.PaginationOption.GetLimit() * *req.PaginationOption.Page
	data, err := s.Moment.ListMoment(ctx, request)
	if err != nil {
		return nil, err
	}

	resp.Total = data.Total
	resp.Moments = make([]*core_api.Moment, 0, pageSize)
	err = copier.Copy(&resp.Moments, data.Moments)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(data.Moments); i++ {
		user, err := s.User.GetUser(ctx, &pb2.GetUserReq{UserId: data.Moments[i].UserId})
		if err == nil {
			resp.Moments[i].UserId = &user1.UserPreview{
				Id:        user.User.Id,
				Nickname:  user.User.Nickname,
				AvatarUrl: user.User.AvatarUrl,
			}
		}
	}
	return resp, nil
}

func (s *MomentService) NewMoment(ctx context.Context, req *core_api.NewMomentReq) (*core_api.NewMomentResp, error) {
	resp := new(core_api.NewMomentResp)
	m := new(pb.Moment)
	//openId := ctx.Value("openId").(string)
	//
	//err = util.MsgSecCheck(l.ctx, l.svcCtx, req.Title+"\n"+req.Text, openId, 2)
	//if err != nil {
	//	return nil, err
	//}

	for i := 0; i < len(req.Photos); i++ {
		var u *url.URL
		u, err := url.Parse(req.Photos[i])
		if err != nil {
			return nil, err
		}
		u.Host = s.Config.CdnHost
		req.Photos[i] = u.String()
	}
	//err = util.PhotoCheck(l.ctx, l.svcCtx, req.Photos)
	//if err != nil {
	//	return nil, err
	//}

	err := copier.Copy(m, req)
	if err != nil {
		return nil, err
	}

	m.UserId = ctx.Value("userId").(string)

	if *req.Id == "" {
		var data *pb.CreateMomentResp
		data, err = s.Moment.CreateMoment(ctx, &pb.CreateMomentReq{Moment: m})
		resp.MomentId = data.MomentId
	} else {
		_, err = s.Moment.UpdateMoment(ctx, &pb.UpdateMomentReq{Moment: m})
		resp.MomentId = *req.Id
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *MomentService) SearchMoment(ctx context.Context, req *core_api.SearchMomentReq) (*core_api.SearchMomentResp, error) {
	resp := new(core_api.SearchMomentResp)
	var data *pb.ListMomentResp

	if req.PaginationOption.Limit == nil {
		req.PaginationOption.Limit = &PageSize
	}
	request := &pb.ListMomentReq{
		SearchOptions: &pb.SearchOptions{Type: &pb.SearchOptions_AllFieldsKey{AllFieldsKey: *req.Keyword}},
		FilterOptions: &pb.FilterOptions{
			OnlyUserId:      req.OnlyUserId,
			OnlyCommunityId: req.CommunityId,
		},
		PaginationOptions: &pb.PaginationOptions{
			Offset:    new(int64),
			Limit:     req.PaginationOption.Limit,
			Backward:  req.PaginationOption.Backward,
			LastToken: req.PaginationOption.LastToken,
		},
	}
	*request.PaginationOptions.Offset = *req.PaginationOption.Limit * *req.PaginationOption.Page
	data, err := s.Moment.ListMoment(ctx, request)
	if err != nil {
		return nil, err
	}

	resp.Total = data.Total
	resp.Moments = make([]*core_api.Moment, 0, PageSize)
	err = copier.Copy(&resp.Moments, data.Moments)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(data.Moments); i++ {
		user, err := s.User.GetUser(ctx, &pb2.GetUserReq{UserId: data.Moments[i].UserId})
		if err == nil {
			resp.Moments[i].UserId = &user1.UserPreview{
				Id:        user.User.Id,
				Nickname:  user.User.Nickname,
				AvatarUrl: user.User.AvatarUrl,
			}
		}
	}
	return resp, err
}

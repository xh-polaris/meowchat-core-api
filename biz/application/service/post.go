package service

import (
	"context"
	"net/url"

	"github.com/google/wire"
	"github.com/samber/lo"
	"github.com/xh-polaris/gopkg/errors"
	genbasic "github.com/xh-polaris/service-idl-gen-go/kitex_gen/basic"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowchat/content"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/platform/sts"

	"github.com/xh-polaris/meowchat-core-api/biz/application/dto/basic"
	"github.com/xh-polaris/meowchat-core-api/biz/application/dto/meowchat/core_api"
	"github.com/xh-polaris/meowchat-core-api/biz/domain/service"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/config"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/meowchat_content"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/platform_sts"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/util"
)

type IPostService interface {
	DeletePost(ctx context.Context, req *core_api.DeletePostReq) (*core_api.DeletePostResp, error)
	GetPostDetail(ctx context.Context, req *core_api.GetPostDetailReq, userMeta *genbasic.UserMeta) (*core_api.GetPostDetailResp, error)
	GetPostPreviews(ctx context.Context, req *core_api.GetPostPreviewsReq) (*core_api.GetPostPreviewsResp, error)
	NewPost(ctx context.Context, req *core_api.NewPostReq, user *genbasic.UserMeta) (*core_api.NewPostResp, error)
	SetOfficial(ctx context.Context, req *core_api.SetOfficialReq) (*core_api.SetOfficialResp, error)
}

type PostService struct {
	Config            *config.Config
	PostDomainService service.IPostDomainService
	MeowchatContent   meowchat_content.IMeowchatContent
	PlatformSts       platform_sts.IPlatformSts
}

var PostServiceSet = wire.NewSet(
	wire.Struct(new(PostService), "*"),
	wire.Bind(new(IPostService), new(*PostService)),
)

func (s *PostService) DeletePost(ctx context.Context, req *core_api.DeletePostReq) (*core_api.DeletePostResp, error) {
	resp := new(core_api.DeletePostResp)

	_, err := s.MeowchatContent.DeletePost(ctx, &content.DeletePostReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *PostService) GetPostDetail(ctx context.Context, req *core_api.GetPostDetailReq, userMeta *genbasic.UserMeta) (*core_api.GetPostDetailResp, error) {
	resp := new(core_api.GetPostDetailResp)

	data, err := s.MeowchatContent.RetrievePost(ctx, &content.RetrievePostReq{PostId: req.PostId})
	if err != nil {
		return nil, err
	}

	p := &core_api.Post{
		Id:         data.Post.Id,
		CreateAt:   data.Post.CreateAt,
		Title:      data.Post.Title,
		Text:       data.Post.Text,
		CoverUrl:   lo.EmptyableToPtr(data.Post.CoverUrl),
		Tags:       data.Post.Tags,
		IsOfficial: data.Post.IsOfficial,
	}
	util.ParallelRun([]func(){
		func() {
			_ = s.PostDomainService.LoadAuthor(ctx, p, data.Post.UserId)
		},
		func() {
			_ = s.PostDomainService.LoadLikeCount(ctx, p)
		},
		func() {
			_ = s.PostDomainService.LoadCommentCount(ctx, p)
		},
		func() {
			_ = s.PostDomainService.LoadIsCurrentUserLiked(ctx, p, userMeta.UserId)
		},
	})
	resp.Post = p
	return resp, nil
}

func (s *PostService) GetPostPreviews(ctx context.Context, req *core_api.GetPostPreviewsReq) (*core_api.GetPostPreviewsResp, error) {
	resp := new(core_api.GetPostPreviewsResp)

	data, err := s.MeowchatContent.ListPost(ctx, s.makeRequest(req))
	if err != nil {
		return nil, err
	}
	resp.Total = data.Total
	resp.Token = data.Token
	resp.Posts = make([]*core_api.Post, len(data.Posts))
	util.ParallelRun(lo.Map(data.Posts, func(val *content.Post, i int) func() {
		return func() {
			p := &core_api.Post{
				Id:         data.Posts[i].Id,
				CreateAt:   data.Posts[i].CreateAt,
				Title:      data.Posts[i].Title,
				Text:       data.Posts[i].Text,
				CoverUrl:   lo.EmptyableToPtr(data.Posts[i].CoverUrl),
				Tags:       data.Posts[i].Tags,
				IsOfficial: data.Posts[i].IsOfficial,
			}
			util.ParallelRun([]func(){
				func() {
					_ = s.PostDomainService.LoadAuthor(ctx, p, data.Posts[i].UserId)
				},
				func() {
					_ = s.PostDomainService.LoadLikeCount(ctx, p)
				},
				func() {
					_ = s.PostDomainService.LoadCommentCount(ctx, p)
				},
			})
			resp.Posts[i] = p
		}
	}))

	return resp, nil
}

func (s *PostService) NewPost(ctx context.Context, req *core_api.NewPostReq, user *genbasic.UserMeta) (*core_api.NewPostResp, error) {
	resp := new(core_api.NewPostResp)

	r, err := s.PlatformSts.TextCheck(ctx, &sts.TextCheckReq{
		Text:  req.Text,
		User:  user,
		Scene: 2,
		Title: &req.Title,
	})
	if err != nil {
		return nil, err
	}
	if r.Pass == false {
		return nil, errors.NewBizError(10001, "TextCheck don't pass")
	}

	if req.GetCoverUrl() != "" {
		var u *url.URL
		u, err = url.Parse(req.GetCoverUrl())
		if err != nil {
			return resp, err
		}
		u.Host = s.Config.CdnHost
		req.CoverUrl = lo.ToPtr(u.String())

		res, err := s.PlatformSts.PhotoCheck(ctx, &sts.PhotoCheckReq{
			User: user,
			Url:  []string{req.GetCoverUrl()},
		})
		if err != nil {
			return nil, err
		}
		if res.Pass == false {
			return nil, errors.NewBizError(10002, "PhotoCheck don't pass")
		}
	}

	if req.GetId() == "" {
		res, err := s.MeowchatContent.CreatePost(ctx, &content.CreatePostReq{
			Title:    req.Title,
			Text:     req.Text,
			CoverUrl: req.GetCoverUrl(),
			Tags:     req.Tags,
			UserId:   user.UserId,
		})
		if err != nil {
			return nil, err
		}
		if res.GetGetFish() == true {
			_, err = s.MeowchatContent.AddUserFish(ctx, &content.AddUserFishReq{
				UserId: user.UserId,
				Fish:   s.Config.Fish.Content,
			})
		}
		resp.GetFish = res.GetFish
		resp.GetFishTimes = res.GetFishTimes
		resp.PostId = res.PostId
	} else {
		_, err = s.MeowchatContent.UpdatePost(ctx, &content.UpdatePostReq{
			Id:       *req.Id,
			Title:    req.Title,
			Text:     req.Text,
			CoverUrl: req.GetCoverUrl(),
			Tags:     req.Tags,
		})
		if err != nil {
			return nil, err
		}
		resp.PostId = *req.Id
		resp.GetFish = false
		resp.GetFishTimes = 0
	}

	return resp, nil
}

func (s *PostService) SetOfficial(ctx context.Context, req *core_api.SetOfficialReq) (*core_api.SetOfficialResp, error) {
	resp := new(core_api.SetOfficialResp)
	_, err := s.MeowchatContent.SetOfficial(ctx, &content.SetOfficialReq{
		PostId:   req.PostId,
		IsRemove: *req.IsRemove,
	})
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *PostService) makeRequest(req *core_api.GetPostPreviewsReq) *content.ListPostReq {
	r := &content.ListPostReq{}

	if req.SearchOptions != nil {
		if req.SearchOptions.Key != nil {
			r.SearchOptions = &content.SearchOptions{
				Type: &content.SearchOptions_AllFieldsKey{
					AllFieldsKey: *req.SearchOptions.Key,
				},
			}
		} else {
			r.SearchOptions = &content.SearchOptions{
				Type: &content.SearchOptions_MultiFieldsKey{
					MultiFieldsKey: &content.SearchField{
						Text:  req.SearchOptions.Text,
						Title: req.SearchOptions.Title,
						Tag:   req.SearchOptions.Tag,
					},
				},
			}
		}
	}
	if req.PaginationOption == nil {
		req.PaginationOption = &basic.PaginationOptions{}
	}
	if req.PaginationOption.Limit == nil {
		req.PaginationOption.Limit = lo.ToPtr[int64](10)
	}
	r.PaginationOptions = &genbasic.PaginationOptions{
		Offset:    req.PaginationOption.Offset,
		Limit:     req.PaginationOption.Limit,
		Backward:  req.PaginationOption.Backward,
		LastToken: req.PaginationOption.LastToken,
	}
	r.FilterOptions = &content.PostFilterOptions{
		OnlyOfficial: req.OnlyOfficial,
		OnlyUserId:   req.OnlyUserId,
	}

	return r
}

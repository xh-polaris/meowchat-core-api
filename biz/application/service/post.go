package service

import (
	"context"
	"net/url"

	"github.com/google/wire"
	"github.com/xh-polaris/gopkg/errors"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/basic"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowchat/content"
	genuser "github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowchat/user"
	gencomment "github.com/xh-polaris/service-idl-gen-go/kitex_gen/platform/comment"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/platform/sts"

	"github.com/xh-polaris/meowchat-core-api/biz/application/dto/meowchat/core_api"
	user1 "github.com/xh-polaris/meowchat-core-api/biz/application/dto/meowchat/user"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/config"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/meowchat_content"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/meowchat_user"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/platform_comment"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/platform_sts"
)

type IPostService interface {
	DeletePost(ctx context.Context, req *core_api.DeletePostReq) (*core_api.DeletePostResp, error)
	GetPostDetail(ctx context.Context, req *core_api.GetPostDetailReq) (*core_api.GetPostDetailResp, error)
	GetPostPreviews(ctx context.Context, req *core_api.GetPostPreviewsReq) (*core_api.GetPostPreviewsResp, error)
	NewPost(ctx context.Context, req *core_api.NewPostReq, user *basic.UserMeta) (*core_api.NewPostResp, error)
	SetOfficial(ctx context.Context, req *core_api.SetOfficialReq) (*core_api.SetOfficialResp, error)
}

type PostService struct {
	Config  *config.Config
	Content meowchat_content.IMeowchatContent
	User    meowchat_user.IMeowchatUser
	Comment platform_comment.IPlatformCommment
	Sts     platform_sts.IPlatformSts
}

var PostServiceSet = wire.NewSet(
	wire.Struct(new(PostService), "*"),
	wire.Bind(new(IPostService), new(*PostService)),
)

func (s *PostService) DeletePost(ctx context.Context, req *core_api.DeletePostReq) (*core_api.DeletePostResp, error) {
	resp := new(core_api.DeletePostResp)

	_, err := s.Content.DeletePost(ctx, &content.DeletePostReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *PostService) GetPostDetail(ctx context.Context, req *core_api.GetPostDetailReq) (*core_api.GetPostDetailResp, error) {
	resp := new(core_api.GetPostDetailResp)

	data, err := s.Content.RetrievePost(ctx, &content.RetrievePostReq{PostId: req.PostId})
	if err != nil {
		return nil, err
	}

	respPost, _ := s.toRespPost(ctx, data.Post)
	resp.Post = &respPost

	return resp, nil
}

func (s *PostService) GetPostPreviews(ctx context.Context, req *core_api.GetPostPreviewsReq) (*core_api.GetPostPreviewsResp, error) {
	resp := new(core_api.GetPostPreviewsResp)

	data, err := s.Content.ListPost(ctx, s.makeRequest(req))
	if err != nil {
		return nil, err
	}
	resp.Total = data.Total
	resp.Posts = make([]*core_api.Post, len(data.Posts))
	for i, val := range data.Posts {
		respPost, _ := s.toRespPost(ctx, val)
		resp.Posts[i] = &respPost
	}
	resp.Token = data.Token

	return resp, nil
}

func (s *PostService) NewPost(ctx context.Context, req *core_api.NewPostReq, user *basic.UserMeta) (*core_api.NewPostResp, error) {
	resp := new(core_api.NewPostResp)

	r, err := s.Sts.TextCheck(ctx, &sts.TextCheckReq{
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

	var u *url.URL
	u, err = url.Parse(req.CoverUrl)
	if err != nil {
		return resp, err
	}
	u.Host = s.Config.CdnHost
	req.CoverUrl = u.String()

	var i = []string{req.CoverUrl}

	res, err := s.Sts.PhotoCheck(ctx, &sts.PhotoCheckReq{
		User: user,
		Url:  i,
	})
	if err != nil {
		return nil, err
	}
	if res.Pass == false {
		return nil, errors.NewBizError(10002, "PhotoCheck don't pass")
	}

	if err != nil {
		return nil, err
	}

	if *req.Id == "" {
		res, err := s.Content.CreatePost(ctx, &content.CreatePostReq{
			Title:    req.Title,
			Text:     req.Text,
			CoverUrl: req.CoverUrl,
			Tags:     req.Tags,
			UserId:   user.UserId,
		})
		if err != nil {
			return nil, err
		}
		if res.GetGetFish() == true {
			_, err = s.Content.AddUserFish(ctx, &content.AddUserFishReq{
				UserId: user.UserId,
				Fish:   s.Config.Fish.Content,
			})
		}
		resp.GetFish = res.GetFish
		resp.GetFishTimes = res.GetFishTimes
		resp.PostId = res.PostId
	} else {
		_, err = s.Content.UpdatePost(ctx, &content.UpdatePostReq{
			Id:       *req.Id,
			Title:    req.Title,
			Text:     req.Text,
			CoverUrl: req.CoverUrl,
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
	_, err := s.Content.SetOfficial(ctx, &content.SetOfficialReq{
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
	r.PaginationOptions = &basic.PaginationOptions{
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

func (s *PostService) toRespPost(ctx context.Context, post *content.Post) (resp core_api.Post, err error) {
	resp = core_api.Post{
		Id:         post.Id,
		CreateAt:   post.CreateAt,
		Title:      post.Title,
		Text:       post.Text,
		CoverUrl:   post.CoverUrl,
		Tags:       post.Tags,
		IsOfficial: post.IsOfficial,
	}

	// user preview
	user, err := s.User.GetUser(ctx, &genuser.GetUserReq{UserId: post.UserId})
	if user != nil && err == nil {
		resp.User = &user1.UserPreview{
			Id:        post.UserId,
			Nickname:  user.User.Nickname,
			AvatarUrl: user.User.AvatarUrl,
		}
	}

	// likes
	likes, err := s.User.GetTargetLikes(ctx, &genuser.GetTargetLikesReq{
		TargetId: post.Id,
		Type:     genuser.LikeType_Post,
	})
	if likes != nil && err == nil {
		resp.Likes = likes.Count
	}

	// comments
	data, err := s.Comment.CountCommentByParent(ctx, &gencomment.CountCommentByParentReq{
		Type:     "post",
		ParentId: post.Id,
	})
	if err == nil {
		resp.Comments = data.Total
	}

	return
}

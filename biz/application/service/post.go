package service

import (
	"context"
	"github.com/google/wire"
	"github.com/xh-polaris/meowchat-core-api/biz/application/dto/meowchat/core_api"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/config"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/meowchat_like"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/meowchat_post"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/meowchat_user"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/platform_comment"
)

type IPostService interface {
	DeletePost(ctx context.Context, req *core_api.DeletePostReq) (*core_api.DeletePostResp, error)
	GetPostDetail(ctx context.Context, req *core_api.GetPostDetailReq) (*core_api.GetPostDetailResp, error)
	GetPostPreviews(ctx context.Context, req *core_api.GetPostPreviewsReq) (*core_api.GetPostPreviewsResp, error)
	NewPost(ctx context.Context, req *core_api.NewPostReq) (*core_api.NewPostResp, error)
	SetOfficial(ctx context.Context, req *core_api.SetOfficialReq) (*core_api.SetOfficialResp, error)
}

type PostService struct {
	Config  *config.Config
	Post    meowchat_post.IMeowchatPost
	User    meowchat_user.IMeowchatUser
	Like    meowchat_like.IMeowchatLike
	Comment platform_comment.IPlatformCommment
}

var PostServiceSet = wire.NewSet(
	wire.Struct(new(PostService), "*"),
	wire.Bind(new(IPostService), new(*PostService)),
)

func (s *PostService) DeletePost(ctx context.Context, req *core_api.DeletePostReq) (*core_api.DeletePostResp, error) {
	//TODO implement me
	panic("implement me")
}

func (s *PostService) GetPostDetail(ctx context.Context, req *core_api.GetPostDetailReq) (*core_api.GetPostDetailResp, error) {
	//TODO implement me
	panic("implement me")
}

func (s *PostService) GetPostPreviews(ctx context.Context, req *core_api.GetPostPreviewsReq) (*core_api.GetPostPreviewsResp, error) {
	//TODO implement me
	panic("implement me")
}

func (s *PostService) NewPost(ctx context.Context, req *core_api.NewPostReq) (*core_api.NewPostResp, error) {
	//TODO implement me
	panic("implement me")
}

func (s *PostService) SetOfficial(ctx context.Context, req *core_api.SetOfficialReq) (*core_api.SetOfficialResp, error) {
	//TODO implement me
	panic("implement me")
}

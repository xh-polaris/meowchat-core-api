package service

import (
	"context"
	"errors"
	genbasic "github.com/xh-polaris/service-idl-gen-go/kitex_gen/basic"

	"github.com/google/wire"
	"github.com/samber/lo"
	genuser "github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowchat/user"
	gencomment "github.com/xh-polaris/service-idl-gen-go/kitex_gen/platform/comment"

	"github.com/xh-polaris/meowchat-core-api/biz/application/dto/meowchat/core_api"
	"github.com/xh-polaris/meowchat-core-api/biz/application/dto/meowchat/user"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/meowchat_user"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/platform_comment"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/util/log"
)

type IPostDomainService interface {
	LoadAuthor(ctx context.Context, post *core_api.Post, userId string) error
	LoadCommentCount(ctx context.Context, post *core_api.Post) error
	LoadLikeCount(ctx context.Context, post *core_api.Post) error
	LoadIsCurrentUserLiked(ctx context.Context, post *core_api.Post, userId string) error
}

type PostDomainService struct {
	MeowchatUser    meowchat_user.IMeowchatUser
	PlatformComment platform_comment.IPlatformComment
}

var PostDomainServiceSet = wire.NewSet(
	wire.Struct(new(PostDomainService), "*"),
	wire.Bind(new(IPostDomainService), new(*PostDomainService)),
)

func (s *PostDomainService) LoadAuthor(ctx context.Context, post *core_api.Post, userId string) error {
	if userId == "" {
		return errors.New("userId is empty")
	}
	post.User = &user.UserPreview{
		Id: userId,
	}
	rpcResp, err := s.MeowchatUser.GetUserDetail(ctx, &genuser.GetUserDetailReq{UserId: userId})
	if err == nil {
		post.User.Nickname = rpcResp.User.GetNickname()
		post.User.AvatarUrl = rpcResp.User.GetAvatarUrl()
	}
	return nil
}

func (s *PostDomainService) LoadCommentCount(ctx context.Context, post *core_api.Post) error {
	rpcResp, err := s.PlatformComment.CountCommentByParent(ctx, &gencomment.CountCommentByParentReq{
		ParentId: post.Id,
		Type:     gencomment.CommentType_CommentType_Post,
	})
	if err != nil {
		return err
	}
	post.Comments = lo.ToPtr(rpcResp.GetTotal())
	return nil
}

func (s *PostDomainService) LoadLikeCount(ctx context.Context, post *core_api.Post) error {
	rpcResp, err := s.MeowchatUser.GetLikedUsers(ctx, &genuser.GetLikedUsersReq{
		TargetId: post.Id,
		Type:     genuser.LikeType_Post,
		PaginationOptions: &genbasic.PaginationOptions{
			Page:      nil,
			Limit:     lo.ToPtr(int64(0)),
			LastToken: nil,
			Backward:  nil,
			Offset:    nil,
		},
	})
	if err != nil {
		return err
	}
	post.Likes = lo.ToPtr(int64(len(rpcResp.UserIds)))
	return nil
}

func (s *PostDomainService) LoadIsCurrentUserLiked(ctx context.Context, post *core_api.Post, userId string) error {
	if userId == "" {
		post.IsLiked = lo.ToPtr(false)
		return nil
	}
	rpcResp, err := s.MeowchatUser.GetUserLike(ctx, &genuser.GetUserLikedReq{
		UserId:   userId,
		TargetId: post.Id,
		Type:     genuser.LikeType_Post,
	})
	if err != nil {
		log.CtxError(ctx, "[LoadIsCurrentUserLiked] fail, err=%v", err)
		return err
	}
	post.IsLiked = lo.ToPtr(rpcResp.Liked)
	return nil
}

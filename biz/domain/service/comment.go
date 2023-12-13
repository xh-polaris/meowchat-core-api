package service

import (
	"context"

	"github.com/google/wire"
	"github.com/samber/lo"
	genuser "github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowchat/user"
	gencomment "github.com/xh-polaris/service-idl-gen-go/kitex_gen/platform/comment"

	"github.com/xh-polaris/meowchat-core-api/biz/application/dto/meowchat/core_api"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/meowchat_content"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/meowchat_user"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/platform_comment"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/util/log"
)

type ICommentDomainService interface {
	LoadAuthor(ctx context.Context, comment *core_api.Comment, userId string) error
	LoadReplyUser(ctx context.Context, comment *core_api.Comment, userId string) error
	LoadCommentCount(ctx context.Context, comment *core_api.Comment) error
	LoadLikeCount(ctx context.Context, comment *core_api.Comment) error
	LoadIsCurrentUserLiked(ctx context.Context, comment *core_api.Comment, userId string) error
}

type CommentDomainService struct {
	MeowchatContent  meowchat_content.IMeowchatContent
	MeowchatUser     meowchat_user.IMeowchatUser
	PlatformCommment platform_comment.IPlatformCommment
}

var CommentDomainServiceSet = wire.NewSet(
	wire.Struct(new(CommentDomainService), "*"),
	wire.Bind(new(ICommentDomainService), new(*CommentDomainService)),
)

// LoadAuthor 评论作者信息
func (s *CommentDomainService) LoadAuthor(ctx context.Context, comment *core_api.Comment, userId string) error {
	author := &core_api.User{}
	author.Id = userId
	rpcResp, err := s.MeowchatUser.GetUserDetail(ctx, &genuser.GetUserDetailReq{
		UserId: userId,
	})
	if err != nil {
		log.CtxError(ctx, "[LoadAuthor] fail, err=%v", err)
	}
	if rpcResp != nil && err == nil {
		author.Nickname = rpcResp.User.Nickname
		author.AvatarUrl = rpcResp.User.AvatarUrl
	}
	comment.User = author
	return nil
}

func (s *CommentDomainService) LoadCommentCount(ctx context.Context, comment *core_api.Comment) error {
	rpcResp, err := s.PlatformCommment.CountCommentByParent(ctx, &gencomment.CountCommentByParentReq{
		Type:     gencomment.CommentType_CommentType_Comment,
		ParentId: comment.Id,
	})
	if err != nil {
		log.CtxError(ctx, "[LoadCommentCount] fail, err=%v", err)
		return err
	}
	comment.Comments = lo.ToPtr(rpcResp.GetTotal())
	return nil
}

func (s *CommentDomainService) LoadLikeCount(ctx context.Context, comment *core_api.Comment) error {
	rpcResp, err := s.MeowchatUser.GetLikedUsers(ctx, &genuser.GetLikedUsersReq{
		TargetId: comment.Id,
		Type:     genuser.LikeType_Comment,
	})
	if err != nil {
		log.CtxError(ctx, "[LoadLikeCount] fail, err=%v", err)
		return err
	}
	comment.LikeCount = lo.ToPtr(int64(len(rpcResp.UserIds)))
	return nil
}

// LoadIsCurrentUserLiked 当前用户是否点赞
func (s *CommentDomainService) LoadIsCurrentUserLiked(ctx context.Context, comment *core_api.Comment, userId string) error {
	if userId == "" {
		comment.IsLiked = lo.ToPtr(false)
		return nil
	}
	rpcResp, err := s.MeowchatUser.GetUserLike(ctx, &genuser.GetUserLikedReq{
		UserId:   userId,
		TargetId: comment.Id,
		Type:     genuser.LikeType_Comment,
	})
	if err != nil {
		log.CtxError(ctx, "[LoadIsCurrentUserLiked] fail, err=%v", err)
		return err
	}
	comment.IsLiked = lo.ToPtr(rpcResp.Liked)
	return nil
}

func (s *CommentDomainService) LoadReplyUser(ctx context.Context, comment *core_api.Comment, userId string) error {
	author := &core_api.User{}
	author.Id = userId
	rpcResp, err := s.MeowchatUser.GetUserDetail(ctx, &genuser.GetUserDetailReq{
		UserId: userId,
	})
	if err != nil {
		log.CtxError(ctx, "[LoadAuthor] fail, err=%v", err)
	}
	if rpcResp != nil && err == nil {
		author.Nickname = rpcResp.User.Nickname
		author.AvatarUrl = rpcResp.User.AvatarUrl
	}
	comment.ReplyUser = author
	return nil
}

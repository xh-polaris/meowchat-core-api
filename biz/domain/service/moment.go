package service

import (
	"context"
	"errors"

	"github.com/google/wire"
	"github.com/samber/lo"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowchat/content"
	genuser "github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowchat/user"
	gencomment "github.com/xh-polaris/service-idl-gen-go/kitex_gen/platform/comment"

	"github.com/xh-polaris/meowchat-core-api/biz/application/dto/meowchat/core_api"
	"github.com/xh-polaris/meowchat-core-api/biz/application/dto/meowchat/user"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/meowchat_content"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/meowchat_user"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/platform_comment"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/util"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/util/log"
)

type IMomentDomainService interface {
	LoadCats(ctx context.Context, moment *core_api.Moment, catIds []string) error
	LoadAuthor(ctx context.Context, moment *core_api.Moment, userId string) error
	LoadCommentCount(ctx context.Context, moment *core_api.Moment) error
	LoadLikeCount(ctx context.Context, moment *core_api.Moment) error
	LoadIsCurrentUserLiked(ctx context.Context, moment *core_api.Moment, userId string) error
}

type MomentDomainService struct {
	MeowchatContent  meowchat_content.IMeowchatContent
	MeowchatUser     meowchat_user.IMeowchatUser
	PlatformCommment platform_comment.IPlatformCommment
}

var MomentDomainServiceSet = wire.NewSet(
	wire.Struct(new(MomentDomainService), "*"),
	wire.Bind(new(IMomentDomainService), new(*MomentDomainService)),
)

func (s *MomentDomainService) LoadCats(ctx context.Context, moment *core_api.Moment, catIds []string) error {
	if len(catIds) == 0 {
		return nil
	}
	cats := make([]*core_api.CatPreview, len(catIds))
	util.ParallelRun(lo.Map(catIds, func(catId string, i int) func() {
		return func() {
			rpcResp, err := s.MeowchatContent.RetrieveCat(ctx, &content.RetrieveCatReq{CatId: catId})
			if err == nil && rpcResp != nil {
				cats[i] = &core_api.CatPreview{
					Id:        rpcResp.Cat.Id,
					Name:      rpcResp.Cat.Name,
					AvatarUrl: rpcResp.Cat.Avatars[0],
					Color:     rpcResp.Cat.Color,
					Area:      rpcResp.Cat.Area,
				}
			}
		}
	}))
	moment.Cats = lo.Reject(cats, func(cat *core_api.CatPreview, _ int) bool {
		return cat == nil
	})
	return nil
}

func (s *MomentDomainService) LoadIsCurrentUserLiked(ctx context.Context, moment *core_api.Moment, userId string) error {
	if userId == "" {
		moment.IsLiked = lo.ToPtr(false)
		return nil
	}
	rpcResp, err := s.MeowchatUser.GetUserLike(ctx, &genuser.GetUserLikedReq{
		UserId:   userId,
		TargetId: moment.Id,
		Type:     genuser.LikeType_Moment,
	})
	if err != nil {
		log.CtxError(ctx, "[LoadIsCurrentUserLiked] fail, err=%v", err)
		return err
	}
	moment.IsLiked = lo.ToPtr(rpcResp.Liked)
	return nil
}

func (s *MomentDomainService) LoadCommentCount(ctx context.Context, moment *core_api.Moment) error {
	rpcResp, err := s.PlatformCommment.CountCommentByParent(ctx, &gencomment.CountCommentByParentReq{
		ParentId: moment.Id,
		Type:     gencomment.CommentType_CommentType_Moment,
	})
	if err != nil {
		return err
	}
	moment.CommentCount = lo.ToPtr(rpcResp.GetTotal())
	return nil
}

func (s *MomentDomainService) LoadLikeCount(ctx context.Context, moment *core_api.Moment) error {
	rpcResp, err := s.MeowchatUser.GetLikedUsers(ctx, &genuser.GetLikedUsersReq{
		TargetId: moment.Id,
		Type:     genuser.LikeType_Moment,
	})
	if err != nil {
		return err
	}
	moment.LikeCount = lo.ToPtr(int64(len(rpcResp.UserIds)))
	return nil
}

func (s *MomentDomainService) LoadAuthor(ctx context.Context, moment *core_api.Moment, userId string) error {
	if userId == "" {
		return errors.New("userId is empty")
	}
	moment.User = &user.UserPreview{
		Id: userId,
	}
	rpcResp, err := s.MeowchatUser.GetUser(ctx, &genuser.GetUserReq{UserId: userId})
	if err == nil {
		moment.User.Nickname = rpcResp.User.GetNickname()
		moment.User.AvatarUrl = rpcResp.User.GetAvatarUrl()
	}
	return nil
}

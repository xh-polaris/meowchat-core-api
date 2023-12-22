package service

import (
	"context"

	genbasic "github.com/xh-polaris/service-idl-gen-go/kitex_gen/basic"

	"github.com/google/wire"
	"github.com/samber/lo"
	genuser "github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowchat/user"

	"github.com/xh-polaris/meowchat-core-api/biz/application/dto/meowchat/core_api"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/meowchat_content"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/meowchat_user"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/util/log"
)

type ICatImageDomainService interface {
	LoadLikeCount(ctx context.Context, image *core_api.Image) error
	LoadIsCurrentUserLiked(ctx context.Context, image *core_api.Image, userId string) error
}

type CatImageDomainService struct {
	MeowchatContent meowchat_content.IMeowchatContent
	MeowchatUser    meowchat_user.IMeowchatUser
}

var CatImageDomainServiceSet = wire.NewSet(
	wire.Struct(new(CatImageDomainService), "*"),
	wire.Bind(new(ICatImageDomainService), new(*CatImageDomainService)),
)

func (s *CatImageDomainService) LoadLikeCount(ctx context.Context, image *core_api.Image) error {

	rpcResp, err := s.MeowchatUser.GetLikedUsers(ctx, &genuser.GetLikedUsersReq{
		TargetId: image.Id,
		Type:     genuser.LikeType_CatPhoto,
		PaginationOptions: &genbasic.PaginationOptions{
			Limit: lo.ToPtr(int64(0)),
		},
	})
	if err != nil {
		log.CtxError(ctx, "[LoadLikeCount] fail, err=%v", err)
		return err
	}
	image.Likes = lo.ToPtr(rpcResp.GetTotal())
	return nil
}

// LoadIsCurrentUserLiked 当前用户是否点赞
func (s *CatImageDomainService) LoadIsCurrentUserLiked(ctx context.Context, image *core_api.Image, userId string) error {
	if userId == "" {
		image.IsLiked = lo.ToPtr(false)
		return nil
	}
	rpcResp, err := s.MeowchatUser.GetUserLike(ctx, &genuser.GetUserLikedReq{
		UserId:   userId,
		TargetId: image.Id,
		Type:     genuser.LikeType_CatPhoto,
	})
	if err != nil {
		log.CtxError(ctx, "[LoadIsCurrentUserLiked] fail, err=%v", err)
		return err
	}
	image.IsLiked = lo.ToPtr(rpcResp.Liked)
	return nil
}

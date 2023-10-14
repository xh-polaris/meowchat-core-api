package service

import (
	"context"
	"errors"

	"github.com/google/wire"
	"github.com/samber/lo"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowchat/content"
	gensystem "github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowchat/system"
	genuser "github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowchat/user"

	"github.com/xh-polaris/meowchat-core-api/biz/application/dto/meowchat/core_api"
	"github.com/xh-polaris/meowchat-core-api/biz/application/dto/meowchat/system"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/meowchat_content"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/meowchat_system"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/meowchat_user"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/util"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/util/log"
)

type IUserDomainService interface {
	LoadFollower(ctx context.Context, u *core_api.User) error
	LoadFollowing(ctx context.Context, u *core_api.User) error
	LoadRoles(ctx context.Context, u *core_api.User) error
	LoadEnableDebug(ctx context.Context, u *core_api.User) error
	LoadArticle(ctx context.Context, u *core_api.User) error
}

type UserDomainService struct {
	MeowchatUser    meowchat_user.IMeowchatUser
	MeowchatContent meowchat_content.IMeowchatContent
	MeowchatSystem  meowchat_system.IMeowchatSystem
}

var UserDomainServiceSet = wire.NewSet(
	wire.Struct(new(UserDomainService), "*"),
	wire.Bind(new(IUserDomainService), new(*UserDomainService)),
)

func (s *UserDomainService) LoadFollower(ctx context.Context, u *core_api.User) error {
	follower, err := s.MeowchatUser.GetTargetLikes(ctx, &genuser.GetTargetLikesReq{
		TargetId: u.Id,
		Type:     genuser.LikeType_User,
	})
	if err != nil {
		log.CtxError(ctx, "[LoadFollower] fail, err=%v", err)
		return err
	}
	u.Follower = lo.ToPtr(follower.GetCount())
	return nil
}

func (s *UserDomainService) LoadFollowing(ctx context.Context, u *core_api.User) error {
	followee, err := s.MeowchatUser.GetUserLikes(ctx, &genuser.GetUserLikesReq{
		UserId: u.Id,
		Type:   genuser.LikeType_User,
	})
	if err != nil {
		log.CtxError(ctx, "[LoadFollowing] fail, err=%v", err)
		return err
	}
	u.Following = lo.ToPtr(int64(len(followee.GetLikes())))
	return nil
}

func (s *UserDomainService) LoadRoles(ctx context.Context, u *core_api.User) error {
	rpcResp, err := s.MeowchatSystem.RetrieveUserRole(ctx, &gensystem.RetrieveUserRoleReq{UserId: u.Id})
	if err != nil {
		return err
	}
	u.Roles = lo.Map(rpcResp.GetRoles(), func(role *gensystem.Role, _ int) *system.Role {
		return &system.Role{
			RoleType:    system.RoleType(role.RoleType),
			CommunityId: role.CommunityId,
		}
	})
	return nil
}

func (s *UserDomainService) LoadEnableDebug(ctx context.Context, u *core_api.User) error {
	rpcResp, err := s.MeowchatSystem.RetrieveUserRole(ctx, &gensystem.RetrieveUserRoleReq{UserId: u.Id})
	if err != nil {
		return err
	}
	u.EnableDebug = lo.ToPtr(false)
	for _, role := range rpcResp.Roles {
		if role.RoleType == gensystem.RoleType_TypeDeveloper {
			u.EnableDebug = lo.ToPtr(true)
		}
	}
	return nil
}

func (s *UserDomainService) LoadArticle(ctx context.Context, u *core_api.User) error {
	momentCount := lo.Empty[*int64]()
	postCount := lo.Empty[*int64]()
	util.ParallelRun([]func(){
		func() {
			rpcResp, err := s.MeowchatContent.CountPost(ctx, &content.CountPostReq{
				FilterOptions: &content.PostFilterOptions{
					OnlyUserId: lo.ToPtr(u.Id),
				},
			})
			if err != nil {
				log.CtxError(ctx, "[LoadArticle] load post fail, err=%v", err)
				return
			}
			postCount = lo.ToPtr(rpcResp.GetTotal())
		},
		func() {
			rpcResp, err := s.MeowchatContent.CountMoment(ctx, &content.CountMomentReq{
				FilterOptions: &content.MomentFilterOptions{OnlyUserId: &u.Id},
			})
			if err != nil {
				log.CtxError(ctx, "[LoadArticle] load moment fail, err=%v", err)
				return
			}
			momentCount = lo.ToPtr(rpcResp.GetTotal())
		},
	})
	if momentCount == nil || postCount == nil {
		return errors.New("load article fail")
	}
	u.Article = lo.ToPtr(*momentCount + *postCount)
	return nil
}

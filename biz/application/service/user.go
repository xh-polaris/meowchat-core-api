package service

import (
	"context"
	"net/url"

	"github.com/google/wire"
	"github.com/samber/lo"
	"github.com/xh-polaris/gopkg/errors"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowchat/content"
	genuser "github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowchat/user"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/platform/sts"

	"github.com/xh-polaris/meowchat-core-api/biz/adaptor"
	"github.com/xh-polaris/meowchat-core-api/biz/application/dto/meowchat/core_api"
	"github.com/xh-polaris/meowchat-core-api/biz/domain/service"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/config"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/consts"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/meowchat_content"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/meowchat_user"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/platform_sts"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/util"
)

type IUserService interface {
	GetUserInfo(ctx context.Context, req *core_api.GetUserInfoReq) (*core_api.GetUserInfoResp, error)
	SearchUser(ctx context.Context, req *core_api.SearchUserReq) (*core_api.SearchUserResp, error)
	UpdateUserInfo(ctx context.Context, req *core_api.UpdateUserInfoReq) (*core_api.UpdateUserInfoResp, error)
	CheckIn(ctx context.Context, req *core_api.CheckInReq) (*core_api.CheckInResp, error)
}

type UserService struct {
	Config          *config.Config
	UserService     service.IUserDomainService
	MeowchatUser    meowchat_user.IMeowchatUser
	PlatformSts     platform_sts.IPlatformSts
	MeowchatContent meowchat_content.IMeowchatContent
}

var UserServiceSet = wire.NewSet(
	wire.Struct(new(UserService), "*"),
	wire.Bind(new(IUserService), new(*UserService)),
)

func (s *UserService) GetUserInfo(ctx context.Context, req *core_api.GetUserInfoReq) (*core_api.GetUserInfoResp, error) {
	resp := new(core_api.GetUserInfoResp)

	user := adaptor.ExtractUserMeta(ctx)
	if user.WechatUserMeta != nil {
		_, _ = s.PlatformSts.AddUserAuth(ctx, &sts.AddUserAuthReq{
			UserId:  user.UserId,
			Type:    "wechat",
			UnionId: user.WechatUserMeta.UnionId,
		})
	}

	var userId string
	if req.GetUserId() != "" {
		userId = *req.UserId
	} else if user.GetUserId() == "" {
		return nil, consts.ErrNotAuthentication
	} else {
		userId = user.GetUserId()
	}

	data, err := s.MeowchatUser.GetUserDetail(ctx, &genuser.GetUserDetailReq{UserId: userId})
	if err != nil {
		return nil, err
	}
	resp.User = &core_api.User{
		Id:        data.User.Id,
		Nickname:  data.User.Nickname,
		AvatarUrl: data.User.AvatarUrl,
		Motto:     lo.ToPtr(data.User.Motto),
	}
	util.ParallelRun([]func(){
		func() {
			_ = s.UserService.LoadArticle(ctx, resp.User)
		},
		func() {
			_ = s.UserService.LoadFollowing(ctx, resp.User)
		},
		func() {
			_ = s.UserService.LoadFollower(ctx, resp.User)
		},
		func() {
			_ = s.UserService.LoadEnableDebug(ctx, resp.User)
		},
	})

	return resp, nil
}

func (s *UserService) SearchUser(ctx context.Context, req *core_api.SearchUserReq) (*core_api.SearchUserResp, error) {
	resp := new(core_api.SearchUserResp)
	var pageSize int64 = 10
	if *req.PaginationOption.Limit != 0 {
		pageSize = *req.PaginationOption.Limit
	}
	request := &genuser.SearchUserReq{
		Nickname: req.Keyword,
		Offset:   lo.ToPtr(*req.PaginationOption.Page * pageSize),
		Limit:    lo.ToPtr(pageSize),
	}
	if req.PaginationOption.LastToken != nil && *req.PaginationOption.LastToken != "" {
		request.LastToken = req.PaginationOption.LastToken
	}
	data, err := s.MeowchatUser.SearchUser(ctx, request)
	if err != nil {
		return nil, err
	}
	resp.Total = data.Total
	resp.Token = data.Token
	resp.Users = make([]*core_api.User, len(data.Users))
	util.ParallelRun(lo.Map(data.Users, func(user *genuser.UserPreview, i int) func() {
		return func() {
			u := &core_api.User{
				Id:        user.Id,
				Nickname:  user.Nickname,
				AvatarUrl: user.AvatarUrl,
			}
			_ = s.UserService.LoadRoles(ctx, u)
			resp.Users[i] = u
		}
	}))
	return resp, nil
}

func (s *UserService) UpdateUserInfo(ctx context.Context, req *core_api.UpdateUserInfoReq) (*core_api.UpdateUserInfoResp, error) {
	resp := new(core_api.UpdateUserInfoResp)

	user := adaptor.ExtractUserMeta(ctx)
	if user.GetUserId() == "" {
		return nil, consts.ErrNotAuthentication
	}
	if req.GetNickname() != "" {
		r, err := s.PlatformSts.TextCheck(ctx, &sts.TextCheckReq{
			Text:  *req.Nickname,
			User:  user,
			Scene: 2,
			Title: req.Nickname,
		})
		if err != nil {
			return nil, err
		}
		if r.Pass == false {
			return nil, errors.NewBizError(10001, "TextCheck don't pass")
		}
	}

	if req.GetAvatarUrl() != "" {
		var u *url.URL
		u, err := url.Parse(*req.AvatarUrl)
		if err != nil {
			return resp, nil
		}
		u.Host = s.Config.CdnHost
		*req.AvatarUrl = u.String()
		if err != nil {
			return nil, err
		}
		var i = []string{*req.AvatarUrl}
		r, err := s.PlatformSts.PhotoCheck(ctx, &sts.PhotoCheckReq{
			User: user,
			Url:  i,
		})
		if err != nil {
			return nil, err
		}
		if r.Pass == false {
			return nil, errors.NewBizError(10002, "PhotoCheck don't pass")
		}
	}

	_, err := s.MeowchatUser.UpdateUser(ctx, &genuser.UpdateUserReq{
		User: &genuser.UserDetail{
			Id:        user.UserId,
			AvatarUrl: req.GetAvatarUrl(),
			Nickname:  req.GetNickname(),
			Motto:     req.GetMotto(),
		},
	})
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *UserService) CheckIn(ctx context.Context, req *core_api.CheckInReq) (*core_api.CheckInResp, error) {
	user := adaptor.ExtractUserMeta(ctx)
	if user.GetUserId() == "" {
		return nil, consts.ErrNotAuthentication
	}
	resp := new(core_api.CheckInResp)

	rpcResp, err := s.MeowchatUser.CheckIn(ctx, &genuser.CheckInReq{
		UserId: user.GetUserId(),
	})
	if err != nil {
		return nil, err
	}
	if rpcResp.GetGetFish() == true {
		_, err = s.MeowchatContent.AddUserFish(ctx, &content.AddUserFishReq{
			UserId: user.GetUserId(),
			Fish:   s.Config.Fish.SignIn[rpcResp.GetFishTimes-1],
		})
		if err == nil {
			resp.GetFishNum = s.Config.Fish.SignIn[rpcResp.GetFishTimes-1]
			resp.GetFishTimes = rpcResp.GetFishTimes
		}
	}
	resp.GetFish = rpcResp.GetGetFish()
	return resp, nil
}

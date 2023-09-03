package service

import (
	"context"
	"github.com/google/wire"
	"github.com/xh-polaris/gopkg/errors"
	"github.com/xh-polaris/meowchat-core-api/biz/application/dto/meowchat/core_api"
	"github.com/xh-polaris/meowchat-core-api/biz/application/dto/meowchat/system"
	user1 "github.com/xh-polaris/meowchat-core-api/biz/application/dto/meowchat/user"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/config"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/meowchat_content"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/meowchat_system"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/meowchat_user"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/platform_sts"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/basic"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowchat/content"
	system2 "github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowchat/system"
	genuser "github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowchat/user"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/platform/sts"
	"github.com/zeromicro/go-zero/core/logx"
	"net/url"
	"sync"
)

type IUserService interface {
	GetUserInfo(ctx context.Context, req *core_api.GetUserInfoReq, user *basic.UserMeta) (*core_api.GetUserInfoResp, error)
	SearchUserForAdmin(ctx context.Context, req *core_api.SearchUserForAdminReq) (*core_api.SearchUserForAdminResp, error)
	SearchUser(ctx context.Context, req *core_api.SearchUserReq) (*core_api.SearchUserResp, error)
	UpdateUserInfo(ctx context.Context, req *core_api.UpdateUserInfoReq, user *basic.UserMeta) (*core_api.UpdateUserInfoResp, error)
}

type UserService struct {
	Config *config.Config
	User   meowchat_user.IMeowchatUser
	Moment meowchat_content.IMeowchatContent
	System meowchat_system.IMeowchatSystem
	Sts    platform_sts.IPlatformSts
}

var UserServiceSet = wire.NewSet(
	wire.Struct(new(UserService), "*"),
	wire.Bind(new(IUserService), new(*UserService)),
)

func (s *UserService) GetUserInfo(ctx context.Context, req *core_api.GetUserInfoReq, user *basic.UserMeta) (*core_api.GetUserInfoResp, error) {
	resp := new(core_api.GetUserInfoResp)

	var userId string
	if req.UserId != nil {
		userId = *req.UserId
	} else {
		userId = user.GetUserId()
	}

	data, err := s.User.GetUserDetail(ctx, &genuser.GetUserDetailReq{UserId: userId})
	if err != nil {
		return nil, err
	}
	res, err := s.System.RetrieveUserRole(ctx, &system2.RetrieveUserRoleReq{UserId: userId})
	if err != nil {
		return nil, err
	}
	resp.EnableDebug = false
	if res != nil {
		for _, role := range res.Roles {
			if role.RoleType == system2.RoleType_TypeDeveloper {
				resp.EnableDebug = true
			}
		}
	}
	resp.User = &core_api.User{
		Id:        data.User.Id,
		Nickname:  data.User.Nickname,
		AvatarUrl: data.User.AvatarUrl,
		Motto:     data.User.Motto,
	}

	s.getLessDependentInfo(ctx, resp.User)

	return resp, nil
}

func (s *UserService) SearchUserForAdmin(ctx context.Context, req *core_api.SearchUserForAdminReq) (*core_api.SearchUserForAdminResp, error) {
	resp := new(core_api.SearchUserForAdminResp)
	var pageSize int64 = 10
	if *req.PaginationOption.Limit != 0 {
		pageSize = *req.PaginationOption.Limit
	}
	request := &genuser.SearchUserReq{
		Nickname: req.Keyword,
		Offset:   new(int64),
		Limit:    new(int64),
	}
	if *req.PaginationOption.LastToken != "" {
		request.LastToken = req.PaginationOption.LastToken
	}
	*request.Offset = *req.PaginationOption.Page * pageSize
	*request.Limit = pageSize
	data, err := s.User.SearchUser(ctx, request)
	if err != nil {
		return nil, err
	}
	resp.Total = data.Total
	resp.Token = data.Token
	resp.Users = make([]*core_api.UserPreviewWithRole, 0, len(data.Users))
	for _, user := range data.Users {
		u := core_api.UserPreviewWithRole{
			User: &user1.UserPreview{
				Id:        user.Id,
				Nickname:  user.Nickname,
				AvatarUrl: user.AvatarUrl,
			},
		}

		u.Roles = make([]*system.Role, 0)
		data, err := s.System.RetrieveUserRole(ctx, &system2.RetrieveUserRoleReq{UserId: user.Id})
		if err != nil {
			return nil, err
		}
		for _, role := range data.Roles {
			u.Roles = append(u.Roles, &system.Role{
				RoleType:    system.RoleType(role.RoleType),
				CommunityId: role.CommunityId,
			})
		}

		resp.Users = append(resp.Users, &u)
	}
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
		Offset:   new(int64),
		Limit:    new(int64),
	}
	if *req.PaginationOption.LastToken != "" {
		request.LastToken = req.PaginationOption.LastToken
	}
	*request.Offset = *req.PaginationOption.Page * pageSize
	*request.Limit = pageSize
	data, err := s.User.SearchUser(ctx, request)
	if err != nil {
		return nil, err
	}
	resp.Total = data.Total
	resp.Token = data.Token
	resp.Users = make([]*user1.UserPreview, 0, len(data.Users))
	for _, user := range data.Users {
		resp.Users = append(resp.Users, &user1.UserPreview{
			Id:        user.Id,
			Nickname:  user.Nickname,
			AvatarUrl: user.AvatarUrl,
		})
	}
	return resp, nil
}

func (s *UserService) UpdateUserInfo(ctx context.Context, req *core_api.UpdateUserInfoReq, user *basic.UserMeta) (*core_api.UpdateUserInfoResp, error) {
	resp := new(core_api.UpdateUserInfoResp)

	if req.GetNickname() != "" {
		r, err := s.Sts.TextCheck(ctx, &sts.TextCheckReq{
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

	if *req.AvatarUrl != "" {
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
		r, err := s.Sts.PhotoCheck(ctx, &sts.PhotoCheckReq{
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

	_, err := s.User.UpdateUser(ctx, &genuser.UpdateUserReq{
		User: &genuser.UserDetail{
			Id:        user.UserId,
			AvatarUrl: *req.AvatarUrl,
			Nickname:  *req.Nickname,
			Motto:     *req.Motto,
		},
	})
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *UserService) getLessDependentInfo(ctx context.Context, user *core_api.User) {
	wg := &sync.WaitGroup{}
	wg.Add(4)

	go func() {
		defer wg.Done()
		momentCount, err := s.Moment.CountMoment(ctx, &content.CountMomentReq{
			FilterOptions: &content.MomentFilterOptions{OnlyUserId: &user.Id},
		})
		if err != nil {
			logx.Error(err)
		}
		user.Article += momentCount.GetTotal()
	}()

	go func() {
		defer wg.Done()
		postCount, err := s.Moment.CountPost(ctx, &content.CountPostReq{
			FilterOptions: &content.PostFilterOptions{
				OnlyUserId: &user.Id,
			},
		})
		if err != nil {
			logx.Error(err)
		}
		// TODO 偶尔有并发问题
		user.Article += postCount.GetTotal()
	}()

	go func() {
		defer wg.Done()
		follower, err := s.User.GetTargetLikes(ctx, &genuser.GetTargetLikesReq{
			TargetId: user.Id,
			Type:     genuser.LikeType_User,
		})
		if err != nil {
			logx.Error(err)
		}
		user.Follower = follower.GetCount()
	}()

	go func() {
		defer wg.Done()
		followee, err := s.User.GetUserLikes(ctx, &genuser.GetUserLikesReq{
			UserId: user.Id,
			Type:   genuser.LikeType_User,
		})
		if err != nil {
			logx.Error(err)
		}
		user.Following = int64(len(followee.GetLikes()))
	}()

	wg.Wait()
}

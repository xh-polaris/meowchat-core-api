package service

import (
	"context"
	"github.com/google/wire"
	"github.com/xh-polaris/meowchat-core-api/biz/application/dto/meowchat/core_api"
	user1 "github.com/xh-polaris/meowchat-core-api/biz/application/dto/meowchat/user"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/config"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/meowchat_like"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/meowchat_moment"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/meowchat_post"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/meowchat_user"
	"github.com/xh-polaris/meowchat-like-rpc/likerpc"
	pb4 "github.com/xh-polaris/meowchat-like-rpc/pb"
	pb2 "github.com/xh-polaris/meowchat-moment-rpc/pb"
	pb3 "github.com/xh-polaris/meowchat-post-rpc/pb"
	"github.com/xh-polaris/meowchat-user-rpc/pb"
	"github.com/zeromicro/go-zero/core/logx"
	"net/url"
	"sync"
)

type IUserService interface {
	GetUserInfo(ctx context.Context, req *core_api.GetUserInfoReq) (*core_api.GetUserInfoResp, error)
	SearchUserForAdmin(ctx context.Context, req *core_api.SearchUserForAdminReq) (*core_api.SearchUserForAdminResp, error)
	SearchUser(ctx context.Context, req *core_api.SearchUserReq) (*core_api.SearchUserResp, error)
	UpdateUserInfo(ctx context.Context, req *core_api.UpdateUserInfoReq) (*core_api.UpdateUserInfoResp, error)
}

type UserService struct {
	Config *config.Config
	User   meowchat_user.IMeowchatUser
	Moment meowchat_moment.IMeowchatMoment
	Like   meowchat_like.IMeowchatLike
	Post   meowchat_post.IMeowchatPost
}

var UserServiceSet = wire.NewSet(
	wire.Struct(new(UserService), "*"),
	wire.Bind(new(IUserService), new(*UserService)),
)

func (s *UserService) GetUserInfo(ctx context.Context, req *core_api.GetUserInfoReq) (*core_api.GetUserInfoResp, error) {
	resp := new(core_api.GetUserInfoResp)

	var userId string
	if req.UserId != nil {
		userId = *req.UserId
	} else {
		userId = ctx.Value("userId").(string)
	}

	data, err := s.User.GetUserDetail(ctx, &pb.GetUserDetailReq{UserId: userId})
	if err != nil {
		return nil, err
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
	//resp := new(core_api.SearchUserForAdminResp)
	//var pageSize int64 = 10
	//if *req.PaginationOption.Limit != 0 {
	//	pageSize = *req.PaginationOption.Limit
	//}
	//request := &pb.SearchUserReq{
	//	Nickname: req.Keyword,
	//	Offset:   new(int64),
	//	Limit:    new(int64),
	//}
	//if *req.PaginationOption.LastToken != "" {
	//	request.LastToken = req.PaginationOption.LastToken
	//}
	//*request.Offset = *req.PaginationOption.Page * pageSize
	//*request.Limit = pageSize
	//data, err := s.User.SearchUser(ctx, request)
	//if err != nil {
	//	return nil, err
	//}
	//resp.Total = data.Total
	//resp.Token = data.Token
	//resp.Users = make([]*core_api.UserPreviewWithRole, 0, len(data.Users))
	//for _, user := range data.Users {
	//	u := core_api.UserPreviewWithRole{
	//		User: &user1.UserPreview{
	//			Id:        user.Id,
	//			Nickname:  user.Nickname,
	//			AvatarUrl: user.AvatarUrl,
	//		},
	//	}
	//
	//	u.Roles = make([]*system.Role, 0)
	//	data, err := s.System.RetrieveUserRole(l.ctx, &pb2.RetrieveUserRoleReq{UserId: user.Id})
	//	if err != nil {
	//		return nil, err
	//	}
	//	for _, role := range data.Roles {
	//		u.Roles = append(u.Roles, core_api.Role{
	//			RoleType:    role.Type,
	//			CommunityId: role.CommunityId,
	//		})
	//	}
	//
	//	resp.Users = append(resp.Users, u)
	//}
	return nil, nil
}

func (s *UserService) SearchUser(ctx context.Context, req *core_api.SearchUserReq) (*core_api.SearchUserResp, error) {
	resp := new(core_api.SearchUserResp)
	var pageSize int64 = 10
	if *req.PaginationOption.Limit != 0 {
		pageSize = *req.PaginationOption.Limit
	}
	request := &pb.SearchUserReq{
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

func (s *UserService) UpdateUserInfo(ctx context.Context, req *core_api.UpdateUserInfoReq) (*core_api.UpdateUserInfoResp, error) {
	resp := new(core_api.UpdateUserInfoResp)
	userId := ctx.Value("userId").(string)
	//openId := ctx.Value("openId").(string)
	//
	//err = util.MsgSecCheck(l.ctx, l.svcCtx, req.Nickname, openId, 2)
	//if err != nil {
	//	return nil, err
	//}

	if *req.AvatarUrl != "" {
		var u *url.URL
		u, err := url.Parse(*req.AvatarUrl)
		if err != nil {
			return resp, nil
		}
		u.Host = s.Config.CdnHost
		*req.AvatarUrl = u.String()
		//var r = []string{*req.AvatarUrl}
		//err = util.PhotoCheck(ctx, s, r)
		if err != nil {
			return nil, err
		}
	}

	_, err := s.User.UpdateUser(ctx, &pb.UpdateUserReq{
		User: &pb.UserDetail{
			Id:        userId,
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
		momentCount, err := s.Moment.CountMoment(ctx, &pb2.CountMomentReq{
			FilterOptions: &pb2.FilterOptions{OnlyUserId: &user.Id},
		})
		if err != nil {
			logx.Error(err)
		}
		user.Article += momentCount.Total
	}()

	go func() {
		defer wg.Done()
		postCount, err := s.Post.CountPost(ctx, &pb3.CountPostReq{
			FilterOptions: &pb3.FilterOptions{
				OnlyUserId: &user.Id,
			},
		})
		if err != nil {
			logx.Error(err)
		}
		// TODO 偶尔有并发问题
		user.Article += postCount.Total
	}()

	go func() {
		defer wg.Done()
		follower, err := s.Like.GetTargetLikes(ctx, &pb4.GetTargetLikesReq{
			TargetId: user.Id,
			Type:     likerpc.TargetTypeUser,
		})
		if err != nil {
			logx.Error(err)
		}
		user.Follower = follower.Count
	}()

	go func() {
		defer wg.Done()
		followee, err := s.Like.GetUserLikes(ctx, &pb4.GetUserLikesReq{
			UserId:     user.Id,
			TargetType: likerpc.TargetTypeUser,
		})
		if err != nil {
			logx.Error(err)
		}
		user.Following = int64(len(followee.Likes))
	}()

	wg.Wait()
}

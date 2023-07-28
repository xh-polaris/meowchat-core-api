package service

import (
	"context"
	"github.com/google/wire"
	"github.com/xh-polaris/meowchat-core-api/biz/application/dto/meowchat/core_api"
	"github.com/xh-polaris/meowchat-core-api/biz/application/dto/meowchat/user"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/config"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/meowchat_user"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/basic"
	genlike "github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowchat/user"
	"github.com/zeromicro/go-zero/core/logx"
)

type ILikeService interface {
	DoLike(ctx context.Context, req *core_api.DoLikeReq, user *basic.UserMeta) (*core_api.DoLikeResp, error)
	GetLikedCount(ctx context.Context, req *core_api.GetLikedCountReq) (*core_api.GetLikedCountResp, error)
	GetLikedUsers(ctx context.Context, req *core_api.GetLikedUsersReq) (*core_api.GetLikedUsersResp, error)
	GetUserLiked(ctx context.Context, req *core_api.GetUserLikedReq, user *basic.UserMeta) (*core_api.GetUserLikedResp, error)
	GetUserLikes(ctx context.Context, req *core_api.GetUserLikesReq) (*core_api.GetUserLikesResp, error)
}

type LikeService struct {
	Config *config.Config
	User   meowchat_user.IMeowchatUser
}

var LikeServiceSet = wire.NewSet(
	wire.Struct(new(LikeService), "*"),
	wire.Bind(new(ILikeService), new(*LikeService)),
)

func (s *LikeService) DoLike(ctx context.Context, req *core_api.DoLikeReq, user *basic.UserMeta) (*core_api.DoLikeResp, error) {
	resp := new(core_api.DoLikeResp)

	userId := user.UserId

	_, err := s.User.DoLike(ctx, &genlike.DoLikeReq{
		UserId:       userId,
		TargetId:     req.TargetId,
		Type:         req.TargetType,
		AssociatedId: "",
	})

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *LikeService) GetLikedCount(ctx context.Context, req *core_api.GetLikedCountReq) (*core_api.GetLikedCountResp, error) {
	resp := new(core_api.GetLikedCountResp)

	likes, err := s.User.GetTargetLikes(ctx, &genlike.GetTargetLikesReq{
		TargetId: req.TargetId,
		Type:     req.TargetType,
	})
	if err != nil {
		return nil, err
	}

	resp.Count = likes.Count

	return resp, nil
}

func (s *LikeService) GetLikedUsers(ctx context.Context, req *core_api.GetLikedUsersReq) (*core_api.GetLikedUsersResp, error) {
	resp := new(core_api.GetLikedUsersResp)
	data, err := s.User.GetLikedUsers(ctx, &genlike.GetLikedUsersReq{
		TargetId:   req.TargetId,
		TargetType: req.TargetType,
	})
	if err != nil {
		return nil, err
	}
	resp.Users = make([]*user.UserPreview, 0, len(data.UserIds))
	for _, userId := range data.UserIds {
		res, err := s.User.GetUser(ctx, &genlike.GetUserReq{UserId: userId})
		if err != nil {
			logx.Error(err)
		}
		resp.Users = append(resp.Users, &user.UserPreview{
			Id:        res.User.Id,
			Nickname:  res.User.Nickname,
			AvatarUrl: res.User.AvatarUrl,
		})
	}
	return resp, nil
}

func (s *LikeService) GetUserLiked(ctx context.Context, req *core_api.GetUserLikedReq, user *basic.UserMeta) (*core_api.GetUserLikedResp, error) {
	resp := new(core_api.GetUserLikedResp)

	userId := user.UserId
	like, err := s.User.GetUserLike(ctx, &genlike.GetUserLikedReq{
		UserId:   userId,
		TargetId: req.TargetId,
		Type:     req.TargetType,
	})
	if err != nil {
		return nil, err
	}

	resp.Liked = like.Liked

	return resp, nil
}

func (s *LikeService) GetUserLikes(ctx context.Context, req *core_api.GetUserLikesReq) (*core_api.GetUserLikesResp, error) {
	resp := new(core_api.GetUserLikesResp)
	data, err := s.User.GetUserLikes(ctx, &genlike.GetUserLikesReq{
		UserId:     req.UserId,
		TargetType: req.TargetType,
	})
	if err != nil {
		return nil, err
	}
	resp.Likes = make([]*user.Like, 0, len(data.Likes))
	for _, like := range data.Likes {
		resp.Likes = append(resp.Likes, &user.Like{
			TargetId:     like.TargetId,
			AssociatedId: like.AssociatedId,
		})
	}
	return resp, nil
}

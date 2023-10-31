package service

import (
	"context"

	"github.com/google/wire"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/basic"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowchat/content"
	genlike "github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowchat/user"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/xh-polaris/meowchat-core-api/biz/application/dto/meowchat/core_api"
	"github.com/xh-polaris/meowchat-core-api/biz/application/dto/meowchat/user"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/config"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/meowchat_content"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/meowchat_user"
)

type ILikeService interface {
	DoLike(ctx context.Context, req *core_api.DoLikeReq, user *basic.UserMeta) (*core_api.DoLikeResp, error)
	GetLikedCount(ctx context.Context, req *core_api.GetLikedCountReq) (*core_api.GetLikedCountResp, error)
	GetLikedUsers(ctx context.Context, req *core_api.GetLikedUsersReq) (*core_api.GetLikedUsersResp, error)
	GetUserLiked(ctx context.Context, req *core_api.GetUserLikedReq, user *basic.UserMeta) (*core_api.GetUserLikedResp, error)
	GetUserLikes(ctx context.Context, req *core_api.GetUserLikesReq) (*core_api.GetUserLikesResp, error)
}

type LikeService struct {
	Config  *config.Config
	User    meowchat_user.IMeowchatUser
	Content meowchat_content.IMeowchatContent
}

var LikeServiceSet = wire.NewSet(
	wire.Struct(new(LikeService), "*"),
	wire.Bind(new(ILikeService), new(*LikeService)),
)

func (s *LikeService) DoLike(ctx context.Context, req *core_api.DoLikeReq, user *basic.UserMeta) (*core_api.DoLikeResp, error) {
	resp := new(core_api.DoLikeResp)

	userId := user.UserId

	r, err := s.User.DoLike(ctx, &genlike.DoLikeReq{
		UserId:       userId,
		TargetId:     req.TargetId,
		Type:         genlike.LikeType(req.TargetType),
		AssociatedId: "",
	})

	if err != nil {
		return nil, err
	}
	if r.GetIsFirst() == true {
		_, err = s.Content.AddUserFish(ctx, &content.AddUserFishReq{
			UserId: user.UserId,
			Fish:   s.Config.Fish.Like,
		})
		if err == nil {
			resp.GetFishNum = s.Config.Fish.Like
		}
	}
	resp.IsFirst = r.IsFirst
	return resp, nil
}

func (s *LikeService) GetLikedCount(ctx context.Context, req *core_api.GetLikedCountReq) (*core_api.GetLikedCountResp, error) {
	resp := new(core_api.GetLikedCountResp)

	likes, err := s.User.GetTargetLikes(ctx, &genlike.GetTargetLikesReq{
		TargetId: req.TargetId,
		Type:     genlike.LikeType(req.TargetType),
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
		TargetId: req.TargetId,
		Type:     genlike.LikeType(req.TargetType),
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
		Type:     genlike.LikeType(req.TargetType),
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
		UserId: req.UserId,
		Type:   genlike.LikeType(req.TargetType),
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

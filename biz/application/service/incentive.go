package service

import (
	"context"

	"github.com/google/wire"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowchat/content"

	"github.com/xh-polaris/meowchat-core-api/biz/adaptor"
	"github.com/xh-polaris/meowchat-core-api/biz/application/dto/meowchat/core_api"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/config"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/consts"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/meowchat_content"
)

type IIncentiveService interface {
	CheckIn(ctx context.Context, req *core_api.CheckInReq) (*core_api.CheckInResp, error)
	GetMission(ctx context.Context, req *core_api.GetMissionReq) (*core_api.GetMissionResp, error)
}

type IncentiveService struct {
	Config          *config.Config
	MeowchatContent meowchat_content.IMeowchatContent
}

var IncentiveServiceSet = wire.NewSet(
	wire.Struct(new(IncentiveService), "*"),
	wire.Bind(new(IIncentiveService), new(*IncentiveService)),
)

func (s *IncentiveService) CheckIn(ctx context.Context, req *core_api.CheckInReq) (*core_api.CheckInResp, error) {
	user := adaptor.ExtractUserMeta(ctx)
	if user.GetUserId() == "" {
		return nil, consts.ErrNotAuthentication
	}
	resp := new(core_api.CheckInResp)

	rpcResp, err := s.MeowchatContent.CheckIn(ctx, &content.CheckInReq{
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

func (s *IncentiveService) GetMission(ctx context.Context, req *core_api.GetMissionReq) (*core_api.GetMissionResp, error) {
	resp := new(core_api.GetMissionResp)

	user := adaptor.ExtractUserMeta(ctx)

	mission, err := s.MeowchatContent.GetMission(ctx, &content.GetMissionReq{UserId: user.UserId})
	if err != nil {
		return nil, err
	}
	resp.CommentTime = mission.GetCommentTime()
	resp.LikeTime = mission.GetLikeTime()
	resp.SignInTime = mission.GetSignInTime()
	resp.ContentTime = mission.GetContentTime()
	resp.SignInFishes = s.Config.Fish.SignIn
	resp.CommentFishes = s.Config.Fish.Comment
	resp.ContentFishes = s.Config.Fish.Content
	resp.LikeFishes = s.Config.Fish.Like
	return resp, nil
}

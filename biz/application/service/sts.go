package service

import (
	"context"
	"net/http"

	"github.com/google/wire"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/basic"

	"github.com/google/uuid"

	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/platform/sts"

	"github.com/xh-polaris/meowchat-core-api/biz/application/dto/meowchat/core_api"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/consts"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/platform_sts"
)

type IStsService interface {
	ApplySignedUrl(ctx context.Context, req *core_api.ApplySignedUrlReq, user *basic.UserMeta) (*core_api.ApplySignedUrlResp, error)
	ApplySignedUrlAsCommunity(ctx context.Context, req *core_api.ApplySignedUrlAsCommunityReq, user *basic.UserMeta) (*core_api.ApplySignedUrlAsCommunityResp, error)
}

type StsService struct {
	PlatformSts platform_sts.IPlatformSts
}

var StsServiceSet = wire.NewSet(
	wire.Struct(new(StsService), "*"),
	wire.Bind(new(IStsService), new(*StsService)),
)

func (s *StsService) ApplySignedUrl(ctx context.Context, req *core_api.ApplySignedUrlReq, user *basic.UserMeta) (*core_api.ApplySignedUrlResp, error) {
	if user.GetUserId() == "" {
		return nil, consts.ErrNotAuthentication
	}
	resp := new(core_api.ApplySignedUrlResp)
	userId := user.GetUserId()
	data, err := s.PlatformSts.GenCosSts(ctx, &sts.GenCosStsReq{Path: "users/" + userId + "/*"})
	if err != nil {
		return nil, err
	}
	resp.SessionToken = data.SessionToken
	if req.Prefix != nil {
		*req.Prefix += "/"
	}
	data2, err := s.PlatformSts.GenSignedUrl(ctx, &sts.GenSignedUrlReq{
		SecretId:  data.SecretId,
		SecretKey: data.SecretKey,
		Method:    http.MethodPut,
		Path:      "users/" + userId + "/" + req.GetPrefix() + uuid.New().String() + req.GetSuffix(),
	})
	resp.Url = data2.SignedUrl
	return resp, nil
}

func (s *StsService) ApplySignedUrlAsCommunity(ctx context.Context, req *core_api.ApplySignedUrlAsCommunityReq, user *basic.UserMeta) (*core_api.ApplySignedUrlAsCommunityResp, error) {
	if user.GetUserId() == "" {
		return nil, consts.ErrNotAuthentication
	}
	resp := new(core_api.ApplySignedUrlAsCommunityResp)
	data, err := s.PlatformSts.GenCosSts(ctx, &sts.GenCosStsReq{Path: "communities/" + req.CommunityId + "/*"})
	if err != nil {
		return nil, err
	}
	resp.SessionToken = data.SessionToken
	if req.Prefix != "" {
		req.Prefix += "/"
	}
	data2, err := s.PlatformSts.GenSignedUrl(ctx, &sts.GenSignedUrlReq{
		SecretId:  data.SecretId,
		SecretKey: data.SecretKey,
		Method:    http.MethodPut,
		Path:      "communities/" + req.CommunityId + "/" + req.Prefix + uuid.New().String() + req.Suffix,
	})
	resp.Url = data2.SignedUrl
	return resp, nil
}

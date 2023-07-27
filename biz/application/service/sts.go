package service

import (
	"context"
	"net/http"

	"github.com/google/uuid"

	"github.com/xh-polaris/meowchat-core-api/biz/application/dto/meowchat/core_api"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/platform_sts"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/platform/sts"
)

type IStsService interface {
	ApplySignedUrl(ctx context.Context, req *core_api.ApplySignedUrlReq) (*core_api.ApplySignedUrlResp, error)
	ApplySignedUrlAsCommunity(ctx context.Context, req *core_api.ApplySignedUrlAsCommunityReq) (*core_api.ApplySignedUrlAsCommunityResp, error)
}

type StsService struct {
	PlatformSts platform_sts.IPlatformSts
}

func (s *StsService) ApplySignedUrl(ctx context.Context, req *core_api.ApplySignedUrlReq) (*core_api.ApplySignedUrlResp, error) {
	resp := new(core_api.ApplySignedUrlResp)
	userId := ctx.Value("userId").(string)
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

func (s *StsService) ApplySignedUrlAsCommunity(ctx context.Context, req *core_api.ApplySignedUrlAsCommunityReq) (*core_api.ApplySignedUrlAsCommunityResp, error) {
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

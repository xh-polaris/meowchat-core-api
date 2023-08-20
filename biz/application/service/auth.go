package service

import (
	"context"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/meowchat_content"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowchat/content"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/wire"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/basic"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/platform/sts"

	"github.com/xh-polaris/meowchat-core-api/biz/application/dto/meowchat/core_api"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/config"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/platform_sts"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/util"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/util/log"
)

type IAuthService interface {
	SignIn(ctx context.Context, req *core_api.SignInReq) (*core_api.SignInResp, error)
	SetPassword(ctx context.Context, req *core_api.SetPasswordReq, user *basic.UserMeta) (*core_api.SetPasswordResp, error)
	SendVerifyCode(ctx context.Context, req *core_api.SendVerifyCodeReq) (*core_api.SendVerifyCodeResp, error)
}

type AuthService struct {
	Config  *config.Config
	Sts     platform_sts.IPlatformSts
	Content meowchat_content.IMeowchatContent
}

var AuthServiceSet = wire.NewSet(
	wire.Struct(new(AuthService), "*"),
	wire.Bind(new(IAuthService), new(*AuthService)),
)

func (s *AuthService) SignIn(ctx context.Context, req *core_api.SignInReq) (*core_api.SignInResp, error) {
	resp := new(core_api.SignInResp)
	rpcResp, err := s.Sts.SignIn(ctx, &sts.SignInReq{
		AuthType:   req.GetAuthType(),
		AuthId:     req.GetAuthId(),
		Password:   req.Password,
		VerifyCode: req.VerifyCode,
	})
	if err != nil {
		return nil, err
	}

	auth := s.Config.Auth
	resp.AccessToken, resp.AccessExpire, err = generateJwtToken(req, rpcResp, auth.AccessSecret, auth.AccessExpire)
	if err != nil {
		log.CtxError(ctx, "[generateJwtToken] fail, err=%v, config=%s, resp=%s", err, util.JSONF(s.Config.Auth), util.JSONF(rpcResp))
		return nil, err
	}
	if rpcResp.GetIsFirst() == true {
		_, err = s.Content.AddUserFish(ctx, &content.AddUserFishReq{
			UserId: rpcResp.UserId,
			Fish:   s.Config.Fish.SignIn,
		})
	}
	resp.IsFirst = rpcResp.GetIsFirst()
	resp.UserId = rpcResp.GetUserId()
	return resp, nil
}

func generateJwtToken(req *core_api.SignInReq, resp *sts.SignInResp, secret string, expire int64) (string, int64, error) {
	iat := time.Now().Unix()
	exp := iat + expire
	claims := make(jwt.MapClaims)
	claims["exp"] = exp
	claims["iat"] = iat
	claims["userId"] = resp.GetUserId()
	claims["appId"] = req.GetAppId()
	claims["deviceId"] = req.GetDeviceId()
	claims["wechatUserMeta"] = &basic.WechatUserMeta{
		AppId:   resp.GetAppId(),
		OpenId:  resp.GetOpenId(),
		UnionId: resp.GetUnionId(),
	}
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", 0, err
	}
	return tokenString, exp, nil
}

func (s *AuthService) SetPassword(ctx context.Context, req *core_api.SetPasswordReq, user *basic.UserMeta) (*core_api.SetPasswordResp, error) {
	resp := new(core_api.SetPasswordResp)
	_, err := s.Sts.SetPassword(ctx, &sts.SetPasswordReq{
		UserId:   user.UserId,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *AuthService) SendVerifyCode(ctx context.Context, req *core_api.SendVerifyCodeReq) (*core_api.SendVerifyCodeResp, error) {
	resp := new(core_api.SendVerifyCodeResp)
	_, err := s.Sts.SendVerifyCode(ctx, &sts.SendVerifyCodeReq{
		AuthType: req.AuthType,
		AuthId:   req.AuthId,
	})
	if err != nil {
		return nil, err
	}

	return resp, nil
}

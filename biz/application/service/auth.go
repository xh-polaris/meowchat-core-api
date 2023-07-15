package service

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/wire"
	"github.com/xh-polaris/auth-rpc/pb"

	"github.com/xh-polaris/meowchat-core-api/biz/application/dto/meowchat/core_api"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/config"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/platform_authentication"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/util"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/util/log"
)

type IAuthService interface {
	SignIn(ctx context.Context, req *core_api.SignInReq) (*core_api.SignInResp, error)
	SetPassword(ctx context.Context, req *core_api.SetPasswordReq) (*core_api.SetPasswordResp, error)
	SendVerifyCode(ctx context.Context, req *core_api.SendVerifyCodeReq) (*core_api.SendVerifyCodeResp, error)
}

type AuthService struct {
	Config         *config.Config
	Authentication platform_authentication.IPlatformAuthentication
}

var AuthServiceSet = wire.NewSet(
	wire.Struct(new(AuthService), "*"),
	wire.Bind(new(IAuthService), new(*AuthService)),
)

func (s *AuthService) SignIn(ctx context.Context, req *core_api.SignInReq) (*core_api.SignInResp, error) {
	resp := new(core_api.SignInResp)
	rpcResp, err := s.Authentication.SignIn(ctx, &pb.SignInReq{
		AuthType: req.AuthType,
		AuthId:   req.AuthId,
		Password: req.GetPassword(),
		Params:   req.Params,
	})
	if err != nil {
		return nil, err
	}

	auth := s.Config.Auth
	resp.AccessToken, resp.AccessExpire, err = generateJwtToken(rpcResp.User, auth.AccessSecret, auth.AccessExpire)
	if err != nil {
		log.CtxError(ctx, "[generateJwtToken] fail, err=%v, config=%s, user=%s", err, util.JSONF(s.Config.Auth), util.JSONF(rpcResp.User))
		return nil, err
	}
	resp.UserId = rpcResp.User.UserId
	return resp, nil
}

func generateJwtToken(user *pb.User, secret string, expire int64) (string, int64, error) {
	iat := time.Now().Unix()
	exp := iat + expire
	claims := make(jwt.MapClaims)
	claims["exp"] = exp
	claims["iat"] = iat
	claims["userID"] = user.UserId
	claims["sessionUserID"] = user.UserId
	claims["appID"] = user.AppId
	claims["sessionAppID"] = user.AppId
	claims["unionID"] = user.UnionId
	claims["openID"] = user.OpenId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", 0, err
	}
	return tokenString, exp, nil
}

func (s *AuthService) SetPassword(ctx context.Context, req *core_api.SetPasswordReq) (*core_api.SetPasswordResp, error) {
	resp := new(core_api.SetPasswordResp)
	_, err := s.Authentication.SetPassword(ctx, &pb.SetPasswordReq{
		UserId:   req.User.UserID,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *AuthService) SendVerifyCode(ctx context.Context, req *core_api.SendVerifyCodeReq) (*core_api.SendVerifyCodeResp, error) {
	resp := new(core_api.SendVerifyCodeResp)
	_, err := s.Authentication.SendVerifyCode(ctx, &pb.SendVerifyCodeReq{
		AuthType: req.AuthType,
		AuthId:   req.AuthId,
	})
	if err != nil {
		return nil, err
	}

	return resp, nil
}

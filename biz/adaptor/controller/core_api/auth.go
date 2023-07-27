// Code generated by hertz generator.

package core_api

import (
	"context"

	"github.com/xh-polaris/meowchat-core-api/biz/adaptor"
	"github.com/xh-polaris/meowchat-core-api/biz/application/dto/meowchat/core_api"
	"github.com/xh-polaris/meowchat-core-api/provider"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// SignIn .
// @router /auth/sign_in [POST]
func SignIn(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.SignInReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	p := provider.Get()
	resp, err := p.AuthService.SignIn(ctx, &req)
	adaptor.Return(ctx, c, &req, resp, err)
}

// SendVerifyCode .
// @router /auth/send_verify_code [POST]
func SendVerifyCode(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.SendVerifyCodeReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	p := provider.Get()
	resp, err := p.AuthService.SendVerifyCode(ctx, &req)
	adaptor.Return(ctx, c, &req, resp, err)
}

// SetPassword .
// @router /auth/set_password [POST]
func SetPassword(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.SetPasswordReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	p := provider.Get()
	resp, err := p.AuthService.SetPassword(ctx, &req, adaptor.ExtractUserMeta(ctx, c))
	adaptor.Return(ctx, c, &req, resp, err)
}

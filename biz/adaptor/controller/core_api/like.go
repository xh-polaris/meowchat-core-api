// Code generated by hertz generator.

package core_api

import (
	"context"
	"github.com/xh-polaris/meowchat-core-api/biz/adaptor"
	"github.com/xh-polaris/meowchat-core-api/provider"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/xh-polaris/meowchat-core-api/biz/application/dto/meowchat/core_api"
)

// DoLike .
// @router /like/do_like [POST]
func DoLike(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.DoLikeReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	p := provider.Get()
	resp, err := p.LikeService.DoLike(ctx, &req, adaptor.ExtractUserMeta(ctx, c))
	adaptor.Return(ctx, c, &req, resp, err)
}

// GetUserLiked .
// @router /like/get_user_liked [GET]
func GetUserLiked(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.GetUserLikedReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	p := provider.Get()
	resp, err := p.LikeService.GetUserLiked(ctx, &req, adaptor.ExtractUserMeta(ctx, c))
	adaptor.Return(ctx, c, &req, resp, err)
}

// GetLikedCount .
// @router /like/get_count [GET]
func GetLikedCount(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.GetLikedCountReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	p := provider.Get()
	resp, err := p.LikeService.GetLikedCount(ctx, &req)
	adaptor.Return(ctx, c, &req, resp, err)
}

// GetLikedUsers .
// @router /like/get_liked_users [GET]
func GetLikedUsers(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.GetLikedUsersReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	p := provider.Get()
	resp, err := p.LikeService.GetLikedUsers(ctx, &req)
	adaptor.Return(ctx, c, &req, resp, err)
}

// GetUserLikes .
// @router /like/get_user_likes [GET]
func GetUserLikes(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.GetUserLikesReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	p := provider.Get()
	resp, err := p.LikeService.GetUserLikes(ctx, &req)
	adaptor.Return(ctx, c, &req, resp, err)
}

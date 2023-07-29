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

// GetPostPreviews .
// @router /post/get_post_previews [POST]
func GetPostPreviews(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.GetPostPreviewsReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	p := provider.Get()
	resp, err := p.PostService.GetPostPreviews(ctx, &req)
	adaptor.Return(ctx, c, &req, resp, err)
}

// GetPostDetail .
// @router /post/get_post_detail [GET]
func GetPostDetail(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.GetPostDetailReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	p := provider.Get()
	resp, err := p.PostService.GetPostDetail(ctx, &req)
	adaptor.Return(ctx, c, &req, resp, err)
}

// NewPost .
// @router /post/new_post [POST]
func NewPost(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.NewPostReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	p := provider.Get()
	resp, err := p.PostService.NewPost(ctx, &req, adaptor.ExtractUserMeta(ctx, c))
	adaptor.Return(ctx, c, &req, resp, err)
}

// DeletePost .
// @router /post/delete_post [POST]
func DeletePost(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.DeletePostReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	p := provider.Get()
	resp, err := p.PostService.DeletePost(ctx, &req)
	adaptor.Return(ctx, c, &req, resp, err)
}

// SetOfficial .
// @router /post/set_official [POST]
func SetOfficial(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.SetOfficialReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	p := provider.Get()
	resp, err := p.PostService.SetOfficial(ctx, &req)
	adaptor.Return(ctx, c, &req, resp, err)
}

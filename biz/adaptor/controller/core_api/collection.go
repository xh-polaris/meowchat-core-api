// Code generated by hertz generator.

package core_api

import (
	"context"

	"github.com/xh-polaris/meowchat-core-api/biz/adaptor"
	"github.com/xh-polaris/meowchat-core-api/biz/application/dto/base"
	"github.com/xh-polaris/meowchat-core-api/biz/application/dto/meowchat/core_api"
	"github.com/xh-polaris/meowchat-core-api/provider"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// GetCatPreviews .
// @router /collection/get_cat_previews [GET]
func GetCatPreviews(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.GetCatPreviewsReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	p := provider.Get()
	resp, err := p.CollectionService.GetCatPreviews(ctx, &req)
	resp.Status = new(base.Status)
	adaptor.Return(ctx, c, &req, resp, err)
}

// GetCatDetail .
// @router /collection/get_cat_detail [GET]
func GetCatDetail(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.GetCatDetailReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	p := provider.Get()
	resp, err := p.CollectionService.GetCatDetail(ctx, &req)
	resp.Status = new(base.Status)
	adaptor.Return(ctx, c, &req, resp, err)
}

// NewCat .
// @router /collection/new_cat [POST]
func NewCat(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.NewCatReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	p := provider.Get()
	resp, err := p.CollectionService.NewCat(ctx, &req)
	resp.Status = new(base.Status)
	adaptor.Return(ctx, c, &req, resp, err)
}

// DeleteCat .
// @router /collection/delete_cat [POST]
func DeleteCat(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.DeleteCatReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	p := provider.Get()
	resp, err := p.CollectionService.DeleteCat(ctx, &req)
	resp.Status = new(base.Status)
	adaptor.Return(ctx, c, &req, resp, err)
}

// SearchCat .
// @router /collection/search_cat [GET]
func SearchCat(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.SearchCatReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	p := provider.Get()
	resp, err := p.CollectionService.SearchCat(ctx, &req)
	resp.Status = new(base.Status)
	adaptor.Return(ctx, c, &req, resp, err)
}

// CreateImage .
// @router /collection/create_image [POST]
func CreateImage(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.CreateImageReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	p := provider.Get()
	resp, err := p.CollectionService.CreateImage(ctx, &req)
	resp.Status = new(base.Status)
	adaptor.Return(ctx, c, &req, resp, err)
}

// DeleteImage .
// @router /collection/delete_image [POST]
func DeleteImage(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.DeleteImageReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	p := provider.Get()
	resp, err := p.CollectionService.DeleteImage(ctx, &req)
	resp.Status = new(base.Status)
	adaptor.Return(ctx, c, &req, resp, err)
}

// GetImageByCat .
// @router /collection/get_image_by_cat [GET]
func GetImageByCat(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.GetImageByCatReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	p := provider.Get()
	resp, err := p.CollectionService.GetImageByCat(ctx, &req)
	resp.Status = new(base.Status)
	adaptor.Return(ctx, c, &req, resp, err)
}

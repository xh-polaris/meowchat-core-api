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

// GetPlanPreviews .
// @router /plan/get_plan_previews [GET]
func GetPlanPreviews(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.GetPlanPreviewsReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	p := provider.Get()
	resp, err := p.PlanService.GetPlanPreviews(ctx, &req)
	adaptor.PostProcess(ctx, c, &req, resp, err)
}

// GetPlanDetail .
// @router /plan/get_plan_detail [GET]
func GetPlanDetail(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.GetPlanDetailReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	p := provider.Get()
	resp, err := p.PlanService.GetPlanDetail(ctx, &req)
	adaptor.PostProcess(ctx, c, &req, resp, err)
}

// NewPlan .
// @router /plan/new_plan [POST]
func NewPlan(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.NewPlanReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	p := provider.Get()
	resp, err := p.PlanService.NewPlan(ctx, &req, adaptor.ExtractUserMeta(ctx, c))
	adaptor.PostProcess(ctx, c, &req, resp, err)
}

// DeletePlan .
// @router /plan/delete_plan [POST]
func DeletePlan(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.DeletePlanReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	p := provider.Get()
	resp, err := p.PlanService.DeletePlan(ctx, &req)
	adaptor.PostProcess(ctx, c, &req, resp, err)
}

// DonateFish .
// @router /plan/donate_fish [GET]
func DonateFish(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.DonateFishReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	p := provider.Get()
	resp, err := p.PlanService.DonateFish(ctx, &req, adaptor.ExtractUserMeta(ctx, c))
	adaptor.PostProcess(ctx, c, &req, resp, err)
}

// GetUserFish .
// @router /plan/get_user_fish [GET]
func GetUserFish(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.GetUserFishReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	p := provider.Get()
	resp, err := p.PlanService.GetUserFish(ctx, &req, adaptor.ExtractUserMeta(ctx, c))
	adaptor.PostProcess(ctx, c, &req, resp, err)
}

// ListFishByPlan .
// @router /plan/list_fish_by_plan [GET]
func ListFishByPlan(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.ListFishByPlanReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	p := provider.Get()
	resp, err := p.PlanService.ListFishByPlan(ctx, &req)
	adaptor.PostProcess(ctx, c, &req, resp, err)
}

// ListDonateByUser .
// @router /plan/list_donate_by_user [GET]
func ListDonateByUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.ListDonateByUserReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	p := provider.Get()
	resp, err := p.PlanService.ListDonateByUser(ctx, &req, adaptor.ExtractUserMeta(ctx, c))
	adaptor.PostProcess(ctx, c, &req, resp, err)
}

// CountDonateByUser .
// @router /plan/count_donate_by_user [GET]
func CountDonateByUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.CountDonateByUserReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	p := provider.Get()
	resp, err := p.PlanService.CountDonateByUser(ctx, &req, adaptor.ExtractUserMeta(ctx, c))
	adaptor.PostProcess(ctx, c, &req, resp, err)
}

// CountDonateByPlan .
// @router /plan/count_donate_by_plan [GET]
func CountDonateByPlan(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.CountDonateByPlanReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	p := provider.Get()
	resp, err := p.PlanService.CountDonateByPlan(ctx, &req)
	adaptor.PostProcess(ctx, c, &req, resp, err)
}

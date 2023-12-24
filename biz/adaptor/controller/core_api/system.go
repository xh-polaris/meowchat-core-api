// Code generated by hertz generator.

package core_api

import (
	"context"
	"time"

	"github.com/bytedance/gopkg/cloud/metainfo"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/samber/lo"
	"github.com/xh-polaris/gopkg/kitex/client"

	"github.com/xh-polaris/meowchat-core-api/biz/adaptor"
	"github.com/xh-polaris/meowchat-core-api/biz/application/dto/basic"
	"github.com/xh-polaris/meowchat-core-api/biz/application/dto/meowchat/core_api"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/util"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/util/log"
	"github.com/xh-polaris/meowchat-core-api/provider"
)

// GetAdmins .
// @router /notice/get_admins [GET]
func GetAdmins(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.GetAdminsReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	p := provider.Get()
	resp, err := p.SystemService.GetAdmins(ctx, &req)
	adaptor.PostProcess(ctx, c, &req, resp, err)
}

// NewAdmin .
// @router /notice/new_admin [POST]
func NewAdmin(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.NewAdminReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	p := provider.Get()
	resp, err := p.SystemService.NewAdmin(ctx, &req)
	adaptor.PostProcess(ctx, c, &req, resp, err)
}

// DeleteAdmin .
// @router /notice/delete_admin [POST]
func DeleteAdmin(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.DeleteAdminReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	p := provider.Get()
	resp, err := p.SystemService.DeleteAdmin(ctx, &req)
	adaptor.PostProcess(ctx, c, &req, resp, err)
}

// ListApply .
// @router /notice/list_apply [POST]
func ListApply(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.ListApplyReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	p := provider.Get()
	resp, err := p.SystemService.ListApply(ctx, &req)
	adaptor.PostProcess(ctx, c, &req, resp, err)
}

// HandleApply .
// @router /notice/handle_apply [POST]
func HandleApply(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.HandleApplyReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	p := provider.Get()
	resp, err := p.SystemService.HandleApply(ctx, &req)
	adaptor.PostProcess(ctx, c, &req, resp, err)
}

// GetNews .
// @router /notice/get_news [GET]
func GetNews(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.GetNewsReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	p := provider.Get()
	resp, err := p.SystemService.GetNews(ctx, &req)
	adaptor.PostProcess(ctx, c, &req, resp, err)
}

// NewNews .
// @router /notice/new_news [POST]
func NewNews(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.NewNewsReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	p := provider.Get()
	resp, err := p.SystemService.NewNews(ctx, &req)
	adaptor.PostProcess(ctx, c, &req, resp, err)
}

// DeleteNews .
// @router /notice/remove_news [POST]
func DeleteNews(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.DeleteNewsReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	p := provider.Get()
	resp, err := p.SystemService.DeleteNews(ctx, &req)
	adaptor.PostProcess(ctx, c, &req, resp, err)
}

// GetNotices .
// @router /notice/get_notices [GET]
func GetNotices(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.GetNoticesReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	p := provider.Get()
	resp, err := p.SystemService.GetNotices(ctx, &req)
	adaptor.PostProcess(ctx, c, &req, resp, err)
}

// NewNotice .
// @router /notice/new_notice [POST]
func NewNotice(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.NewNoticeReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	p := provider.Get()
	resp, err := p.SystemService.NewNotice(ctx, &req)
	adaptor.PostProcess(ctx, c, &req, resp, err)
}

// DeleteNotice .
// @router /notice/remove_notice [POST]
func DeleteNotice(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.DeleteNoticeReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	p := provider.Get()
	resp, err := p.SystemService.DeleteNotice(ctx, &req)
	adaptor.PostProcess(ctx, c, &req, resp, err)
}

// ListCommunity .
// @router /community/list_community [GET]
func ListCommunity(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.ListCommunityReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	p := provider.Get()
	resp, err := p.SystemService.ListCommunity(ctx, &req)
	adaptor.PostProcess(ctx, c, &req, resp, err)
}

// NewCommunity .
// @router /community/new_community [POST]
func NewCommunity(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.NewCommunityReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	p := provider.Get()
	resp, err := p.SystemService.NewCommunity(ctx, &req)
	adaptor.PostProcess(ctx, c, &req, resp, err)
}

// DeleteCommunity .
// @router /community/delete_community [POST]
func DeleteCommunity(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.DeleteCommunityReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	p := provider.Get()
	resp, err := p.SystemService.DeleteCommunity(ctx, &req)
	adaptor.PostProcess(ctx, c, &req, resp, err)
}

// GetUserRoles .
// @router /role/get_user_roles [GET]
func GetUserRoles(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.GetUserRolesReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	p := provider.Get()
	resp, err := p.SystemService.GetUserRoles(ctx, &req)
	adaptor.PostProcess(ctx, c, &req, resp, err)
}

// UpdateCommunityAdmin .
// @router /role/update_community_admin [POST]
func UpdateCommunityAdmin(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.UpdateCommunityAdminReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	p := provider.Get()
	resp, err := p.SystemService.UpdateCommunityAdmin(ctx, &req)
	adaptor.PostProcess(ctx, c, &req, resp, err)
}

// UpdateSuperAdmin .
// @router /role/update_super_admin [POST]
func UpdateSuperAdmin(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.UpdateSuperAdminReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	p := provider.Get()
	resp, err := p.SystemService.UpdateSuperAdmin(ctx, &req)
	adaptor.PostProcess(ctx, c, &req, resp, err)
}

// GetUserByRole .
// @router /role/get_user_by_role [GET]
func GetUserByRole(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.RetrieveUserPreviewReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	p := provider.Get()
	resp, err := p.SystemService.GetUserByRole(ctx, &req)
	adaptor.PostProcess(ctx, c, &req, resp, err)
}

// CreateApply .
// @router /role/create_apply [POST]
func CreateApply(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.CreateApplyReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	p := provider.Get()
	resp, err := p.SystemService.CreateApply(ctx, &req)
	adaptor.PostProcess(ctx, c, &req, resp, err)
}

// UpdateRole .
// @router /role/update_role [POST]
func UpdateRole(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.UpdateRoleReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	p := provider.Get()
	resp, err := p.SystemService.UpdateRole(ctx, &req)
	adaptor.PostProcess(ctx, c, &req, resp, err)
}

// Prefetch .
// @router /prefetch [GET]
func Prefetch(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.PrefetchReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(core_api.PrefetchResp)
	resp.Token = req.Token
	resp.Timestamp = time.Now().UnixMilli()
	p := provider.Get()
	params := new(struct {
		CommunityId string `json:"communityId"`
		UserId      string `json:"userId"`
		Env         string `json:"env"`
	})

	if req.Token != nil {
		err = sonic.UnmarshalString(*req.Token, params)
		if err != nil {
			log.CtxError(ctx, "[Prefetch] unmarshal token failed, err=%v", err)
		}
	}
	if params.CommunityId == "" {
		params.CommunityId = p.Config.DefaultCommunityId
	}
	ctx = metainfo.WithPersistentValue(ctx, client.EnvHeader, params.Env)
	util.ParallelRun(
		func() {
			if req.Code != nil {
				resp.SignInResp, err = p.AuthService.SignIn(ctx, &core_api.SignInReq{
					AuthType:   "wechat",
					AuthId:     req.Appid,
					VerifyCode: req.Code,
					AppId:      basic.APP_Meowchat,
				})
				if err != nil || resp.GetSignInResp().GetUserId() == "" {
					log.CtxError(ctx, "[Prefetch] sign in failed, err=%v", err)
					return
				}
				resp.GetUserInfoResp, err = p.UserService.GetUserInfo(ctx, &core_api.GetUserInfoReq{UserId: lo.ToPtr(resp.GetSignInResp().GetUserId())})
				if err != nil {
					log.CtxError(ctx, "[Prefetch] get user info failed, err=%v", err)
				}
			}
		},
		func() {
			resp.ListCommunityResp, err = p.SystemService.ListCommunity(ctx, &core_api.ListCommunityReq{})
			if err != nil {
				log.CtxError(ctx, "[Prefetch] list community failed, err=%v", err)
			}
		},
		func() {
			if params.UserId != "" {
				resp.GetUserInfoResp, err = p.UserService.GetUserInfo(ctx, &core_api.GetUserInfoReq{UserId: lo.ToPtr(params.UserId)})
				if err != nil {
					log.CtxError(ctx, "[Prefetch] get user info failed, err=%v", err)
				}
			}
		},
		func() {
			resp.FirstMomentPreviewsResp, err = p.MomentService.GetMomentPreviews(ctx, &core_api.GetMomentPreviewsReq{
				CommunityId: lo.ToPtr(params.CommunityId),
			})
			if err != nil {
				log.CtxError(ctx, "[Prefetch] get moment previews failed, err=%v", err)
			}
		},
		func() {
			resp.FirstPostPreviewsResp, err = p.PostService.GetPostPreviews(ctx, &core_api.GetPostPreviewsReq{})
			if err != nil {
				log.CtxError(ctx, "[Prefetch] get post previews failed, err=%v", err)
			}
		},
		func() {
			resp.FirstCatPreviewsResp, err = p.CollectionService.GetCatPreviews(ctx, &core_api.GetCatPreviewsReq{
				CommunityId: params.CommunityId,
			})
			if err != nil {
				log.CtxError(ctx, "[Prefetch] get cat previews failed, err=%v", err)
			}
		},
		func() {
			resp.GetNewsResp, err = p.SystemService.GetNews(ctx, &core_api.GetNewsReq{
				CommunityId: params.CommunityId,
			})
			if err != nil {
				log.CtxError(ctx, "[Prefetch] get news failed, err=%v", err)
			}
		})

	adaptor.PostProcess(ctx, c, &req, resp, err)
}

// ListNotification .
// @router /notification/list_notification [GET]
func ListNotification(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.ListNotificationReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	p := provider.Get()
	resp, err := p.SystemService.ListNotification(ctx, &req)
	adaptor.PostProcess(ctx, c, &req, resp, err)
}

// CleanNotification .
// @router /notification/clean_notification [GET]
func CleanNotification(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.CleanNotificationReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	p := provider.Get()
	resp, err := p.SystemService.CleanNotification(ctx, &req)
	adaptor.PostProcess(ctx, c, &req, resp, err)
}

// CountNotification .
// @router /notification/count_notification [GET]
func CountNotification(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.CountNotificationReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	p := provider.Get()
	resp, err := p.SystemService.CountNotification(ctx, &req)
	adaptor.PostProcess(ctx, c, &req, resp, err)
}

// GetMinVersion .
// @router /get_min_version [GET]
func GetMinVersion(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.GetMinVersionReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	p := provider.Get()
	resp, err := p.SystemService.GetMinVersion(ctx, &req)
	adaptor.PostProcess(ctx, c, &req, resp, err)
}

// ReadNotification .
// @router /notification/read_notification [GET]
func ReadNotification(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.ReadNotificationReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	p := provider.Get()
	resp, err := p.SystemService.ReadNotification(ctx, &req)
	adaptor.PostProcess(ctx, c, &req, resp, err)
}

// ReadRangeNotification .
// @router /notification/read_range_notification [GET]
func ReadRangeNotification(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.ReadRangeNotificationReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	p := provider.Get()
	resp, err := p.SystemService.ReadRangeNotification(ctx, &req)
	adaptor.PostProcess(ctx, c, &req, resp, err)
}

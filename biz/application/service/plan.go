package service

import (
	"context"
	"net/url"

	"github.com/google/wire"
	"github.com/jinzhu/copier"
	"github.com/samber/lo"
	"github.com/xh-polaris/gopkg/errors"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/platform/sts"

	"github.com/xh-polaris/meowchat-core-api/biz/adaptor"
	content2 "github.com/xh-polaris/meowchat-core-api/biz/application/dto/meowchat/content"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/consts"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/meowchat_content"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/platform_sts"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/util"

	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/basic"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowchat/content"
	genuser "github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowchat/user"

	basic2 "github.com/xh-polaris/meowchat-core-api/biz/application/dto/basic"
	"github.com/xh-polaris/meowchat-core-api/biz/application/dto/meowchat/core_api"
	user1 "github.com/xh-polaris/meowchat-core-api/biz/application/dto/meowchat/user"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/config"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/meowchat_user"
)

type IPlanService interface {
	DeletePlan(ctx context.Context, req *core_api.DeletePlanReq) (*core_api.DeletePlanResp, error)
	GetPlanDetail(ctx context.Context, req *core_api.GetPlanDetailReq) (*core_api.GetPlanDetailResp, error)
	GetPlanPreviews(ctx context.Context, req *core_api.GetPlanPreviewsReq) (*core_api.GetPlanPreviewsResp, error)
	NewPlan(ctx context.Context, req *core_api.NewPlanReq) (*core_api.NewPlanResp, error)
	DonateFish(ctx context.Context, req *core_api.DonateFishReq) (*core_api.DonateFishResp, error)
	GetUserFish(ctx context.Context, req *core_api.GetUserFishReq) (*core_api.GetUserFishResp, error)
	ListFishByPlan(ctx context.Context, req *core_api.ListFishByPlanReq) (*core_api.ListFishByPlanResp, error)
	ListDonateByUser(ctx context.Context, req *core_api.ListDonateByUserReq) (*core_api.ListDonateByUserResp, error)
	CountDonateByPlan(ctx context.Context, req *core_api.CountDonateByPlanReq) (*core_api.CountDonateByPlanResp, error)
	CountDonateByUser(ctx context.Context, req *core_api.CountDonateByUserReq) (*core_api.CountDonateByUserResp, error)
}

type PlanService struct {
	Config *config.Config
	Plan   meowchat_content.IMeowchatContent
	User   meowchat_user.IMeowchatUser
	Sts    platform_sts.IPlatformSts
}

var PlanServiceSet = wire.NewSet(
	wire.Struct(new(PlanService), "*"),
	wire.Bind(new(IPlanService), new(*PlanService)),
)

func (s *PlanService) DonateFish(ctx context.Context, req *core_api.DonateFishReq) (*core_api.DonateFishResp, error) {
	user := adaptor.ExtractUserMeta(ctx)
	if user.GetUserId() == "" {
		return nil, consts.ErrNotAuthentication
	}
	resp := new(core_api.DonateFishResp)
	_, err := s.Plan.DonateFish(ctx, &content.DonateFishReq{
		UserId: user.UserId,
		PlanId: req.PlanId,
		Fish:   req.Fish,
	})
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *PlanService) GetUserFish(ctx context.Context, req *core_api.GetUserFishReq) (*core_api.GetUserFishResp, error) {
	resp := new(core_api.GetUserFishResp)
	user := adaptor.ExtractUserMeta(ctx)
	var uid string
	if req.GetUserId() != "" {
		uid = req.GetUserId()
	} else if user.GetUserId() == "" {
		return nil, consts.ErrNotAuthentication
	} else {
		uid = user.GetUserId()
	}
	data, err := s.Plan.RetrieveUserFish(ctx, &content.RetrieveUserFishReq{
		UserId: uid,
	})
	if err != nil {
		return nil, err

	}
	resp.Fish = data.Fish
	return resp, nil
}

func (s *PlanService) ListFishByPlan(ctx context.Context, req *core_api.ListFishByPlanReq) (*core_api.ListFishByPlanResp, error) {
	resp := new(core_api.ListFishByPlanResp)

	if req.PaginationOption == nil {
		req.PaginationOption = &basic2.PaginationOptions{}
	}
	if req.PaginationOption.Limit == nil {
		req.PaginationOption.Limit = &PageSize
	}

	request := &content.ListFishByPlanReq{
		PlanId: req.PlanId,
		PaginationOptions: &basic.PaginationOptions{
			Offset:    new(int64),
			Limit:     req.PaginationOption.Limit,
			Backward:  req.PaginationOption.Backward,
			LastToken: req.PaginationOption.LastToken,
		},
	}
	*request.PaginationOptions.Offset = req.PaginationOption.GetLimit() * *req.PaginationOption.Page

	data, err := s.Plan.ListFishByPlan(ctx, request)
	if err != nil {
		return nil, err
	}

	users := make([]*user1.UserPreview, 0, len(data.UserIds))
	for _, userId := range data.UserIds {
		user, err := s.User.GetUserDetail(ctx, &genuser.GetUserDetailReq{UserId: userId})
		if err == nil {
			users = append(users, &user1.UserPreview{
				Id:        user.User.Id,
				Nickname:  user.User.Nickname,
				AvatarUrl: user.User.AvatarUrl,
			})
		}
	}
	resp.Total = data.GetTotal()
	resp.Users = users
	resp.FishMap = data.FishMap
	return resp, nil
}

func (s *PlanService) ListDonateByUser(ctx context.Context, req *core_api.ListDonateByUserReq) (*core_api.ListDonateByUserResp, error) {
	resp := new(core_api.ListDonateByUserResp)
	user := adaptor.ExtractUserMeta(ctx)
	if req.PaginationOption == nil {
		req.PaginationOption = &basic2.PaginationOptions{}
	}
	if req.PaginationOption.Limit == nil {
		req.PaginationOption.Limit = &PageSize
	}

	request := &content.ListDonateByUserReq{
		PaginationOptions: &basic.PaginationOptions{
			Offset:    new(int64),
			Limit:     req.PaginationOption.Limit,
			Backward:  req.PaginationOption.Backward,
			LastToken: req.PaginationOption.LastToken,
		},
	}
	if req.GetUserId() != "" {
		request.UserId = req.GetUserId()
	} else {
		request.UserId = user.UserId
	}
	if req.PaginationOption.LastToken == nil {
		request.PaginationOptions.Offset = lo.EmptyableToPtr(req.PaginationOption.GetLimit() * req.PaginationOption.GetPage())
	}

	data, err := s.Plan.ListDonateByUser(ctx, request)
	if err != nil {
		return nil, err
	}

	p := make([]*core_api.Donation, 0)
	for _, _donation := range data.Donations {
		temp := new(core_api.Plan)
		err = copier.Copy(temp, _donation.Plan)
		if err != nil {
			return nil, err
		}
		if _donation.Plan.CatId != "" {
			_cat, err := s.Plan.RetrieveCat(ctx, &content.RetrieveCatReq{CatId: _donation.Plan.CatId})
			if err == nil {
				c := new(content2.Cat)
				err = copier.Copy(c, _cat.Cat)
				if err == nil {
					temp.Cat = c
				}
			}
		}
		user, err := s.User.GetUserDetail(ctx, &genuser.GetUserDetailReq{UserId: _donation.Plan.InitiatorId})
		if err == nil {
			temp.User = &user1.UserPreview{
				Id:        user.User.Id,
				Nickname:  user.User.Nickname,
				AvatarUrl: user.User.AvatarUrl,
			}
		}
		p = append(p, &core_api.Donation{
			Plan:       temp,
			DonateTime: _donation.DonateTime,
			DonateNum:  _donation.DonateNum,
			DonateId:   _donation.DonateId,
		})
	}

	resp.Total = data.GetTotal()
	resp.Token = data.GetToken()
	resp.Donations = p

	return resp, nil
}

func (s *PlanService) DeletePlan(ctx context.Context, req *core_api.DeletePlanReq) (*core_api.DeletePlanResp, error) {
	user := adaptor.ExtractUserMeta(ctx)
	if user.GetUserId() == "" {
		return nil, consts.ErrNotAuthentication
	}
	resp := new(core_api.DeletePlanResp)
	_, err := s.Plan.DeletePlan(ctx, &content.DeletePlanReq{
		PlanId: req.PlanId,
	})
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *PlanService) GetPlanDetail(ctx context.Context, req *core_api.GetPlanDetailReq) (*core_api.GetPlanDetailResp, error) {
	resp := new(core_api.GetPlanDetailResp)
	data, err := s.Plan.RetrievePlan(ctx, &content.RetrievePlanReq{PlanId: req.PlanId})
	if err != nil {
		return nil, err
	}

	resp.Plan = new(core_api.Plan)
	err = copier.Copy(resp.Plan, data.Plan)
	if err != nil {
		return nil, err
	}
	util.ParallelRun(
		func() {
			if data.GetPlan().GetCatId() == "" {
				return
			}
			_cat, err := s.Plan.RetrieveCat(ctx, &content.RetrieveCatReq{CatId: data.Plan.CatId})
			if err != nil {
				return
			}
			c := new(content2.Cat)
			err = copier.Copy(c, _cat.Cat)
			if err == nil {
				resp.Plan.Cat = c
			}
		},
		func() {
			user, err := s.User.GetUserDetail(ctx, &genuser.GetUserDetailReq{UserId: data.Plan.InitiatorId})
			if err != nil {
				return
			}
			resp.Plan.User = &user1.UserPreview{
				Id:        user.User.Id,
				Nickname:  user.User.Nickname,
				AvatarUrl: user.User.AvatarUrl,
			}
		})
	return resp, nil
}

func (s *PlanService) GetPlanPreviews(ctx context.Context, req *core_api.GetPlanPreviewsReq) (*core_api.GetPlanPreviewsResp, error) {
	resp := new(core_api.GetPlanPreviewsResp)
	var data *content.ListPlanResp

	if req.PaginationOption == nil {
		req.PaginationOption = &basic2.PaginationOptions{}
	}
	if req.PaginationOption.Limit == nil {
		req.PaginationOption.Limit = &PageSize
	}
	request := &content.ListPlanReq{
		FilterOptions: &content.PlanFilterOptions{
			OnlyUserId:      req.OnlyUserId,
			OnlyCatId:       req.CatId,
			OnlyCommunityId: req.OnlyCommunityId,
			IncludeGlobal:   req.IncludeGlobal,
		},
		PaginationOptions: &basic.PaginationOptions{
			Offset:    new(int64),
			Limit:     req.PaginationOption.Limit,
			Backward:  req.PaginationOption.Backward,
			LastToken: req.PaginationOption.LastToken,
		},
	}
	if req.GetKeyword() != "" {
		request.SearchOptions = &content.SearchOptions{Type: &content.SearchOptions_AllFieldsKey{AllFieldsKey: req.GetKeyword()}}
	}
	if req.PaginationOption.LastToken == nil {
		request.PaginationOptions.Offset = lo.EmptyableToPtr(req.PaginationOption.GetLimit() * req.PaginationOption.GetPage())
	}
	data, err := s.Plan.ListPlan(ctx, request)
	if err != nil {
		return nil, err
	}

	resp.Total = data.Total
	resp.Plans = make([]*core_api.Plan, 0, len(data.Plans))
	resp.Token = data.Token
	err = copier.Copy(&resp.Plans, data.Plans)
	if err != nil {
		return nil, err
	}

	util.ParallelRun(lo.Map(data.Plans, func(plan *content.Plan, i int) func() {
		return func() {
			util.ParallelRun(
				func() {
					user, err := s.User.GetUserDetail(ctx, &genuser.GetUserDetailReq{UserId: plan.InitiatorId})
					if err != nil {
						return
					}
					resp.Plans[i].User = &user1.UserPreview{
						Id:        user.User.Id,
						Nickname:  user.User.Nickname,
						AvatarUrl: user.User.AvatarUrl,
					}
				},
				func() {
					if plan.GetCatId() == "" {
						return
					}
					_cat, err := s.Plan.RetrieveCat(ctx, &content.RetrieveCatReq{CatId: plan.CatId})
					if err != nil {
						return
					}
					c := new(content2.Cat)
					err = copier.Copy(c, _cat.Cat)
					if err == nil {
						resp.Plans[i].Cat = c
					}
				})
		}
	})...)
	return resp, nil
}

func (s *PlanService) NewPlan(ctx context.Context, req *core_api.NewPlanReq) (*core_api.NewPlanResp, error) {
	user := adaptor.ExtractUserMeta(ctx)
	if user.GetUserId() == "" {
		return nil, consts.ErrNotAuthentication
	}
	resp := new(core_api.NewPlanResp)
	m := new(content.Plan)

	if req.GetName()+req.GetDescription() != "" {
		r, err := s.Sts.TextCheck(ctx, &sts.TextCheckReq{
			Text:  req.GetName() + req.GetDescription(),
			User:  user,
			Scene: 2,
			Title: req.Name,
		})
		if err != nil {
			return nil, err
		}
		if r.Pass == false {
			return nil, errors.NewBizError(10001, "TextCheck don't pass")
		}
	}
	if len(req.ImageUrls) != 0 {
		urls := make([]string, len(req.ImageUrls))
		for i := 0; i < len(req.ImageUrls); i++ {
			var u *url.URL
			u, err := url.Parse(req.ImageUrls[i])
			if err != nil {
				return nil, err
			}
			u.Host = s.Config.CdnHost
			req.ImageUrls[i] = u.String()
			urls[i] = req.ImageUrls[i]
		}
		res, err := s.Sts.PhotoCheck(ctx, &sts.PhotoCheckReq{
			User: user,
			Url:  urls,
		})
		if err != nil {
			return nil, err
		}
		if res.Pass == false {
			return nil, errors.NewBizError(10002, "PhotoCheck don't pass")
		}

	}

	err := copier.Copy(m, req)
	if err != nil {
		return nil, err
	}

	m.InitiatorId = user.GetUserId()

	if req.GetId() == "" {
		var data *content.CreatePlanResp
		data, err = s.Plan.CreatePlan(ctx, &content.CreatePlanReq{Plan: m})
		resp.PlanId = data.PlanId
	} else {
		_, err = s.Plan.UpdatePlan(ctx, &content.UpdatePlanReq{Plan: m})
		resp.PlanId = *req.Id
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *PlanService) CountDonateByPlan(ctx context.Context, req *core_api.CountDonateByPlanReq) (*core_api.CountDonateByPlanResp, error) {
	total, err := s.Plan.CountDonateByPlan(ctx, &content.CountDonateByPlanReq{PlanId: req.PlanId})
	if err != nil {
		return nil, err
	}
	return &core_api.CountDonateByPlanResp{Total: total.Total}, nil
}

func (s *PlanService) CountDonateByUser(ctx context.Context, req *core_api.CountDonateByUserReq) (*core_api.CountDonateByUserResp, error) {
	user := adaptor.ExtractUserMeta(ctx)
	userId := ""
	if req.GetUserId() != "" {
		userId = req.GetUserId()
	} else {
		userId = user.UserId
	}

	total, err := s.Plan.CountDonateByUser(ctx, &content.CountDonateByUserReq{UserId: userId})
	if err != nil {
		return nil, err
	}
	return &core_api.CountDonateByUserResp{Total: total.Total}, nil
}

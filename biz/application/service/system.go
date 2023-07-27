package service

import (
	"context"
	"github.com/google/wire"
	"github.com/jinzhu/copier"
	"github.com/xh-polaris/meowchat-core-api/biz/application/dto/meowchat/core_api"
	system2 "github.com/xh-polaris/meowchat-core-api/biz/application/dto/meowchat/system"
	user1 "github.com/xh-polaris/meowchat-core-api/biz/application/dto/meowchat/user"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/config"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/meowchat_system"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/meowchat_user"
	"github.com/xh-polaris/meowchat-system-rpc/common/constant"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowchat/system"
	genuser "github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowchat/user"
	"sync"
)

type ISystemService interface {
	CreateApply(ctx context.Context, req *core_api.CreateApplyReq) (*core_api.CreateApplyResp, error)
	DeleteAdmin(ctx context.Context, req *core_api.DeleteAdminReq) (*core_api.DeleteAdminResp, error)
	DeleteCommunity(ctx context.Context, req *core_api.DeleteCommunityReq) (*core_api.DeleteCommunityResp, error)
	DeleteNews(ctx context.Context, req *core_api.DeleteNewsReq) (*core_api.DeleteNewsResp, error)
	DeleteNotice(ctx context.Context, req *core_api.DeleteNoticeReq) (*core_api.DeleteNoticeResp, error)
	GetAdmins(ctx context.Context, req *core_api.GetAdminsReq) (*core_api.GetAdminsResp, error)
	GetNews(ctx context.Context, req *core_api.GetNewsReq) (*core_api.GetNewsResp, error)
	GetNotices(ctx context.Context, req *core_api.GetNoticesReq) (*core_api.GetNoticesResp, error)
	GetUserByRole(ctx context.Context, req *core_api.RetrieveUserPreviewReq) (*core_api.RetrieveUserPreviewResp, error)
	GetUserRoles(ctx context.Context, req *core_api.GetUserRolesReq) (*core_api.GetUserRolesResp, error)
	HandleApply(ctx context.Context, req *core_api.HandleApplyReq) (*core_api.HandleApplyResp, error)
	ListApply(ctx context.Context, req *core_api.ListApplyReq) (*core_api.ListApplyResp, error)
	ListCommunity(ctx context.Context, req *core_api.ListCommunityReq) (*core_api.ListCommunityResp, error)
	NewAdmin(ctx context.Context, req *core_api.NewAdminReq) (*core_api.NewAdminResp, error)
	NewCommunity(ctx context.Context, req *core_api.NewCommunityReq) (*core_api.NewCommunityResp, error)
	NewNews(ctx context.Context, req *core_api.NewNewsReq) (*core_api.NewNewsResp, error)
	NewNotice(ctx context.Context, req *core_api.NewNoticeReq) (*core_api.NewNoticeResp, error)
	UpdateCommunityAdmin(ctx context.Context, req *core_api.UpdateCommunityAdminReq) (*core_api.UpdateCommunityAdminResp, error)
	UpdateSuperAdmin(ctx context.Context, req *core_api.UpdateSuperAdminReq) (*core_api.UpdateSuperAdminResp, error)
}

type SystemService struct {
	Config *config.Config
	System meowchat_system.IMeowchatSystem
	User   meowchat_user.IMeowchatUser
}

var SystemServiceSet = wire.NewSet(
	wire.Struct(new(SystemService), "*"),
	wire.Bind(new(ISystemService), new(*SystemService)),
)

func (s *SystemService) CreateApply(ctx context.Context, req *core_api.CreateApplyReq) (*core_api.CreateApplyResp, error) {
	resp := new(core_api.CreateApplyResp)
	ApplicantId := ctx.Value("userId").(string)
	_, err := s.System.CreateApply(ctx, &system.CreateApplyReq{
		ApplicantId: ApplicantId,
		CommunityId: req.CommunityId,
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SystemService) DeleteAdmin(ctx context.Context, req *core_api.DeleteAdminReq) (*core_api.DeleteAdminResp, error) {
	resp := new(core_api.DeleteAdminResp)
	_, err := s.System.DeleteAdmin(ctx, &system.DeleteAdminReq{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SystemService) DeleteCommunity(ctx context.Context, req *core_api.DeleteCommunityReq) (*core_api.DeleteCommunityResp, error) {
	resp := new(core_api.DeleteCommunityResp)

	_, err := s.System.DeleteCommunity(ctx, &system.DeleteCommunityReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *SystemService) DeleteNews(ctx context.Context, req *core_api.DeleteNewsReq) (*core_api.DeleteNewsResp, error) {
	resp := new(core_api.DeleteNewsResp)

	_, err := s.System.DeleteNews(ctx, &system.DeleteNewsReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *SystemService) DeleteNotice(ctx context.Context, req *core_api.DeleteNoticeReq) (*core_api.DeleteNoticeResp, error) {
	resp := new(core_api.DeleteNoticeResp)

	_, err := s.System.DeleteNotice(ctx, &system.DeleteNoticeReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *SystemService) GetAdmins(ctx context.Context, req *core_api.GetAdminsReq) (*core_api.GetAdminsResp, error) {
	resp := new(core_api.GetAdminsResp)
	resp.Admins = make([]*core_api.Admin, 0)

	data, err := s.System.ListAdmin(ctx, &system.ListAdminReq{CommunityId: req.CommunityId})
	if err != nil {
		return nil, err
	}

	err = copier.Copy(&resp.Admins, &data.Admins)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *SystemService) GetNews(ctx context.Context, req *core_api.GetNewsReq) (*core_api.GetNewsResp, error) {
	resp := new(core_api.GetNewsResp)
	resp.News = make([]*core_api.News, 0)

	data, err := s.System.ListNews(ctx, &system.ListNewsReq{CommunityId: req.CommunityId})
	if err != nil {
		return nil, err
	}

	err = copier.Copy(&resp.News, &data.News)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *SystemService) GetNotices(ctx context.Context, req *core_api.GetNoticesReq) (*core_api.GetNoticesResp, error) {
	resp := new(core_api.GetNoticesResp)
	resp.Notices = make([]*core_api.Notice, 0)

	data, err := s.System.ListNotice(ctx, &system.ListNoticeReq{CommunityId: req.CommunityId})
	if err != nil {
		return nil, err
	}

	err = copier.Copy(&resp.Notices, &data.Notices)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *SystemService) GetUserByRole(ctx context.Context, req *core_api.RetrieveUserPreviewReq) (*core_api.RetrieveUserPreviewResp, error) {
	resp := new(core_api.RetrieveUserPreviewResp)

	Userid, err := s.System.ListUserIdByRole(ctx, &system.ListUserIdByRoleReq{RoleType: req.RoleType, CommunityId: req.CommunityId})
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup
	resp.Users = make([]*user1.UserPreview, len(Userid.UserId))

	wg.Add(len(Userid.UserId))
	var errorChannel = make(chan error, len(Userid.UserId))

	for i, userid := range Userid.UserId {
		go s.GetOneUser(userid, &wg, i, resp.Users, errorChannel, ctx)
	}

	wg.Wait()
	if len(errorChannel) != 0 {
		return nil, <-errorChannel
	}

	return resp, nil
}

func (s *SystemService) GetOneUser(userid string, wg *sync.WaitGroup, i int, Users []*user1.UserPreview, chan1 chan error, ctx context.Context) (err error) {
	defer wg.Done()
	request := &genuser.GetUserReq{
		UserId: userid,
	}
	data, err := s.User.GetUser(ctx, request)
	if err != nil {
		chan1 <- err
		return err
	}
	Users[i] = &user1.UserPreview{
		Id:        data.User.Id,
		Nickname:  data.User.Nickname,
		AvatarUrl: data.User.AvatarUrl,
	}
	return nil
}

func (s *SystemService) GetUserRoles(ctx context.Context, req *core_api.GetUserRolesReq) (*core_api.GetUserRolesResp, error) {
	resp := new(core_api.GetUserRolesResp)
	data, err := s.System.RetrieveUserRole(ctx, &system.RetrieveUserRoleReq{UserId: ctx.Value("userId").(string)})
	if err != nil {
		return nil, err
	}
	for _, role := range data.Roles {
		resp.Roles = append(resp.Roles, &system2.Role{
			RoleType:    role.RoleType,
			CommunityId: role.CommunityId,
		})
	}
	return resp, nil
}

func (s *SystemService) HandleApply(ctx context.Context, req *core_api.HandleApplyReq) (*core_api.HandleApplyResp, error) {
	resp := new(core_api.HandleApplyResp)
	_, err := s.System.HandleApply(ctx, &system.HandleApplyReq{
		ApplyId:    req.ApplyId,
		IsRejected: req.IsRejected,
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SystemService) ListApply(ctx context.Context, req *core_api.ListApplyReq) (*core_api.ListApplyResp, error) {
	resp := new(core_api.ListApplyResp)
	apply, err := s.System.ListApply(ctx, &system.ListApplyReq{CommunityId: req.CommunityId})
	if err != nil {
		return nil, err
	}
	applyInfo := make([]*core_api.ApplyInfo, 0, len(apply.Apply))
	for _, x := range apply.Apply {
		user, _ := s.User.GetUser(ctx, &genuser.GetUserReq{UserId: x.ApplicantId})
		userPreview := user1.UserPreview{
			Id:        user.User.Id,
			Nickname:  user.User.Nickname,
			AvatarUrl: user.User.AvatarUrl,
		}
		applyInfo = append(applyInfo, &core_api.ApplyInfo{
			User:    &userPreview,
			ApplyId: x.ApplyId,
		})
	}
	resp.ApplyInfo = applyInfo
	return resp, nil
}

func (s *SystemService) ListCommunity(ctx context.Context, req *core_api.ListCommunityReq) (*core_api.ListCommunityResp, error) {
	resp := new(core_api.ListCommunityResp)

	data, err := s.System.ListCommunity(ctx, &system.ListCommunityReq{ParentId: req.ParentId, Size: -1})
	if err != nil {
		return nil, err
	}
	resp.Communities = make([]*core_api.Community, len(data.Communities))

	err = copier.Copy(&resp.Communities, &data.Communities)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *SystemService) NewAdmin(ctx context.Context, req *core_api.NewAdminReq) (*core_api.NewAdminResp, error) {
	resp := new(core_api.NewAdminResp)
	if req.Id == "" {
		data, err := s.System.CreateAdmin(ctx, &system.CreateAdminReq{
			CommunityId: req.CommunityId,
			Name:        req.Name,
			Title:       req.Title,
			Phone:       req.Phone,
			Wechat:      req.Wechat,
			AvatarUrl:   req.AvatarUrl,
		})
		if err != nil {
			return nil, err
		}
		resp.Id = data.Id
	} else {
		resp.Id = req.Id
		_, err := s.System.UpdateAdmin(ctx, &system.UpdateAdminReq{
			Id:          req.Id,
			CommunityId: req.CommunityId,
			Name:        req.Name,
			Title:       req.Title,
			Phone:       req.Phone,
			Wechat:      req.Wechat,
			AvatarUrl:   req.AvatarUrl,
		})
		if err != nil {
			return nil, err
		}
	}
	return resp, nil
}

func (s *SystemService) NewCommunity(ctx context.Context, req *core_api.NewCommunityReq) (*core_api.NewCommunityResp, error) {
	resp := new(core_api.NewCommunityResp)

	if req.Id == "" {
		data, err := s.System.CreateCommunity(ctx, &system.CreateCommunityReq{
			Name:     req.Name,
			ParentId: req.ParentId,
		})
		if err != nil {
			return nil, err
		}
		resp.Id = data.Id
	} else {
		resp.Id = req.Id
		_, err := s.System.UpdateCommunity(ctx, &system.UpdateCommunityReq{
			Id:       req.Id,
			Name:     req.Name,
			ParentId: req.ParentId,
		})
		if err != nil {
			return nil, err
		}
	}
	return resp, nil
}

func (s *SystemService) NewNews(ctx context.Context, req *core_api.NewNewsReq) (*core_api.NewNewsResp, error) {

	resp := new(core_api.NewNewsResp)

	if req.Id == "" {
		data, err := s.System.CreateNews(ctx, &system.CreateNewsReq{
			CommunityId: req.CommunityId,
			ImageUrl:    req.ImageUrl,
			LinkUrl:     req.LinkUrl,
			Type:        req.Type,
			IsPublic:    req.IsPublic,
		})
		if err != nil {
			return nil, err
		}
		resp.Id = data.Id
	} else {
		resp.Id = req.Id
		_, err := s.System.UpdateNews(ctx, &system.UpdateNewsReq{
			Id:       req.Id,
			ImageUrl: req.ImageUrl,
			LinkUrl:  req.LinkUrl,
			Type:     req.Type,
			IsPublic: req.IsPublic,
		})
		if err != nil {
			return nil, err
		}
	}
	return resp, nil
}

func (s *SystemService) NewNotice(ctx context.Context, req *core_api.NewNoticeReq) (*core_api.NewNoticeResp, error) {
	resp := new(core_api.NewNoticeResp)

	if req.Id == "" {
		data, err := s.System.CreateNotice(ctx, &system.CreateNoticeReq{
			Text:        req.Text,
			CommunityId: req.CommunityId,
		})
		if err != nil {
			return nil, err
		}
		resp.Id = data.Id
	} else {
		resp.Id = req.Id
		_, err := s.System.UpdateNotice(ctx, &system.UpdateNoticeReq{
			Id:   req.Id,
			Text: req.Text,
		})
		if err != nil {
			return nil, err
		}
	}

	return resp, nil
}

func (s *SystemService) UpdateCommunityAdmin(ctx context.Context, req *core_api.UpdateCommunityAdminReq) (*core_api.UpdateCommunityAdminResp, error) {
	resp := new(core_api.UpdateCommunityAdminResp)
	data, err := s.System.RetrieveUserRole(ctx, &system.RetrieveUserRoleReq{UserId: req.UserId})
	if err != nil {
		return nil, err
	}
	if !req.IsRemove {
		for _, role := range data.Roles {
			if role.RoleType == constant.RoleCommunityAdmin && *role.CommunityId == req.CommunityId {
				return nil, err
				// TODO 应当返回错误
			}
		}
		_, err = s.System.UpdateUserRole(ctx, &system.UpdateUserRoleReq{
			UserId: req.UserId,
			Roles: append(data.Roles, &system.Role{
				RoleType:    constant.RoleCommunityAdmin,
				CommunityId: &req.CommunityId,
			}),
		})
		if err != nil {
			return nil, err
		}
	} else {
		roles := make([]*system.Role, 0, len(data.Roles))
		for _, role := range data.Roles {
			if role.RoleType != constant.RoleCommunityAdmin || *role.CommunityId != req.CommunityId {
				roles = append(roles, role)
			}
		}
		if len(roles) == len(data.Roles) {
			return nil, err
			// TODO 应当返回错误
		}
		_, err = s.System.UpdateUserRole(ctx, &system.UpdateUserRoleReq{
			UserId: req.UserId,
			Roles:  roles,
		})
		if err != nil {
			return nil, err
		}
	}
	return resp, nil
}

func (s *SystemService) UpdateSuperAdmin(ctx context.Context, req *core_api.UpdateSuperAdminReq) (*core_api.UpdateSuperAdminResp, error) {
	resp := new(core_api.UpdateSuperAdminResp)
	data, err := s.System.RetrieveUserRole(ctx, &system.RetrieveUserRoleReq{UserId: req.UserId})
	if err != nil {
		return nil, err
	}
	if !req.IsRemove {
		for _, role := range data.Roles {
			if role.RoleType == constant.RoleSuperAdmin {
				return nil, err
				// TODO 应当返回错误
			}
		}
		_, err = s.System.UpdateUserRole(ctx, &system.UpdateUserRoleReq{
			UserId: req.UserId,
			Roles: append(data.Roles, &system.Role{
				RoleType: constant.RoleSuperAdmin,
			}),
		})
		if err != nil {
			return nil, err
		}
	} else {
		roles := make([]*system.Role, 0, len(data.Roles))
		for _, role := range data.Roles {
			if role.RoleType != constant.RoleSuperAdmin {
				roles = append(roles, role)
			}
		}
		if len(roles) == len(data.Roles) {
			return nil, err
			// TODO 应当返回错误
		}
		_, err = s.System.UpdateUserRole(ctx, &system.UpdateUserRoleReq{
			UserId: req.UserId,
			Roles:  roles,
		})
		if err != nil {
			return nil, err
		}
	}
	return resp, nil
}

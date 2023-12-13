package service

import (
	"context"
	"sync"

	"github.com/google/wire"
	"github.com/jinzhu/copier"
	"github.com/samber/lo"
	genbasic "github.com/xh-polaris/service-idl-gen-go/kitex_gen/basic"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowchat/system"
	genuser "github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowchat/user"

	"github.com/xh-polaris/meowchat-core-api/biz/application/dto/basic"
	"github.com/xh-polaris/meowchat-core-api/biz/application/dto/meowchat/core_api"
	system2 "github.com/xh-polaris/meowchat-core-api/biz/application/dto/meowchat/system"
	user1 "github.com/xh-polaris/meowchat-core-api/biz/application/dto/meowchat/user"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/config"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/consts"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/meowchat_system"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/meowchat_user"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/util"
)

type ISystemService interface {
	CreateApply(ctx context.Context, req *core_api.CreateApplyReq, user *genbasic.UserMeta) (*core_api.CreateApplyResp, error)
	DeleteAdmin(ctx context.Context, req *core_api.DeleteAdminReq, user *genbasic.UserMeta) (*core_api.DeleteAdminResp, error)
	DeleteCommunity(ctx context.Context, req *core_api.DeleteCommunityReq, user *genbasic.UserMeta) (*core_api.DeleteCommunityResp, error)
	DeleteNews(ctx context.Context, req *core_api.DeleteNewsReq, user *genbasic.UserMeta) (*core_api.DeleteNewsResp, error)
	DeleteNotice(ctx context.Context, req *core_api.DeleteNoticeReq, user *genbasic.UserMeta) (*core_api.DeleteNoticeResp, error)
	GetAdmins(ctx context.Context, req *core_api.GetAdminsReq) (*core_api.GetAdminsResp, error)
	GetNews(ctx context.Context, req *core_api.GetNewsReq) (*core_api.GetNewsResp, error)
	GetNotices(ctx context.Context, req *core_api.GetNoticesReq) (*core_api.GetNoticesResp, error)
	GetUserByRole(ctx context.Context, req *core_api.RetrieveUserPreviewReq) (*core_api.RetrieveUserPreviewResp, error)
	GetUserRoles(ctx context.Context, req *core_api.GetUserRolesReq, user *genbasic.UserMeta) (*core_api.GetUserRolesResp, error)
	HandleApply(ctx context.Context, req *core_api.HandleApplyReq, user *genbasic.UserMeta) (*core_api.HandleApplyResp, error)
	ListApply(ctx context.Context, req *core_api.ListApplyReq) (*core_api.ListApplyResp, error)
	ListCommunity(ctx context.Context, req *core_api.ListCommunityReq) (*core_api.ListCommunityResp, error)
	NewAdmin(ctx context.Context, req *core_api.NewAdminReq, user *genbasic.UserMeta) (*core_api.NewAdminResp, error)
	NewCommunity(ctx context.Context, req *core_api.NewCommunityReq, user *genbasic.UserMeta) (*core_api.NewCommunityResp, error)
	NewNews(ctx context.Context, req *core_api.NewNewsReq, user *genbasic.UserMeta) (*core_api.NewNewsResp, error)
	NewNotice(ctx context.Context, req *core_api.NewNoticeReq, user *genbasic.UserMeta) (*core_api.NewNoticeResp, error)
	UpdateCommunityAdmin(ctx context.Context, req *core_api.UpdateCommunityAdminReq, user *genbasic.UserMeta) (*core_api.UpdateCommunityAdminResp, error)
	UpdateSuperAdmin(ctx context.Context, req *core_api.UpdateSuperAdminReq, user *genbasic.UserMeta) (*core_api.UpdateSuperAdminResp, error)
	UpdateRole(ctx context.Context, req *core_api.UpdateRoleReq, user *genbasic.UserMeta) (*core_api.UpdateRoleResp, error)
	GetMinVersion(ctx context.Context, req *core_api.GetMinVersionReq) (*core_api.GetMinVersionResp, error)
	ListNotification(ctx context.Context, req *core_api.ListNotificationReq, user *genbasic.UserMeta) (*core_api.ListNotificationResp, error)
	ReadNotification(ctx context.Context, req *core_api.ReadNotificationReq, user *genbasic.UserMeta) (*core_api.ReadNotificationResp, error)
	CountNotification(ctx context.Context, req *core_api.CountNotificationReq, user *genbasic.UserMeta) (*core_api.CountNotificationResp, error)
	CleanNotification(ctx context.Context, req *core_api.CleanNotificationReq, user *genbasic.UserMeta) (*core_api.CleanNotificationResp, error)
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

func (s *SystemService) ListNotification(ctx context.Context, req *core_api.ListNotificationReq, user *genbasic.UserMeta) (*core_api.ListNotificationResp, error) {
	resp := new(core_api.ListNotificationResp)

	if req.PaginationOption == nil {
		req.PaginationOption = &basic.PaginationOptions{}
	}
	if req.PaginationOption.Limit == nil {
		req.PaginationOption.Limit = lo.ToPtr[int64](10)
	}
	request := &system.ListNotificationReq{
		UserId: user.UserId,
		PaginationOptions: &genbasic.PaginationOptions{
			Limit:     req.PaginationOption.Limit,
			Backward:  req.PaginationOption.Backward,
			LastToken: req.PaginationOption.LastToken,
		},
	}
	if req.PaginationOption.LastToken == nil {
		request.PaginationOptions.Offset = lo.EmptyableToPtr(req.PaginationOption.GetLimit() * req.PaginationOption.GetPage())
	}

	data, err := s.System.ListNotification(ctx, request)
	if err != nil {
		return nil, err
	}

	resp.Total = data.Total
	resp.NotRead = data.NotRead
	resp.Notifications = make([]*core_api.Notification, 0, len(data.Notifications))
	err = copier.Copy(&resp.Notifications, data.Notifications)
	if err != nil {
		return nil, err
	}

	util.ParallelRun(lo.Map(data.Notifications, func(notification *system.Notification, i int) func() {
		return func() {
			user, err := s.User.GetUserDetail(ctx, &genuser.GetUserDetailReq{UserId: notification.SourceUserId})
			if err == nil {
				resp.Notifications[i].User = &user1.UserPreview{
					Id:        user.User.Id,
					Nickname:  user.User.Nickname,
					AvatarUrl: user.User.AvatarUrl,
				}
			}
		}
	}))

	return resp, nil
}

func (s *SystemService) ReadNotification(ctx context.Context, req *core_api.ReadNotificationReq, user *genbasic.UserMeta) (*core_api.ReadNotificationResp, error) {
	if user.GetUserId() == "" {
		return nil, consts.ErrNotAuthentication
	}
	_, err := s.System.ReadNotification(ctx, &system.ReadNotificationReq{
		NotificationId: req.NotificationId,
	})
	if err != nil {
		return nil, err
	}
	return &core_api.ReadNotificationResp{}, nil
}

func (s *SystemService) CountNotification(ctx context.Context, req *core_api.CountNotificationReq, user *genbasic.UserMeta) (*core_api.CountNotificationResp, error) {
	data, err := s.System.CountNotification(ctx, &system.CountNotificationReq{
		UserId: user.UserId,
	})
	if err != nil {
		return nil, err
	}
	return &core_api.CountNotificationResp{
		NotRead: data.NotificationCount,
	}, nil
}

func (s *SystemService) CleanNotification(ctx context.Context, req *core_api.CleanNotificationReq, user *genbasic.UserMeta) (*core_api.CleanNotificationResp, error) {
	if user.GetUserId() == "" {
		return nil, consts.ErrNotAuthentication
	}
	_, err := s.System.CleanNotification(ctx, &system.CleanNotificationReq{
		UserId: user.UserId,
	})
	if err != nil {
		return nil, err
	}
	return &core_api.CleanNotificationResp{}, nil
}

func (s *SystemService) CreateApply(ctx context.Context, req *core_api.CreateApplyReq, user *genbasic.UserMeta) (*core_api.CreateApplyResp, error) {
	if user.GetUserId() == "" {
		return nil, consts.ErrNotAuthentication
	}
	resp := new(core_api.CreateApplyResp)
	ApplicantId := user.UserId
	_, err := s.System.CreateApply(ctx, &system.CreateApplyReq{
		ApplicantId: ApplicantId,
		CommunityId: req.CommunityId,
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SystemService) DeleteAdmin(ctx context.Context, req *core_api.DeleteAdminReq, user *genbasic.UserMeta) (*core_api.DeleteAdminResp, error) {
	if user.GetUserId() == "" {
		return nil, consts.ErrNotAuthentication
	}
	resp := new(core_api.DeleteAdminResp)
	_, err := s.System.DeleteAdmin(ctx, &system.DeleteAdminReq{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SystemService) DeleteCommunity(ctx context.Context, req *core_api.DeleteCommunityReq, user *genbasic.UserMeta) (*core_api.DeleteCommunityResp, error) {
	if user.GetUserId() == "" {
		return nil, consts.ErrNotAuthentication
	}
	resp := new(core_api.DeleteCommunityResp)

	_, err := s.System.DeleteCommunity(ctx, &system.DeleteCommunityReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *SystemService) DeleteNews(ctx context.Context, req *core_api.DeleteNewsReq, user *genbasic.UserMeta) (*core_api.DeleteNewsResp, error) {
	if user.GetUserId() == "" {
		return nil, consts.ErrNotAuthentication
	}
	resp := new(core_api.DeleteNewsResp)

	_, err := s.System.DeleteNews(ctx, &system.DeleteNewsReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *SystemService) DeleteNotice(ctx context.Context, req *core_api.DeleteNoticeReq, user *genbasic.UserMeta) (*core_api.DeleteNoticeResp, error) {
	if user.GetUserId() == "" {
		return nil, consts.ErrNotAuthentication
	}
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

	resp.Admins = make([]*core_api.Admin, 0, PageSize)
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

	resp.News = make([]*core_api.News, 0, PageSize)
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

	resp.Notices = make([]*core_api.Notice, 0, PageSize)
	err = copier.Copy(&resp.Notices, &data.Notices)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *SystemService) GetUserByRole(ctx context.Context, req *core_api.RetrieveUserPreviewReq) (*core_api.RetrieveUserPreviewResp, error) {
	resp := new(core_api.RetrieveUserPreviewResp)

	Userid, err := s.System.ListUserIdByRole(ctx, &system.ListUserIdByRoleReq{Role: &system.Role{
		RoleType:    system.RoleType(req.RoleType),
		CommunityId: req.CommunityId,
	}})
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
	request := &genuser.GetUserDetailReq{
		UserId: userid,
	}
	data, err := s.User.GetUserDetail(ctx, request)
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

func (s *SystemService) GetUserRoles(ctx context.Context, req *core_api.GetUserRolesReq, user *genbasic.UserMeta) (*core_api.GetUserRolesResp, error) {
	resp := new(core_api.GetUserRolesResp)
	data, err := s.System.RetrieveUserRole(ctx, &system.RetrieveUserRoleReq{UserId: user.UserId})
	if err != nil {
		return nil, err
	}
	for _, role := range data.Roles {
		resp.Roles = append(resp.Roles, &system2.Role{
			RoleType:    system2.RoleType(role.RoleType),
			CommunityId: role.CommunityId,
		})
	}
	return resp, nil
}

func (s *SystemService) HandleApply(ctx context.Context, req *core_api.HandleApplyReq, user *genbasic.UserMeta) (*core_api.HandleApplyResp, error) {
	if user.GetUserId() == "" {
		return nil, consts.ErrNotAuthentication
	}
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
		user, _ := s.User.GetUserDetail(ctx, &genuser.GetUserDetailReq{UserId: x.ApplicantId})
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

	data, err := s.System.ListCommunity(ctx, &system.ListCommunityReq{ParentId: req.GetParentId(), PageSize: -1})
	if err != nil {
		return nil, err
	}
	resp.Communities = make([]*core_api.Community, len(data.Communities))

	resp.Communities = make([]*core_api.Community, 0)
	err = copier.Copy(&resp.Communities, &data.Communities)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *SystemService) NewAdmin(ctx context.Context, req *core_api.NewAdminReq, user *genbasic.UserMeta) (*core_api.NewAdminResp, error) {
	if user.GetUserId() == "" {
		return nil, consts.ErrNotAuthentication
	}
	resp := new(core_api.NewAdminResp)
	if req.GetId() == "" {
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
		resp.Id = req.GetId()
		_, err := s.System.UpdateAdmin(ctx, &system.UpdateAdminReq{
			Id:          req.GetId(),
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

func (s *SystemService) NewCommunity(ctx context.Context, req *core_api.NewCommunityReq, user *genbasic.UserMeta) (*core_api.NewCommunityResp, error) {
	if user.GetUserId() == "" {
		return nil, consts.ErrNotAuthentication
	}
	resp := new(core_api.NewCommunityResp)

	if req.GetId() == "" {
		data, err := s.System.CreateCommunity(ctx, &system.CreateCommunityReq{
			Name:     req.Name,
			ParentId: req.GetParentId(),
		})
		if err != nil {
			return nil, err
		}
		resp.Id = data.Id
	} else {
		resp.Id = req.GetId()
		_, err := s.System.UpdateCommunity(ctx, &system.UpdateCommunityReq{
			Id:       req.GetId(),
			Name:     req.Name,
			ParentId: req.GetParentId(),
		})
		if err != nil {
			return nil, err
		}
	}
	return resp, nil
}

func (s *SystemService) NewNews(ctx context.Context, req *core_api.NewNewsReq, user *genbasic.UserMeta) (*core_api.NewNewsResp, error) {
	if user.GetUserId() == "" {
		return nil, consts.ErrNotAuthentication
	}
	resp := new(core_api.NewNewsResp)

	if req.GetId() == "" {
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
		resp.Id = req.GetId()
		_, err := s.System.UpdateNews(ctx, &system.UpdateNewsReq{
			Id:       req.GetId(),
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

func (s *SystemService) NewNotice(ctx context.Context, req *core_api.NewNoticeReq, user *genbasic.UserMeta) (*core_api.NewNoticeResp, error) {
	if user.GetUserId() == "" {
		return nil, consts.ErrNotAuthentication
	}
	resp := new(core_api.NewNoticeResp)

	if req.GetId() == "" {
		data, err := s.System.CreateNotice(ctx, &system.CreateNoticeReq{
			Text:        req.Text,
			CommunityId: req.CommunityId,
		})
		if err != nil {
			return nil, err
		}
		resp.Id = data.Id
	} else {
		resp.Id = req.GetId()
		_, err := s.System.UpdateNotice(ctx, &system.UpdateNoticeReq{
			Id:   req.GetId(),
			Text: req.Text,
		})
		if err != nil {
			return nil, err
		}
	}

	return resp, nil
}

func (s *SystemService) UpdateCommunityAdmin(ctx context.Context, req *core_api.UpdateCommunityAdminReq, user *genbasic.UserMeta) (*core_api.UpdateCommunityAdminResp, error) {
	if user.GetUserId() == "" {
		return nil, consts.ErrNotAuthentication
	}
	resp := new(core_api.UpdateCommunityAdminResp)
	data, err := s.System.RetrieveUserRole(ctx, &system.RetrieveUserRoleReq{UserId: req.UserId})
	if err != nil {
		return nil, err
	}
	if !req.IsRemove {
		for _, role := range data.Roles {
			if role.RoleType == system.RoleType_TypeCommunityAdmin && *role.CommunityId == req.CommunityId {
				return nil, err
				// TODO 应当返回错误
			}
		}
		_, err = s.System.UpdateUserRole(ctx, &system.UpdateUserRoleReq{
			UserId: req.UserId,
			Roles: append(data.Roles, &system.Role{
				RoleType:    system.RoleType_TypeCommunityAdmin,
				CommunityId: &req.CommunityId,
			}),
		})
		if err != nil {
			return nil, err
		}
	} else {
		roles := make([]*system.Role, 0, len(data.Roles))
		for _, role := range data.Roles {
			if role.RoleType != system.RoleType_TypeCommunityAdmin || *role.CommunityId != req.CommunityId {
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

func (s *SystemService) UpdateSuperAdmin(ctx context.Context, req *core_api.UpdateSuperAdminReq, user *genbasic.UserMeta) (*core_api.UpdateSuperAdminResp, error) {
	if user.GetUserId() == "" {
		return nil, consts.ErrNotAuthentication
	}
	resp := new(core_api.UpdateSuperAdminResp)
	data, err := s.System.RetrieveUserRole(ctx, &system.RetrieveUserRoleReq{UserId: req.UserId})
	if err != nil {
		return nil, err
	}
	if !req.IsRemove {
		for _, role := range data.Roles {
			if role.RoleType == system.RoleType_TypeSuperAdmin {
				return nil, err
				// TODO 应当返回错误
			}
		}
		_, err = s.System.UpdateUserRole(ctx, &system.UpdateUserRoleReq{
			UserId: req.UserId,
			Roles: append(data.Roles, &system.Role{
				RoleType: system.RoleType_TypeSuperAdmin,
			}),
		})
		if err != nil {
			return nil, err
		}
	} else {
		roles := make([]*system.Role, 0, len(data.Roles))
		for _, role := range data.Roles {
			if role.RoleType != system.RoleType_TypeSuperAdmin {
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

func (s *SystemService) UpdateRole(ctx context.Context, req *core_api.UpdateRoleReq, user *genbasic.UserMeta) (*core_api.UpdateRoleResp, error) {
	resp := new(core_api.UpdateRoleResp)

	roles := make([]*system.Role, 0, len(req.Roles))
	for _, role := range req.Roles {
		var r system.Role
		r.RoleType = system.RoleType(role.RoleType)
		r.CommunityId = role.CommunityId
		roles = append(roles, &r)
	}
	_, err := s.System.UpdateUserRole(ctx, &system.UpdateUserRoleReq{
		UserId: req.UserId,
		Roles:  roles,
	})
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *SystemService) GetMinVersion(ctx context.Context, req *core_api.GetMinVersionReq) (*core_api.GetMinVersionResp, error) {
	version := s.Config.MinVersion
	version = version[1:]
	res := new(core_api.GetMinVersionResp)
	res.MinVersion = 0
	time, tmp, n := 1, 0, 1
	for i := len(version) - 1; i >= 0; i-- {
		if version[i] == '.' {
			res.MinVersion += int64(time * tmp)
			time *= 100
			n = 1
			tmp = 0
		} else {
			tmp += int(version[i]-'0') * n
			n *= 10
			println(tmp)
		}
	}
	res.MinVersion += int64(time * tmp)
	return res, nil
}

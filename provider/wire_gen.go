// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package provider

import (
	service2 "github.com/xh-polaris/meowchat-core-api/biz/application/service"
	"github.com/xh-polaris/meowchat-core-api/biz/domain/service"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/config"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/meowchat_content"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/meowchat_system"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/meowchat_user"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/platform_comment"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/platform_sts"
)

// Injectors from wire.go:

func NewProvider() (*Provider, error) {
	configConfig, err := config.NewConfig()
	if err != nil {
		return nil, err
	}
	client := meowchat_content.NewMeowchatContent(configConfig)
	meowchatContent := &meowchat_content.MeowchatContent{
		Client: client,
	}
	stsserviceClient := platform_sts.NewPlatformSts(configConfig)
	platformSts := &platform_sts.PlatformSts{
		Client: stsserviceClient,
	}
	userserviceClient := meowchat_user.NewMeowchatUser(configConfig)
	meowchatUser := &meowchat_user.MeowchatUser{
		Client: userserviceClient,
	}
	catImageDomainService := &service.CatImageDomainService{
		MeowchatContent: meowchatContent,
		MeowchatUser:    meowchatUser,
	}
	collectionService := &service2.CollectionService{
		MeowchatContent:       meowchatContent,
		Config:                configConfig,
		PlatformSts:           platformSts,
		CatImageDomainService: catImageDomainService,
	}
	authService := &service2.AuthService{
		Config:  configConfig,
		Sts:     platformSts,
		Content: meowchatContent,
	}
	commentserviceClient := platform_comment.NewPlatformComment(configConfig)
	platformComment := &platform_comment.PlatformComment{
		Client: commentserviceClient,
	}
	commentDomainService := &service.CommentDomainService{
		MeowchatContent: meowchatContent,
		MeowchatUser:    meowchatUser,
		PlatformComment: platformComment,
	}
	systemrpcClient := meowchat_system.NewMeowchatSystem(configConfig)
	meowchatSystem := &meowchat_system.MeowchatSystem{
		Client: systemrpcClient,
	}
	commentService := &service2.CommentService{
		Config:               configConfig,
		CommentDomainService: commentDomainService,
		PlatformComment:      platformComment,
		PlatformSts:          platformSts,
		MeowchatContent:      meowchatContent,
		MeowchatSystem:       meowchatSystem,
	}
	userDomainService := &service.UserDomainService{
		MeowchatUser:    meowchatUser,
		MeowchatContent: meowchatContent,
		MeowchatSystem:  meowchatSystem,
	}
	userService := &service2.UserService{
		Config:       configConfig,
		UserService:  userDomainService,
		MeowchatUser: meowchatUser,
		PlatformSts:  platformSts,
	}
	momentDomainService := &service.MomentDomainService{
		MeowchatContent: meowchatContent,
		MeowchatUser:    meowchatUser,
		PlatformComment: platformComment,
	}
	momentService := &service2.MomentService{
		Config:              configConfig,
		MomentDomainService: momentDomainService,
		MeowchatContent:     meowchatContent,
		MeowchatUser:        meowchatUser,
		PlatformSts:         platformSts,
	}
	postDomainService := &service.PostDomainService{
		MeowchatUser:    meowchatUser,
		PlatformComment: platformComment,
	}
	postService := &service2.PostService{
		Config:            configConfig,
		PostDomainService: postDomainService,
		MeowchatContent:   meowchatContent,
		PlatformSts:       platformSts,
	}
	likeService := &service2.LikeService{
		Config:               configConfig,
		MeowchatUser:         meowchatUser,
		MeowchatContent:      meowchatContent,
		PlatformComment:      platformComment,
		UserDomainService:    userDomainService,
		PostDomainService:    postDomainService,
		MomentDomainService:  momentDomainService,
		CommentDomainService: commentDomainService,
		MeowchatSystem:       meowchatSystem,
	}
	stsService := &service2.StsService{
		PlatformSts: platformSts,
	}
	systemService := &service2.SystemService{
		Config: configConfig,
		System: meowchatSystem,
		User:   meowchatUser,
	}
	planService := &service2.PlanService{
		Config: configConfig,
		Plan:   meowchatContent,
		User:   meowchatUser,
		Sts:    platformSts,
	}
	incentiveService := &service2.IncentiveService{
		Config:          configConfig,
		MeowchatContent: meowchatContent,
	}
	providerProvider := &Provider{
		Config:            configConfig,
		CollectionService: collectionService,
		AuthService:       authService,
		CommentService:    commentService,
		UserService:       userService,
		MomentService:     momentService,
		PostService:       postService,
		LikeService:       likeService,
		StsService:        stsService,
		SystemService:     systemService,
		PlanService:       planService,
		IncentiveService:  incentiveService,
	}
	return providerProvider, nil
}

// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package provider

import (
	"github.com/xh-polaris/meowchat-core-api/biz/application/service"
	service2 "github.com/xh-polaris/meowchat-core-api/biz/domain/service"
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
	collectionService := &service.CollectionService{
		Collection: meowchatContent,
		Config:     configConfig,
		Sts:        platformSts,
	}
	authService := &service.AuthService{
		Config:  configConfig,
		Sts:     platformSts,
		Content: meowchatContent,
	}
	userserviceClient := meowchat_user.NewMeowchatUser(configConfig)
	meowchatUser := &meowchat_user.MeowchatUser{
		Client: userserviceClient,
	}
	commentserviceClient := platform_comment.NewPlatformComment(configConfig)
	platformComment := &platform_comment.PlatformComment{
		Client: commentserviceClient,
	}
	commentDomainService := &service2.CommentDomainService{
		MeowchatContent:  meowchatContent,
		MeowchatUser:     meowchatUser,
		PlatformCommment: platformComment,
	}
	commentService := &service.CommentService{
		Config:               configConfig,
		CommentDomainService: commentDomainService,
		PlatformComment:      platformComment,
		PlatformSts:          platformSts,
		MeowchatContent:      meowchatContent,
	}
	systemrpcClient := meowchat_system.NewMeowchatSystem(configConfig)
	meowchatSystem := &meowchat_system.MeowchatSystem{
		Client: systemrpcClient,
	}
	userDomainService := &service2.UserDomainService{
		MeowchatUser:    meowchatUser,
		MeowchatContent: meowchatContent,
		MeowchatSystem:  meowchatSystem,
	}
	userService := &service.UserService{
		Config:          configConfig,
		UserService:     userDomainService,
		MeowchatUser:    meowchatUser,
		PlatformSts:     platformSts,
		MeowchatContent: meowchatContent,
	}
	momentDomainService := &service2.MomentDomainService{
		MeowchatContent:  meowchatContent,
		MeowchatUser:     meowchatUser,
		PlatformCommment: platformComment,
	}
	momentService := &service.MomentService{
		Config:              configConfig,
		MomentDomainService: momentDomainService,
		MeowchatContent:     meowchatContent,
		MeowchatUser:        meowchatUser,
		PlatformCommment:    platformComment,
		PlatformSts:         platformSts,
	}
	postDomainService := &service2.PostDomainService{
		MeowchatUser:     meowchatUser,
		PlatformCommment: platformComment,
	}
	postService := &service.PostService{
		Config:            configConfig,
		PostDomainService: postDomainService,
		MeowchatContent:   meowchatContent,
		PlatformSts:       platformSts,
	}
	likeService := &service.LikeService{
		Config:  configConfig,
		User:    meowchatUser,
		Content: meowchatContent,
	}
	stsService := &service.StsService{
		PlatformSts: platformSts,
	}
	systemService := &service.SystemService{
		Config: configConfig,
		System: meowchatSystem,
		User:   meowchatUser,
	}
	planService := &service.PlanService{
		Config: configConfig,
		Plan:   meowchatContent,
		User:   meowchatUser,
		Sts:    platformSts,
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
	}
	return providerProvider, nil
}

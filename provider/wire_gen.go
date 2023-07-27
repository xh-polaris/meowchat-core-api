// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package provider

import (
	"github.com/xh-polaris/meowchat-core-api/biz/application/service"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/config"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/meowchat_collection"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/meowchat_like"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/meowchat_moment"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/meowchat_post"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/meowchat_user"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/platform_authentication"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/platform_comment"
)

// Injectors from wire.go:

func NewProvider() (*Provider, error) {
	configConfig, err := config.NewConfig()
	if err != nil {
		return nil, err
	}
	client := meowchat_collection.NewMeowchatCollection(configConfig)
	meowchatCollection := &meowchat_collection.MeowchatCollection{
		Client: client,
	}
	collectionService := &service.CollectionService{
		Collection: meowchatCollection,
		Config:     configConfig,
	}
	auth := platform_authentication.NewPlatformAuthentication(configConfig)
	platformAuthentication := &platform_authentication.PlatformAuthentication{
		Auth: auth,
	}
	authService := &service.AuthService{
		Config:         configConfig,
		Authentication: platformAuthentication,
	}
	commentRpc := platform_comment.NewPlatformComment(configConfig)
	platformComment := &platform_comment.PlatformComment{
		CommentRpc: commentRpc,
	}
	userRpc := meowchat_user.NewMeowchatUser(configConfig)
	meowchatUser := &meowchat_user.MeowchatUser{
		UserRpc: userRpc,
	}
	commentService := &service.CommentService{
		Config:  configConfig,
		Comment: platformComment,
		User:    meowchatUser,
	}
	momentRpc := meowchat_moment.NewMeowchatMoment(configConfig)
	meowchatMoment := &meowchat_moment.MeowchatMoment{
		MomentRpc: momentRpc,
	}
	likerpc := meowchat_like.NewMeowchatLike(configConfig)
	meowchatLike := &meowchat_like.MeowchatLike{
		Likerpc: likerpc,
	}
	postRpc := meowchat_post.NewMeowchatPost(configConfig)
	meowchatPost := &meowchat_post.MeowchatPost{
		PostRpc: postRpc,
	}
	userService := &service.UserService{
		Config: configConfig,
		User:   meowchatUser,
		Moment: meowchatMoment,
		Like:   meowchatLike,
		Post:   meowchatPost,
	}
	momentService := &service.MomentService{
		Config: configConfig,
		Moment: meowchatMoment,
		User:   meowchatUser,
	}
	postService := service.PostService{
		Config:  configConfig,
		Post:    meowchatPost,
		User:    meowchatUser,
		Like:    meowchatLike,
		Comment: platformComment,
	}
	likeService := &service.LikeService{
		Config: configConfig,
		Like:   meowchatLike,
		User:   meowchatUser,
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
	}
	return providerProvider, nil
}

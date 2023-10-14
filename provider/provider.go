package provider

import (
	"github.com/google/wire"

	"github.com/xh-polaris/meowchat-core-api/biz/application/service"
	domainservice "github.com/xh-polaris/meowchat-core-api/biz/domain/service"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/config"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/meowchat_content"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/meowchat_system"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/meowchat_user"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/platform_comment"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/platform_sts"
)

var provider *Provider

func Init() {
	var err error
	provider, err = NewProvider()
	if err != nil {
		panic(err)
	}
}

// Provider 提供controller依赖的对象
type Provider struct {
	Config            *config.Config
	CollectionService service.ICollectionService
	AuthService       service.IAuthService
	CommentService    service.ICommentService
	UserService       service.IUserService
	MomentService     service.IMomentService
	PostService       service.IPostService
	LikeService       service.ILikeService
	StsService        service.IStsService
	SystemService     service.ISystemService
	PlanService       service.IPlanService
}

func Get() *Provider {
	return provider
}

var RPCSet = wire.NewSet(
	meowchat_content.MeowchatContentSet,
	platform_sts.PlatformStsSet,
	platform_comment.PlatformCommentSet,
	meowchat_user.MeowchatUserSet,
	meowchat_system.MeowchatSystemSet,
)

var ApplicationSet = wire.NewSet(
	service.CollectionServiceSet,
	service.AuthServiceSet,
	service.CommentServiceSet,
	service.UserServiceSet,
	service.MomentServiceSet,
	service.LikeServiceSet,
	service.PostServiceSet,
	service.SystemServiceSet,
	service.StsServiceSet,
	service.PlanServiceSet,
)

var DomainSet = wire.NewSet(
	domainservice.UserDomainServiceSet,
	domainservice.CommentDomainServiceSet,
	domainservice.MomentDomainServiceSet,
	domainservice.PostDomainServiceSet,
)

var InfrastructureSet = wire.NewSet(
	config.NewConfig,
	RPCSet,
)

var AllProvider = wire.NewSet(
	ApplicationSet,
	DomainSet,
	InfrastructureSet,
)

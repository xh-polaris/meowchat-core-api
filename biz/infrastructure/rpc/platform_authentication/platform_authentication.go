package platform_authentication

import (
	"context"

	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/config"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/util"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/util/log"

	"github.com/google/wire"
	"github.com/xh-polaris/auth-rpc/auth"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type IPlatformAuthentication interface {
	auth.Auth
}

type PlatformAuthentication struct {
	auth.Auth
}

var PlatformAuthenticationSet = wire.NewSet(
	NewPlatformAuthentication,
	wire.Struct(new(PlatformAuthentication), "*"),
	wire.Bind(new(IPlatformAuthentication), new(*PlatformAuthentication)),
)

func NewPlatformAuthentication(config *config.Config) auth.Auth {
	return auth.NewAuth(zrpc.MustNewClient(
		config.AuthRPC,
		zrpc.WithUnaryClientInterceptor(func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
			err := invoker(ctx, method, req, reply, cc)
			log.CtxInfo(ctx, "[%s] req=%s, resp=%s, err=%v", method, util.JSONF(req), util.JSONF(reply), err)
			return err
		}),
	))
}

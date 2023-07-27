package meowchat_user

import (
	"context"
	"github.com/google/wire"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/config"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/util"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/util/log"
	"github.com/xh-polaris/meowchat-user-rpc/userrpc"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type IMeowchatUser interface {
	userrpc.UserRpc
}

type MeowchatUser struct {
	userrpc.UserRpc
}

var MeowchatUserSet = wire.NewSet(
	NewMeowchatUser,
	wire.Struct(new(MeowchatUser), "*"),
	wire.Bind(new(IMeowchatUser), new(*MeowchatUser)),
)

func NewMeowchatUser(config *config.Config) userrpc.UserRpc {
	return userrpc.NewUserRpc(zrpc.MustNewClient(
		config.UserRPC,
		zrpc.WithUnaryClientInterceptor(func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
			err := invoker(ctx, method, req, reply, cc)
			log.CtxInfo(ctx, "[%s] req=%s, resp=%s, err=%v", method, util.JSONF(req), util.JSONF(reply), err)
			return err
		}),
	))
}

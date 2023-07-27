package meowchat_moment

import (
	"context"
	"github.com/google/wire"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/config"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/util"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/util/log"
	"github.com/xh-polaris/meowchat-moment-rpc/momentrpc"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type IMeowchatMoment interface {
	momentrpc.MomentRpc
}

type MeowchatMoment struct {
	momentrpc.MomentRpc
}

var MeowchatMomentSet = wire.NewSet(
	NewMeowchatMoment,
	wire.Struct(new(MeowchatMoment), "*"),
	wire.Bind(new(IMeowchatMoment), new(*MeowchatMoment)),
)

func NewMeowchatMoment(config *config.Config) momentrpc.MomentRpc {
	return momentrpc.NewMomentRpc(zrpc.MustNewClient(
		config.MomentRPC,
		zrpc.WithUnaryClientInterceptor(func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
			err := invoker(ctx, method, req, reply, cc)
			log.CtxInfo(ctx, "[%s] req=%s, resp=%s, err=%v", method, util.JSONF(req), util.JSONF(reply), err)
			return err
		}),
	))
}

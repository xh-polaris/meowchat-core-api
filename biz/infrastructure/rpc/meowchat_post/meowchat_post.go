package meowchat_post

import (
	"context"
	"github.com/google/wire"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/config"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/util"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/util/log"
	"github.com/xh-polaris/meowchat-post-rpc/postrpc"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type IMeowchatPost interface {
	postrpc.PostRpc
}

type MeowchatPost struct {
	postrpc.PostRpc
}

var MeowchatPostSet = wire.NewSet(
	NewMeowchatPost,
	wire.Struct(new(MeowchatPost), "*"),
	wire.Bind(new(IMeowchatPost), new(*MeowchatPost)),
)

func NewMeowchatPost(config *config.Config) postrpc.PostRpc {
	return postrpc.NewPostRpc(zrpc.MustNewClient(
		config.PostRPC,
		zrpc.WithUnaryClientInterceptor(func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
			err := invoker(ctx, method, req, reply, cc)
			log.CtxInfo(ctx, "[%s] req=%s, resp=%s, err=%v", method, util.JSONF(req), util.JSONF(reply), err)
			return err
		}),
	))
}

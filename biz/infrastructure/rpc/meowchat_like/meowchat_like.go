package meowchat_like

import (
	"context"
	"github.com/google/wire"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/config"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/util"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/util/log"
	"github.com/xh-polaris/meowchat-like-rpc/likerpc"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type IMeowchatLike interface {
	likerpc.Likerpc
}

type MeowchatLike struct {
	likerpc.Likerpc
}

var MeowchatLikeSet = wire.NewSet(
	NewMeowchatLike,
	wire.Struct(new(MeowchatLike), "*"),
	wire.Bind(new(IMeowchatLike), new(*MeowchatLike)),
)

func NewMeowchatLike(config *config.Config) likerpc.Likerpc {
	return likerpc.NewLikerpc(zrpc.MustNewClient(
		config.LikeRPC,
		zrpc.WithUnaryClientInterceptor(func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
			err := invoker(ctx, method, req, reply, cc)
			log.CtxInfo(ctx, "[%s] req=%s, resp=%s, err=%v", method, util.JSONF(req), util.JSONF(reply), err)
			return err
		}),
	))
}

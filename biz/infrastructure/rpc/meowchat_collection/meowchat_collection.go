package meowchat_collection

import (
	"context"

	"github.com/google/wire"
	"github.com/xh-polaris/meowchat-collection-rpc/collectionrpc"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"

	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/config"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/util"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/util/log"
)

type IMeowchatCollection interface {
	collectionrpc.CollectionRpc
}

type MeowchatCollection struct {
	collectionrpc.CollectionRpc
}

var MeowchatCollectionSet = wire.NewSet(
	NewMeowchatCollection,
	wire.Struct(new(MeowchatCollection), "*"),
	wire.Bind(new(IMeowchatCollection), new(*MeowchatCollection)),
)

func NewMeowchatCollection(config *config.Config) collectionrpc.CollectionRpc {
	return collectionrpc.NewCollectionRpc(zrpc.MustNewClient(
		config.CollectionRPC,
		zrpc.WithUnaryClientInterceptor(func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
			err := invoker(ctx, method, req, reply, cc)
			log.CtxInfo(ctx, "[%s] req=%s, resp=%s, err=%v", method, util.JSONF(req), util.JSONF(reply), err)
			return err
		}),
	))
}

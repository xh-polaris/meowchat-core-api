package rpc

import (
	"github.com/google/wire"
	"github.com/xh-polaris/meowchat-collection-rpc/collectionrpc"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/config"
	"github.com/zeromicro/go-zero/zrpc"
)

type IMeowchatCollection interface {
	collectionrpc.CollectionRpc
}

type MeowchatCollection struct {
	collectionrpc.CollectionRpc
}

var MeowchatCollectionSet = wire.NewSet(
	NewCollectionRPC,
	wire.Struct(new(MeowchatCollection), "*"),
	wire.Bind(new(IMeowchatCollection), new(*MeowchatCollection)),
)

func NewCollectionRPC(config *config.Config) collectionrpc.CollectionRpc {
	return collectionrpc.NewCollectionRpc(zrpc.MustNewClient(config.CollectionRPC))
}

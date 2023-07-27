package platform_comment

import (
	"context"

	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/config"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/util"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/util/log"

	"github.com/google/wire"
	"github.com/xh-polaris/meowchat-comment-rpc/commentrpc"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type IPlatformCommment interface {
	commentrpc.CommentRpc
}

type PlatformComment struct {
	commentrpc.CommentRpc
}

var PlatformCommentSet = wire.NewSet(
	NewPlatformComment,
	wire.Struct(new(PlatformComment), "*"),
	wire.Bind(new(IPlatformCommment), new(*PlatformComment)),
)

func NewPlatformComment(config *config.Config) commentrpc.CommentRpc {
	return commentrpc.NewCommentRpc(zrpc.MustNewClient(
		config.CommentRPC,
		zrpc.WithUnaryClientInterceptor(func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
			err := invoker(ctx, method, req, reply, cc)
			log.CtxInfo(ctx, "[%s] req=%s, resp=%s, err=%v", method, util.JSONF(req), util.JSONF(reply), err)
			return err
		}),
	))
}

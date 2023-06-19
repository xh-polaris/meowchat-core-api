package adaptor

import (
	"context"
	"github.com/google/wire"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

type IHandler interface {
	HandlerError(ctx context.Context, err error, c *app.RequestContext)
}

type Handler struct{}

var HandlerSet = wire.NewSet(
	wire.Struct(new(Handler), "*"),
	wire.Bind(new(IHandler), new(*Handler)),
)

func (h *Handler) HandlerError(ctx context.Context, err error, c *app.RequestContext) {
	switch err.(type) {
	case nil:
		return
	default:
		code := consts.StatusInternalServerError
		c.String(code, consts.StatusMessage(code))
	}
}

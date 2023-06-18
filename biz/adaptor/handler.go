package adaptor

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func HandlerError(ctx context.Context, err error, c *app.RequestContext) {
	switch err.(type) {
	case nil:
		return
	default:
		code := consts.StatusInternalServerError
		c.String(code, consts.StatusMessage(code))
	}
}

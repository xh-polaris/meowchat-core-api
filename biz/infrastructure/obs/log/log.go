package log

import (
	"context"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	hertzlogrus "github.com/hertz-contrib/obs-opentelemetry/logging/logrus"
)

func Init() {
	hlog.SetLogger(hertzlogrus.NewLogger())
}

func CtxInfo(ctx context.Context, format string, v ...any) {
	hlog.CtxInfof(ctx, format, v)
}

func CtxError(ctx context.Context, format string, v ...any) {
	hlog.CtxErrorf(ctx, format, v)
}

func CtxWarn(ctx context.Context, format string, v ...any) {
	hlog.CtxWarnf(ctx, format, v)
}

func CtxDebug(ctx context.Context, format string, v ...any) {
	hlog.CtxDebugf(ctx, format, v)
}

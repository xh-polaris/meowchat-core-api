package log

import (
	"context"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	hertzlogrus "github.com/hertz-contrib/obs-opentelemetry/logging/logrus"
)

var levelMap = map[string]hlog.Level{
	"debug": hlog.LevelDebug,
	"info":  hlog.LevelInfo,
	"error": hlog.LevelError,
}

func Init() {
	hlog.SetLogger(hertzlogrus.NewLogger())
}

func SetLevel(level string) {
	hlog.SetLevel(levelMap[level])
}

func CtxInfo(ctx context.Context, format string, v ...any) {
	hlog.CtxInfof(ctx, format, v...)
}

func Info(format string, v ...any) {
	hlog.CtxInfof(context.Background(), format, v...)
}

func CtxError(ctx context.Context, format string, v ...any) {
	hlog.CtxErrorf(ctx, format, v...)
}

func Error(format string, v ...any) {
	hlog.CtxErrorf(context.Background(), format, v...)
}

func CtxWarn(ctx context.Context, format string, v ...any) {
	hlog.CtxWarnf(ctx, format, v...)
}

func CtxDebug(ctx context.Context, format string, v ...any) {
	hlog.CtxDebugf(ctx, format, v...)
}

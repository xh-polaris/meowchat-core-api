package util

import (
	"github.com/cloudwego/hertz/pkg/common/json"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/util/log"
	"strconv"
)

func JSONF(v any) string {
	data, err := json.Marshal(v)
	if err != nil {
		log.Error("JSONF fail, v=%v, err=%v", v, err)
	}
	return string(data)
}

func ParseInt(s string) int64 {
	appID, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	}
	return appID
}

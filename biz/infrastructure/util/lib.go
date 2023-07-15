package util

import (
	"github.com/cloudwego/hertz/pkg/common/json"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/util/log"
)

func JSONF(v any) string {
	data, err := json.Marshal(v)
	if err != nil {
		log.Error("JSONF fail, v=%v, err=%v", v, err)
	}
	return string(data)
}

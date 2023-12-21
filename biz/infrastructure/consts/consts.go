package consts

import (
	"errors"

	"github.com/xh-polaris/meowchat-core-api/biz/application/dto/basic"
)

const DefaultPageSize int64 = 10

var APPMap = map[basic.APP]string{
	basic.APP_Meowchat:        "wxd39cebf05e21d3b6",
	basic.APP_MeowchatManager: "wx40ab73e6ebd6e636",
}

var ErrNotAuthentication = errors.New("not authentication")
var ErrForbidden = errors.New("forbidden")
var PageSize int64 = 10

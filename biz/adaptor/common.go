package adaptor

import (
	"context"
	"errors"
	"reflect"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server/binding"
	"github.com/cloudwego/hertz/pkg/common/json"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/golang-jwt/jwt/v4"

	"github.com/xh-polaris/meowchat-core-api/biz/application/dto/meowchat"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/util"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/util/log"
	"github.com/xh-polaris/meowchat-core-api/provider"
)

func Init() {
	binding.SetLooseZeroMode(true)
	binding.MustRegTypeUnmarshal(reflect.TypeOf(meowchat.UserMeta{}), func(v string, emptyAsZero bool) (reflect.Value, error) {
		if v == "" && emptyAsZero {
			return reflect.ValueOf(meowchat.UserMeta{}), nil
		}
		token, err := jwt.Parse(v, func(_ *jwt.Token) (interface{}, error) {
			return []byte(provider.Get().Config.Auth.AccessSecret), nil
		})
		if err != nil {
			return reflect.ValueOf(meowchat.UserMeta{}), err
		}
		if !token.Valid {
			return reflect.ValueOf(meowchat.UserMeta{}), errors.New("token is not valid")
		}
		data, err := json.Marshal(token.Claims)
		if err != nil {
			return reflect.ValueOf(meowchat.UserMeta{}), err
		}
		user := new(meowchat.UserMeta)
		err = json.Unmarshal(data, user)
		if err != nil {
			return reflect.ValueOf(meowchat.UserMeta{}), err
		}
		return reflect.ValueOf(*user), nil
	})
}

func LogAndReturn(ctx context.Context, c *app.RequestContext, req, resp any, err error) {
	log.CtxInfo(ctx, "[%s] req=%s, resp=%s, err=%v", c.Path(), util.JSONF(req), util.JSONF(resp), err)
	switch err.(type) {
	case nil:
		c.JSON(consts.StatusOK, resp)
	default:
		log.CtxError(ctx, "internal error, err=%s", err.Error())
		code := consts.StatusInternalServerError
		c.String(code, consts.StatusMessage(code))
	}
}

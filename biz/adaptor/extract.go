package adaptor

import (
	"context"
	"errors"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/json"
	"github.com/golang-jwt/jwt/v4"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/basic"

	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/util"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/util/log"
	"github.com/xh-polaris/meowchat-core-api/provider"
)

func ExtractMeta(ctx context.Context, c *app.RequestContext) (*basic.UserMeta, *basic.Extra) {
	return ExtractUserMeta(ctx, c), ExtractExtra(ctx, c)
}

func ExtractUserMeta(ctx context.Context, c *app.RequestContext) (user *basic.UserMeta) {
	user = new(basic.UserMeta)
	var err error
	defer func() {
		if err != nil {
			log.CtxError(ctx, "extract user meta fail, err=%v", err)
		}
	}()
	tokenString := c.GetHeader("Authorization")
	token, err := jwt.Parse(string(tokenString), func(_ *jwt.Token) (interface{}, error) {
		return []byte(provider.Get().Config.Auth.AccessSecret), nil
	})
	if err != nil {
		return
	}
	if !token.Valid {
		err = errors.New("token is not valid")
		return
	}
	data, err := json.Marshal(token.Claims)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, user)
	if err != nil {
		return
	}
	if user.SessionUserId == "" {
		user.SessionUserId = user.UserId
	}
	if user.SessionAppId == 0 {
		user.SessionAppId = user.AppId
	}
	if user.SessionDeviceId == "" {
		user.SessionDeviceId = user.DeviceId
	}
	log.CtxInfo(ctx, "userMeta=%s", util.JSONF(user))
	return
}

func ExtractExtra(ctx context.Context, c *app.RequestContext) *basic.Extra {
	extra := new(basic.Extra)
	extra.ClientIP = c.ClientIP()
	err := c.Bind(extra)
	if err != nil {
		log.CtxError(ctx, "extract extra fail, err=%v", err)
		return nil
	}
	log.CtxInfo(ctx, "extra=%s", util.JSONF(extra))
	return extra
}

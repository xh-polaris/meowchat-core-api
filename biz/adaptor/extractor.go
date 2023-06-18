package adaptor

import (
	"context"
	"errors"

	"github.com/xh-polaris/meowchat-core-api/biz/application/dto/meowchat"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/config"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/obs/log"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/json"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/wire"
)

const AuthorizationHeader = "Authorization"

type IExtractor interface {
	Extract(ctx context.Context, c *app.RequestContext) (*meowchat.UserMeta, *meowchat.Extra)
	ExtractUserMeta(ctx context.Context, c *app.RequestContext) (user *meowchat.UserMeta)
	ExtractExtra(ctx context.Context, c *app.RequestContext) *meowchat.Extra
}

type Extractor struct {
	Config *config.Config
}

var ExtractorSet = wire.NewSet(
	wire.Struct(new(Extractor), "*"),
	wire.Bind(new(IExtractor), new(*Extractor)),
)

func (e *Extractor) Extract(ctx context.Context, c *app.RequestContext) (*meowchat.UserMeta, *meowchat.Extra) {
	return e.ExtractUserMeta(ctx, c), e.ExtractExtra(ctx, c)
}

func (e *Extractor) ExtractUserMeta(ctx context.Context, c *app.RequestContext) (user *meowchat.UserMeta) {
	user = new(meowchat.UserMeta)
	var err error
	defer func() {
		if err != nil {
			log.CtxError(ctx, "extract user meta fail, err=%v", err)
		}
	}()
	tokenString := c.GetHeader(AuthorizationHeader)
	token, err := jwt.Parse(string(tokenString), func(_ *jwt.Token) (interface{}, error) {
		return []byte(e.Config.Auth.AccessSecret), nil
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
	return
}

func (e *Extractor) ExtractExtra(ctx context.Context, c *app.RequestContext) *meowchat.Extra {
	extra := new(meowchat.Extra)
	extra.ClientIP = c.ClientIP()
	err := c.Bind(extra)
	if err != nil {
		log.CtxError(ctx, "extract extra fail, err=%v", err)
		return nil
	}
	return extra
}

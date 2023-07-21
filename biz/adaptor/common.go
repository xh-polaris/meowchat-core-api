package adaptor

import (
	"context"
	"errors"
	"net/http"
	"reflect"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server/binding"
	"github.com/cloudwego/hertz/pkg/common/json"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/golang-jwt/jwt/v4"
	"go.opentelemetry.io/contrib/propagators/b3"
	"go.opentelemetry.io/otel/propagation"
	"google.golang.org/grpc/status"

	"github.com/xh-polaris/meowchat-core-api/biz/application/dto/base"
	"github.com/xh-polaris/meowchat-core-api/biz/application/dto/meowchat/basic"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/util"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/util/log"
	"github.com/xh-polaris/meowchat-core-api/provider"
)

func Init() {
	binding.SetLooseZeroMode(true)
	binding.MustRegTypeUnmarshal(reflect.TypeOf(basic.UserMeta{}), func(v string, emptyAsZero bool) (reflect.Value, error) {
		if v == "" && emptyAsZero {
			return reflect.ValueOf(basic.UserMeta{}), nil
		}
		token, err := jwt.Parse(v, func(_ *jwt.Token) (interface{}, error) {
			return []byte(provider.Get().Config.Auth.AccessSecret), nil
		})
		if err != nil {
			return reflect.ValueOf(basic.UserMeta{}), err
		}
		if !token.Valid {
			return reflect.ValueOf(basic.UserMeta{}), errors.New("token is not valid")
		}
		data, err := json.Marshal(token.Claims)
		if err != nil {
			return reflect.ValueOf(basic.UserMeta{}), err
		}
		user := new(basic.UserMeta)
		err = json.Unmarshal(data, user)
		if err != nil {
			return reflect.ValueOf(basic.UserMeta{}), err
		}
		return reflect.ValueOf(*user), nil
	})
}

var _ propagation.TextMapCarrier = &headerProvider{}

type headerProvider struct {
	headers *protocol.ResponseHeader
}

// Get a value from metadata by key
func (m *headerProvider) Get(key string) string {
	return m.headers.Get(key)
}

// Set a value to metadata by k/v
func (m *headerProvider) Set(key, value string) {
	m.headers.Set(key, value)
}

// Keys Iteratively get all keys of metadata
func (m *headerProvider) Keys() []string {
	out := make([]string, 0)

	m.headers.VisitAll(func(key, value []byte) {
		out = append(out, string(key))
	})

	return out
}

func Return(ctx context.Context, c *app.RequestContext, req, resp any, err error) {
	log.CtxInfo(ctx, "[%s] req=%s, resp=%s, err=%v", c.Path(), util.JSONF(req), util.JSONF(resp), err)
	b3.New().Inject(ctx, &headerProvider{headers: &c.Response.Header})

	switch err.(type) {
	case nil:
		c.JSON(consts.StatusOK, resp)
	default:
		if s, ok := status.FromError(err); ok {
			c.JSON(http.StatusBadRequest, &base.Status{
				Code: int64(int(s.Code())),
				Msg:  s.Message(),
			})
		} else {
			log.CtxError(ctx, "internal error, err=%s", err.Error())
			code := consts.StatusInternalServerError
			c.String(code, consts.StatusMessage(code))
		}
	}
}

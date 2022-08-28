package auth

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/lffwl/utility/jwt"
	"time"
	"www.lffwl.com/internal/consts"
	"www.lffwl.com/utility/response"
)

var Jwt *jwt.Jwt

func NewJwt(ctx context.Context) *jwt.Jwt {
	return jwt.NewJwt(jwt.Config{
		ExpiryAt:       g.Cfg("auth").MustGet(ctx, "jwt.expiryAt").Duration() * time.Hour,
		RefreshTokenAt: g.Cfg("auth").MustGet(ctx, "jwt.refreshTokenAt").Duration() * time.Hour,
		EncryptKey:     []byte(g.Cfg("auth").MustGet(ctx, "jwt.encryptKey").String()),
		TokenKey:       g.Cfg("auth").MustGet(ctx, "jwt.tokenKey").String(),
		TokenLookup:    g.Cfg("auth").MustGet(ctx, "jwt.tokenLookup").String(),
		Issuer:         g.Cfg("auth").MustGet(ctx, "jwt.issuer").String(),
	})
}

func Auth(r *ghttp.Request) {

	verifyAuth := true

	for _, val := range g.Cfg("auth").MustGet(r.Context(), "notLogin").Strings() {
		if r.URL.Path == val {
			verifyAuth = false
			break
		}
	}

	if verifyAuth {

		data, err := Jwt.VerifyToken(r.Request)
		if err.Error != nil {
			response.JsonExit(r, err.Code, err.Error.Error())
			return
		}

		r.SetCtxVar(consts.AuthCtxAdminIdKey, data.(map[string]interface{})["Id"])
	}

	r.Middleware.Next()
}

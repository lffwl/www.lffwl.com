package auth

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/lffwl/utility/jwt"
	"time"
)

func Jwt(ctx context.Context) *jwt.Jwt {
	return jwt.NewJwt(jwt.Config{
		ExpiryAt:       g.Cfg("jwt").MustGet(ctx, "jwt.expiryAt").Duration() * time.Hour,
		RefreshTokenAt: g.Cfg("jwt").MustGet(ctx, "jwt.refreshTokenAt").Duration() * time.Hour,
		EncryptKey:     []byte(g.Cfg("jwt").MustGet(ctx, "jwt.encryptKey").String()),
		TokenKey:       g.Cfg("jwt").MustGet(ctx, "jwt.tokenKey").String(),
		TokenLookup:    g.Cfg("jwt").MustGet(ctx, "jwt.tokenLookup").String(),
		Issuer:         g.Cfg("jwt").MustGet(ctx, "jwt.issuer").String(),
	})
}

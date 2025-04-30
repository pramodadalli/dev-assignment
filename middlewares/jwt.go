package middlewares

import (
	"net/http"
	"strings"

	"github.com/beego/beego/v2/server/web/context"
	"github.com/golang-jwt/jwt/v4"
)

var JwtKey = []byte("supersecretkey")

func JWTMiddleware(ctx *context.Context) {
	authHeader := ctx.Input.Header("Authorization")
	if authHeader == "" {
		ctx.Output.SetStatus(http.StatusUnauthorized)
		ctx.Output.Body([]byte("Missing Authorization Header"))
		return
	}
	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

	claims := &jwt.RegisteredClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})

	if err != nil || !token.Valid {
		ctx.Output.SetStatus(http.StatusUnauthorized)
		ctx.Output.Body([]byte("Invalid token"))
		return
	}

	ctx.Input.SetData("username", claims.Subject)
}

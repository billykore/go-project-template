package middleware

import (
	"context"

	"github.com/billykore/go-service-tmpl/internal/adapter/http/response"
	"github.com/billykore/go-service-tmpl/internal/domain/user"
	"github.com/billykore/go-service-tmpl/internal/pkg/config"
	"github.com/billykore/go-service-tmpl/internal/pkg/constant"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

// AuthenticateUser returns a middleware function that validates token from headers
// and extract user information.
func AuthenticateUser(cfg *config.Configs) echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		ContextKey:     constant.UserContextKey.String(),
		SigningKey:     []byte(cfg.Token.Secret),
		SuccessHandler: successHandler,
		ErrorHandler:   errorHandler,
	})
}

// successHandler extract user information from token
// and save the information in the request context.
func successHandler(ctx echo.Context) {
	t := ctx.Get(constant.UserContextKey.String()).(*jwt.Token)
	usr := userFromToken(t)
	c := ctx.Request().Context()
	c = ContextWithUser(c, &usr)
	ctx.SetRequest(ctx.Request().WithContext(c))
}

// errorHandler returns an unauthorized response if there is an authentication error.
func errorHandler(ctx echo.Context, err error) error {
	return ctx.JSON(response.Unauthorized(err))
}

// ContextWithUser set user data to the ctx context.
func ContextWithUser(ctx context.Context, user *user.User) context.Context {
	return context.WithValue(ctx, constant.UserContextKey, user)
}

// userFromToken returns user information from JWT token.
func userFromToken(token *jwt.Token) user.User {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return user.User{}
	}
	return user.User{
		CIF:   claims["cif"].(string),
		ID:    claims["userID"].(int),
		Name:  claims["sub"].(string),
		Email: claims["email"].(string),
	}
}

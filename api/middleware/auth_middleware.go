package middleware

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"

	"gitlab.com/soy-app/stock-api/api/schema"

	"go.uber.org/zap"

	"context"
	"errors"

	"gitlab.com/soy-app/stock-api/domain/entity"
	"gitlab.com/soy-app/stock-api/log"
	"gitlab.com/soy-app/stock-api/usecase/interactor"
)

var (
	ErrNoAuthorizationHeader = errors.New("no authorization header passed")
	ErrNotSystemAdmin        = errors.New("you are not system admin")
)

type AuthMiddleware struct {
	userUC interactor.IUserUseCase
}

func NewAuthMiddleware(userUC interactor.IUserUseCase) *AuthMiddleware {
	return &AuthMiddleware{userUC}
}

// Authenticate
// tokenを取得して、認証するmiddlewareの例
func (m *AuthMiddleware) Authenticate(next echo.HandlerFunc) echo.HandlerFunc {
	logger, _ := log.NewLogger()

	return func(c echo.Context) error {
		// Get JWT Token From Header
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, schema.TokenType+" ") {
			logger.Error("Failed to authenticate", zap.Error(ErrNoAuthorizationHeader))
			return echo.NewHTTPError(http.StatusUnauthorized)
		}
		token := strings.TrimPrefix(authHeader, schema.TokenType+" ")

		// Authenticate
		userID, err := m.userUC.Authenticate(token)
		if err != nil {
			logger.Error("Failed to authenticate", zap.Error(err))
			return echo.NewHTTPError(http.StatusUnauthorized)
		}

		// set user detail to context
		user, err := m.userUC.FindByID(userID)
		if err != nil {
			logger.Error("Failed to find me", zap.Error(err))
			return echo.NewHTTPError(http.StatusUnauthorized)
		}
		c = SetToContext(c, user)

		return next(c)
	}
}

// AdminAuth
// 管理者権限を持っているかどうかを確認するmiddlewareの例
func (m *AuthMiddleware) AdminAuth(next echo.HandlerFunc) echo.HandlerFunc {
	logger, _ := log.NewLogger()

	return func(c echo.Context) error {
		ctx := c.Request().Context()
		user, err := GetUserFromContext(ctx)
		if err != nil {
			logger.Error("Failed to get isSystemAdmin from context", zap.Error(err))
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		if !user.IsSystemAdmin() {
			logger.Error("you are not system admin", zap.Error(ErrNotSystemAdmin))
			return echo.NewHTTPError(http.StatusForbidden, ErrNotSystemAdmin.Error())
		}

		return next(c)
	}
}

func SetToContext(c echo.Context, user entity.User) echo.Context {
	ctx := c.Request().Context()
	ctx = SetUserToContext(ctx, user)
	c.SetRequest(c.Request().WithContext(ctx))
	return c
}

type ContextKey string

var (
	userKey ContextKey = "userKey"
)

func SetUserToContext(ctx context.Context, user entity.User) context.Context {
	return context.WithValue(ctx, userKey, user)
}

func GetUserFromContext(ctx context.Context) (entity.User, error) {
	v := ctx.Value(userKey)
	user, ok := v.(entity.User)
	if !ok {
		return entity.User{}, errors.New("no user found in context")
	}
	return user, nil
}

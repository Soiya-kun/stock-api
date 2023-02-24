package handler

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"

	"gitlab.com/soy-app/stock-api/api/middleware"
	"gitlab.com/soy-app/stock-api/api/schema"
	log "gitlab.com/soy-app/stock-api/log"
	"gitlab.com/soy-app/stock-api/usecase/interactor"

	"go.uber.org/zap"

	"gitlab.com/soy-app/stock-api/adapter/authentication"
)

type AuthHandler struct {
	UserUC interactor.IUserUseCase
}

func NewAuthHandler(userUC interactor.IUserUseCase) *AuthHandler {
	return &AuthHandler{UserUC: userUC}
}

// Login POST /auth/access-token
func (h *AuthHandler) Login(c echo.Context) error {
	logger, _ := log.NewLogger()

	req := &schema.LoginReq{}
	if err := c.Bind(req); err != nil {
		logger.Error("Failed to bind request", zap.Error(err))
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	user, token, err := h.UserUC.Login(req.Email, req.Password)
	if err != nil {
		logger.Error("Failed to login", zap.Error(err))
		switch {
		case errors.Is(err, interactor.ErrUserNotFound):
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		case errors.Is(err, authentication.ErrWrongPassword):
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	loginUser := &schema.LoginResUser{
		UserId: user.UserID,
		Email:  user.Email,
	}

	return c.JSON(http.StatusOK, &schema.LoginRes{
		AccessToken: token,
		TokenType:   schema.TokenType,
		User:        *loginUser,
	})
}

// ResetPassword POST /auth/reset-password
func (h *AuthHandler) ResetPassword(c echo.Context) error {
	logger, _ := log.NewLogger()

	req := &schema.ResetPasswordReq{}
	if err := c.Bind(req); err != nil {
		logger.Error("Failed to bind request", zap.Error(err))
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	err := h.UserUC.SendResetPasswordMail(req.Email)
	if err != nil {
		logger.Error("Failed to reset password", zap.Error(err))
		switch {
		case errors.Is(err, interactor.ErrUserNotFound):
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	return c.NoContent(http.StatusOK)
}

// ChangePassword POST /auth/change-password
func (h *AuthHandler) ChangePassword(c echo.Context) error {
	logger, _ := log.NewLogger()

	req := &schema.UpdatePasswordReq{}
	if err := c.Bind(req); err != nil {
		logger.Error("Failed to bind request", zap.Error(err))
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	ctx := c.Request().Context()
	user, err := middleware.GetUserFromContext(ctx)
	if err != nil {
		logger.Error("Failed to get user from context", zap.Error(err))
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	err = h.UserUC.UpdatePassword(interactor.UserUpdatePassword{
		UserID:      user.UserID,
		NewPassword: req.Password,
	})
	if err != nil {
		logger.Error("Failed to change password", zap.Error(err))
		switch {
		case errors.Is(err, interactor.ErrUserNotFound):
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	return c.NoContent(http.StatusOK)
}

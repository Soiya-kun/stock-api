package handler

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"

	"go.uber.org/zap"

	"gitlab.com/soy-app/stock-api/api/middleware"
	"gitlab.com/soy-app/stock-api/api/schema"
	"gitlab.com/soy-app/stock-api/log"
	"gitlab.com/soy-app/stock-api/usecase/interactor"
)

type UserHandler struct {
	UserUC interactor.IUserUseCase
}

func NewUserHandler(userUC interactor.IUserUseCase) *UserHandler {
	return &UserHandler{UserUC: userUC}
}

func (h *UserHandler) FindMe(c echo.Context) error {
	logger, _ := log.NewLogger()

	ctx := c.Request().Context()
	user, err := middleware.GetUserFromContext(ctx) // トークンからuserIDを取得
	if err != nil {
		logger.Error("Failed to get id from context", zap.Error(err))
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	res, err := h.UserUC.FindByID(user.UserID)
	if err != nil {
		logger.Error("Failed to find me", zap.Error(err))
		switch {
		case errors.Is(err, interactor.ErrUserNotFound):
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	return c.JSON(http.StatusOK, schema.UserResFromEntity(res))
}

func (h *UserHandler) FindById(c echo.Context) error {
	logger, _ := log.NewLogger()

	var id string
	if err := echo.PathParamsBinder(c).MustString("user-id", &id).BindError(); err != nil {
		logger.Error("Failed to bind path param id", zap.Error(err))
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	res, err := h.UserUC.FindByID(id)
	if err != nil {
		logger.Error("Failed to find me", zap.Error(err))
		switch {
		case errors.Is(err, interactor.ErrUserNotFound):
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	return c.JSON(http.StatusOK, schema.UserResFromEntity(res))
}

func (h *UserHandler) FindByIdForAdmin(c echo.Context) error {
	logger, _ := log.NewLogger()

	var id string
	if err := echo.PathParamsBinder(c).MustString("user-id", &id).BindError(); err != nil {
		logger.Error("Failed to bind path param id", zap.Error(err))
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	res, err := h.UserUC.FindByIDByAdmin(id)
	if err != nil {
		logger.Error("Failed to find me by admin", zap.Error(err))
		switch {
		case errors.Is(err, interactor.ErrUserNotFound):
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	return c.JSON(http.StatusOK, schema.UserResFromEntity(res))
}

func (h *UserHandler) Create(c echo.Context) error {
	logger, _ := log.NewLogger()

	req := &schema.CreateUserReq{}
	if err := c.Bind(req); err != nil {
		logger.Error("Failed to bind request", zap.Error(err))
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	res, err := h.UserUC.Create(interactor.UserCreate{
		Email:    req.Email,
		UserType: req.UserType,
		Password: req.Password,
		Name:     req.Name,
	})
	if err != nil {
		logger.Error("Failed to create user", zap.Error(err))
		switch {
		case errors.Is(err, interactor.ErrUserAlreadyExists):
			return echo.NewHTTPError(http.StatusConflict, err.Error())
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	return c.JSON(http.StatusCreated, schema.UserResFromEntity(res))
}

func (h *UserHandler) Update(c echo.Context) error {
	logger, _ := log.NewLogger()

	req := &schema.UpdateUserReq{}
	if err := c.Bind(req); err != nil {
		logger.Error("Failed to bind request", zap.Error(err))
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	res, err := h.UserUC.Update(interactor.UserUpdate{
		UserID:   req.UserID,
		UserType: req.UserType,
		Name:     req.Name,
	})
	if err != nil {
		logger.Error("Failed to update user", zap.Error(err))
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, schema.UserResFromEntity(res))
}

func (h *UserHandler) Search(c echo.Context) error {
	logger, _ := log.NewLogger()

	req := &schema.UserSearchQueryReq{}
	if err := c.Bind(req); err != nil {
		logger.Error("Failed to bind request", zap.Error(err))
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	list, total, err := h.UserUC.Search(
		req.Query,
		req.UserType,
		req.Skip,
		req.Limit,
	)
	if err != nil {
		logger.Error("Failed by invalid request", zap.Error(err))
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, schema.UsersResFromSearchResult(list, total))
}

func (h *UserHandler) Delete(c echo.Context) error {
	logger, _ := log.NewLogger()

	var id string
	if err := echo.PathParamsBinder(c).MustString("user-id", &id).BindError(); err != nil {
		logger.Error("Failed to bind path param id", zap.Error(err))
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	res, err := h.UserUC.Delete(id)
	if err != nil {
		logger.Error("Failed to delete", zap.Error(err))
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, schema.UserResFromEntity(res))
}

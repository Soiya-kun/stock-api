package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"

	"gitlab.com/soy-app/stock-api/api/middleware"

	"gitlab.com/soy-app/stock-api/api/schema"
	"gitlab.com/soy-app/stock-api/log"
	"gitlab.com/soy-app/stock-api/usecase/interactor"
)

type StockHandler struct {
	StockUseCase interactor.IStockUseCase
}

func NewStockHandler(stockUseCase interactor.IStockUseCase) *StockHandler {
	return &StockHandler{StockUseCase: stockUseCase}
}

// Create creates stocks
func (h *StockHandler) Create(c echo.Context) error {
	logger, _ := log.NewLogger()

	req := &schema.StockCreateListReq{}
	if err := c.Bind(req); err != nil {
		logger.Error("Failed to bind request", zap.Error(err))
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	res, err := h.StockUseCase.CreateStocks(req.StockCreate())
	if err != nil {
		logger.Error("Failed to create stock", zap.Error(err))
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, schema.StocksResFromEntity(&res))
}

func (h *StockHandler) FindBySC(c echo.Context) error {
	logger, _ := log.NewLogger()

	var sc string
	if err := echo.PathParamsBinder(c).MustString("sc", &sc).BindError(); err != nil {
		logger.Error("Failed to bind path param id", zap.Error(err))
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	res, err := h.StockUseCase.FindByStockCode(sc)
	if err != nil {
		logger.Error("Failed to find stock", zap.Error(err))
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, schema.StocksResFromEntity(&res))
}

func (h *StockHandler) FindByRandom(c echo.Context) error {
	logger, _ := log.NewLogger()

	res, err := h.StockUseCase.FindByRandom()
	if err != nil {
		logger.Error("Failed to find stock", zap.Error(err))
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, schema.StocksResFromEntity(&res))
}

func (h *StockHandler) SaveSC(c echo.Context) error {
	logger, _ := log.NewLogger()

	req := &schema.SaveSCReq{}
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

	err = h.StockUseCase.SaveStockCode(req.StockCode, user)
	if err != nil {
		logger.Error("Failed to save stock", zap.Error(err))
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, nil)
}

func (h *StockHandler) ListSC(c echo.Context) error {
	logger, _ := log.NewLogger()

	ret, err := h.StockUseCase.ListSC()
	if err != nil {
		logger.Error("Failed to delete stock", zap.Error(err))
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, schema.StockCodeListRes{
		StockCodes: ret,
	})
}

func (h *StockHandler) ListSCByThreshold(c echo.Context) error {
	logger, _ := log.NewLogger()

	req := &schema.StockCodeByThresholdReq{}
	if err := c.Bind(req); err != nil {
		logger.Error("Failed to bind request", zap.Error(err))
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	ret, err := h.StockUseCase.ListSCByThreshold(req.MinTradeValue, req.Date)
	if err != nil {
		logger.Error("Failed to delete stock", zap.Error(err))
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, schema.StockCodeListRes{
		StockCodes: ret,
	})

}

func (h *StockHandler) ListSCSavedByUser(c echo.Context) error {
	logger, _ := log.NewLogger()

	ctx := c.Request().Context()
	user, err := middleware.GetUserFromContext(ctx)
	if err != nil {
		logger.Error("Failed to get user from context", zap.Error(err))
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	ret, err := h.StockUseCase.ListSCSavedByUser(user)
	if err != nil {
		logger.Error("Failed to delete stock", zap.Error(err))
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, schema.StockCodeListRes{
		StockCodes: ret,
	})
}

func (h *StockHandler) CreateSplit(c echo.Context) error {
	logger, _ := log.NewLogger()

	req := &schema.StockSplitReq{}
	if err := c.Bind(req); err != nil {
		logger.Error("Failed to bind request", zap.Error(err))
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	err := h.StockUseCase.CreateStockSplit(interactor.StockSplitCreate{
		StockCode:  req.StockCode,
		Date:       req.Date,
		SplitRatio: req.SplitRatio,
	})
	if err != nil {
		logger.Error("Failed to create stock split", zap.Error(err))
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, nil)
}

func (h *StockHandler) SaveSearchStockPattern(c echo.Context) error {
	logger, _ := log.NewLogger()

	req := &schema.SaveSearchConditionReq{}
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

	err = h.StockUseCase.SaveSearchCondition(req.UsecaseArg(), user)
	if err != nil {
		logger.Error("Failed to save search condition", zap.Error(err))
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, nil)
}

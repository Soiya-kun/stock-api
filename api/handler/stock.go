package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"

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

	return c.JSON(http.StatusCreated, schema.StockListResFromEntity(res))
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

	return c.JSON(http.StatusOK, schema.StocksResFromEntity(res))
}

func (h *StockHandler) FindByRandom(c echo.Context) error {
	logger, _ := log.NewLogger()

	res, err := h.StockUseCase.FindByRandom()
	if err != nil {
		logger.Error("Failed to find stock", zap.Error(err))
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, schema.StocksResFromEntity(res))
}

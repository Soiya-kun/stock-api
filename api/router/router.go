package router

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/labstack/echo/v4/middleware"

	"gitlab.com/soy-app/stock-api/api/handler"
	apiMiddleware "gitlab.com/soy-app/stock-api/api/middleware"
	"gitlab.com/soy-app/stock-api/api/proxy"
	"gitlab.com/soy-app/stock-api/usecase/interactor"
)

func NewServer(
	userUC interactor.IUserUseCase,
	stockUC interactor.IStockUseCase,
) *echo.Echo {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete, http.MethodPatch},
	}))

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Pre(middleware.RemoveTrailingSlash())

	authHandler := handler.NewAuthHandler(userUC)
	userHandler := handler.NewUserHandler(userUC)
	stockHandler := handler.NewStockHandler(stockUC)
	postCodeJPProxyHandler := proxy.NewPostCodeJPProxyHandler()

	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	api := e.Group("/api")
	api.POST("/auth/access-token", authHandler.Login)
	api.POST("/auth/reset-password", authHandler.ResetPassword)
	api.POST("/auth/change-password", authHandler.ChangePassword)

	// auth
	// 認可の例
	auth := api.Group("", apiMiddleware.NewAuthMiddleware(userUC).Authenticate)

	// system admin
	// 管理者権限が必要な場合の認可の例
	admin := auth.Group("/system-admin", apiMiddleware.NewAuthMiddleware(userUC).AdminAuth)

	// user for admin
	userForAdmin := admin.Group("/user")
	userForAdmin.GET("/:user-id", userHandler.FindByIdForAdmin)

	// user
	user := auth.Group("/user")
	user.GET("", userHandler.Search)
	user.POST("", userHandler.Create)
	user.GET("/me", userHandler.FindMe)
	user.GET("/:user-id", userHandler.FindById)
	user.PATCH("/:user-id", userHandler.Update)
	user.DELETE("/:user-id", userHandler.Delete)

	// stock
	stockApi := api.Group("/stocks")
	stockApi.GET("/random", stockHandler.FindByRandom)
	stockApi.GET("/stock-codes", stockHandler.ListSC)

	stockAuth := auth.Group("/stocks")
	stockAuth.POST("", stockHandler.Create)
	stockAuth.POST("/stock-splits", stockHandler.CreateSplit)
	stockAuth.GET("/:sc", stockHandler.FindBySC)
	stockAuth.POST("/stock-codes", stockHandler.SaveSC)
	stockAuth.POST("/save-search-stock-patterns", stockHandler.SaveSearchStockPattern)

	// post code jp proxy
	postCodeJP := auth.Group("/address")
	postCodeJP.GET("/:post-code", postCodeJPProxyHandler.SearchAboutPostCode)

	return e
}

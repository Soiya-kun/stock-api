package main

import (
	"fmt"
	"os"
	"time"

	"go.uber.org/zap"

	"gitlab.com/soy-app/stock-api/adapter/database"
	"gitlab.com/soy-app/stock-api/adapter/ulid"
	"gitlab.com/soy-app/stock-api/interface/repository"
	"gitlab.com/soy-app/stock-api/log"
	"gitlab.com/soy-app/stock-api/usecase/interactor"
)

func main() {
	logger, err := log.NewLogger()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Failed to create logger : %v\n", err)
		os.Exit(1)
	}

	db, err := database.NewDB(logger)
	if err != nil {
		logger.Error("Failed to connect to database", zap.Error(err))
		return
	}

	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			logger.Error("Failed to get sql.DB", zap.Error(err))
		}
		err = sqlDB.Close()
		if err != nil {
			logger.Error("Failed to close database connection", zap.Error(err))
		}
	}()
	ulidDriver := ulid.NewULID()
	stockRepo := repository.NewStockRepository(db)
	searchStockPatternRepo := repository.NewSearchStockPatternRepository(db)
	searchedStockRepo := repository.NewSearchedStockPatternRepository(db)
	stockUC := interactor.NewStockUseCase(ulidDriver, stockRepo, searchStockPatternRepo, searchedStockRepo)

	ret, err := stockUC.SearchByCondition(interactor.SearchReq{
		SearchPatternID: "01H735CS045WM440G2KZKSDFFM",
		EndDate:         time.Date(2023, 8, 4, 0, 0, 0, 0, time.UTC),
	})

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ret)
}

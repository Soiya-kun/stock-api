package main

import (
	"gitlab.com/soy-app/stock-api/adapter/database"
	"gitlab.com/soy-app/stock-api/adapter/file"
	"gitlab.com/soy-app/stock-api/interface/repository"
	"gitlab.com/soy-app/stock-api/log"
	"go.uber.org/zap"
	"os"
)

// コマンドライン引数にstockデータのcsvファイルを格納したパスを取る
func main() {
	logger, err := log.NewLogger()

	db, err := database.NewDB(logger)
	if err != nil {
		logger.Error("Failed to connect to database", zap.Error(err))
		return
	}
	stockRepo := repository.NewStockRepository(db)

	fileDriver := file.NewFileDriverOnLocal()
	csvFilePath, err := fileDriver.GetCSVPath()
	if err != nil {
		return
	}

	for i, path := range csvFilePath {
		stockCSV, err := fileDriver.GetCSVFileReader(path)
		if err != nil {
			logger.Error("Failed to get csv file reader", zap.Error(err))
			return
		}

		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				return
			}
		}(stockCSV.File)

		stocks, err := stockRepo.ReadCSV(stockCSV.Reader)
		if err != nil {
			logger.Error("Failed to read csv", zap.Error(err))
			return
		}

		err = stockRepo.Create(stocks)
		if err != nil {
			logger.Error("Failed to create stocks", zap.Error(err))
			return
		}
		if i == 3 {
			break
		}
	}
}

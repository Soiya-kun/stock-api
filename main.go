package main

import (
	"fmt"
	"os"

	"gitlab.com/soy-app/stock-api/adapter/email"

	"gitlab.com/soy-app/stock-api/adapter/aws"

	"go.uber.org/zap"

	"gitlab.com/soy-app/stock-api/adapter/authentication"
	"gitlab.com/soy-app/stock-api/adapter/database"
	"gitlab.com/soy-app/stock-api/adapter/ulid"
	"gitlab.com/soy-app/stock-api/api/router"
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

	err = database.Migrate(db)
	if err != nil {
		logger.Error("Failed to migrate", zap.Error(err))
		return
	}

	awsCli := aws.NewCli()
	mailDriver := email.NewEmailDriver(awsCli)
	ulidDriver := ulid.NewULID()

	userRepo := repository.NewUserRepository(db, ulidDriver)
	userAuth := authentication.NewUserAuth()
	userUC := interactor.NewUserUseCase(mailDriver, ulidDriver, userAuth, userRepo)

	stockRepo := repository.NewStockRepository(db)
	stockUC := interactor.NewStockUseCase(ulidDriver, stockRepo)

	s := router.NewServer(
		userUC,
		stockUC,
	)

	if err := s.Start(":80"); err != nil {
		logger.Error("Failed to start server", zap.Error(err))
		os.Exit(1)
	}
}

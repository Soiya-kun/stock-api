package main

import (
	"fmt"

	"gitlab.com/soy-app/stock-api/domain/entity"

	"gitlab.com/soy-app/stock-api/adapter/authentication"
	"gitlab.com/soy-app/stock-api/adapter/database"
	"gitlab.com/soy-app/stock-api/log"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func connectDB() (*gorm.DB, error) {
	logger, err := log.NewLogger()
	if err != nil {
		return nil, err
	}

	db, err := database.NewDB(logger)
	if err != nil {
		logger.Error("Failed to connect to database", zap.Error(err))
		return nil, err
	}
	return db, nil
}

func DropDB(db *gorm.DB) {
	_ = db.Migrator().DropTable(&entity.User{})
}

func CreateDevSeedData(db *gorm.DB) error { //nolint
	// system-adminのみをまず生成する
	hp, _ := authentication.HashBcryptPassword("pass")
	u := &entity.User{
		UserID:         "system-admin",
		Email:          "admin@test.com",
		HashedPassword: hp,
		IsDeleted:      false,
	}
	if err := db.Create(&u).Error; err != nil {
		return err
	}

	return nil
}

func main() {
	db, err := connectDB()
	if err != nil {
		fmt.Println("error:", err)
	}

	DropDB(db)

	err = database.Migrate(db)
	if err != nil {
		fmt.Println("error:", err)
	}

	// 開発用テストデータ作成
	err = CreateDevSeedData(db)
	if err != nil {
		fmt.Println("error:", err)
	}
}

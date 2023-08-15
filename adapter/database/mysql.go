package database

import (
	"fmt"
	"math"
	"time"

	"gitlab.com/soy-app/stock-api/domain/entity"

	"gitlab.com/soy-app/stock-api/config"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB(logger *zap.Logger) (*gorm.DB, error) {
	dsn := config.DSN()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to open MySQL: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get sql.DB: %w", err)
	}

	sqlDB.SetConnMaxIdleTime(100)
	sqlDB.SetMaxOpenConns(100)

	// Check connection
	const retryMax = 10
	for i := 0; i < retryMax; i++ {
		err = sqlDB.Ping()
		if err == nil {
			break
		}
		if i == retryMax-1 {
			return nil, fmt.Errorf("failed to connect to database: %w", err)
		}
		duration := time.Millisecond * time.Duration(math.Pow(1.5, float64(i))*1000)
		logger.Warn("failed to connect to database retrying", zap.Error(err), zap.Duration("sleepSeconds", duration))
		time.Sleep(duration)
	}

	return db, nil
}

func Migrate(db *gorm.DB) error {
	if err := db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&entity.User{}); err != nil {
		return err
	}
	if err := db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&entity.Stock{}); err != nil {
		return err
	}
	if err := db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&entity.SavedStockCode{}); err != nil {
		return err
	}
	if err := db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&entity.StockSplit{}); err != nil {
		return err
	}
	if err := db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&entity.SearchStockPattern{}); err != nil {
		return err
	}
	if err := db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&entity.VolumePattern{}); err != nil {
		return err
	}
	if err := db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&entity.PricePattern{}); err != nil {
		return err
	}
	if err := db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&entity.MaXUpDownPattern{}); err != nil {
		return err
	}
	if err := db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&entity.SearchedStockPatternCode{}); err != nil {
		return err
	}
	return nil
}

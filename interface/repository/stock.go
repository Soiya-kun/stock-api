package repository

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm"

	"gitlab.com/soy-app/stock-api/domain/entity"
	"gitlab.com/soy-app/stock-api/usecase/port"
)

type StockRepository struct {
	db *gorm.DB
}

func NewStockRepository(db *gorm.DB) port.StockRepository {
	return &StockRepository{db: db}
}

func (r *StockRepository) Create(stocks []entity.Stock) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		for _, stock := range stocks {
			err := tx.Create(&stock).Error
			if err != nil {
				fmt.Println("failed to create stocks: %w", err)
				return port.ErrCreateStock
			}
		}
		return nil
	})
}

func (r *StockRepository) FindByStockCode(s string) ([]entity.Stock, error) {
	var stocks []entity.Stock
	err := r.db.Where("stock_code = ?", s).Find(&stocks).Error
	if err != nil {
		return nil, err
	}
	return stocks, nil
}

func stringToFloatPointer(s string) *float64 {
	if strings.TrimSpace(s) == "-" {
		return nil
	}
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return nil
	}
	return &f
}

func (r *StockRepository) ReadCSV(reader *csv.Reader) ([]entity.Stock, error) {
	var stocks []entity.Stock
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		// 1行目をスルー
		if record[0] == "SC" {
			continue
		}

		date, err := time.Parse("20060102", record[4])
		if err != nil {
			return nil, err
		}

		stocks = append(stocks, entity.Stock{
			StockCode:     record[0],
			StockName:     record[1],
			Market:        record[2],
			Industry:      record[3],
			Date:          date,
			Price:         stringToFloatPointer(record[5]),
			Change:        stringToFloatPointer(record[6]),
			ChangePercent: stringToFloatPointer(record[7]),
			PreviousClose: stringToFloatPointer(record[8]),
			OpenedPrice:   stringToFloatPointer(record[9]),
			High:          stringToFloatPointer(record[10]),
			Low:           stringToFloatPointer(record[11]),
			Volume:        stringToFloatPointer(record[12]),
			TradingValue:  stringToFloatPointer(record[13]),
			MarketCap:     stringToFloatPointer(record[14]),
			LowerLimit:    stringToFloatPointer(record[15]),
			UpperLimit:    stringToFloatPointer(record[16]),
		})
	}

	return stocks, nil
}

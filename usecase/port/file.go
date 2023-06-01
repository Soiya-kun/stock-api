package port

import (
	"encoding/csv"
	"os"
)

type StockCSV struct {
	File   *os.File
	Reader *csv.Reader
}

type FileDriver interface {
	GetCSVPath() ([]string, error)
	GetCSVFileReader(path string) (StockCSV, error)
}

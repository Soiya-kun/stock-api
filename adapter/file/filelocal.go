package file

import (
	"bufio"
	"encoding/csv"
	"os"
	"path/filepath"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"

	"gitlab.com/soy-app/stock-api/usecase/port"
)

type FileLocal struct{}

func NewFileDriverOnLocal() port.FileDriver {
	return &FileLocal{}
}

func (f FileLocal) getCSVDir() string {
	return "/app/stocks/"
}

func (f FileLocal) GetCSVPath() ([]string, error) {
	files, err := os.ReadDir(f.getCSVDir())
	if err != nil {
		return nil, err
	}

	var paths []string
	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".csv" {
			paths = append(paths, filepath.Join(f.getCSVDir(), file.Name()))
		}
	}

	return paths, nil
}

func (f FileLocal) GetCSVFileReader(path string) (port.StockCSV, error) {
	file, err := os.Open(path)
	if err != nil {
		return port.StockCSV{}, err
	}

	reader := csv.NewReader(transform.NewReader(bufio.NewReader(file), japanese.ShiftJIS.NewDecoder()))
	reader.TrimLeadingSpace = true

	return port.StockCSV{
		File:   file,
		Reader: reader,
	}, nil
}

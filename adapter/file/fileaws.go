package file

import (
	"gitlab.com/soy-app/stock-api/usecase/port"

	awsDriver "gitlab.com/soy-app/stock-api/adapter/aws"
)

type File struct {
	awsCli *awsDriver.Cli
}

func NewFileDriver(awsCli *awsDriver.Cli) port.FileDriver {
	return &File{
		awsCli: awsCli,
	}
}

func (f File) GetCSVPath() ([]string, error) {
	//TODO implement me
	panic("implement me")
}

func (f File) GetCSVFileReader(path string) (port.StockCSV, error) {
	//TODO implement me
	panic("implement me")
}

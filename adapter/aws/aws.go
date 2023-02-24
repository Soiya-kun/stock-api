package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"

	"gitlab.com/soy-app/stock-api/config"
)

type Cli struct {
	awsAccessKey  string
	awsPrivateKey string
	awsRegion     string
}

func NewCli() *Cli {
	return &Cli{
		awsAccessKey:  config.AWSAccessKey(),
		awsPrivateKey: config.AWSPrivateKey(),
		awsRegion:     config.AWSRegion(),
	}
}

func (cli *Cli) CreateSession() (*session.Session, error) {
	creds := credentials.NewStaticCredentials(cli.awsAccessKey, cli.awsPrivateKey, "")
	sess, err := session.NewSession(&aws.Config{
		Credentials: creds,
		Region:      aws.String(cli.awsRegion),
	})
	if err != nil {
		return nil, err
	}

	return sess, nil
}

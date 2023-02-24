package config

import (
	"log"
	"os"
)

var (
	awsAccessKey    string
	awsPrivateKey   string
	awsRegion       string
	env             string
	frontendURL     string
	postCodeJPToken string
	s3Bucket        string
	sesSender       string
	sigKey          string // JWTトークンの署名
)

func init() {
	sigKey = "XXX"
	env = os.Getenv("ENV")
	if env == "" {
		log.Fatal("ENV environment variable is empty")
	}
	awsAccessKey = os.Getenv("AWS_ACCESS_KEY")
	if awsAccessKey == "" {
		log.Fatal("AWS_ACCESS_KEY environment variable is empty")
	}
	awsPrivateKey = os.Getenv("AWS_PRIVATE_KEY")
	if awsPrivateKey == "" {
		log.Fatal("AWS_PRIVATE_KEY environment variable is empty")
	}
	awsRegion = os.Getenv("AWS_REGION")
	if awsRegion == "" {
		log.Fatal("AWS_REGION environment variable is empty")
	}
	s3Bucket = os.Getenv("S3_BUCKET")
	if s3Bucket == "" {
		log.Fatal("S3_BUCKET environment variable is empty")
	}
	sesSender = os.Getenv("SES_SENDER")
	if sesSender == "" {
		log.Fatal("SES_SENDER environment variable is empty")
	}
	postCodeJPToken = os.Getenv("POST_CODE_JP_TOKEN")
	if postCodeJPToken == "" {
		log.Fatal("POST_CODE_JP_TOKEN environment variable is empty")
	}
	frontendURL = os.Getenv("FRONTEND_URL")
	if frontendURL == "" {
		log.Fatal("FRONTEND_URL environment variable is empty")
	}
}

func IsDevelopment() bool {
	return env == "development"
}

func IsTest() bool {
	return env == "test"
}

func IsGitLabCI() bool {
	return env == "gitlab-ci"
}

func AWSAccessKey() string {
	return awsAccessKey
}

func AWSPrivateKey() string {
	return awsPrivateKey
}

func AWSRegion() string {
	return awsRegion
}

func SigKey() string {
	return sigKey
}

func SESSender() string {
	return sesSender
}

func S3Bucket() string {
	return s3Bucket
}

func PostCodeJPToken() string {
	return postCodeJPToken
}

func FrontendURL() string {
	return frontendURL
}

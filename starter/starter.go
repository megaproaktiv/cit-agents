package starter

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

var ec2Client EC2CreateInstanceAPI



func init() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error, " + err.Error())
	}
	ec2Client = ec2.NewFromConfig(cfg)

	Logger = &logrus.Logger{
		Out:   os.Stderr,
		Level: logrus.TraceLevel,
		Formatter: &logrus.TextFormatter{
			DisableColors:   true,
			TimestampFormat: "2006-01-02 15:04:05",
			FullTimestamp:   true,
		},
	}
	Logger.Info("Init")
}


func getAccount() *string {
	var account = "795048271754"
	return &account
}

func getCDKBucket() *string {
	var bucket = "cdk-hnb659fds-assets-795048271754-eu-central-1"
	return &bucket
}
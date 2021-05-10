package main

import (
	"github.com/goochi/configs"
	"github.com/sirupsen/logrus"
)

func main() {
	c := configs.NewConfigProvider(logrus.New(), false, "./configs/.test.env", "./configs/.env")

	logrus.New().Infof("NAME: %v", c.Get("NAME"))
	logrus.New().Infof("DB_NAME: %v", c.ExpandEnv("DB_NAME"))
}

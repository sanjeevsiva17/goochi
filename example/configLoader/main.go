package main

import (
	"github.com/goochi/configs"
	"github.com/goochi/log"
)

func main() {
	logger := log.NewLogger(log.Info)
	c := configs.NewConfigProvider(false, "./configs/.test.env", "./configs/.env")

	logger.Infof("NAME: %v", c.Get("NAME"))
	logger.Infof("DB_NAME: %v", c.ExpandEnv("DB_NAME"))
}

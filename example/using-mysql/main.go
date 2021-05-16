package main

import (
	"github.com/goochi/datastores"
	"github.com/goochi/log"
)

func main() {
	l := log.NewLogger(log.Info)
	_, err := datastores.NewMySQL(&datastores.Config{
		HostName: "localhost",
		Username: "root",
		Password: "password",
		Database: "mysql",
		Port:     2001,
	})
	if err != nil {
		l.Fatal(err)
	}

	l.Info("Connected successfully")
}

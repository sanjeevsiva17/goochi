package main

import (
	"github.com/goochi/datastores"
	"github.com/goochi/log"
)

func main() {
	l := log.NewLogger(log.Info)

	_, err := datastores.NewSQL(&datastores.Config{
		HostName: "localhost",
		Username: "postgres",
		Password: "root123",
		Database: "postgres",
		Port:     "2005",
		Dialect:  "postgres",
	})

	if err != nil {
		l.Fatal(err)
	}

	l.Info("Connected successfully")
}

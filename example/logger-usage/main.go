package main

import "github.com/goochi/log"

func main() {
	logger := log.NewLogger(log.Debug)

	logger.Info(Student{
		Name: "goochi",
		Age:  1,
	})

	logger.Debug("Debug")
	logger.Info("Info")
	logger.Warn("Warn")
	logger.Error("error")
}

type Student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

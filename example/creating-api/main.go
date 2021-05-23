package main

import (
	"net/http"

	"github.com/goochi"
	"github.com/goochi/log"
)

func main() {
	logger := log.NewLogger(log.Info)
	r := goochi.NewRouter()

	r.GET("/name", func(r *http.Request) (statusCode int, data map[string]interface{}) {
		return 200, map[string]interface{}{
			"name": "goochi",
		}
	})

	server := &goochi.HTTPServer{
		Router: r,
		Port:   7000,
	}

	server.Start(logger)
}

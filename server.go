package goochi

import (
	"log"
	"net/http"
	"strconv"
)

type server interface {
	Start(logger *log.Logger)
}

type HTTPServer struct {
	Router *Router
	Port   int
}

func (h *HTTPServer) Start(logger *log.Logger) {
	addr := ":" + strconv.Itoa(h.Port)
	logger.Printf("starting http server at %s", addr)
	err := http.ListenAndServe(addr, h.Router)
	if err != nil {
		logger.Printf("error in starting http server at %s: %s", addr, err)
	}
}

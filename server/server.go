package server

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type server struct {
	httpServer *http.Server
}

func NewServer(port string, handler http.Handler) *server {
	return &server{
		httpServer: &http.Server{
			Addr:         fmt.Sprintf(":%s", port),
			Handler:      handler,
			IdleTimeout:  time.Minute,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 30 * time.Second,
		},
	}
}

func (s *server) Run() {
	log.Printf("Starting server on %s", s.httpServer.Addr)

	if err := s.httpServer.ListenAndServe(); err != nil {
		log.Println("Error with server running: " + err.Error())
	}
}

package internals

import (
	"fmt"
	"net/http"
	"time"
)

type Server struct {
	host string
	port string
}

func NewServer() *http.Server {
	NewServer := &Server{
		host: "dev.skat-service.ru",
		port: "8080",
	}

	server := &http.Server{
		Addr:         fmt.Sprintf("%s:%s", NewServer.host, NewServer.port),
		Handler:      Logging(NewServer.RegisterRoutes()),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}

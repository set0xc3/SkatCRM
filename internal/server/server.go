package server

import (
	"fmt"
	"net/http"
	"time"
)

type Server struct {
	host string
	port string
}

func New() *http.Server {
	ctx := &Server{
		host: "localhost",
		port: "8080",
	}

	ret := &http.Server{
		Addr:         fmt.Sprintf("%s:%s", ctx.host, ctx.port),
		Handler:      Logging(ctx.RegisterRoutes()),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return ret
}

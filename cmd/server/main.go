package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/set0xc3/htmx/internal"
	"github.com/set0xc3/htmx/internal/server"
	"github.com/set0xc3/htmx/internal/server/handlers"
)

func main() {
	s := server.New()

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.Static("/", "static")
	e.GET("/", handlers.GetIndexPage)

	// Start server
	internal.GracefulShutdown(e, s.Port)
}

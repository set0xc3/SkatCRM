package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/set0xc3/htmx/internal"
	"github.com/set0xc3/htmx/internal/server"
	"github.com/set0xc3/htmx/internal/server/handlers"
	"github.com/set0xc3/htmx/internal/views/cms/pages"
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
	e.GET("/clients", handlers.GetClientsPage)

	e.GET("/cms", func(c echo.Context) error {
		return handlers.Render(c, cms.IndexPage())
	})

	// API
	e.GET("/api/v1/clients", handlers.RedirectToDB)
	e.DELETE("/api/v1/client/:id", handlers.StatusOK)

	// Start server
	internal.GracefulShutdown(e, s.Port)
}

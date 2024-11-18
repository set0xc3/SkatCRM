package main

import (
	"math"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/set0xc3/htmx/internal"
	"github.com/set0xc3/htmx/internal/api"
	"github.com/set0xc3/htmx/internal/frontend"
	"github.com/set0xc3/htmx/internal/frontend/views"
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

	// Pages
	e.GET("/", func(c echo.Context) error {
		return handlers.Render(c, frontend.Layout(views.Home()))
	})
	e.GET("/test", func(c echo.Context) error {
		return handlers.Render(c, frontend.Layout(views.Test()))
	})
	e.GET("/clients", func(c echo.Context) error {
		page, _ := strconv.Atoi(c.QueryParam("page"))
		pageMax := int(float64(api.FetchClientCount() / 10))

		if page <= 0 {
			page = 1
		}

		page = int(math.Min(float64(page), float64(pageMax)))

		return handlers.Render(c, frontend.Layout(views.Clients(page)))
	})

	// Views
	e.GET("/views/", func(c echo.Context) error {
		return handlers.Render(c, views.Home())
	})
	e.GET("/views/clients", func(c echo.Context) error {
		page, _ := strconv.Atoi(c.QueryParam("page"))
		pageMax := int(float64(api.FetchClientCount() / 10))

		if page <= 0 {
			page = 1
		}

		page = int(math.Min(float64(page), float64(pageMax)))

		return handlers.Render(c, views.Clients(page))
	})

	// Misc
	e.GET("/status/ok", handlers.StatusOK)

	// API
	// e.GET("/api/v1/clients", handlers.RedirectToDB)
	// e.GET("/api/v1/clients/:count", handlers.RedirectToDB)
	// e.DELETE("/api/v1/client/:id", handlers.StatusOK)

	// e.GET("/api/v1/products", handlers.RedirectToDB)
	// e.DELETE("/api/v1/product/:id", handlers.StatusOK)

	// Start server
	internal.GracefulShutdown(e, s.Port)
}

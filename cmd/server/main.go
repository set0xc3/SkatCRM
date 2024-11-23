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

func getValidPageNumber(page int) int {
	pageLimit := 10
	if page <= 0 {
		page = 1
	}

	pageMax := int(math.Ceil(float64(api.FetchClientCount()) / float64(pageLimit)))
	page = int(math.Min(float64(page), float64(pageMax)))
	return page
}

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
		page = getValidPageNumber(page)
		return handlers.Render(c, frontend.Layout(views.Clients(page)))
	})
	e.GET("/orders", func(c echo.Context) error {
		page, _ := strconv.Atoi(c.QueryParam("page"))
		page = getValidPageNumber(page)
		return handlers.Render(c, frontend.Layout(views.Orders(page)))
	})
	e.GET("/calls", func(c echo.Context) error {
		page, _ := strconv.Atoi(c.QueryParam("page"))
		page = getValidPageNumber(page)
		return handlers.Render(c, frontend.Layout(views.Calls(page)))
	})

	// Views
	e.GET("/views/", func(c echo.Context) error {
		return handlers.Render(c, views.Home())
	})
	e.GET("/views/clients", func(c echo.Context) error {
		page, _ := strconv.Atoi(c.QueryParam("page"))
		page = getValidPageNumber(page)
		return handlers.Render(c, views.Clients(page))
	})
	e.GET("/views/orders", func(c echo.Context) error {
		page, _ := strconv.Atoi(c.QueryParam("page"))
		page = getValidPageNumber(page)
		return handlers.Render(c, views.Orders(page))
	})
	e.GET("/views/calls", func(c echo.Context) error {
		page, _ := strconv.Atoi(c.QueryParam("page"))
		page = getValidPageNumber(page)
		return handlers.Render(c, views.Calls(page))
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

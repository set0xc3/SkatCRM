package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/set0xc3/htmx/internal"
	"github.com/set0xc3/htmx/internal/model"
	"github.com/set0xc3/htmx/internal/server"
	"github.com/set0xc3/htmx/internal/server/handlers"
	"github.com/set0xc3/htmx/internal/views/icons"
	"github.com/set0xc3/htmx/internal/views/pages"
)

func InitModelData() model.Context {
	return model.Context{
		SidebarItems: []model.MenuItem{
			{
				Name: "Главная",
				URL:  "",
				Icon: icons.HomeIcon,
			},
			{
				Name: "Клиенты",
				URL:  "clients",
				Icon: icons.ClientsIcon,
			},
			{
				Name: "Звонки",
				URL:  "calls",
				Icon: icons.CallsIcon,
			},
			{
				Name: "Заказы",
				URL:  "orders",
				Icon: icons.OrdersIcon,
			},
			{
				Name: "Отчёты",
				URL:  "reports",
				Icon: icons.ReportsIcon,
			},
			{
				Name: "Товары",
				URL:  "products",
				Icon: icons.WarehouseIcon,
			},
		},
	}

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
	e.GET("/", func(c echo.Context) error {
		data := InitModelData()
		data.SidebarItems[0].IsHot = true
		return handlers.Render(c, pages.Home(data))
	})

	e.GET("/clients", func(c echo.Context) error {
		data := InitModelData()
		data.SidebarItems[1].IsHot = true
		return handlers.Render(c, pages.Clients(data))
	})

	// e.GET("/products", func(c echo.Context) error {
	// 	return handlers.Render(c, pages.ProductsPage())
	// })

	// e.GET("/cms", func(c echo.Context) error {
	// 	return handlers.Render(c, cms.IndexPage())
	// })
	// e.GET("/test", func(c echo.Context) error {
	// 	return handlers.Render(c, cms.TEST_CMS_IndexPage())
	// })
	e.GET("/status/ok", handlers.StatusOK)

	// API
	e.GET("/api/v1/clients", handlers.RedirectToDB)
	e.GET("/api/v1/clients/:count", handlers.RedirectToDB)
	e.DELETE("/api/v1/client/:id", handlers.StatusOK)

	e.GET("/api/v1/products", handlers.RedirectToDB)
	e.DELETE("/api/v1/product/:id", handlers.StatusOK)

	// Start server
	internal.GracefulShutdown(e, s.Port)
}

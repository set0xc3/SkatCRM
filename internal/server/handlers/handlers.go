package handlers

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/set0xc3/htmx/internal/views/pages"
)

func Render(c echo.Context, cmp templ.Component) error {
	return cmp.Render(c.Request().Context(), c.Response().Writer)
}

func GetIndexPage(c echo.Context) error {
	return Render(c, pages.IndexPage())
}

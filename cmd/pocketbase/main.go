package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"

	"github.com/set0xc3/htmx/internal/db"
)

func main() {
	app := pocketbase.New()

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.GET("/api/v1/products", func(c echo.Context) error {
			products := []db.Product{}
			err := app.Dao().DB().
				NewQuery("SELECT id, name, price, quantity FROM product LIMIT 100").
				All(&products)

			if err != nil {
				return c.JSON(http.StatusInternalServerError, "")
			}

			return c.JSON(http.StatusOK, products)
		}, apis.ActivityLogger(app))

		e.Router.GET("/api/v1/product/:id", func(c echo.Context) error {
			id := c.PathParam("id")
			product := db.Product{}

			err := app.Dao().DB().
				NewQuery("SELECT id, name, price, quantity FROM product WHERE id == {:id}").
				Bind(dbx.Params{
					"id": id,
				}).
				One(&product)

			if err != nil {
				return c.JSON(http.StatusInternalServerError, "")
			}

			return c.JSON(http.StatusOK, product)
		}, apis.ActivityLogger(app))

		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}

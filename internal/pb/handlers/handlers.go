package handlers

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v5"

	"github.com/pocketbase/dbx"
	pb "github.com/pocketbase/pocketbase"

	"github.com/set0xc3/htmx/internal/db"
)

func FetchProducts(ctx *pb.PocketBase) echo.HandlerFunc {
	return func(c echo.Context) error {
		products := []db.Product{}
		err := ctx.Dao().DB().
			NewQuery("SELECT id, name, price, quantity FROM product LIMIT 100").
			All(&products)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Not Found"})
		}

		return c.JSON(http.StatusOK, products)
	}
}

func FetchProduct(ctx *pb.PocketBase) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.PathParam("id")
		product := db.Product{}

		err := ctx.Dao().DB().
			NewQuery("SELECT id, name, price, quantity FROM product WHERE id = {:id}").
			Bind(dbx.Params{"id": id}).
			One(&product)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Not Found"})
		}

		return c.JSON(http.StatusOK, product)
	}
}

func FetchClients(ctx *pb.PocketBase) echo.HandlerFunc {
	return func(c echo.Context) error {
		clients := []db.Client{}
		err := ctx.Dao().DB().
			NewQuery("SELECT id, name, phone FROM client LIMIT 100").
			All(&clients)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Not Found"})
		}

		return c.JSON(http.StatusOK, clients)
	}
}

func FetchClient(ctx *pb.PocketBase) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.PathParam("id")
		client := db.Client{}

		err := ctx.Dao().DB().
			NewQuery("SELECT id, name, phone FROM client WHERE id = {:id}").
			Bind(dbx.Params{"id": id}).
			One(&client)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Not Found"})
		}

		return c.JSON(http.StatusOK, client)
	}
}

func DeleteClient(ctx *pb.PocketBase) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.PathParam("id")

		// Выполняем запрос на удаление клиента
		_, err := ctx.Dao().DB().
			NewQuery("DELETE FROM client WHERE id = {:id}").
			Bind(dbx.Params{"id": id}).
			Execute()

		if err != nil {
			// Логируем ошибку для диагностики
			log.Printf("Error deleting client with id %s: %v\n", id, err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error deleting client"})
		}

		return c.NoContent(http.StatusNoContent)
	}
}

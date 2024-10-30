package pb

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v5"

	pb "github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
	"github.com/set0xc3/htmx/internal/db"
	"github.com/set0xc3/htmx/internal/pb/handlers"
)

var ctx *pb.PocketBase

func Context() *pb.PocketBase {
	return ctx
}

func Run() {
	ctx = pb.New()
	ctx.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.GET("/api/v1/products", handlers.FetchProducts(ctx), apis.ActivityLogger(ctx))
		e.Router.GET("/api/v1/product/:id", handlers.FetchProduct(ctx), apis.ActivityLogger(ctx))
		e.Router.GET("/api/v1/clients", handlers.FetchClients(ctx), apis.ActivityLogger(ctx))
		e.Router.GET("/api/v1/client/:id", handlers.FetchClient(ctx), apis.ActivityLogger(ctx))
		e.Router.DELETE("/api/v1/client/:id", handlers.DeleteClient(ctx), apis.ActivityLogger(ctx))

		e.Router.GET("/api/v1/test", func(c echo.Context) error {
			query := ctx.Dao().RecordQuery("client")

			record := models.Record{}
			if err := query.One(&record); err != nil {
				return c.NoContent(http.StatusInternalServerError)
			}

			// https://github.com/pocketbase/pocketbase/discussions/5192
			client := db.Client{
				Id:    record.GetString("id"),
				Id2:   record.GetString("_id"),
				Name:  record.GetString("name"),
				Phone: record.GetString("phone"),
			}
			log.Println(client)

			return c.NoContent(http.StatusNoContent)
		}, apis.ActivityLogger(ctx))

		e.Router.GET("/api/v1/add-test", func(c echo.Context) error {
			collection, err := ctx.Dao().FindCollectionByNameOrId("client")

			if err != nil {
				return err
			}

			record := models.NewRecord(collection)
			record.Set("_id", "892")
			record.Set("name", "Евгений")

			if err := ctx.Dao().SaveRecord(record); err != nil {
				return c.NoContent(http.StatusNoContent)
			}

			return c.NoContent(http.StatusNoContent)
		}, apis.ActivityLogger(ctx))

		return nil
	})

	if err := ctx.Start(); err != nil {
		log.Fatal(err)
	}
}

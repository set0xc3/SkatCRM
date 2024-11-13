package pb

import (
	"log"

	pb "github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"

	"github.com/set0xc3/htmx/internal/pb/handlers"
)

var ctx *pb.PocketBase

func Context() *pb.PocketBase {
	return ctx
}

func Run() {
	ctx = pb.New()
	ctx.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.GET("/api/v1/products", handlers.GetProducts(ctx), apis.ActivityLogger(ctx))
		e.Router.GET("/api/v1/product/:id", handlers.GetProduct(ctx), apis.ActivityLogger(ctx))

		e.Router.POST("/api/v1/product", handlers.PostAddProduct(ctx), apis.ActivityLogger(ctx))
		e.Router.DELETE("/api/v1/product/:id", handlers.DeleteProduct(ctx), apis.ActivityLogger(ctx))

		e.Router.GET("/api/v1/clients", handlers.GetClients(ctx), apis.ActivityLogger(ctx))
		e.Router.GET("/api/v1/client/:id", handlers.GetClient(ctx), apis.ActivityLogger(ctx))

		e.Router.POST("/api/v1/client", handlers.PostAddClient(ctx), apis.ActivityLogger(ctx))
		e.Router.DELETE("/api/v1/client/:id", handlers.DeleteClient(ctx), apis.ActivityLogger(ctx))

		return nil
	})

	if err := ctx.Start(); err != nil {
		log.Fatal(err)
	}
}

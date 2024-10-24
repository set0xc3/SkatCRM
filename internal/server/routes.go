package server

import (
	// "encoding/json"
	// "log"
	"net/http"

	"github.com/a-h/templ"
	// "github.com/aarol/reload"

	"github.com/set0xc3/htmx"
	"github.com/set0xc3/htmx/views"
)

func (s *Server) RegisterRoutes() http.Handler {
	fileServer := http.FileServer(http.FS(efs.Files))

	mux := http.NewServeMux()
	mux.Handle("/static/", fileServer)

	mux.Handle("/", templ.Handler(views.Home()))
	mux.Handle("/catalog", templ.Handler(views.Catalog()))
	mux.Handle("/contacts", templ.Handler(views.Empty()))
	mux.Handle("/about", templ.Handler(views.Empty()))

	mux.HandleFunc("GET /api/v1/products", GetProducts)


	return mux
}

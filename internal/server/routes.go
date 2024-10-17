package server

import (
	// "encoding/json"
	// "log"
	"net/http"

	"github.com/a-h/templ"
	// "github.com/aarol/reload"

	"github.com/set0xc3/htmx/cmd/web"
)

func (s *Server) RegisterRoutes() http.Handler {
	fileServer := http.FileServer(http.FS(web.Files))

	mux := http.NewServeMux()
	mux.Handle("/", templ.Handler(web.Base()))
	mux.Handle("/static/", fileServer)

	return mux
}

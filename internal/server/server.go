package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kissejau/brawlrate/internal/server/routes"
)

func Run(r *mux.Router) {
	routes.RegisterRoutes(r)
	http.Handle("/", r)
}

func Config(r *mux.Router) {
	r.Use(mux.CORSMethodMiddleware(r))
}

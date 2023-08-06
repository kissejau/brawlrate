package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kissejau/brawlrate/internal/server/handlers"
)

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/steamId", handlers.SteamIdHandler).Methods("POST")
	r.HandleFunc("/info", handlers.InfoHandler).Methods("POST")
	r.HandleFunc("/stats", handlers.StatsHandler).Methods("POST")
	r.HandleFunc("/ranked", handlers.RankedHandler).Methods("POST")

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("home"))
	}).Methods("GET")
}

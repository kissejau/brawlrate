package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/kissejau/brawlhalla-search/internal/server/services"
)

func SteamIdHandler(w http.ResponseWriter, r *http.Request) {
	headers := r.Header.Values("steam_url")
	if len(headers) <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	steamData, err := services.GetSteamId(headers[0])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	data, _ := json.MarshalIndent(steamData, "", "\t")

	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}

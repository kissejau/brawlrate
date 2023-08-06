package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/kissejau/brawlhalla-search/internal/server/handlers/utils"
	u "github.com/kissejau/brawlhalla-search/internal/utils"
)

func SteamIdHandler(w http.ResponseWriter, r *http.Request) {
	var data []byte

	rd, err := utils.GetRequestData(r)
	if err != nil {
		u.Respond(w, http.StatusBadRequest, []byte(err.Error()))
		return
	}

	data, _ = json.Marshal(rd.SteamData)
	u.Respond(w, http.StatusOK, data)
}

func InfoHandler(w http.ResponseWriter, r *http.Request) {
	var data []byte

	rd, err := utils.GetRequestData(r)
	if err != nil {
		u.Respond(w, http.StatusBadRequest, []byte(err.Error()))
		return
	}

	data, _ = json.Marshal(rd.Info)
	u.Respond(w, http.StatusAccepted, data)
}

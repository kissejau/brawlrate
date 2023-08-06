package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/kissejau/brawlhalla-search/internal/server/handlers/utils"
	"github.com/kissejau/brawlhalla-search/internal/server/services"
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

func StatsHandler(w http.ResponseWriter, r *http.Request) {
	var data []byte

	rd, err := utils.GetRequestData(r)
	if err != nil {
		u.Respond(w, http.StatusBadRequest, []byte(err.Error()))
		return
	}

	stats, err := services.GetStats(strconv.Itoa(rd.Info.Id))
	if err != nil {
		u.Respond(w, http.StatusBadRequest, []byte(err.Error()))
		return
	}

	data, _ = json.Marshal(stats)
	u.Respond(w, http.StatusAccepted, data)
}

func RankedHandler(w http.ResponseWriter, r *http.Request) {
	var data []byte

	rd, err := utils.GetRequestData(r)
	if err != nil {
		u.Respond(w, http.StatusBadRequest, []byte(err.Error()))
		return
	}

	ranked, err := services.GetRanked(strconv.Itoa(rd.Info.Id))
	if err != nil {
		u.Respond(w, http.StatusBadRequest, []byte(err.Error()))
		return
	}

	data, _ = json.Marshal(ranked)
	u.Respond(w, http.StatusAccepted, data)
}

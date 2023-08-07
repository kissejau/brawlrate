package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kissejau/brawlrate/internal/server/handlers/utils"
	"github.com/kissejau/brawlrate/internal/server/services"
	u "github.com/kissejau/brawlrate/internal/utils"
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

func RankingsHandler(w http.ResponseWriter, r *http.Request) {
	var (
		data   []byte
		braket string
		region string
		page   int
	)

	vars := mux.Vars(r)
	fmt.Println(vars)
	braket = r.URL.Query().Get("braket")
	region = r.URL.Query().Get("region")
	page, _ = strconv.Atoi(r.URL.Query().Get("page"))

	rankings, err := services.GetRankings(braket, region, page)
	if err != nil {
		u.Respond(w, http.StatusBadRequest, []byte(err.Error()))
		return
	}

	data, _ = json.Marshal(rankings)
	u.Respond(w, http.StatusAccepted, data)
}

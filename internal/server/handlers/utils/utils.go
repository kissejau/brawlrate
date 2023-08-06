package utils

import (
	"fmt"
	"net/http"

	"github.com/kissejau/brawlhalla-search/internal/brawlhalla/models"
	"github.com/kissejau/brawlhalla-search/internal/server/services"
	"github.com/kissejau/brawlhalla-search/internal/utils"
)

type RequestData struct {
	SteamData *utils.SteamData
	Info      *models.Info
}

func GetRequestData(r *http.Request) (*RequestData, error) {
	rd := &RequestData{}

	steamId := r.Header.Values("steam_id")
	steamUrl := r.Header.Values("steam_url")

	if len(steamId)+len(steamUrl) <= 0 {
		return nil, fmt.Errorf("steamID or steamURL were not given")
	}

	if len(steamId) > 0 {
		info, err := services.GetInfoBySteamId(steamId[0])
		if err != nil {
			return nil, err
		}
		rd.Info = info
	}

	if len(steamUrl) > 0 {
		steamData, err := services.GetSteamId(steamUrl[0])
		if err != nil {
			return nil, err
		}
		rd.SteamData = steamData

		info, err := services.GetInfoBySteamId(steamData.SteamId)
		if err != nil {
			return nil, err
		}
		rd.Info = info
	}

	return rd, nil
}

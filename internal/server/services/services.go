package services

import (
	"fmt"

	"github.com/kissejau/brawlhalla-search/internal/brawlhalla"
	"github.com/kissejau/brawlhalla-search/internal/brawlhalla/models"
	"github.com/kissejau/brawlhalla-search/internal/config"
	"github.com/kissejau/brawlhalla-search/internal/utils"
)

// idk how to except it
var bh brawlhalla.BrawlhallaService = brawlhalla.NewBrawlhallaService(cfg.ApiKey)
var cfg, cfgError = config.NewConfig()

func GetSteamId(url string) (*utils.SteamData, error) {
	data, err := utils.GetSteamDataByUrl(url)
	if err != nil {
		return nil, fmt.Errorf("not correct steam url: %s", err.Error())
	}
	return data, nil
}

func GetInfoBySteamId(id string) (*models.Info, error) {
	info, err := bh.GetInfo(id)
	if err != nil {
		return nil, fmt.Errorf("not correct steamID: %s", err)
	}

	return info, nil
}

func init() {
	if cfgError != nil {
		panic(cfgError)
	}
}

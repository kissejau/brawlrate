package services

import (
	"fmt"

	"github.com/kissejau/brawlrate/internal/brawlhalla"
	"github.com/kissejau/brawlrate/internal/brawlhalla/models"
	"github.com/kissejau/brawlrate/internal/config"
	u "github.com/kissejau/brawlrate/internal/server/services/utils"
	"github.com/kissejau/brawlrate/internal/utils"
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

func GetStats(id string) (*models.Stats, error) {
	stats, err := bh.GetStats(id)
	if err != nil {
		return nil, fmt.Errorf("not correct brawlhallaID: %s", err)
	}

	return stats, nil
}

func GetRanked(id string) (*models.Ranked, error) {
	ranked, err := bh.GetRanked(id)
	if err != nil {
		return nil, fmt.Errorf("not correct brawlhallaID: %s", err)
	}

	return ranked, nil
}

func GetRankings(braket string, region string, page int) ([]models.Ranking, error) {
	var rankings []models.Ranking
	rd, err := u.NewRankingsData(braket, region, page)
	if err != nil {
		return nil, err
	}

	rankings, err = bh.GetRankings(rd.Braket, rd.Region, rd.Page)
	if err != nil {
		return nil, err
	}

	return rankings, nil
}

func init() {
	if cfgError != nil {
		panic(cfgError)
	}
}

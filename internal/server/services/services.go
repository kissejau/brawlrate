package services

import (
	"fmt"

	"github.com/kissejau/brawlhalla-search/internal/brawlhalla"
	"github.com/kissejau/brawlhalla-search/internal/config"
	"github.com/kissejau/brawlhalla-search/internal/utils"
)

var bh brawlhalla.BrawlhallaService
var cfg *config.Config // idk how to except it

func GetSteamId(url string) (*utils.SteamData, error) {
	data, err := utils.GetSteamDataByUrl(url)
	if err != nil {
		return nil, fmt.Errorf("not correct steam url: %s", err.Error())
	}
	return data, nil
}

// func init() {
// 	cfg, err := config.NewConfig()
// 	if err != nil {
// 		panic(err)
// 	}

// 	bh := brawlhalla.NewBrawlhallaService(cfg.ApiKey)
// }

package brawlhalla

import "github.com/kissejau/brawlhalla-search/internal/brawlhalla/models"

type BrawlhallaService interface {
	GetInfo(steamId string) (*models.Info, error)
	GetRankings(braket string, region string, page int) ([]models.Ranking, error)
	GetStats(id string) (*models.Stats, error)
	GetRanked(id string) (*models.Ranked, error)
	GetClan(id string) (*models.Clan, error)
}

type brawlhallaService struct {
	url    string
	apiKey string
}

func NewBrawlhallaService(apiKey string) BrawlhallaService {
	brawlhallaService := &brawlhallaService{
		url:    "https://api.brawlhalla.com",
		apiKey: apiKey,
	}
	return brawlhallaService
}

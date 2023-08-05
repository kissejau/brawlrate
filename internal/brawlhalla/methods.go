package brawlhalla

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/kissejau/brawlhalla-search/internal/brawlhalla/models"
)

// param is steamid in format steamID64
func (bs brawlhallaService) GetInfo(steamId string) (*models.Info, error) {
	info := &models.Info{}

	perform, err := url.JoinPath(bs.url, "search")
	if err != nil {
		return nil, err
	}

	urlValues := url.Values{}
	urlValues.Add("steamid", steamId)
	urlValues.Add("api_key", bs.apiKey)

	perform = fmt.Sprintf("%s?%s", perform, urlValues.Encode())

	res, err := http.Get(perform)
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, info)
	if err != nil {
		return nil, err
	}

	return info, nil
}

func (bs brawlhallaService) GetRankings(braket string, region string, page int) ([]models.Ranking, error) {
	rankings := []models.Ranking{}

	perform, err := url.JoinPath(bs.url, "rankings", braket, region, strconv.Itoa(page))
	if err != nil {
		return nil, err
	}

	urlValues := url.Values{}
	urlValues.Add("api_key", bs.apiKey)

	perform = fmt.Sprintf("%s?%s", perform, urlValues.Encode())

	res, err := http.Get(perform)
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &rankings)
	if err != nil {
		return nil, err
	}

	return rankings, nil
}

func (bs brawlhallaService) GetStats(id string) (*models.Stats, error) {
	stats := &models.Stats{}

	perform, err := url.JoinPath(bs.url, "player", id, "stats")
	if err != nil {
		return nil, err
	}

	urlValues := url.Values{}
	urlValues.Add("api_key", bs.apiKey)

	perform = fmt.Sprintf("%s?%s", perform, urlValues.Encode())

	var data []byte
	res, err := http.Get(perform)
	if err != nil {
		return nil, err
	}

	data, err = io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, stats)
	if err != nil {
		return nil, err
	}

	return stats, nil
}

func (bs brawlhallaService) GetRanked(id string) (*models.Ranked, error) {
	ranked := &models.Ranked{}

	perform, err := url.JoinPath(bs.url, "player", id, "ranked")
	if err != nil {
		return nil, err
	}

	urlValues := url.Values{}
	urlValues.Add("api_key", bs.apiKey)

	perform = fmt.Sprintf("%s?%s", perform, urlValues.Encode())

	var data []byte
	res, err := http.Get(perform)
	if err != nil {
		return nil, err
	}

	data, err = io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, ranked)
	if err != nil {
		return nil, err
	}

	return ranked, nil
}

func (bs brawlhallaService) GetClan(id string) (*models.Clan, error) {
	clan := &models.Clan{}

	perform, err := url.JoinPath(bs.url, "clan", id)
	if err != nil {
		return nil, err
	}

	urlValues := url.Values{}
	urlValues.Add("api_key", bs.apiKey)

	perform = fmt.Sprintf("%s?%s", perform, urlValues.Encode())

	var data []byte
	res, err := http.Get(perform)
	if err != nil {
		return nil, err
	}

	data, err = io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, clan)
	if err != nil {
		return nil, err
	}

	return clan, nil
}

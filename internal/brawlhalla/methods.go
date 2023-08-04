package brawlhalla

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/kissejau/brawlhalla-search/internal/brawlhalla/models"
)

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

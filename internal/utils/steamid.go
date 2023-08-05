package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
)

type SteamData struct {
	SteamId string `json:"steamid"`
}

func GetSteamDataByUrl(url string) (*SteamData, error) {
	steamData := &SteamData{}
	prefix := "g_rgProfileData = "

	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	r, err := regexp.Compile(prefix + `{(.*)}`)
	if err != nil {
		return nil, err
	}

	target := r.FindString(string(data))

	target, found := strings.CutPrefix(target, prefix)
	if !found {
		return nil, fmt.Errorf("cant get steam data")
	}

	json.Unmarshal([]byte(target), steamData)

	return steamData, nil
}

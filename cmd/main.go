package main

import (
	"fmt"

	"github.com/kissejau/brawlhalla-search/internal/config"
)

func main() {
	config, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	fmt.Println(config.ApiKey)
}

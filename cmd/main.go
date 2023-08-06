package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kissejau/brawlrate/internal/config"
	"github.com/kissejau/brawlrate/internal/server"
)

func main() {
	config, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	router := mux.NewRouter()
	server.Config(router)
	server.Run(router)
	http.ListenAndServe(fmt.Sprintf("%s:%s", config.Host, config.Port), router)
}

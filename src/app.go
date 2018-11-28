package main

import (
	"fmt"
	"github.com/UPrefer/StorageService/api"
	"github.com/UPrefer/StorageService/config"
	"github.com/jinzhu/configor"
	"log"
	"net/http"
)

type App struct {
	configuration *config.Configuration
}

func (app App) Run() {
	app.configuration = &config.Configuration{}
	configor.Load(app.configuration)

	var listeningAddr = ":" + app.configuration.Port

	fmt.Printf("Starting  StorageService API on address : %s\n", listeningAddr)
	err := http.ListenAndServe(listeningAddr, api.Handlers(app.configuration))
	if err != nil {
		log.Fatal(listeningAddr, err)
	}
	fmt.Printf("Started ...")
}

func main() {
	app := new(App)
	app.Run()
}

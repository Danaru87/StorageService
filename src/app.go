package main

import (
	"fmt"
	"github.com/UPrefer/StorageService/api"
	"log"
	"net/http"
)

type App struct{}

func (app App) Run(listeningAddr string) {
	fmt.Printf("Starting  StorageService API on address : %s\n", listeningAddr)
	err := http.ListenAndServe(listeningAddr, api.Handlers())
	if err != nil {
		log.Fatal(listeningAddr, err)
	}
	fmt.Printf("Started ...")
}

func main() {
	app := new(App)
	app.Run(":8082")
}

package main

import (
	"math/rand"
	"net/http"
	"slava0135/nanoservice/handlers"
	"time"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("starting up")
	rand.Seed(time.Now().UnixNano())

	http.HandleFunc("/generate", handlers.GenerateGameLayout)
	http.HandleFunc("/generate/image", handlers.GenerateGameLayoutImage)
	http.HandleFunc("/generate/grid", handlers.GenerateGameLayoutGrid)
	http.HandleFunc("/validate", handlers.ValidateShipPlacement)
	http.HandleFunc("/replay", handlers.ReplayGame)

	port := ":8080"
	log.Info("listening on port ", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

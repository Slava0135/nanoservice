package main

import (
	"math/rand"
	"net/http"
	"slava0135/nanoservice/handlers"
	"time"
	"log"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	http.HandleFunc("/generate", handlers.GenerateGameLayout)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

package main

import (
	"math/rand"
	"net/http"
	"slava0135/nanoservice/handlers"
	"time"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info("starting up")
	rand.Seed(time.Now().UnixNano())

	port := ":8080"
	http.HandleFunc("/generate", handlers.GenerateGameLayout)
	logrus.Info("listening on port ", port)
	logrus.Fatal(http.ListenAndServe(port, nil))
}

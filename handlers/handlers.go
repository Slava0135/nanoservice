package handlers

import (
	"fmt"
	"io"
	"net/http"
	"slava0135/nanoservice/generate"
	"slava0135/nanoservice/layout"
	"slava0135/nanoservice/shots"
	"slava0135/nanoservice/validate"
	"strings"

	log "github.com/sirupsen/logrus"
)

func GenerateGameLayout(w http.ResponseWriter, req *http.Request) {
	log.Info("handle generate game layout request")
	_, ships := generate.NewGameLayout()
	for _, s := range ships {
		fmt.Fprintf(w, "%v\n", s)
	}
}

func ValidateShipPlacement(w http.ResponseWriter, req *http.Request) {
	log.Info("handle validate ship placement request")
	b, err := io.ReadAll(req.Body)
	if err != nil {
		log.Info("error when reading body: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	gameLayout, _, err := layout.ParseLayout(string(b))
	if err != nil {
		log.Info("error when reading body: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "%v\n", validate.Validate(gameLayout))
}

func ReplayGame(w http.ResponseWriter, req *http.Request) {
	log.Info("handle replay game request")
	b, err := io.ReadAll(req.Body)
	if err != nil {
		log.Info("error when reading body: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	s := string(b)
	bodyParts := strings.Split(s, "@")
	if len(bodyParts) != 2 {
		err := fmt.Errorf("more than 2 parts (ships and shots) detected")
		log.Info("error when reading body: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	gameLayout, ships, err := layout.ParseLayout(bodyParts[0])
	if err != nil {
		log.Info("error when reading body: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if !validate.Validate(gameLayout) {
		log.Info("user provided invalid ship layout")
		http.Error(w, "invalid ship layout", http.StatusBadRequest)
		return
	}
	gameShots, err := shots.ParseShots(bodyParts[1])
	if err != nil {
		log.Info("user provided invalid shots data: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "%v\n", shots.ReplayGame(ships, gameShots))
}
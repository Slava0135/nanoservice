package handlers

import (
	"fmt"
	"io"
	"net/http"
	"slava0135/nanoservice/generate"
	"slava0135/nanoservice/layout"
	"slava0135/nanoservice/validate"

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

package handlers

import (
	"fmt"
	"net/http"
	"slava0135/nanoservice/generate"

	log "github.com/sirupsen/logrus"
)

func GenerateGameLayout(w http.ResponseWriter, req *http.Request) {
	log.Info("handle generate game layout request")
	_, ships := generate.NewGameLayout()
	for _, s := range ships {
		fmt.Fprintf(w, "%v\n", s)
	}
}
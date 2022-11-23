package handlers

import (
	"fmt"
	"net/http"
	"slava0135/nanoservice/generate"

	"github.com/sirupsen/logrus"
)

func GenerateGameLayout(w http.ResponseWriter, req *http.Request) {
	logrus.Info("handle generate game layout request")
	_, ships := generate.NewGameLayout()
	for _, s := range ships {
		fmt.Fprintf(w, "%v\n", s)
	}
}
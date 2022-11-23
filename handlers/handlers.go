package handlers

import (
	"fmt"
	"net/http"
	"slava0135/nanoservice/generate"
)

func GenerateGameLayout(w http.ResponseWriter, req *http.Request) {
	_, ships := generate.NewGameLayout()
	for _, s := range ships {
		fmt.Fprintf(w, "%v\n", s)
	}
}
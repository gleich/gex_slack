package main

import (
	"net/http"

	"github.com/gleich/gex_slack/pkg/api"
	"github.com/gleich/lumber"
)

func main() {
	lumber.ErrNilCheck = false
	lumber.Info("Booted Up")
	http.HandleFunc("/slash", api.HandleSlash)

	lumber.Info("Starting server on port 80")
	if err := http.ListenAndServe(":80", nil); err != nil {
		lumber.Error(err, "Failed to start server")
	}
}

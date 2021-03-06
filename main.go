package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gleich/gex_slack/pkg/api"
	"github.com/gleich/lumber/v2"
	"github.com/go-co-op/gocron"
)

func main() {
	lumber.Info("Booted Up")
	http.HandleFunc("/slash", api.HandleSlash)
	http.HandleFunc("/", api.HandleRoot)
	rand.Seed(time.Now().UnixNano())

	lumber.Info("Starting server on port 80")
	go func() {
		if err := http.ListenAndServe(":80", nil); err != nil {
			lumber.Error(err, "Failed to start server")
		}
	}()

	s := gocron.NewScheduler(time.UTC)
	_, err := s.Every(1).Day().At("19:00").Do(func() {
		api.SendQuote("C02ERUU1N0K") // # gex-test
		api.SendQuote("C02EPHVHVNJ") // # gex-quotes
	})
	if err != nil {
		lumber.Error(err, "Failed to start new daily update")
	}
	lumber.Success("Setup scheduler")
	s.StartBlocking()
}

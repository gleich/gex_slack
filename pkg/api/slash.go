package api

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/gleich/gex_slack/pkg/quotes"
	"github.com/gleich/lumber/v2"
	"github.com/slack-go/slack"
)

func HandleSlash(w http.ResponseWriter, r *http.Request) {
	verifier, err := slack.NewSecretsVerifier(r.Header, os.Getenv("SLACK_SIGNING_SECRET"))
	if err != nil {
		lumber.Error(err, "Failed to verify secrets")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	r.Body = io.NopCloser(io.TeeReader(r.Body, &verifier))
	s, err := slack.SlashCommandParse(r)
	if err != nil {
		lumber.Error(err, "Failed to parse slack command")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err = verifier.Ensure(); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	switch s.Command {
	case "/gex":
		params := &slack.Msg{Text: quotes.RandomQuote(), ResponseType: slack.ResponseTypeInChannel}
		b, err := json.Marshal(params)
		if err != nil {
			lumber.Error(err, "Failed to marshal json for /gex data")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(b)
		if err != nil {
			lumber.Error(err, "Failed to write response for /gex data")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	default:
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	lumber.Success("Handled slash command")
}

package api

import (
	"os"

	"github.com/gleich/gex_slack/pkg/quotes"
	"github.com/gleich/lumber/v2"
	"github.com/slack-go/slack"
)

func SendQuote(channel string) {
	sl := slack.New(os.Getenv("SLACK_OAUTH_TOKEN"))
	quote := quotes.RandomQuote()
	_, _, err := sl.PostMessage(
		channel,
		slack.MsgOptionText(quote, false),
	)
	if err != nil {
		lumber.Error(err, "Failed to post slack message to channel with id of:", channel)
	}
	lumber.Success("Sent the quote:", quote, "to", channel)
}

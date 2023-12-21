package lib

import (
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"log/slog"
	"os"
)

func NewLineBot() (bot *messaging_api.MessagingApiAPI, err error) {
	bot, err = messaging_api.NewMessagingApiAPI(os.Getenv("CHANNEL_TOKEN"))
	if err != nil {
		slog.Error(err.Error(), "LineInit")
		return
	}
	return
}

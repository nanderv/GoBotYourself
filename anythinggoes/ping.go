package anythinggoes

import (
	"gopkg.in/telegram-bot-api.v4"
	"strings"
	"bot/gofuckyourself"
)

func run(update tgbotapi.Update, _ *tgbotapi.BotAPI) string {
	str := ""
	for _, field := range strings.Fields(update.Message.Text) {
		if field == "ping" {
			str = str + " pong"
		} else {
			str = str + " " + field
		}
	}
	return str
}

var Ping Module = Module{"ping",
	"Says pong",
	gofuckyourself.Contains("ping"),
	gofuckyourself.Replier(run)}

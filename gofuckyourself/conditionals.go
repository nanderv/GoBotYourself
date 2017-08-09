package gofuckyourself

import (
	"gopkg.in/telegram-bot-api.v4"
	"strings"
)

func Contains(check string) func(tgbotapi.Update, *tgbotapi.BotAPI) bool {
	return func(u tgbotapi.Update, _ *tgbotapi.BotAPI) bool {
		for _, field := range strings.Fields(u.Message.Text) {
			if field == check {
				return true
			}
		}
		return false
	}
}
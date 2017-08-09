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

func Or(f1 func(tgbotapi.Update, *tgbotapi.BotAPI) bool, f2 func(tgbotapi.Update, *tgbotapi.BotAPI) bool) func(tgbotapi.Update, *tgbotapi.BotAPI) bool {
	return func(u tgbotapi.Update, b *tgbotapi.BotAPI) bool {
		return f1(u, b) || f2(u, b)
	}
}
func RunAlways(tgbotapi.Update, *tgbotapi.BotAPI) bool {
	return true
}
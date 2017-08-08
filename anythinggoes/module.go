package anythinggoes

import "gopkg.in/telegram-bot-api.v4"

var MOD map[string] Module = map[string] Module {}
type Module struct{
	Name string
	Help string
	Condition func(tgbotapi.Update, *tgbotapi.BotAPI) bool
	Run func(tgbotapi.Update, *tgbotapi.BotAPI)
}
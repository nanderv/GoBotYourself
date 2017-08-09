package anythinggoes

import (
	"gopkg.in/telegram-bot-api.v4"
	"bot/gofuckyourself"
	"fmt"
)

func store(update tgbotapi.Update, _ *tgbotapi.BotAPI) string {

	var variable string
	var variablerec string
	var text string
	int1, _ := fmt.Sscanf(update.Message.Text, "/store %v %v", &variable, &text)
	int2, _ := fmt.Sscanf(update.Message.Text, "/receive %v", &variablerec)

	switch true{
	case int1 == 2:        fmt.Println("store")
	case int2 == 1:        fmt.Println("retreive")
	}
	return ""
}

var Store Module = Module{"store",
	"store value using /store, retrieve using /retreive",
	gofuckyourself.RunAlways,
	gofuckyourself.Replier(store)}

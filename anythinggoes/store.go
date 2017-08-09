package anythinggoes

import (
	"gopkg.in/telegram-bot-api.v4"
	"bot/gofuckyourself"
	"fmt"
	"bot/sqli"
)

func store(update tgbotapi.Update, _ *tgbotapi.BotAPI) string {

	var variable string
	var variablerec string
	var text string
	int1, _ := fmt.Sscanf(update.Message.Text, "/store %v %v", &variable, &text)
	int2, _ := fmt.Sscanf(update.Message.Text, "/receive %v", &variablerec)

	switch true{
	case int1 == 2:
		d := sql.Data{0, false, update.Message.Chat.ID, "store", variable, text}
		d.Save()
		return "Data stored"
	case int2 == 1:   d := sql.Data{}.LoadData(update.Message.Chat.ID, Store.Name, variablerec)
		return d.Data
	}

	return ""
}

var Store Module = Module{"store",
	"store value using /store, retrieve using /retreive",
	gofuckyourself.RunAlways,
	gofuckyourself.Replier(store)}

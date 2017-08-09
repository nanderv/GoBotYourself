package anythinggoes

import (
	"gopkg.in/telegram-bot-api.v4"
	"bot/gofuckyourself"
	"fmt"
	"bot/sqli"
)

func store(update tgbotapi.Update, _ *tgbotapi.BotAPI) (string, bool) {

	var variable string
	var variablerec string
	var text string
	int1, _ := fmt.Sscanf(update.Message.Text, "/store %v %v", &variable, &text)
	int2, _ := fmt.Sscanf(update.Message.Text, "/receive %v", &variablerec)

	switch{
	case int1 == 2:
		d := sql.Data{0, false, update.Message.Chat.ID, "store", variable, text}
		d.Save()
		return "Data stored", true
	case int2 == 1:   d := sql.Data{}.LoadData(update.Message.Chat.ID, "store", variablerec)
		return d.Data, true
	}

	return "", false
}

var Store Module = Module{"store",
	"store value using /store, retrieve using /retreive",
	gofuckyourself.RunAlways,
	gofuckyourself.IFReplier(store)}

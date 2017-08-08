package anythinggoes

import ("gopkg.in/telegram-bot-api.v4"
	"strings"
)

func condition(u tgbotapi.Update, bot *tgbotapi.BotAPI) bool{
	for _, field := range strings.Fields(u.Message.Text){
		if field =="ping"{
			return true
		}
	}
	return false
}
func run(update tgbotapi.Update, bot *tgbotapi.BotAPI){
	str := ""
	for _, field := range strings.Fields(update.Message.Text){
		if field =="ping"{
			str = str +" pong"
		} else{
			str = str + " "+field
		}
	}
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, str)
	msg.ReplyToMessageID = update.Message.MessageID
	bot.Send(msg)
}
var Ping Module = Module{"ping",
  "Says pong",
  condition,
  run}

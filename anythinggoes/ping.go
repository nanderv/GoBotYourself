package anythinggoes

import ("gopkg.in/telegram-bot-api.v4"
	"strings"
	"bot/gofuckyourself"
)

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
  gofuckyourself.Contains("ping"),
  run}

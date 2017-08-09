package gofuckyourself

import "gopkg.in/telegram-bot-api.v4"

func Replier(fun func(tgbotapi.Update, *tgbotapi.BotAPI) string) func(tgbotapi.Update, *tgbotapi.BotAPI) {
	return func(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
		str := fun(update, bot)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, str)
		msg.ReplyToMessageID = update.Message.MessageID
		bot.Send(msg)
	}
}
package main

import (
	"log"
	"gopkg.in/telegram-bot-api.v4"
	"bot/sqli"
	"fmt"
	"reflect"
)

func main() {
	dba := sql.Getdb()
	fooType := reflect.TypeOf(dba)
	for i := 0; i < fooType.NumMethod(); i++ {
    		method := fooType.Method(i)
    		fmt.Println(method.Name)
	}
	dba.Emptydb()
	
	fmt.Println(fooType)
	fmt.Println(dba)
	bot, err := tgbotapi.NewBotAPI("281611909:AAEszBrE92Ok5W7WL1Qxcx6rY2zNNS5lGkw")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		if update.Message.Chat.ID != -1001050885996{
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
		}
	}
}

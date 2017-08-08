package main

import (
	"log"
	"gopkg.in/telegram-bot-api.v4"
	"bot/sqli"
	"bot/settings"
	"fmt"
	"io/ioutil"
	"os"
)
type Settings struct {
    Api string
}
func main() {
	dba := sql.Getdb()
	fmt.Println(dba)
	    file, e := ioutil.ReadFile("./config.json")
    if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }
    fmt.Printf("%s\n", string(file))

    	//m := new(Dispatch)
    	//var m interface{}
	botSetup := settings.GetSettings()
	bot, err := tgbotapi.NewBotAPI(botSetup.Api)
	
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

package main

import (
	"log"
	"encoding/json"
	"gopkg.in/telegram-bot-api.v4"
	"bot/sqli"
	"fmt"
	"io/ioutil"
	"os"
)
type Settings struct {
    Api string
}
func main() {
	    file, e := ioutil.ReadFile("./config.json")
    if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }
    fmt.Printf("%s\n", string(file))

    	//m := new(Dispatch)
    	//var m interface{}
    	var settings Settings
	json.Unmarshal(file, &settings)
	fmt.Printf("Results: %v\n", settings)
	dba := sql.Getdb()

	dba.Emptydb()
	
	fmt.Println(dba)
	bot, err := tgbotapi.NewBotAPI(settings.Api)
	
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

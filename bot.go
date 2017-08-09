package main

import (
	"log"
	"gopkg.in/telegram-bot-api.v4"
	"bot/sqli"
	"bot/settings"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"bot/anythinggoes"
)

func main() {
	dba := sql.Getdb()
	dba.Emptydb()
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
	//anythinggoes.MOD["ping"] = anythinggoes.Ping
	anythinggoes.MOD["store"] = anythinggoes.Store
	anythinggoes.MOD["nsapi"] = anythinggoes.NSApi

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
		m := sql.Message{update.Message.Chat.ID, update.Message.MessageID, update.Message.From.UserName, update.Message.Text}
		fmt.Println(reflect.TypeOf(update))
		m.Save()
			for _, Mod := range anythinggoes.MOD {
				if Mod.Condition(update, bot) {
					Mod.Run(update, bot)
				}
			}

	}
}

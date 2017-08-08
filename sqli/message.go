package sql

import (
	"log"
)

type Message struct {
	ChatID    int64
	MessageID int
	UserName  string
	Text      string
}

func (m Message) Save() {
	db := Getdb().db

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare("insert into Messages(chatID, messageID, userName, content) values(?, ?, ?, ?)")
	defer stmt.Close()

	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(m.ChatID, m.MessageID, m.UserName, m.Text)
	if err != nil {
		log.Fatal(err)
	}
	tx.Commit()
}
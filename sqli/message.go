package sql

import (
	"fmt"
	"log"
)
type Message struct{
	ChatID int64
	MessageID int
	UserName string
	Text string
}

func (m Message) Save(){
	db := Getdb().db
	fmt.Println(db)
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(m)
	stmt, err := tx.Prepare("insert into Messages(chatID, messageID, userName, content) values(?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(m.ChatID, m.MessageID,m.UserName, m.Text)
	if err != nil {
		log.Fatal(err)
	}
	tx.Commit()
}
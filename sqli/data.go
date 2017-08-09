package sql

import (
	"log"
	"fmt"
)

type Data struct {
	ID         int
	Indb       bool
	ChatID     int64
	ModuleName string
	Variable   string
	Data       string
}

func (d Data) Save() {
	db := Getdb().db

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	if !d.Indb {
		stmt, err := tx.Prepare("insert into Data(chatID, moduleName, variable, data) values(?, ?, ?, ?)")
		defer stmt.Close()

		if err != nil {
			log.Fatal(err)
		}
		_, err = stmt.Exec(d.ChatID, d.ModuleName, d.Variable, d.Data)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// update
		stmt, err := tx.Prepare("update Data SET data = ? WHERE chatID = ? and moduleName = ? and variable = ?)")
		defer stmt.Close()

		if err != nil {
			log.Fatal(err)
		}
		_, err = stmt.Exec(d.Data, d.ChatID, d.ModuleName, d.Variable)
		if err != nil {
			log.Fatal(err)
		}
	}
	tx.Commit()
}

func (d Data) LoadData(chatID int64, moduleName string, variable string) Data {
	db := Getdb().db
	fmt.Print(moduleName)

	stmt, err := db.Prepare("select id, data from Data where moduleName = ? AND chatID = ? AND variable = ?")
	defer stmt.Close()

	if err != nil {
		log.Fatal(err)
	}
	var data string
	var id int
	err = stmt.QueryRow(moduleName, chatID, variable).Scan(&id, &data)
	if err != nil {
		log.Fatal(err)
	}
	return Data{id, true, chatID, moduleName, variable, data}
}
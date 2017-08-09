package sql

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	//"fmt"
	//"reflect"
)

type DB struct {
	name string
	db   *sql.DB
}

func (dba DB) Emptydb() {
	db := dba.db

	sqlStmt := `
	create table IF NOT EXISTS Messages (id integer not null primary key, messageID integer not null, chatID integer not Null, userName text, content text);
	create table IF NOT EXISTS Data (id integer not null primary key, chatID integer not Null, moduleName text, variable text, data text);

	`
	_, err := db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}

}
func (dba DB) StoreMessage() {

}

var myDB DB

func Getdb() DB {
	if myDB.name != "" {
		return myDB
	}

	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		log.Fatal(err)
	}

	myDB = DB{"simpleDB", db}
	return myDB
}

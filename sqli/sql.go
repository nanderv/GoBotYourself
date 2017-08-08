package sql

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	//"fmt"
	//"reflect"
)
type DB struct{
	name string
	db *sql.DB
}
func (r DB) Perim() string {
    return r.name
}
func (dba DB) Emptydb(){
 	os.Remove("./foo.db")
	db := dba.db	

	sqlStmt := `
	create table msgs (id integer not null primary key, messageID integer not null, chatID integer not Null, userName text, content text);
	delete from msgs;
	`
	_, err := db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}

}
func Getdb() DB {
	os.Remove("./foo.db")

	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil{
		log.Fatal(err)
	}

	//dbb := DB{"simpleDB",db}
	//fooType := reflect.TypeOf(dbb)
	//for i := 0; i < fooType.NumMethod(); i++ {
	 //   method := fooType.Method(i)
	 //   fmt.Println(method.Name)
	//}
	return DB{"simpleDB",db}
}

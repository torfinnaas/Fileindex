// Taken from Linux Magazine #215
package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./files.db")
	if err != nil {
		panic(err)
	}
	
	_, err = db.Exec("CREATE TABLE files (path text, modified int, size int)")
	if err != nil {
		panic(err)
	}
}


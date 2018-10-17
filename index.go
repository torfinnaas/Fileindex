package main

import (
	"database/sql"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

type Walker struct {
	Db *sql.DB
}

func main() {
	if len(os.Args) != 2 {
		panic("usage: " + os.Args[0] + " start_dir")
	}
	root := os.Args[1]

	db, err := sql.Open("sqlite3", "./files.db")

	w := &Walker{
		Db: db,
	}

	err = filepath.Walk(root, w.Visit)
	checkErr(err)

	db.Close()
}

func (w *Walker) Visit(path string, f os.FileInfo, err error) error {
	stmt, err := w.Db.Prepare("INSERT INTO files VALUES (?,?,?)")
	checkErr(err)

	_, err = stmt.Exec(path, f.ModTime().Unix(), f.Size())
	checkErr(err)

	return nil
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

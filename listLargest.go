// listLargest.go
// Taken from Linux Magazine #215

package main

import (
	"database/sql"
	"os"
  "flag"
  "fmt"
	"strconv"
	_ "github.com/mattn/go-sqlite3"
)



func main() {
	max_files := flag.Int("max-files", 10, "max number of files")

	db, err := sql.Open("sqlite3", "./files.db")
  checkErr(err)

  flag.Parse()
  if len(flag.Args()) != 0 {
    panic("Usage: " + os.Args[0])
  }

  rows, err := db.Query("SELECT path, size FROM files ORDER BY size DESC LIMIT " + strconv.Itoa(*max_files))
	checkErr(err)

  var path string
  var size string

  for rows.Next() {
    err = rows.Scan(&path, &size)
    checkErr(err)
    fmt.Printf("%s %s\n", path, size)
  }
}


func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

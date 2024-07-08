package main

import (
	"database/sql"

	"github.com/mattn/so-sqlite3"
)


func main() {

  db, err := sql.Open(sqlite3, OMCdataSourceName strins)
  if err != nil {
    panic(err.Error())

    } 



}

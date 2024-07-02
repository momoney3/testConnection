package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3" // Import the SQLite driver
)

func main() {
	dataPath := "./OMCDateBase.db"
	db, err := sql.Open("sqlite3", dataPath)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Update tables
	sqlInsert := `
  INSERT INTO DataSources (Name)
  VALUES (?)`
	_, err = db.Exec(sqlInsert, "DS-Test")
	if err != nil {
		panic(err)
	}

	fmt.Println("Insertion successful")
}

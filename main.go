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

	// Create a new tabel
	t
	// Update tables
	sqlInsert := `
  INSERT INTO DataSources (Name)
  VALUES (?);`
	_, err = db.Exec(sqlInsert, "DSeTest")
	if err != nil {
		panic(err)
	}

	fmt.Println("Insertion successful")
}

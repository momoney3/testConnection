package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3" // Import the SQLite driver
)

func main() {
	dataPath := "./OMCDataBase.db"
	db, err := sql.Open("sqlite3", dataPath)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Alter table
	sqlAlter := `
  ALTER TABLE tt
  ADD COLUMN`
	_, err = db.Exec(sqlInsert)
	if err != nil {
		panic(err)
	}

	// Retrieve data
	sqlAlter := `
  SELECT * FROM DataSources 
  ADD COLUMN;`
	_, err = db.Exec(sqlInsert)
	if err != nil {
		panic(err)
	}

	// Insert data in tables
	sqlInsert := `
  INSERT INTO DataSources (ID, Name, Agency, DataType)
  VALUES ($1, $2, $3, $4);`
	_, err = db.Exec(sqlInsert, 4244, "DS-test-go", "goTest", "smb")
	if err != nil {
		panic(err)
	}

	// Select
	sqlSelect := `
  SELECT * FROM ConnetionsAndScan WHERE LastTest > 1;`
	_, err = db.Exec(sqlSelect)
	if err != nil {
		panic(err)
	}

	fmt.Println("Insertion successful")
}

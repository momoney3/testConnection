package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3" // Import the SQLite driver
)

func main() {
	dataPath := "./OMCDataBase.db"
	db, err := sql.Open("sqlite3", dataPath)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	fmt.Println("Insertion successful")
}

func displayDataSource(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM datasources ORDER BY name")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() { // iterate and fetch the records from result cursor
		var id int
		var name, agency, datatype string

		err := rows.Scan(&id, &name, &agency, &datatype)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("DataSource:", name, agency, datatype)
	}

	// check for any error encountered during iteration
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}

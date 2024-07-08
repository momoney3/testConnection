package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3" // Import the SQLite driver
	"google.golang.org/protobuf/compiler/protogen"
)

func main() {
	dataPath := "./OMCDataBase.db"
	db, err := sql.Open("sqlite3", dataPath)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Insert data in tables
	//sqlInsert := `
	//INSERT INTO DataSources (ID, Name, Agency, DataType)
	//VALUES (?, ?, ?, ?);`
	//_, err = db.Exec(sqlInsert, 655, "DS-rtest-gov2", "goTest", "smb")
	//if err != nil {
	//	panic(err)
	//}

	// Show data
	sqlSelect := `
  SELECT * FROM DataSources;`
	_, err = db.Exec(sqlSelect)
	if err != nil {
		panic(err)
	}
  
func displayDataSource(db *sql.DB) {
    rows, err := db.Query("select * from datasources order by name")
    if err != nil {
      log.Fatal(err)
    }
      defer rows.Close()

      for rows.Next() { // Iterate and fetch the records from result cursor}
        var ID int
        var Name, Agency, DataType string

        err := rows.Scan(&ID, &Name, &Agency, &DataType)
        if err != nil {
         log.Fatal(err)
       }
        log.Println("DataSources: ", Name, " ", Agency, " ", DataType)
    }
    // Check for any error encountered during iteration
    if err := rows.Err(); err != nil {
      log.Fatal(err)
    }

	fmt.Println("Insertion successful")
}

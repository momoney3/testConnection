package main

import (
	"database/sql"
)

// NOTE: Creat a function to dispaly the data.
// func displayData(db)
func main() {
	dataPath := "./OMCDataBase.db"
	db, err := sql.Open("sqlite3", dataPath)
	if err != nil {
		panic(err)
	}
	defer db.Close()



  rows, err 



  
}

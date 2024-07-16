package main

import (
	"database/sql"
)

// NOTE: Creat a function to dispaly the data.
// TODO: fix to make sour I am using the right sentix

func displayData(db *sql.DB) ([]dataSourceName, error) {
	rows, err := db.Query("SELECT * FROM DataSources ")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Slice to hold data from returned rows.
	var nameDS []namesDS

	// Loop through rows, using scan to assign colum data to struct files.
	for rows.Next() {
		var nds DataSources
		if err := rows.Scan(&alb.ID, &alb.Name, &alb.Agency,
			&alb.DataType); err != nil {
			return nameDS, err
		}
		namesDS = append(nameDS, alb)
	}
	if err = rows.Err(); err != nil {
		return namesDS, err
	}
	return pnamesDS, nil
}

func main() {
	dataPath := "./OMCDataBase.db"
	db, err := sql.Open("sqlite3", dataPath)
	if err != nil {
		panic(err)
	}
	defer db.Close()
}

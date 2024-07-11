package main

import (
	"database/sql"
)

// NOTE: Creat a function to dispaly the data.

func displayData(db *sql.DB) ([]dataSourceName, error) {
	rows, err := db.Query("SELECT * FROM album", "dataSource.db")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Slice to hold data from returned rows.
	var albums []Album

	// Loop through rows, using scan to assign colum data to struct files.
	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist,
			&alb.Price); err != nil {
			return albums, err
		}
		albums = append(albums, alb)
	}
	if err = rows.Err(); err != nil {
		return albums, err
	}
	return albums, nil
}

func main() {
	dataPath := "./OMCDataBase.db"
	db, err := sql.Open("sqlite3", dataPath)
	if err != nil {
		panic(err)
	}
	defer db.Close()
}

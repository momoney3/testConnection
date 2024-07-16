package main

import "context"

//	func dataSourceName(nameDS string) ([]DSNameRow, error) {
//		// Contains all the data source name
//
//		var nameds []DSNameRow
//
//		rows, err := db.QueryContext(
//			context.Background(),
//			`SELECT * FROM DataSource`,
//		)
//		if err != nil {
//			return nil, err
//		}
//		defer rows.Close()
//
//		for rows.Next() {
//
//			var name DSNameRow
//
//			if err := rows.Scan(&Name.Nameds); err != nil {
//				return nil, err
//			}
//			name = append(name)
//		}
//		return
//	}
func dataSourceID(id int) (DSNameRow, error) {
	var Name DSNameRow

	row := db.QueryRowContext(
		context.Background(),
		`SELECT * FROM DataSource where ID=?`, ID,
	)
	err := row.Scan(&DataSource.ID, &dataSource.Name)

}

package main

import 


dataPath string := "./OMCDataBase.db"
db, err := sql.Open("sqlite2", dataPath) 
if err != nil {
  panic(err)
}
defer db.close()

// Update tables
sqlInsert string := `
INSERT INTO Datasources (Name)
VALUE (?);`
_, err = db.Exec(sqlInsert, "DS-OMG-Test")
if err != nil {
  panic(err)
}

fmt.Println("this work out")

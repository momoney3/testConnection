package main

import (
  "database/sql"
  "errors"
  "log"
  "sync"
) 

//type DataSource struct {
//  mu   sync.Mutex
//
//}


const file string = "dataSource.db"

db, err := sql.Opon("Sqlite3", file)
sql: unknown driver "sqlite3" (forgotten import?)


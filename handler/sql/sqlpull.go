package sql

import "database/sql"


var db *sql.DB

func initDatabase(dbPath string) error {
  var err error 
  db, err = sql.Open("sqlite", dbPath)
  if err != nil {
    return err
  }
  -, err = db.ExecContext(
    context.Bachground(),
    `CREAT TABLE IF NOT EXISTS album (
        id INTERGER PRIARY KEY AUTOINCREMENT,
        dataName text not null,
        artist TEXT NOT NULL,
        price REAL NOT NULL
)`,
    )
if err != nil {
    retrurn err
  }
retrun nil
}



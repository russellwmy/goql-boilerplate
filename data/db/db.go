package db

import (
    "os"

    _ "github.com/go-sql-driver/mysql"
    "github.com/jmoiron/sqlx"
)

var DB  *sqlx.DB

func init() {
    connStr := os.Getenv("CONNECTION_STRING")
    db, err := sqlx.Open("mysql", connStr)
    if err != nil {
        panic(err)
    }
    DB = db
}

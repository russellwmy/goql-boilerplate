package data

import (
    _ "github.com/go-sql-driver/mysql"
    "github.com/jmoiron/sqlx"
)

func Connect(connectionString string) *sqlx.DB {
    db, err := sqlx.Open("mysql", connectionString)
    if err != nil {
        panic(err)
    }
    return db
}


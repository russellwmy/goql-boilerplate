package data

import (
    _ "github.com/go-sql-driver/mysql"
    "github.com/jmoiron/sqlx"
)

var DB *sqlx.DB;

func CheckErr(err error) {
    if err != nil {
        panic(err)
    }
}

func ConnectDB(connectionString string) {
    db, err := sqlx.Open("mysql", connectionString)
    CheckErr(err)
    DB = db
}


package graph

import (
    "./todo"

    "github.com/jmoiron/sqlx"
)

func Init (db *sqlx.DB) {
    todo.Init(db)
}

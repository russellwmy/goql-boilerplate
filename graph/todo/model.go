package todo

import (
    "fmt"
    "github.com/jmoiron/sqlx"
)

var tableName = "todos";

var schema = fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %s (
    id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    text text,
    done bool
);`, tableName )

type Todo struct {
    ID   int `json:"id" db:"id"`
    Text string `json:"text" db:"text"`
    Done bool   `json:"done" db:"done"`
}

func Init (db *sqlx.DB) {
    db.MustExec(schema);
}

func CreateTodo (db *sqlx.DB, todo Todo){
    query := fmt.Sprintf("INSERT INTO %s (text , done) VALUES (?, ?);", tableName)
    tx := db.MustBegin()
    tx.MustExec (query, todo.Text, todo.Done)
    tx.Commit()
}

func DeleteTodo (db *sqlx.DB, id int) {
    query := fmt.Sprintf("DELETE FROM %s WHERE id = ?;", tableName)
    db.MustExec(query, id);
}

func FindOne (db *sqlx.DB, id int) Todo {
    query := fmt.Sprintf("SELECT * FROM %s WHERE id = ?;", tableName)
    todo := Todo{}
    err := db.Get(&todo, query, id)
    if err != nil {
        fmt.Println(err)
        return Todo{}
    }
    return todo
}

func FindAll(db *sqlx.DB) []Todo {
    query := fmt.Sprintf("SELECT * FROM %s;", tableName)
    todos := []Todo{}
    err := db.Select(&todos, query)
    if err != nil {
        fmt.Println(err)
        return []Todo{}
    }
    return todos
}
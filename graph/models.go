package graph

import (
    "../data"

    "fmt"
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

func Init () {
    data.DB.MustExec(schema);
}

func CreateTodo (todo Todo){
    query := fmt.Sprintf("INSERT INTO %s (text , done) VALUES (?, ?);", tableName)
    tx := data.DB.MustBegin()
    tx.MustExec (query, todo.Text, todo.Done)
    tx.Commit()
}

func DeleteTodo (id int) {
    query := fmt.Sprintf("DELETE FROM %s WHERE id = ?;", tableName)
    data.DB.MustExec(query, id);
}

func FindOne (id int) Todo {
    query := fmt.Sprintf("SELECT * FROM %s WHERE id = ?;", tableName)
    todo := Todo{}
    err := data.DB.Get(&todo, query, id)
    if err != nil {
        fmt.Println(err)
        return Todo{}
    }
    return todo
}

func FindAll() []Todo {
    query := fmt.Sprintf("SELECT * FROM %s;", tableName)
    todos := []Todo{}
    err := data.DB.Select(&todos, query)
    if err != nil {
        fmt.Println(err)
        return []Todo{}
    }
    return todos
}
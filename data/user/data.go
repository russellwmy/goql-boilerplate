package user

import (
    "fmt"
    "golang.org/x/crypto/bcrypt"
    "github.com/russellwmy/goql-boilerplate/data/db"
)
var tableName = "users";

var schema = fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %s (
    id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name        varchar(255),
    email       varchar(255),
    password    varchar(255),
    created_at  timestamp NOT NULL DEFAULT current_timestamp,
    updated_at  timestamp NOT NULL DEFAULT now() on update now() 

);`, tableName )

func init() {
    db.DB.MustExec(schema);
}

func FindUserByID ( id int) *User {
    var user User
    query := fmt.Sprintf("SELECT id, name, email, updated_at FROM %s WHERE id=? LIMIT 1;", tableName)
    stmt, err := db.DB.Preparex(query)
    if err != nil {
        panic(err)
    }
    err = stmt.Get(&user, id)
    if err != nil {
        panic(err)
    }
    return &user
}

func CreateUser ( userInput *UserInput ) *User {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
    if err != nil {
        panic(err)
    }
    userInput.Password = string(hashedPassword)
    query := fmt.Sprintf("INSERT INTO %s (name, email, password) VALUES (:name, :email, :password)", tableName)
    result, _ := db.DB.NamedExec(query, &userInput)
    id, _:= result.LastInsertId()
    user := FindUserByID(int(id))
    fmt.Println(user)
    return user
}
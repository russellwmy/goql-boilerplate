package user

import (
    "time"
)

type User struct {
    ID        int     `db:"id"`
    Name      string  `db:"name"`
    Email     string  `db:"email"`
    UpdatedAt time.Time `db:"updated_at"`
}

type UserInput struct {
    Name        string
    Email       string
    Password    string
}
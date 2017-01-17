package data

import (
    "github.com/russellwmy/goql-boilerplate/data/user"
)

type Resolver struct{}

func (r *Resolver) CreateUser(args *struct { User  *user.UserInput }) *user.UserResolver {    
    User := user.CreateUser(args.User)
    return &user.UserResolver{User}
}

func (r *Resolver) User(args *struct{ ID int }) *user.UserResolver {
    User := user.FindUserByID(args.ID)
    return &user.UserResolver{User}
}
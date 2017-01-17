package user

import (
    "time"
)

type UserResolver struct {
    User *User
}

func (r *UserResolver) ID() int {
    return r.User.ID
}

func (r *UserResolver) Name() string {
    return r.User.Name
}

func (r *UserResolver) Email() string {
    return r.User.Email
}

func (r *UserResolver) UpdatedAt() *time.Time {
    return &r.User.UpdatedAt
}
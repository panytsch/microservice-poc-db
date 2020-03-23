package models

import "github.com/panytsch/microservice-poc-db/go/pkg/core"

type User struct {
	ID       uint64
	Name     string
	Password string
}

func (u *User) Create() {
	core.DB.Create(u)
}

func (*User) TableName() string {
	return "users"
}

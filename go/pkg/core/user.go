package core

import (
	"errors"
	"github.com/panytsch/microservice-poc-db/go/pkg/db"
)

const TestToken = "test"

type User struct {
	db.User
}

func CreateUser(name string, pass string) (*User, error) {
	user := new(User)
	user.User.Name = name
	user.User.Password = pass
	user.User.Create()
	if user.ID == 0 {
		return user, errors.New("user wasn't created")
	}
	return user, nil
}

package core

import (
	"errors"
	"github.com/panytsch/microservice-poc-db/go/pkg/db"
)

const TestToken = "test"

func CreateUser(name string, pass string) (*db.User, error) {
	user := new(db.User)
	user.Name = name
	user.Password = pass
	user.Create()
	if user.ID == 0 {
		return user, errors.New("user wasn't created")
	}
	return user, nil
}

func GetUserByID(id uint) (*db.User, error) {
	user := new(db.User)
	user.ID = id
	db.DB.Unscoped().Where(user).Find(user)
	if user.Name == "" {
		return nil, errors.New("user not found")
	}
	return user, nil
}

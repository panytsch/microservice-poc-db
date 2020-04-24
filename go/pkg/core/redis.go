package core

import (
	"github.com/panytsch/microservice-poc-db/go/pkg/db"
)

func GetUserByToken(token string) *db.User {
	user := new(db.User)
	if token == TestToken {
		db.DB.Unscoped().First(user)
	}
	return user
}

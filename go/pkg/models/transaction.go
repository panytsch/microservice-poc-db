package models

import (
	"github.com/jinzhu/gorm"
	"github.com/panytsch/microservice-poc-db/go/pkg/core"
)

type Transaction struct {
	gorm.Model
	Status int
	User   User
	UserID uint
}

func (t *Transaction) Create() {
	core.DB.Create(t)
}

package models

import (
	"github.com/jinzhu/gorm"
	"github.com/panytsch/microservice-poc-db/go/pkg/core"
)

type IModel interface {
	Create()
}

type Model struct {
	gorm.Model
}

func (m *Model) Create() {
	core.DB.Create(m)
}

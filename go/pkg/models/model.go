package models

import (
	"github.com/jinzhu/gorm"
)

type IModel interface {
	Create()
}

type Model struct {
	gorm.Model
}

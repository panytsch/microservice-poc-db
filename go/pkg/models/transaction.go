package models

import "github.com/jinzhu/gorm"

type Transaction struct {
	gorm.Model
	Status int
	User   User
	UserID uint
}

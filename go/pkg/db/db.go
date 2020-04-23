package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"log"
)

const (
	MaxIdleCon  = 3
	MaxOpenConn = 30
	server      = "mssql"
	port        = 1433
	user        = "sa"
	password    = "Qwerty1234"
	database    = "master"
)

func init() {
	connectDB()
	DB.LogMode(true)
}

var DB *gorm.DB

func connectDB() {
	var err error
	connString := fmt.Sprintf("sqlserver://%s:%s@%s:%v?database=%s",
		user, password, server, port, database)
	DB, err = gorm.Open("mssql", connString)
	if err != nil {
		log.Fatalf("Got error: %v", err)
	}
	DB.DB().SetMaxIdleConns(MaxIdleCon)
	DB.DB().SetMaxOpenConns(MaxOpenConn)
}

type Model struct {
	gorm.Model
	CreatedAt struct{} `gorm:"-" sql:"-"`
	UpdatedAt struct{} `gorm:"-" sql:"-"`
	DeletedAt struct{} `gorm:"-" sql:"-"`
}

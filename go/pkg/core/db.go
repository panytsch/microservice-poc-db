package core

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"log"
)

const (
	MaxIdleCon  = 3
	MaxOpenConn = 30
	server      = "localhost"
	port        = 1433
	user        = "sa"
	password    = "Qwerty1234"
	database    = "master"
)

var DB *gorm.DB

func ConnectDB() {
	connString := fmt.Sprintf("sqlserver://%s:%s@%s:%v?database=%s",
		user, password, server, port, database)
	DB, err := gorm.Open("mssql", connString)
	if err != nil {
		log.Fatalf("Got error: %v", err)
	}
	DB.DB().SetMaxIdleConns(MaxIdleCon)
	DB.DB().SetMaxOpenConns(MaxOpenConn)
}

func CloseDB() {
	_ = DB.Close()
}

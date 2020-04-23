package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"log"
	"os"
	"strconv"
)

func init() {
	connectDB()
	DB.LogMode(true)
}

var DB *gorm.DB

func connectDB() {
	var err error
	connString := getConnectionString()
	DB, err = gorm.Open("mssql", connString)
	if err != nil {
		log.Fatalf("Got error: %v", err)
	}
	MaxIdleCon, _ := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNECTIONS"))
	MaxOpenConn, _ := strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONNECTIONS"))
	DB.DB().SetMaxIdleConns(MaxIdleCon)
	DB.DB().SetMaxOpenConns(MaxOpenConn)
}

func getConnectionString() string {
	server := os.Getenv("DB_SERVER")
	user := os.Getenv("DB_USER")
	password := os.Getenv("MSSQL_SA_PASSWORD")
	database := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	return fmt.Sprintf("sqlserver://%s:%s@%s:%v?database=%s",
		user, password, server, port, database)
}

type Model struct {
	gorm.Model
	CreatedAt struct{} `gorm:"-" sql:"-"`
	UpdatedAt struct{} `gorm:"-" sql:"-"`
	DeletedAt struct{} `gorm:"-" sql:"-"`
}

package db

func init()  {
	
}

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"github.com/panytsch/go_poc/mssql/go/pkg/procedures/test"
	"log"
)

var server = "localhost"
var port = 1433
var user = "sa"
var password = "Qwerty1234"
var database = "master"

func connect() {
	connString := fmt.Sprintf("sqlserver://%s:%s@%s:%v?database=%s",
		user, password, server, port, database)
	db, err := gorm.Open("mssql", connString)
	if err != nil {
		log.Fatalf("Got error: %v", err)
	}
	defer db.Close()
	
	// sp := test.NewTwoDataSetsProcedure(db)
	// res := sp.Run()
	// log.Printf("sp result %v", res)
}

func baseUsage(db *sql.DB) {
	rows, err := db.Query("exec twoDataSets")
	if err != nil {
		log.Fatalf("after rows %v", err)
	}
	defer rows.Close()

	rows.Next()
	str := &firstDataSet{}
	err = rows.Scan(&str.One, &str.Two, &str.Three)
	if err != nil {
		log.Fatalf("while scan %v", err)
	}
	log.Printf("first result %v", str)
	rows.Next()

	if !rows.NextResultSet() {
		log.Println("no sets left")
		return
	}

	rows.Next()
	var returnCode int
	err = rows.Scan(&returnCode)
	if err != nil {
		log.Fatalf("while second scan %v", err)
	}
	log.Printf("result %v", returnCode)
}

func usingGorm(db *gorm.DB) {
	rows, err := db.Raw("exec twoDataSets").Rows()

	if err != nil {
		log.Fatal("after exec", err, rows)
	}
	//rows.NextResultSet()
	rows.Next()
	log.Printf("got row %v", rows)
	str := &firstDataSet{}
	//err = rows.Scan(&str.One, &str.Two, &str.Three)
	err = db.ScanRows(rows, str)
	if err != nil {
		log.Printf("after scan error %v", err)
	} else {
		log.Printf("got struct %v", str)
	}
	rows.Next()

	log.Printf("exist %v", rows.NextResultSet())
	log.Printf("rows %v", rows)
}

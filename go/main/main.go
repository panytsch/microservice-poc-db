package main

import (
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"github.com/panytsch/microservice-poc-db/go/pkg/core"
	"github.com/panytsch/microservice-poc-db/go/pkg/procedures/test"
	"github.com/panytsch/microservice-poc-db/go/routes"
	"log"
)

func main() {
	core.ConnectDB()
	defer core.CloseDB()

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/user/create", routes.CreateNewUserHandler)

	sp := test.NewTwoDataSetsProcedure(core.DB)
	res := sp.Run()
	log.Printf("sp result %v", res)
}

package main

import (
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"github.com/panytsch/microservice-poc-db/go/pkg/core"
	"github.com/panytsch/microservice-poc-db/go/routes"
	"log"
	"net/http"
	"time"
)

func main() {
	core.ConnectDB()
	defer core.CloseDB()

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/user/create", routes.CreateNewUserHandler)

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:80",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

	//sp := test.NewTwoDataSetsProcedure(core.DB)
	//res := sp.Run()
	//log.Printf("sp result %v", res)
}

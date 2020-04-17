// Package classification DB microservice REST API.
//
// the purpose of this application is to provide an application
// that is using plain go code to define an API
//
// This should demonstrate all the possible comment annotations
// that are available to turn go code into a fully compliant swagger 2.0 spec
//
// Terms Of Service:
//
// there are no TOS at this moment, use at your own risk we take no responsibility
//
//     Schemes: http
//     Host: localhost
//     BasePath: /rest/v1
//     Version: 1.0.0
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: Roman Panasiuk<gfyroman@gmail.com>
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package server

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"github.com/panytsch/microservice-poc-db/go/rest/routes"
	"log"
	"net/http"
	"sync"
	"time"
)

func RunRestServer(wg *sync.WaitGroup) {
	router := mux.NewRouter()
	restV1Router := router.PathPrefix("/rest/v1").Subrouter()
	collectRoutes(restV1Router)
	setDocsRoute(router)

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:80",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("SERVER LISTEN ERROR: %v\n", srv.ListenAndServe())
	wg.Done()
}

func collectRoutes(router *mux.Router) {
	router.HandleFunc("/user", routes.CreateNewUserHandler)
}

func setDocsRoute(router *mux.Router) {
	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)
	router.Handle("/docs", sh)
	log.Println(http.Dir("."))
	router.Handle("/swagger.yaml", http.FileServer(http.Dir(".")))
}

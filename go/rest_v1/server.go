// Package classification DB microservice REST API.
//
// the purpose of this application is to provide an application
// to work with DB like with service
//
//     Schemes: http
//     Version: 1.0.0
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: Roman Panasiuk<gfyroman@gmail.com>
//
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package rest_v1

import (
	"encoding/json"
	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"sync"
	"time"
)

func RunRestServer(wg *sync.WaitGroup) {
	router := mux.NewRouter()
	restV1Router := router.PathPrefix("/rest/v1").Subrouter()
	collectRestRoutes(restV1Router)
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

func collectRestRoutes(router *mux.Router) {
	router.Use(restHeaders)
	usersRouter := router.PathPrefix("/users").Subrouter()
	usersRouter.HandleFunc("", CreateNewUserHandler).Methods(http.MethodPost)
	usersRouter.HandleFunc("/get", GetUserHandler).Methods(http.MethodPost)

	transactionsRouter := router.PathPrefix("/transactions").Subrouter()
	transactionsRouter.HandleFunc("", MakeTransactionHandler).Methods(http.MethodPost)
	transactionsRouter.HandleFunc("", GetTransactionsHandler).Methods(http.MethodGet).Queries(
		"limit", "{limit}",
		"offset", "{offset}",
	)
	transactionsRouter.HandleFunc("/{TransactionID}", GetTransactionHandler).Methods(http.MethodGet)
	transactionsRouter.Use(checkUserTokenMiddleware)
}

func setDocsRoute(router *mux.Router) {
	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)
	router.Handle("/docs", sh)
	router.Handle("/swagger.yaml", http.FileServer(http.Dir(".")))
}

// send Json response
func SendJSON(i interface{}, w http.ResponseWriter) error {
	e := json.NewEncoder(w)
	return e.Encode(i)
}

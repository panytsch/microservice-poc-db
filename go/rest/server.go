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
package rest

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
	router.HandleFunc("/user", CreateNewUserHandler).Methods(http.MethodPost)
	router.HandleFunc("/user/get", GetUserHandler).Methods(http.MethodPost)
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

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
	"os"
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
		Addr:    ":" + os.Getenv("APP_PORT"),
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
	usersRouter.HandleFunc("", GetUserByTokenHandler).Methods(http.MethodGet)
	usersRouter.HandleFunc("/get", GetUserHandler).Methods(http.MethodPost)

	PaymentsRouter := router.PathPrefix("/Payments").Subrouter()
	PaymentsRouter.HandleFunc("", MakePaymentHandler).Methods(http.MethodPost)
	PaymentsRouter.HandleFunc("", GetPaymentsHandler).Methods(http.MethodGet).Queries(
		"limit", "{limit:[0-9]+}",
		"offset", "{offset:[0-9]+}",
	)
	PaymentsRouter.HandleFunc("/{PaymentID}", GetPaymentHandler).Methods(http.MethodGet)

	refundsRouter := router.PathPrefix("/refunds").Subrouter()
	refundsRouter.HandleFunc("", MakeRefundHandler).Methods(http.MethodPost)
	refundsRouter.HandleFunc("", GetRefundsHandler).Methods(http.MethodGet).Queries(
		"limit", "{limit:[0-9]+}",
		"offset", "{offset:[0-9]+}",
	)
	refundsRouter.HandleFunc("/{RefundID}", GetRefundHandler).Methods(http.MethodGet)

	refundsRouter.Use(checkUserTokenMiddleware)
	PaymentsRouter.Use(checkUserTokenMiddleware)
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

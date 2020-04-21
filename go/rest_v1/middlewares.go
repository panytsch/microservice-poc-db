package rest_v1

import (
	"github.com/panytsch/microservice-poc-db/go/pkg/core"
	"log"
	"net/http"
)

func checkUserTokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token != core.TestToken {
			w.WriteHeader(http.StatusUnauthorized)
			log.Printf("token %v didn't pass verification\n", token)
			_ = SendJSON(ErrorResponse{
				Message: "Wrong Token provided",
				Code:    WrongToken,
			}, w)
		} else {
			log.Printf("token %v passed verification\n", token)
			next.ServeHTTP(w, r)
		}
	})
}

func restHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

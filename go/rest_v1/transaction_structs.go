package rest_v1

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/panytsch/microservice-poc-db/go/pkg/db"
	"net/http"
	"strconv"
)

//region createTransaction

//swagger:parameters createTransaction
type SwaggerMakeTransactionRequest struct {
	//User token
	//in:header
	Authorization string

	//in:body
	Body MakeTransactionRequest
}

type MakeTransactionRequest struct {
	//Amount of transaction
	//required:true
	Amount db.TransactionAmount
}

func (req *MakeTransactionRequest) capture(r *http.Request) error {
	decoder := json.NewDecoder(r.Body)
	er := decoder.Decode(&req)
	if er != nil {
		return errors.New("error while decode request")
	}
	return nil
}

//swagger:response createTransaction
type SwaggerMakeTransactionResponse struct {
	//in:body
	Body MakeTransactionResponse
}

type MakeTransactionResponse struct {
	ID     uint
	Status db.TransactionStatus
	Amount db.TransactionAmount
}

//endregion
//region getTransaction

//swagger:parameters getTransaction
type SwaggerGetTransactionRequest struct {
	//User token
	//in:header
	//required:true
	Authorization string

	//in:path
	//required:true
	TransactionID uint
}

//swagger:response getTransaction
type SwaggerGetTransactionResponse struct {
	//in:body
	Body GetTransactionResponse
}

type GetTransactionResponse struct {
	ID     uint
	Status db.TransactionStatus
	Amount db.TransactionAmount
}

//endregion
//region getTransactions

//swagger:parameters getTransactions
type SwaggerGetTransactionsRequest struct {
	//User token
	//in:header
	//required:true
	Authorization string

	//in:query
	//required:true
	limit uint

	//default value is 0
	//in:query
	//required:true
	offset uint
}

type getTransactionsRequest struct {
	limit  uint
	offset uint
}

func (req *getTransactionsRequest) isValid() bool {
	return req.limit != 0
}

func (req *getTransactionsRequest) capture(r *http.Request) *getTransactionsRequest {
	vars := mux.Vars(r)
	parsedUint, _ := strconv.ParseUint(vars["limit"], 10, 64)
	req.limit = uint(parsedUint)
	parsedUint, _ = strconv.ParseUint(vars["offset"], 10, 64)
	req.offset = uint(parsedUint)
	return req
}

//swagger:response getTransactions
type SwaggerGetTransactionsResponse struct {
	//in:body
	Body []GetTransactionsResponse
}

type GetTransactionsResponse struct {
	ID     uint
	Status db.TransactionStatus
	Amount db.TransactionAmount
}

//endregion

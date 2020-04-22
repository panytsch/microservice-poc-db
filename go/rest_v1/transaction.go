package rest_v1

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/panytsch/microservice-poc-db/go/pkg/core"
	"github.com/panytsch/microservice-poc-db/go/pkg/db"
	"net/http"
	"strconv"
)

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
	ID     uint64
	Status db.TransactionStatus
	Amount db.TransactionAmount
}

// swagger:route POST /rest/v1/transactions transaction createTransaction
//
// Create new Transaction
//     Responses:
//       201: createTransaction
//       400: errorResponse
//       401: errorResponse
func MakeTransactionHandler(w http.ResponseWriter, r *http.Request) {
	user := core.GetUserByToken(r.Header.Get("Authorization"))
	if user.ID == 0 {
		sendBadResponse(w, "User's token probably wrong. User not found", WrongToken)
		return
	}
	req := new(MakeTransactionRequest)
	err := req.capture(r)
	if err != nil {
		sendBadResponse(w, err.Error(), ParsingRequestError)
		return
	}
	transaction := core.CreateTransaction(user.ID, req.Amount)
	if !transaction.IsSuccess() {
		sendBadResponse(w, "Can't create transaction", GeneralBad)
		return
	}
	w.WriteHeader(http.StatusCreated)
	_ = SendJSON(MakeTransactionResponse{
		ID:     transaction.Result.ID,
		Status: transaction.Result.Status,
		Amount: transaction.Result.Amount,
	}, w)
}

func sendBadResponse(w http.ResponseWriter, message string, internalCode ErrorCode) {
	w.WriteHeader(http.StatusBadRequest)
	_ = SendJSON(ErrorResponse{
		Message: message,
		Code:    internalCode,
	}, w)
}

// swagger:route GET /rest/v1/transactions/{TransactionID} transaction getTransaction
//
// Get one Transaction
//     Responses:
//       200: getTransaction
//       400: errorResponse
//       401: errorResponse
func GetTransactionHandler(w http.ResponseWriter, r *http.Request) {
	transactionID, _ := strconv.ParseUint(mux.Vars(r)["TransactionID"], 10, 64)
	user := core.GetUserByToken(r.Header.Get("Authorization"))
	transaction, err := core.GetTransactionByIDAndUserID(transactionID, user.ID)
	if err != nil { //not found
		w.WriteHeader(http.StatusBadRequest)
		sendBadResponse(w, err.Error(), NoDataFound)
	} else {
		w.WriteHeader(http.StatusOK)
		_ = SendJSON(GetTransactionResponse{
			ID:     transaction.ID,
			Status: transaction.Status,
			Amount: transaction.Amount,
		}, w)
	}
}

//swagger:parameters getTransaction
type SwaggerGetTransactionRequest struct {
	//User token
	//in:header
	//required:true
	Authorization string

	//in:path
	//required:true
	TransactionID uint64
}

//swagger:response getTransaction
type SwaggerGetTransactionResponse struct {
	//in:body
	Body GetTransactionResponse
}

type GetTransactionResponse struct {
	ID     uint64
	Status db.TransactionStatus
	Amount db.TransactionAmount
}

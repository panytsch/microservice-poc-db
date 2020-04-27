package rest_v1

import (
	"github.com/gorilla/mux"
	"github.com/panytsch/microservice-poc-db/go/pkg/core"
	"net/http"
	"strconv"
)

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
		w.WriteHeader(http.StatusBadRequest)
		sendBadResponse(w, "User's token probably wrong. User not found", WrongToken)
		return
	}
	req := new(MakeTransactionRequest)
	err := req.capture(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		sendBadResponse(w, err.Error(), ParsingRequestError)
		return
	}
	transaction := core.CreateTransaction(user.ID, req.Amount)
	if !transaction.IsSuccess() {
		w.WriteHeader(http.StatusBadRequest)
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

// swagger:route GET /rest/v1/transactions/{TransactionID} transaction getTransaction
//
// Get one Transaction
//     Responses:
//       200: getTransaction
//       400: errorResponse
//       401: errorResponse
func GetTransactionHandler(w http.ResponseWriter, r *http.Request) {
	transactionID, _ := strconv.ParseUint(mux.Vars(r)["TransactionID"], 10, 64)
	if transactionID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		sendBadResponse(w, "transaction id weren't provided", NoDataFound)
		return
	}
	user := core.GetUserByToken(r.Header.Get("Authorization"))
	transaction, err := core.GetTransactionByIDAndUserID(uint(transactionID), user.ID)
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

// swagger:route GET /rest/v1/transactions transaction getTransactions
//
// Get few Transaction
//     Responses:
//       200: getTransactions
//       400: errorResponse
//       401: errorResponse
func GetTransactionsHandler(w http.ResponseWriter, r *http.Request) {
	req := new(getTransactionsRequest).capture(r)
	if !req.isValid() {
		w.WriteHeader(http.StatusBadRequest)
		sendBadResponse(w, "mandatory data weren't provided", BadRequest)
		return
	}
	user := core.GetUserByToken(r.Header.Get("Authorization"))
	transactions, err := core.GetLastTransactions(req.limit, req.offset, user.ID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		sendBadResponse(w, err.Error(), GeneralBad)
	}
	res := make([]GetTransactionsResponse, 0)
	for _, t := range transactions {
		res = append(res, GetTransactionsResponse{
			ID:     t.ID,
			Status: t.Status,
			Amount: t.Amount,
		})
	}
	w.WriteHeader(http.StatusOK)
	_ = SendJSON(res, w)
}

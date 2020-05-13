package rest_v1

import (
	"github.com/gorilla/mux"
	"github.com/panytsch/microservice-poc-db/go/pkg/core"
	"net/http"
	"strconv"
)

// swagger:route POST /rest/v1/Payments Payment createPayment
//
// Create new Payment
//     Responses:
//       201: createPayment
//       400: errorResponse
//       401: errorResponse
func MakePaymentHandler(w http.ResponseWriter, r *http.Request) {
	user := core.GetUserByToken(r.Header.Get("Authorization"))
	if user.ID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		sendBadResponse(w, "User's token probably wrong. User not found", WrongToken)
		return
	}
	req := new(MakePaymentRequest)
	err := req.capture(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		sendBadResponse(w, err.Error(), ParsingRequestError)
		return
	}
	Payment := core.CreatePayment(user.ID, req.Amount)
	if !Payment.IsSuccess() {
		w.WriteHeader(http.StatusBadRequest)
		sendBadResponse(w, "Can't create Payment", GeneralBad)
		return
	}
	w.WriteHeader(http.StatusCreated)
	_ = SendJSON(MakePaymentResponse{
		ID:     Payment.Result.ID,
		Status: Payment.Result.Status,
		Amount: Payment.Result.Amount,
	}, w)
}

// swagger:route GET /rest/v1/Payments/{PaymentID} Payment getPayment
//
// Get one Payment
//     Responses:
//       200: getPayment
//       400: errorResponse
//       401: errorResponse
func GetPaymentHandler(w http.ResponseWriter, r *http.Request) {
	PaymentID, _ := strconv.ParseUint(mux.Vars(r)["PaymentID"], 10, 64)
	if PaymentID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		sendBadResponse(w, "Payment id weren't provided", NoDataFound)
		return
	}
	user := core.GetUserByToken(r.Header.Get("Authorization"))
	Payment, err := core.GetPaymentByIDAndUserID(uint(PaymentID), user.ID)
	if err != nil { //not found
		w.WriteHeader(http.StatusBadRequest)
		sendBadResponse(w, err.Error(), NoDataFound)
	} else {
		w.WriteHeader(http.StatusOK)
		_ = SendJSON(GetPaymentResponse{
			ID:     Payment.ID,
			Status: Payment.Status,
			Amount: Payment.Amount,
			UserID: Payment.UserID,
		}, w)
	}
}

// swagger:route GET /rest/v1/Payments Payment getPayments
//
// Get few Payment
//     Responses:
//       200: getPayments
//       400: errorResponse
//       401: errorResponse
func GetPaymentsHandler(w http.ResponseWriter, r *http.Request) {
	req := new(getPaymentsRequest).capture(r)
	if !req.isValid() {
		w.WriteHeader(http.StatusBadRequest)
		sendBadResponse(w, "mandatory data weren't provided", BadRequest)
		return
	}
	user := core.GetUserByToken(r.Header.Get("Authorization"))
	Payments, err := core.GetLastPayments(req.limit, req.offset, user.ID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		sendBadResponse(w, err.Error(), GeneralBad)
	}
	res := make([]GetPaymentsResponse, 0)
	for _, t := range Payments {
		res = append(res, GetPaymentsResponse{
			ID:     t.ID,
			Status: t.Status,
			Amount: t.Amount,
		})
	}
	w.WriteHeader(http.StatusOK)
	_ = SendJSON(res, w)
}


// swagger:route GET /rest/v1/Payments/{PaymentID} Payment getPayment
//
// Get one Payment
//     Responses:
//       200: getPayment
//       400: errorResponse
//       401: errorResponse
func UpdatePaymentHandler(w http.ResponseWriter, r *http.Request) {
	PaymentID, _ := strconv.ParseUint(mux.Vars(r)["PaymentID"], 10, 64)
	if PaymentID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		sendBadResponse(w, "Payment id weren't provided", NoDataFound)
		return
	}
	user := core.GetUserByToken(r.Header.Get("Authorization"))
	Payment, err := core.GetPaymentByIDAndUserID(uint(PaymentID), user.ID)
	if err != nil { //not found
		w.WriteHeader(http.StatusBadRequest)
		sendBadResponse(w, err.Error(), NoDataFound)
		return
	}
}
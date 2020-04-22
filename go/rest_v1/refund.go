package rest_v1

import (
	"github.com/gorilla/mux"
	"github.com/panytsch/microservice-poc-db/go/pkg/core"
	"net/http"
	"strconv"
)

// swagger:route GET /rest/v1/refunds refund getRefunds
//
// Get few Refunds
//     Responses:
//       200: getRefunds
//       400: errorResponse
//       401: errorResponse
func GetRefundsHandler(w http.ResponseWriter, r *http.Request) {
	req := new(getRefundsRequest).capture(r)
	if !req.isValid() {
		w.WriteHeader(http.StatusBadRequest)
		sendBadResponse(w, "mandatory data weren'refund provided", BadRequest)
		return
	}
	user := core.GetUserByToken(r.Header.Get("Authorization"))
	refunds, err := core.GetLastRefunds(req.limit, req.offset, user.ID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		sendBadResponse(w, err.Error(), GeneralBad)
	}
	res := make([]GetRefundsResponse, 0)
	for _, refund := range refunds {
		res = append(res, GetRefundsResponse{
			ID:     refund.ID,
			Status: refund.Status,
			Amount: refund.Amount,
		})
	}
	w.WriteHeader(http.StatusOK)
	_ = SendJSON(res, w)
}

// swagger:route GET /rest/v1/refunds/{RefundID} refund getRefund
//
// Get one Refund
//     Responses:
//       200: getRefund
//       400: errorResponse
//       401: errorResponse
func GetRefundHandler(w http.ResponseWriter, r *http.Request) {
	refundID, _ := strconv.ParseUint(mux.Vars(r)["RefundID"], 10, 64)
	if refundID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		sendBadResponse(w, "refund id weren't provided", NoDataFound)
		return
	}
	user := core.GetUserByToken(r.Header.Get("Authorization"))
	refund, err := core.GetRefundByIDAndUserID(uint(refundID), user.ID)
	if err != nil { //not found
		w.WriteHeader(http.StatusBadRequest)
		sendBadResponse(w, err.Error(), NoDataFound)
	} else {
		w.WriteHeader(http.StatusOK)
		_ = SendJSON(GetRefundResponse{
			ID:     refund.ID,
			Status: refund.Status,
			Amount: refund.Amount,
		}, w)
	}
}

// swagger:route POST /rest/v1/refunds refund createRefund
//
// Create new Refund
//     Responses:
//       201: createRefund
//       400: errorResponse
//       401: errorResponse
func MakeRefundHandler(w http.ResponseWriter, r *http.Request) {
	user := core.GetUserByToken(r.Header.Get("Authorization"))
	if user.ID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		sendBadResponse(w, "User's token probably wrong. User not found", WrongToken)
		return
	}
	req := new(MakeRefundRequest)
	err := req.capture(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		sendBadResponse(w, err.Error(), ParsingRequestError)
		return
	}
	refund := core.MakeRefund(req.Amount, user.ID)
	if refund.ID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		sendBadResponse(w, "Can't create refund", GeneralBad)
		return
	}
	w.WriteHeader(http.StatusCreated)
	_ = SendJSON(MakeRefundResponse{
		ID:     refund.ID,
		Status: refund.Status,
		Amount: refund.Amount,
	}, w)
}

package rest_v1

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/panytsch/microservice-poc-db/go/pkg/db"
	"net/http"
	"strconv"
)

//region createPayment

//swagger:parameters createPayment
type SwaggerMakePaymentRequest struct {
	//User token
	//in:header
	Authorization string

	//in:body
	Body MakePaymentRequest
}

type MakePaymentRequest struct {
	//Amount of Payment
	//required:true
	Amount db.PaymentAmount
}

func (req *MakePaymentRequest) capture(r *http.Request) error {
	decoder := json.NewDecoder(r.Body)
	er := decoder.Decode(&req)
	if er != nil {
		return errors.New("error while decode request")
	}
	return nil
}

//swagger:response createPayment
type SwaggerMakePaymentResponse struct {
	//in:body
	Body MakePaymentResponse
}

type MakePaymentResponse struct {
	ID     uint
	Status db.PaymentStatus
	Amount db.PaymentAmount
}

//endregion
//region getPayment

//swagger:parameters getPayment
type SwaggerGetPaymentRequest struct {
	//User token
	//in:header
	//required:true
	Authorization string

	//in:path
	//required:true
	PaymentID uint
}

//swagger:response getPayment
type SwaggerGetPaymentResponse struct {
	//in:body
	Body GetPaymentResponse
}

type GetPaymentResponse struct {
	ID     uint
	Status db.PaymentStatus
	Amount db.PaymentAmount
	UserID uint
}

//endregion
//region getPayments

//swagger:parameters getPayments
type SwaggerGetPaymentsRequest struct {
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

type getPaymentsRequest struct {
	limit  uint
	offset uint
}

func (req *getPaymentsRequest) isValid() bool {
	return req.limit != 0
}

func (req *getPaymentsRequest) capture(r *http.Request) *getPaymentsRequest {
	vars := mux.Vars(r)
	parsedUint, _ := strconv.ParseUint(vars["limit"], 10, 64)
	req.limit = uint(parsedUint)
	parsedUint, _ = strconv.ParseUint(vars["offset"], 10, 64)
	req.offset = uint(parsedUint)
	return req
}

//swagger:response getPayments
type SwaggerGetPaymentsResponse struct {
	//in:body
	Body []GetPaymentsResponse
}

type GetPaymentsResponse struct {
	ID     uint
	Status db.PaymentStatus
	Amount db.PaymentAmount
}

//endregion
//region updatePayment

//swagger:parameters updatePayment
type SwaggerUpdatePaymentStatusRequest struct {
	//required:true
	//in:path
	PaymentID uint

	//in:body
	Body updatePaymentStatusRequest
}

type updatePaymentStatusRequest struct {
	Status db.PaymentStatus
	Amount db.PaymentAmount
	UserID uint
}

//swagger:response updatePayment
type SwaggerUpdatePaymentStatusResponse struct {
	//in:body
	Body updatePaymentResponse
}

type updatePaymentResponse struct {
	GetPaymentResponse
}
//endregion
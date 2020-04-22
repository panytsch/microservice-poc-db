package rest_v1

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/panytsch/microservice-poc-db/go/pkg/db"
	"net/http"
	"strconv"
)

//region get-refunds

//swagger:parameters getRefunds
type SwaggerGetRefundsRequest struct {
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

type getRefundsRequest struct {
	limit  uint
	offset uint
}

func (req *getRefundsRequest) isValid() bool {
	return req.limit != 0
}

func (req *getRefundsRequest) capture(r *http.Request) *getRefundsRequest {
	vars := mux.Vars(r)
	parsedUint, _ := strconv.ParseUint(vars["limit"], 10, 64)
	req.limit = uint(parsedUint)
	parsedUint, _ = strconv.ParseUint(vars["offset"], 10, 64)
	req.offset = uint(parsedUint)
	return req
}

//swagger:response getRefunds
type SwaggerGetRefundsResponse struct {
	//in:body
	Body []GetRefundsResponse
}

type GetRefundsResponse struct {
	ID     uint
	Status db.RefundStatus
	Amount db.RefundAmount
}

//endregion

//region get-refund

//swagger:parameters getRefund
type SwaggerGetRefundRequest struct {
	//User token
	//in:header
	//required:true
	Authorization string

	//in:path
	//required:true
	RefundID uint
}

//swagger:response getRefund
type SwaggerGetRefundResponse struct {
	//in:body
	Body GetRefundResponse
}

type GetRefundResponse struct {
	ID     uint
	Status db.RefundStatus
	Amount db.RefundAmount
}

//endregion

//region make-refund

//swagger:parameters createRefund
type SwaggerMakeRefundRequest struct {
	//User token
	//in:header
	Authorization string

	//in:body
	Body MakeRefundRequest
}

type MakeRefundRequest struct {
	//Amount of refund
	//required:true
	Amount db.RefundAmount
}

func (req *MakeRefundRequest) capture(r *http.Request) error {
	decoder := json.NewDecoder(r.Body)
	er := decoder.Decode(&req)
	if er != nil {
		return errors.New("error while decode request")
	}
	return nil
}

//swagger:response createRefund
type SwaggerMakeRefundResponse struct {
	//in:body
	Body MakeRefundResponse
}

type MakeRefundResponse struct {
	ID     uint
	Status db.RefundStatus
	Amount db.RefundAmount
}

//endregion

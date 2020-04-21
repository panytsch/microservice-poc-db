package rest_v1

type ErrorCode uint16

const (
	GeneralBad ErrorCode = iota + 1
	NoDataFound
	BadRequest
	WrongToken
	ParsingRequestError
)

type NoDataFoundResponse struct {
	ErrorResponse
}

//swagger:response noDataFound
type SwaggerNoDataFoundResponse struct {
	//in:body
	Body NoDataFoundResponse
}

type ErrorResponse struct {
	// Error description
	Message string
	// Error Code
	Code ErrorCode
}

//swagger:response errorResponse
type SwaggerErrorResponse struct {
	//in:body
	Body ErrorResponse
}

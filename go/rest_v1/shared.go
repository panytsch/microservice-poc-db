package rest_v1

import "net/http"

func sendBadResponse(w http.ResponseWriter, message string, internalCode ErrorCode) {
	_ = SendJSON(ErrorResponse{
		Message: message,
		Code:    internalCode,
	}, w)
}

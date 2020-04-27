package rest_v1

import (
	"encoding/json"
	"github.com/panytsch/microservice-poc-db/go/pkg/core"
	"github.com/panytsch/microservice-poc-db/go/pkg/db"
	"log"
	"net/http"
)

// swagger:route POST /rest/v1/users user createUser
//
// Create new user
//     Responses:
//       201: createUser
//       400: errorResponse
func CreateNewUserHandler(w http.ResponseWriter, r *http.Request) {
	req := CreateNewUserRequest{}
	decoder := json.NewDecoder(r.Body)
	er := decoder.Decode(&req)
	if er != nil {
		log.Printf("error while decode request: %v\n", er)
		w.WriteHeader(http.StatusBadRequest)
		_ = SendJSON(ErrorResponse{
			Message: "Bad request provided",
			Code:    GeneralBad,
		}, w)
		return
	}
	user, err := core.CreateUser(req.Name, req.Password)
	if err != nil {
		log.Printf("error while saving user: %v\n", er)
		w.WriteHeader(http.StatusBadRequest)
		_ = SendJSON(ErrorResponse{
			Message: "Error while user saving",
			Code:    BadRequest,
		}, w)
		return
	}
	w.WriteHeader(http.StatusCreated)
	_ = SendJSON(CreateNewUserResponse{
		ID:   user.ID,
		Name: user.Name,
	}, w)
}

// swagger:route POST /rest/v1//users/get user getUser
//
// Get user
//     Responses:
//       200: getUser
//       400: noDataFound
//       400: errorResponse
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	req := GetUserRequest{}
	decoder := json.NewDecoder(r.Body)
	er := decoder.Decode(&req)
	if er != nil {
		log.Printf("error while decode request: %v\n", er)
		w.WriteHeader(http.StatusBadRequest)
		_ = SendJSON(ErrorResponse{
			Message: "Bad request provided",
			Code:    GeneralBad,
		}, w)
		return
	}
	user := new(db.User).FindByNameAndPass(req.Name, req.Password)

	if user.ID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		response := new(NoDataFoundResponse)
		response.Code = NoDataFound
		response.Message = "User not found"
		_ = SendJSON(response, w)
	} else {
		w.WriteHeader(http.StatusOK)
		_ = SendJSON(GetUserResponse{
			ID:       user.ID,
			Name:     user.Name,
			Balance:  user.Balance,
			CCNumber: user.CCNumber,
		}, w)
	}
}

// swagger:route GET /rest/v1//users user getUserByToken
//
// Get user
//     Responses:
//       200: getUserByToken
//       400: errorResponse
func GetUserByTokenHandler(w http.ResponseWriter, r *http.Request) {
	sharedUser := core.GetUserByToken(r.Header.Get("Authorization"))
	if sharedUser.ID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		sendBadResponse(w, "no user in shared storage", BadRequest)
		return
	}

	user, err := core.GetUserByID(sharedUser.ID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		sendBadResponse(w, "no user in db", BadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	_ = SendJSON(GetUserResponse{
		ID:       user.ID,
		Name:     user.Name,
		Balance:  user.Balance,
		CCNumber: user.CCNumber,
	}, w)
}

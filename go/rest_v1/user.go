package rest_v1

import (
	"encoding/json"
	"github.com/panytsch/microservice-poc-db/go/pkg/core"
	"log"
	"net/http"
)

//swagger:parameters createUser
type SwaggerCreateNewUserRequest struct {
	//in:body
	Body CreateNewUserRequest
}

type CreateNewUserRequest struct {
	//unique:true
	//required:true
	Name string
	//required:true
	Password string
}

//swagger:response createUser
type SwaggerCreateNewUserResponse struct {
	//in:body
	Body CreateNewUserResponse
}

type CreateNewUserResponse struct {
	ID   uint64
	Name string
}

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

//swagger:parameters getUser
type SwaggerGetUserRequest struct {
	//in: body
	//required:true
	Body GetUserRequest
}

type GetUserRequest struct {
	//required:true
	Name string
	//required:true
	Password string
}

//swagger:response getUser
type SwaggerGetUserResponse struct {
	//in:body
	Body CreateNewUserResponse
}

type GetUserResponse struct {
	ID   uint64
	Name string
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
	user := new(core.User).User.FindByNameAndPass(req.Name, req.Password)

	if user.ID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		response := new(NoDataFoundResponse)
		response.Code = NoDataFound
		response.Message = "User not found"
		_ = SendJSON(response, w)
	} else {
		w.WriteHeader(http.StatusOK)
		_ = SendJSON(GetUserResponse{
			ID:   user.ID,
			Name: user.Name,
		}, w)
	}
}

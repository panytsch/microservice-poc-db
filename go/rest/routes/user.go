package routes

import (
	"encoding/json"
	"errors"
	"github.com/panytsch/microservice-poc-db/go/pkg/models"
	"log"
	"net/http"
)

//swagger:parameters createUser
type SwaggerCreateNewUserRequest struct {
	//New user`s info
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
	ID     uint
	Status models.UserStatus
}

// swagger:route POST /user createUser
//
// Create new user
//     Responses:
//       201: createUser
func CreateNewUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	}
	req := CreateNewUserRequest{}
	decoder := json.NewDecoder(r.Body)
	er := decoder.Decode(&req)
	if er != nil {
		log.Printf("error while decode request: %v\n", er)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if createUser(req.Name, req.Password) != nil {
		log.Printf("error while saving user: %v\n", er)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func createUser(name string, pass string) error {
	user := models.User{
		Name:     name,
		Password: pass,
	}
	user.Create()
	if user.ID == 0 {
		return errors.New("user wasn't created")
	}
	return nil
}

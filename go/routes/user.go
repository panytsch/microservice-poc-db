package routes

import (
	"encoding/json"
	"errors"
	"github.com/panytsch/microservice-poc-db/go/pkg/models"
	"log"
	"net/http"
)

type CreateNewUserRequest struct {
	Name     string
	Password string
}

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

package core

import (
	"errors"
	"github.com/panytsch/microservice-poc-db/go/pkg/db"
)

func CreateTransaction(userID uint64, amount db.TransactionAmount) *db.CreateTransactionProcedureResult {
	return db.NewCreateTransactionProcedure().Run(userID, db.TransactionStatusCreated, amount)
}

func GetTransactionByIDAndUserID(id uint64, userID uint64) (*db.Transaction, error) {
	tr := new(db.Transaction)
	tr.ID = id
	tr.UserID = userID
	db.DB.Find(tr)
	if tr.Status == 0 {
		return tr, errors.New("transaction not found")
	}
	return tr, nil
}

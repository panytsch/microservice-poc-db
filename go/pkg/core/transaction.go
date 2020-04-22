package core

import (
	"errors"
	"github.com/panytsch/microservice-poc-db/go/pkg/db"
)

func CreateTransaction(userID uint, amount db.TransactionAmount) *db.CreateTransactionProcedureResult {
	return db.NewCreateTransactionProcedure().Run(userID, db.TransactionStatusCreated, amount)
}

func GetTransactionByIDAndUserID(id uint, userID uint) (*db.Transaction, error) {
	tr := new(db.Transaction)
	tr.ID = id
	tr.User = &db.User{ID: userID}
	db.DB.Unscoped().Find(tr)
	if tr.Status == 0 {
		return tr, errors.New("transaction not found")
	}
	return tr, nil
}

func GetLastTransactions(limit uint, offset uint, userID uint) ([]*db.Transaction, error) {
	var transactions []*db.Transaction
	db.DB.Find(&transactions).Limit(limit).Offset(offset).Where("user_id = ?", userID)
	return transactions, nil
}

package core

import "github.com/panytsch/microservice-poc-db/go/pkg/db"

func CreateTransaction(userID uint64, amount db.TransactionAmount) *db.CreateTransactionProcedureResult {
	return db.NewCreateTransactionProcedure().Run(userID, db.TransactionStatusCreated, amount)
}

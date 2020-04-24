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
	tr.UserID = userID
	db.DB.Unscoped().Where(tr).Find(tr)
	if tr.Status == 0 {
		return tr, errors.New("transaction not found")
	}
	return tr, nil
}

func GetLastTransactions(limit uint, offset uint, userID uint) ([]*db.Transaction, error) {
	var transactions []*db.Transaction
	emptyTr := db.Transaction{}
	rows, err := db.DB.Table(emptyTr.TableName()).Order("id desc").Where("user_id = ?", userID).Limit(limit).Offset(offset).Rows()
	if err != nil {
		return transactions, errors.New("can't find any transaction")
	}
	for rows.Next() {
		tr := new(db.Transaction)
		err = db.DB.ScanRows(rows, tr)
		if err != nil {
			return transactions, errors.New("can't find any transaction")
		}
		transactions = append(transactions, tr)
	}
	return transactions, nil
}

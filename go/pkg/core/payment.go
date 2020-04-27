package core

import (
	"errors"
	"github.com/panytsch/microservice-poc-db/go/pkg/db"
)

func CreatePayment(userID uint, amount db.PaymentAmount) *db.CreatePaymentProcedureResult {
	return db.NewCreatePaymentProcedure().Run(userID, db.PaymentStatusCreated, amount)
}

func GetPaymentByIDAndUserID(id uint, userID uint) (*db.Payment, error) {
	tr := new(db.Payment)
	tr.ID = id
	tr.UserID = userID
	db.DB.Unscoped().Where(tr).Find(tr)
	if tr.Status == 0 {
		return tr, errors.New("Payment not found")
	}
	return tr, nil
}

func GetLastPayments(limit uint, offset uint, userID uint) ([]*db.Payment, error) {
	var Payments []*db.Payment
	emptyTr := db.Payment{}
	rows, err := db.DB.Table(emptyTr.TableName()).Order("id desc").Where("user_id = ?", userID).Limit(limit).Offset(offset).Rows()
	if err != nil {
		return Payments, errors.New("can't find any Payment")
	}
	for rows.Next() {
		tr := new(db.Payment)
		err = db.DB.ScanRows(rows, tr)
		if err != nil {
			return Payments, errors.New("can't find any Payment")
		}
		Payments = append(Payments, tr)
	}
	return Payments, nil
}

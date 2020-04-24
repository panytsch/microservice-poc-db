package core

import (
	"errors"
	"github.com/panytsch/microservice-poc-db/go/pkg/db"
)

func MakeRefund(amount db.RefundAmount, userId uint) *db.Refund {
	refund := &db.Refund{
		Status: db.RefundStatusCreated,
		UserID: userId,
		Amount: amount,
	}
	refund.Create()
	return refund
}

func GetRefundByIDAndUserID(id uint, userID uint) (*db.Refund, error) {
	tr := new(db.Refund)
	tr.ID = id
	tr.UserID = userID
	db.DB.Unscoped().Where(tr).Find(tr)
	if tr.Status == 0 {
		return tr, errors.New("refund not found")
	}
	return tr, nil
}

func GetLastRefunds(limit uint, offset uint, userID uint) ([]*db.Refund, error) {
	var refunds []*db.Refund
	emptyTr := db.Refund{}
	rows, err := db.DB.Table(emptyTr.TableName()).Order("id desc").Where("user_id = ?", userID).Limit(limit).Offset(offset).Rows()
	if err != nil {
		return refunds, errors.New("can't find any refund")
	}
	for rows.Next() {
		tr := new(db.Refund)
		err = db.DB.ScanRows(rows, tr)
		if err != nil {
			return refunds, errors.New("can't find any refund")
		}
		refunds = append(refunds, tr)
	}
	return refunds, nil
}

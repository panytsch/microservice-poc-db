package db

import "errors"

type RefundAmount int

// possible values will be show in swagger docs

// - 1 RefundStatusCreated
// - 2 RefundStatusInProgress
// - 3 RefundStatusFinished
// - 4 RefundStatusCancelled
type RefundStatus int

// after adding status add comment to type above. In this case new status will be shown in doc
//noinspection GoUnusedConst
const (
	RefundStatusCreated RefundStatus = iota + 1
	RefundStatusInProgress
	RefundStatusFinished
	RefundStatusCancelled
)

type Refund struct {
	Model
	Status RefundStatus
	User   *User
	UserID uint
	Amount RefundAmount
}

func (r *Refund) Create() {
	DB.Omit("created_at", "updated_at", "deleted_at").Create(r)
}

func (*Refund) TableName() string {
	return "refunds"
}

func (r *Refund) ChangeStatus(status RefundStatus) (*Refund, error) {
	r.Status = status
	err := DB.Save(r).Error
	if err != nil {
		return nil, errors.New("can't update refund")
	}
	return r, nil
}

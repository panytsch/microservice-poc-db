package db

type PaymentAmount int

// possible values will be show in swagger docs

// - 1 PaymentStatusCreated
// - 2 PaymentStatusInProgress
// - 3 PaymentStatusInFinished
// - 4 PaymentStatusInCancelled
type PaymentStatus int

// after adding status add comment to type above. In this case new status will be shown in doc
//noinspection GoUnusedConst
const (
	PaymentStatusCreated PaymentStatus = iota + 1
	PaymentStatusInProgress
	PaymentStatusInFinished
	PaymentStatusInCancelled
)

type Payment struct {
	Model
	Status PaymentStatus
	User   *User
	UserID uint
	Amount PaymentAmount
}

func (t *Payment) Create() {
	DB.Omit("created_at", "updated_at", "deleted_at").Create(t)
}

func (*Payment) TableName() string {
	return "Payments"
}

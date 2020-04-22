package db

type TransactionAmount int

// possible values will be show in swagger docs

// - 1 TransactionStatusCreated
// - 2 TransactionStatusInProgress
// - 3 TransactionStatusInFinished
// - 4 TransactionStatusInCancelled
type TransactionStatus int

// after adding status add comment to type above. In this case new status will be shown in doc
//noinspection GoUnusedConst
const (
	TransactionStatusCreated TransactionStatus = iota + 1
	TransactionStatusInProgress
	TransactionStatusInFinished
	TransactionStatusInCancelled
)

type Transaction struct {
	Model
	Status TransactionStatus
	User   *User
	UserID uint
	Amount TransactionAmount
}

func (t *Transaction) Create() {
	DB.Omit("created_at", "updated_at", "deleted_at").Create(t)
}

func (*Transaction) TableName() string {
	return "transactions"
}

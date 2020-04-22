package db

type TransactionAmount int

// possible values will be show in swagger docs

// - 1 - created
type TransactionStatus int

// after adding status add comment to type above. In this case new status will be shown in doc
const (
	TransactionStatusCreated TransactionStatus = iota + 1
)

type Transaction struct {
	ID     uint64
	Status TransactionStatus
	UserID uint64
	Amount TransactionAmount
}

func (t *Transaction) Create() {
	DB.Create(t)
}

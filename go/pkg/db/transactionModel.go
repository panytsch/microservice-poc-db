package db

type TransactionStatus int
type TransactionAmount int

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

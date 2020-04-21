package db

type TransactionStatus int
type TransactionAmount uint64

const (
	TransactionStatusCreated TransactionStatus = iota + 1
	TransactionStatusInProcess
	TransactionStatusFinished
	TransactionStatusCancelled
)

type Transaction struct {
	ID     uint64
	Status TransactionStatus
	User   User
	UserID uint
	Amount TransactionAmount
}

func (t *Transaction) Create() {
	DB.Create(t)
}

package blockchain

type Transaction interface {
	Verify() (bool, error)
}

type TransactionType string

const (
	Send TransactionType = "Send"
)

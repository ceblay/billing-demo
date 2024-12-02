package billing

import "errors"

var ErrUnsupportedTransactionType = errors.New("Unsupported transaction type")

const (
	Deposit      = "DEPOSIT"
	Debit        = "DEBIT"
	Credit       = "CREDIT"
	Transfer     = "TRANSFER"
	Subscription = "SUBSCRIPTION"
)

type TransactionType struct {
	ref string
}

func (tt TransactionType) String() string {
	return tt.ref
}

func (tt TransactionType) IsZero() bool {
	return tt == TransactionType{}
}

func NewTransactionTypeFromString(s string) (TransactionType, error) {
	switch s {
	case Deposit:
		return TransactionType{Deposit}, nil
	case Debit:
		return TransactionType{Debit}, nil
	case Credit:
		return TransactionType{Credit}, nil
	case Transfer:
		return TransactionType{Transfer}, nil
	case Subscription:
		return TransactionType{Subscription}, nil
	default:
		return TransactionType{}, ErrUnsupportedTransactionType
	}
}

package billing

import "errors"

var ErrUnsupportedTransactionStatus = errors.New("Unsupported transaction status")

const (
	Pending   = "PENDING"
	Completed = "COMPLETED"
	Failed    = "FAILED"
	Canceled  = "CANCELED"
)

type TransactionStatus struct {
	ref string
}

func (ts TransactionStatus) String() string {
	return ts.ref
}

func (ts TransactionStatus) IsZero() bool {
	return ts == TransactionStatus{}
}

func NewTransactionStatusFromString(s string) (TransactionStatus, error) {
	switch s {
	case Pending:
		return TransactionStatus{Pending}, nil
	case Completed:
		return TransactionStatus{Completed}, nil
	case Failed:
		return TransactionStatus{Failed}, nil
	case Canceled:
		return TransactionStatus{Canceled}, nil
	default:
		return TransactionStatus{}, ErrUnsupportedTransactionStatus
	}
}

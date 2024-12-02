package billing

type Transaction struct {
	id              string
	wallet          string
	user            string
	currency        string
	kind            TransactionType
	amount          float64
	status          TransactionStatus
	confirmed       bool
	ipaddress       string
	paymentProvider string
	paymentID       string
	modeOfPayment   string
	last4Digits     string
	channel         string
	createdAt       int64
}

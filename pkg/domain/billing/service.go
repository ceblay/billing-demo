package billing

type PaymentGatewayService interface {
	GetPaymentMethodByID() error
	GetDefaultPaymentMethodEnvironment() error
}

type PlanService interface {
	GetByID() error
}

type SubscriptionService interface {
	GetByID() error
	Create() error
}

type TransactionService interface {
	Update() error
	GetByID() error
	UpdateSuccess() error
}

type UserService interface {
	GetByID() error
}

type WalletService interface {
	GetByID() error
	GetUserWallet() error
}

type ServiceBusService interface {
	SendToTopic() error
}

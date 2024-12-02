package billing

type Repository interface {
	GetAllBillingHistory() (string, error)
	// GetAllBillingHistory() ([]*History, error)
	// GetBillingHistoryById(string) (*History, error)
	// GetBillingInfoByUserID(string) (*History, error) // should be info instead of history object
	// CreateBillingHistory(*History) error
	// CreateBillingInfo(*History) error // should be info object instead of history
	// SaveBillingInfo(*History) error   // info instead of history
	// GetByTransactionID(string) (*History, error)
	// UpdateHistory(*History) error
	// GetByID(string) (*History, error)
}

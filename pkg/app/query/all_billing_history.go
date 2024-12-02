package query

type AllBillingHistoryReadModel interface {
	GetAllBillingHistory() (string, error)
}

type AllBillingHistoryHandler struct {
	readModel AllBillingHistoryReadModel
}

func NewAllBillingHistoryHandler(rm AllBillingHistoryReadModel) AllBillingHistoryHandler {
	return AllBillingHistoryHandler{
		readModel: rm,
	}
}

func (h *AllBillingHistoryHandler) Handle() (string, error) {
	return h.readModel.GetAllBillingHistory()
}

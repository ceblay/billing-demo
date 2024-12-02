package app

import (
	"github.com/ceblay/billing-demo/pkg/app/query"
)

type Queries struct {
	AllBillingHistory query.AllBillingHistoryHandler
}

type Commands struct {
}

type Application struct {
	Queries  Queries
	Commands Commands
}

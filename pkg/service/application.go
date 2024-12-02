package service

import (
	"fmt"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/ceblay/billing-demo/pkg/adapters"
	"github.com/ceblay/billing-demo/pkg/app"
	"github.com/ceblay/billing-demo/pkg/app/query"
	bl "github.com/ceblay/billing-demo/pkg/domain/billing"
)

func initializeDatabase() bl.Repository {
	db, err := gorm.Open(sqlite.Open(os.Getenv("DB_NAME")), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Could not connect to database: %v", err))
	}

	repository := adapters.NewSqliteRepository(db)
	return repository
}

func NewApplication() app.Application {
	repository := initializeDatabase()
	return app.Application{
		Queries: app.Queries{
			AllBillingHistory: query.NewAllBillingHistoryHandler(repository),
		},
		Commands: app.Commands{},
	}
}

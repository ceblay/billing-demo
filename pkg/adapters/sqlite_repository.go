package adapters

import (
	"database/sql"
	"log"
	"time"

	"gorm.io/gorm"
)

type billingHistory struct {
	gorm.Model
	Id            string `gorm:"primaryKey"`
	UserID        string
	Status        string // default UNPAID
	TransactionID string
	BillingItem   string
	InvoiceID     sql.NullString
	ItemType      string
}

type billingInformation struct {
	gorm.Model
	Id           string `gorm:"primaryKey"`
	UserID       string
	Firstname    string
	Lastname     string
	Addressline1 string
	Addressline2 string
	City         string
	State        string
	Postcode     string
}

type billingItem struct {
	gorm.Model
	Id                 string `gorm:"primaryKey"`
	BillingHistoryID   string
	ItemType           string
	SubscriptionID     sql.NullString
	PlanID             sql.NullString
	Amount             float64
	BillingPeriodStart time.Time
	BillingPeriodEnd   time.Time
	Liters             float64
	StartConsumption   sql.NullString
	EndConsumption     sql.NullString
	SerianNumber       string
	Currency           string
	Zone               sql.NullString
}

type SqliteRepository struct {
	db *gorm.DB
}

func applySchema(db *gorm.DB) error {
	log.Println("About applying migrations")
	return db.AutoMigrate(
		&billingHistory{},
		&billingInformation{},
		&billingItem{},
	)
}

func NewSqliteRepository(db *gorm.DB) *SqliteRepository {
	err := applySchema(db)
	if err != nil {
		log.Println("Could not perform database automigration")
	}
	return &SqliteRepository{
		db: db,
	}
}

func (s *SqliteRepository) GetAllBillingHistory() (string, error) {
	return "All history is contained herein", nil
}

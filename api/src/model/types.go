package model

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Gift struct {
	gorm.Model
	Name            string
	Price           float64
	Link            string
	Description     string
	Demographic     string
	Category        pq.StringArray `gorm:"type:text[]"`
	Occasion        string
	GiftCollections []*GiftCollection `gorm:"many2many:gift_collection_gifts;"`
}

type GiftRequest struct {
	gorm.Model
	CustomerID         uint
	GiftResponseID     *uint
	RecipientName      string
	RecipientAge       uint
	Occasion           pq.StringArray `gorm:"type:text[]"`
	RecipientInterests pq.StringArray `gorm:"type:text[]"`
	BudgetMax          uint
	BudgetMin          uint
	GiftResponse       *GiftResponse
	DateNeeded         time.Time
}

type GiftCollection struct {
	gorm.Model
	CustomerID     *uint
	Customer       *Customer
	CollectionName string
	Gifts          []*Gift `gorm:"many2many:gift_collection_gifts;"`
}

type GiftResponse struct {
	gorm.Model
	GiftCollection   GiftCollection
	GiftCollectionID uint
	CustomMessage    string
}

type User struct {
	gorm.Model
	Email     string
	FirstName string
	LastName  string
	Password  string
}

type Customer struct {
	gorm.Model
	UserID            uint
	User              User
	AvailableRequests uint
	GiftCollections   []*GiftCollection
	GiftRequests      []*GiftRequest
}

type Admin struct {
	gorm.Model
	UserID uint
	User   User
}

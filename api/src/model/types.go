package model

import (
	"gorm.io/gorm"
)

type ExampleGift struct {
	gorm.Model
	Name  string
	Price int
}

type ExampleGiftInput struct {
	Name  string
	Price int
}

type Gift struct {
	gorm.Model
	Name            string
	Price           float64
	Link            string
	Description     string
	Demographic     string
	GiftCollections []*GiftCollection `gorm:"many2many:gift_request_gifts;"`
}

type GiftRequest struct {
	gorm.Model
	// Need to link to the customer for the customer ID
	GiftResponseID uint
	// RecipientName      string
	// RecipientAge       uint
	// Occasion           []string
	// RecipientInterests []string
	// BudgetMax          float64
	// BudgetMin          float64
	GiftResponse GiftResponse
	// DateNeeded         time.Time
}

type GiftCollection struct {
	gorm.Model
	//CustomerID int, fk
	CollectionName string
	Gifts          []*Gift `gorm:"many2many:gift_request_gifts;"`
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
	UserID uint
	User   User
	//GiftCollections []ExampleGiftCollection
	//GiftRequests    []ExampleGiftRequests
}

type Admin struct {
	gorm.Model
	UserID uint
	User   User
}

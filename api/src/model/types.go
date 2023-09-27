package model

import (
	"time"
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

type GiftInput struct {
	Name        string
	Price       float64
	Link        string
	Description string
	Demographic string
}

type GiftRequest struct {
	gorm.Model
	// Need to link to the customer for the customer ID
	RecipientName      string
	RecipientAge       uint
	Occasion           []string
	RecipientInterests []string
	BudgetMax          float64
	BudgetMin          float64
	GiftResponse       GiftResponse
	DateNeeded         time.Time
}

type GiftRequestInput struct {
	RecipientName      string
	RecipientAge       uint
	Occasion           []string
	RecipientInterests []string
	BudgetMax          float64
	BudgetMin          float64
	DateNeeded         time.Time
}

type GiftCollection struct {
	gorm.Model
	//CustomerID int, fk
	CollectionName   string
	Gifts            []*Gift `gorm:"many2many:gift_request_gifts;"`
}

type GiftCollectionInput struct {
	CollectionName   string
	collectionsGifts GiftCollection
}

type GiftResponse struct {
	GiftCollection GiftCollection
	CustomMessage  string
}

type GiftResponseInput struct {
	CustomMessage string
}

type User struct {
	gorm.Model
	Email     string
	FirstName string 
	LastName  string 
	Password  string
}

type UserInput struct {
	Email     string
	FirstName string
	LastName  string
	Password  string
}

type Customer struct {
	gorm.Model
	UserID uint 
	User User
	//GiftCollections []ExampleGiftCollection
	//GiftRequests    []ExampleGiftRequests
}

type CustomerInput struct {
	UserID uint
	User User
	//GiftCollections []ExampleGiftCollection
	//GiftRequests    []ExampleGiftRequests
}

type Admin struct {
	gorm.Model
	UserID uint
	User User
}

type AdminInput struct {
	UserID uint
	User User
}
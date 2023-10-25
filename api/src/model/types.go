package model

import (
	"errors"
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
	GiftCollections []*GiftCollection `gorm:"many2many:gift_request_gifts;"`
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
	UserID          uint
	User            User
	GiftCollections []*GiftCollection
	GiftRequests    []*GiftRequest
}

type Admin struct {
	gorm.Model
	UserID uint
	User   User
}

func (gift *Gift) BeforeSave(tx *gorm.DB) (err error) {
	if len(gift.Name) >= 10 || len(gift.Name) == 0 {
		err = errors.New("gift name must be between 1 and 9 characters")
		return err
	}

	if gift.Price < 0 {
		err = errors.New("gift must have a price")
		return err
	}

	// ask about other validations
	return
}

func (gc *GiftCollection) BeforeSave(tx *gorm.DB) (err error) {
	if len(gc.CollectionName) >= 10 || len(gc.CollectionName) == 0 {
		err = errors.New("giftCollection name must be between 1 and 9 characters")
		return err
	}

	if gc.CustomerID == nil {
		err = errors.New("giftCollection must have a customerID")
		return err
	}

	// ask about other validations
	return
}

package model

import (
	"errors"
	"regexp"
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

func (user *User) BeforeSave(tx *gorm.DB) (err error) {
	if len(user.FirstName) >= 10 || len(user.FirstName) == 0 {
		err = errors.New("user first name must be between 1 and 9 characters")
		return err
	}

	if len(user.LastName) >= 10 || len(user.LastName) == 0 {
		err = errors.New("user last name must be between 1 and 9 characters")
		return err
	}

	if len(user.Password) >= 10 || len(user.Password) == 0 {
		err = errors.New("user password must be between 1 and 9 characters")
		return err
	}

	if len(user.Email) >= 10 || len(user.Email) == 0 {
		err = errors.New("user email must be between 1 and 9 characters")
		return err
	}

	// use regex to validate email
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	regex := regexp.MustCompile(pattern)

	// Use the MatchString function to check if the text matches the pattern
	if regex.MatchString(user.Email) == false {
		err = errors.New("user email must be a valid email")
		return err
	}

	// ask about other validations
	return
}

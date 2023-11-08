package model

import (
	"errors"
	"regexp"
	"time"

	"github.com/lib/pq"
	"golang.org/x/exp/slices"
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
	if len(gift.Name) == 0 {
		err = errors.New("gift name cannot be empty")
		return err
	}

	if gift.Price < 0 {
		err = errors.New("gift must have a price")
		return err
	}

	// if a gift has no link
	if len(gift.Link) == 0 {
		err = errors.New("gift must have a link")
		return err
	}

	Demographics := []string{
		"For her",
		"For him",
		"For kids",
		"For mom",
		"For dad",
		"For women",
		"For men",
	}

	for _, demographic := range Demographics {
		if !slices.Contains(Demographics, demographic) {
			err = errors.New("gift must have a valid demographic")
			return err
		}
	}

	Occasions := []string{
		"Birthday",
		"Bridal",
		"Get well soon",
		"New baby",
		"Thinking of you",
		"Thank you",
	}

	for _, occasion := range Occasions {
		if !slices.Contains(Occasions, occasion) {
			err = errors.New("gift must have a valid occasion")
			return err
		}
	}

	Categories := []string{
		"Best selling",
		"Fun",
		"Gadgets",
		"Home",
		"Jewelry",
		"Kitchen & bar",
		"Warm and cozy",
	}

	for _, category := range Categories {
		if !slices.Contains(Categories, category) {
			err = errors.New("gift must have a valid category")
			return err
		}
	}

	// if a gift has no description
	if len(gift.Description) == 0 {
		err = errors.New("gift must have a description")
		return err
	}

	return
}

func (gc *GiftCollection) BeforeSave(tx *gorm.DB) (err error) {
	// if collection name is not set
	if len(gc.CollectionName) == 0 {
		err = errors.New("giftCollection must have a name")
		return err
	}

	// if customer is not found
	if gc.Customer == nil {
		err = errors.New("giftCollection must have a customer")
		return err
	}

	return
}

func (user *User) BeforeSave(tx *gorm.DB) (err error) {
	if len(user.FirstName) == 0 {
		err = errors.New("user must have a first name")
		return err
	}

	if len(user.LastName) == 0 {
		err = errors.New("user must have a last name")
		return err
	}

	if len(user.Password) == 0 {
		err = errors.New("Please enter a password")
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

	return
}

func (customer *Customer) BeforeSave(tx *gorm.DB) (err error) {
	// if giftRequests is not empty or not populated
	if customer.GiftRequests == nil {
		err = errors.New("customer must have a giftRequests")
		return err
	}

	return
}

func (giftRequest *GiftRequest) BeforeSave(tx *gorm.DB) (err error) {
	// if recipient name is not set
	if len(giftRequest.RecipientName) == 0 {
		err = errors.New("giftRequest must have a recipient name")
		return err
	}

	// if recipient age is not set
	if giftRequest.RecipientAge < 1 || giftRequest.RecipientAge > 150 {
		err = errors.New("giftRequest must have a valid recipient age")
		return err
	}

	GiftOccasions := []string{
		"Birthday",
		"Bridal",
		"Get well soon",
		"New baby",
		"Thinking of you",
		"Thank you",
	}

	// if occasion is not in GiftOccasions
	if giftRequest.Occasion != nil {
		for _, occasion := range giftRequest.Occasion {
			if !slices.Contains(GiftOccasions, occasion) {
				err = errors.New("giftRequest must have a valid occasion")
				return err
			}
		}
	} else {
		err = errors.New("giftRequest must have an occasion")
		return err
	}

	Interests := []string{
		"Best selling",
		"Fun",
		"Gadgets",
		"Home",
		"Jewelry",
		"Kitchen & bar",
		"Warm and cozy",
	}

	// if interests is not in Interests
	if giftRequest.RecipientInterests != nil {
		for _, interest := range giftRequest.RecipientInterests {
			if !slices.Contains(Interests, interest) {
				err = errors.New("giftRequest must have a valid interest")
				return err
			}
		}
	} else {
		err = errors.New("giftRequest must have an interest")
		return err
	}

	// if either budget is below negative
	if giftRequest.BudgetMax < 0 || giftRequest.BudgetMin < 0 {
		err = errors.New("giftRequest budget cannot be negative")
		return err
	}

	// if max budget is less than min budget
	if giftRequest.BudgetMax <= giftRequest.BudgetMin {
		err = errors.New("giftRequest max budget must be greater than min budget")
		return err
	}

	// if date needed is not set
	if giftRequest.DateNeeded.IsZero() {
		err = errors.New("giftRequest must have a date needed")
		return err
	} else if giftRequest.DateNeeded.Before(time.Now()) {
		// if date needed is in the past
		err = errors.New("giftRequest date needed must be in the future")
		return err
	}

	return
}

func (giftResponse *GiftResponse) BeforeSave(tx *gorm.DB) (err error) {
	if giftResponse.CustomMessage == "" {
		err = errors.New("giftResponse must have a custom message")
		return err
	}

	return
}

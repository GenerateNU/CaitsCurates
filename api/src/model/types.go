package model

import "gorm.io/gorm"

type ExampleGift struct {
	gorm.Model
	Name  string
	Price int
}

type ExampleGiftInput struct {
	Name  string
	Price int
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
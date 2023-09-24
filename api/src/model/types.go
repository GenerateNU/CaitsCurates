package model

import "gorm.io/gorm"

type ExampleGift struct {
	gorm.Model
	Name  string
	Price int
}

type ExampleGiftInput struct {
	Name  string `binding:"required"`
	Price int    `binding:"required"`
}

type User struct {
	gorm.Model
	Email     string `binding:"required, email"`
	FirstName string `binding:"required"`
	LastName  string `binding:"required"`
	Password  string `binding:"required" valid:"length(6|20)"`
}

type UserInput struct {
	Email     string `binding:"required, email"`
	FirstName string `binding:"required"`
	LastName  string `binding:"required"`
	Password  string `binding:"required" valid:"length(6|20)"`
}

type Customer struct {
	gorm.Model
	UserID uint `binding:"required"`
	//GiftCollections []ExampleGiftCollection
	//GiftRequests    []ExampleGiftRequests
}

type CustomerInput struct {
	UserID uint `binding:"required"`
	//GiftCollections []ExampleGiftCollection
	//GiftRequests    []ExampleGiftRequests
}

type Admin struct {
	gorm.Model
	UserID uint `binding:"required"`
}

type AdminInput struct {
	UserID uint `binding:"required"`
}
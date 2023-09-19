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

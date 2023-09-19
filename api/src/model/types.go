package model

import "gorm.io/gorm"

type ExampleGift struct {
	gorm.Model
	Name  string `gorm:"column:name" json:"name"`
	Price int    `gorm:"column:price" json:"price"`
}
type ExampleGiftInput struct {
	Name  string `json:"name" binding:"required"`
	Price int    `json:"price" binding:"required"`
}

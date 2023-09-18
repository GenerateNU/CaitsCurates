package model

import (
	"gorm.io/gorm"
)

// WriteExampleGiftToDb saves the ExampleGift and returns it
func WriteExampleGiftToDb(db *gorm.DB, eg ExampleGift) (ExampleGift, error) {
	if err := db.Create(&eg).Error; err != nil {
		return ExampleGift{}, err
	}
	return eg, nil
}

// GetExampleGiftFromDB fetches an ExampleGift by ID
func GetExampleGiftFromDB(db *gorm.DB, id int64) (ExampleGift, error) {
	var eg ExampleGift
	if err := db.Where("id = ?", id).First(&eg).Error; err != nil {
		return ExampleGift{}, err
	}
	return eg, nil
}

// GetAllExampleGiftsFromDB fetches all ExampleGift
func GetAllExampleGiftsFromDB(db *gorm.DB) ([]ExampleGift, error) {
	var gifts []ExampleGift
	if err := db.Find(&gifts).Error; err != nil {
		return nil, err
	}
	return gifts, nil
}

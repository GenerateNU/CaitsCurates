package model

import (
	"gorm.io/gorm"
)

// WriteExampleGiftToDb saves the ExampleGift and returns it
func WriteExampleGiftToDb(db *gorm.DB, inputGift ExampleGiftInput) (ExampleGift, error) {
	eg := ExampleGift{Name: inputGift.Name, Price: inputGift.Price}
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

func GetIncompleteGiftRequestsFromDB(db *gorm.DB) ([]GiftRequest, error) {
	var requests []GiftRequest
	if err := db.Where("gift_response_id IS NULL").Preload("GiftResponse").Find(&requests).Error; err != nil {
		return nil, err
	}
	return requests, nil
}

func GetCompleteGiftRequestsFromDB(db *gorm.DB) ([]GiftRequest, error) {
	var requests []GiftRequest
	if err := db.Where("gift_response_id IS NOT NULL").Find(&requests).Error; err != nil {
		return nil, err
	}
	return requests, nil
}
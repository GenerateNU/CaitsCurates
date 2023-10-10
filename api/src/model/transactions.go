package model

import (
	"gorm.io/gorm"
)

func WriteRequestToDb(db *gorm.DB, inputRequest GiftRequest) (GiftRequest, error) {
	if err := db.Create(&inputRequest).Error; err != nil {
		return GiftRequest{}, err
	}
	return inputRequest, nil
}
func WriteResponseToDb(db *gorm.DB, inputResponse GiftResponse) (GiftResponse, error) {
	if err := db.Create(&inputResponse).Error; err != nil {
		return GiftResponse{}, err
	}
	return inputResponse, nil
}
func WriteCollectionToDb(db *gorm.DB, inputCollection GiftCollection) (GiftCollection, error) {
	if err := db.Create(&inputCollection).Error; err != nil {
		return GiftCollection{}, err
	}
	return inputCollection, nil
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

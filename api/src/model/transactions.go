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

// UpdateGiftToDb updates the Gift and returns it
func UpdateGiftToDb(db *gorm.DB, id int64, inputGift Gift) (Gift, error) {
	// Fetch Gift Record
	var updatedGift Gift
	if err := db.Where("id = ?", id).First(&updatedGift).Error; err != nil {
		return Gift{}, err
	}

	// Update Gift Record
	updates := map[string]interface{}{
		"Name":            inputGift.Name,
		"Price":           inputGift.Price,
		"Link":            inputGift.Link,
		"Description":     inputGift.Description,
		"Demographic":     inputGift.Demographic,
		"GiftCollections": inputGift.GiftCollections,
	}

	if err := db.Model(&updatedGift).Updates(updates).Error; err != nil {
		return Gift{}, err
	}

	// Return Updated Gift Record
	return updatedGift, nil
}

// DeleteGiftFromDb Delete Gift
func DeleteGiftFromDb(db *gorm.DB, id int64) error {
	var deletedGift Gift
	if err := db.Where("id = ?", id).First(&deletedGift).Error; err != nil {
		return err
	}

	if err := db.Delete(&deletedGift).Error; err != nil {
		return err
	}

	return nil
}

func GetCompleteGiftRequestsFromDB(db *gorm.DB) ([]GiftRequest, error) {
	var requests []GiftRequest
	if err := db.Where("gift_response_id IS NOT NULL").Find(&requests).Error; err != nil {
		return nil, err
	}
	return requests, nil
}

// GetGiftFromDB fetches an Gift by ID
func GetGiftFromDB(db *gorm.DB, id int64) (Gift, error) {
	var gift Gift
	if err := db.Where("id = ?", id).First(&gift).Error; err != nil {
		return Gift{}, err
	}
	return gift, nil
}

// GetAllGiftsFromDB fetches all ExampleGift
// GetAllGiftsFromDB fetches all Gift
func GetAllGiftsFromDB(db *gorm.DB) ([]Gift, error) {
	var gifts []Gift
	if err := db.Find(&gifts).Error; err != nil {
		return nil, err
	}
	return gifts, nil
}

// GetAllResponsesFromDB fetches all GiftResponse
// WriteGiftToDb saves the Gift and returns it
func WriteGiftToDb(db *gorm.DB, inputGift Gift) (Gift, error) {
	if err := db.Create(&inputGift).Error; err != nil {
		return Gift{}, err
	}
	return inputGift, nil
}
func GetAllResponsesFromDB(db *gorm.DB) ([]GiftResponse, error) {
	var response []GiftResponse
	if err := db.Find(&response).Error; err != nil {
		return nil, err
	}
	return response, nil
}

// GetAllCollectionsFromDB fetches all GiftCollection
func GetAllCollectionsFromDB(db *gorm.DB) ([]GiftCollection, error) {
	var collections []GiftCollection
	if err := db.Find(&collections).Error; err != nil {
		return nil, err
	}
	return collections, nil
}

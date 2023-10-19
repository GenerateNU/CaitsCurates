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
func UpdateGiftRequestToDb(db *gorm.DB, inputRequest GiftRequest) (GiftRequest, error) {
	var updatedGiftRequest GiftRequest
	if err := db.Where("id = ?", inputRequest.ID).First(&updatedGiftRequest).Error; err != nil {
		return GiftRequest{}, err
	}

	updates := make(map[string]interface{})

	if inputRequest.RecipientName != "" {
		updates["RecipientName"] = inputRequest.RecipientName
	}
	if inputRequest.RecipientAge != 0 {
		updates["RecipientAge"] = inputRequest.RecipientAge
	}
	if len(inputRequest.Occasion) > 0 {
		updates["Occasion"] = inputRequest.Occasion
	}
	if len(inputRequest.RecipientInterests) > 0 {
		updates["RecipientInterests"] = inputRequest.RecipientInterests
	}
	if inputRequest.BudgetMax != 0 {
		updates["BudgetMax"] = inputRequest.BudgetMax
	}
	if inputRequest.BudgetMin != 0 {
		updates["BudgetMin"] = inputRequest.BudgetMin
	}
	if !inputRequest.DateNeeded.IsZero() {
		updates["DateNeeded"] = inputRequest.DateNeeded
	}
	if inputRequest.GiftResponseID != nil {
		updates["GiftResponseID"] = inputRequest.GiftResponseID
	}

	if err := db.Model(&updatedGiftRequest).Updates(updates).Error; err != nil {
		return GiftRequest{}, err
	}

	// Return Updated Gift Record
	return updatedGiftRequest, nil
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
func UpdateCollectionToDb(db *gorm.DB, inputCollection GiftCollection) (GiftCollection, error) {
	var updatedCollection GiftCollection
	if err := db.Where("id = ?", inputCollection.ID).First(&updatedCollection).Error; err != nil {
		return GiftCollection{}, err
	}

	updates := make(map[string]interface{})

	if inputCollection.CollectionName != "" {
		updates["CollectionName"] = inputCollection.CollectionName
	}
	if inputCollection.Customer != nil {
		updates["Customer"] = inputCollection.Customer
	}
	if inputCollection.CustomerID != nil {
		updates["CustomerID"] = inputCollection.CustomerID
	}
	if inputCollection.Gifts != nil {
		updates["Gifts"] = inputCollection.Gifts
	}
	if err := db.Model(&updatedCollection).Association("Gifts").Clear(); err != nil {
		return GiftCollection{}, err
	}
	if err := db.Model(&updatedCollection).Updates(updates).Error; err != nil {
		return GiftCollection{}, err
	}

	// Return Updated Gift Record
	return updatedCollection, nil
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
	updates := make(map[string]interface{})

	// Check each field in inputGift and add it to the updates map if it is non-zero
	if inputGift.Name != "" {
		updates["Name"] = inputGift.Name
	}
	if inputGift.Price != 0 {
		updates["Price"] = inputGift.Price
	}
	if inputGift.Link != "" {
		updates["Link"] = inputGift.Link
	}
	if inputGift.Description != "" {
		updates["Description"] = inputGift.Description
	}
	if inputGift.Demographic != "" {
		updates["Demographic"] = inputGift.Demographic
	}
	if inputGift.GiftCollections != nil && len(inputGift.GiftCollections) > 0 {
		updates["GiftCollections"] = inputGift.GiftCollections
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

func DeleteGiftCollectionFromDb(db *gorm.DB, id int64) error {
	var deletedGiftCollection GiftCollection
	if err := db.Where("id = ?", id).First(&deletedGiftCollection).Error; err != nil {
		return err
	}

	if err := db.Delete(&deletedGiftCollection).Error; err != nil {
		return err
	}

	return nil
}
func SearchGiftsDb(db *gorm.DB, searchTerm string, minPrice int, maxPrice int) ([]Gift, error) {
	var gifts []Gift

	if err := db.Where("to_tsvector('english', name || ' ' || description) @@ to_tsquery('english', ?)", searchTerm).
		Where("price >= ? AND price <= ?", minPrice, maxPrice).
		Find(&gifts).Error; err != nil {
		return nil, err
	}

	return gifts, nil
}
func GetCompleteGiftRequestsFromDB(db *gorm.DB) ([]GiftRequest, error) {
	var requests []GiftRequest
	if err := db.Where("gift_response_id IS NOT NULL").Preload("GiftResponse").Preload("GiftResponse.GiftCollection").Find(&requests).Error; err != nil {
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
	if err := db.Preload("GiftCollections").Find(&gifts).Error; err != nil {
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
	if err := db.Preload("GiftCollection").Find(&response).Error; err != nil {
		return nil, err
	}
	return response, nil
}

// GetAllCollectionsFromDB fetches all GiftCollection
func GetAllCollectionsFromDB(db *gorm.DB) ([]GiftCollection, error) {
	var collections []GiftCollection
	if err := db.Preload("Gifts").Find(&collections).Error; err != nil {
		return nil, err
	}
	return collections, nil
}

func AddGiftToCollectionFromDB(db *gorm.DB, inputGift Gift, id int64) (GiftCollection, error) {
	var collection GiftCollection
	if err := db.Where("id = ?", id).First(&collection).Error; err != nil {
		return GiftCollection{}, err
	}

	collection.Gifts = append(collection.Gifts, &inputGift)

	if err := db.Save(&collection).Error; err != nil {
		return GiftCollection{}, err
	}

	return collection, nil
}

func DeleteGiftFromCollectionFromDB(db *gorm.DB, giftID int64, giftCollectionID int64) (GiftCollection, error) {
	var collection GiftCollection
	if err := db.Preload("Gifts").First(&collection, giftCollectionID).Error; err != nil {
		return GiftCollection{}, err
	}

	// Create a new GiftCollection array without the inputGift
	var giftRemovedCollection []*Gift
	for _, gift := range collection.Gifts {
		if gift.ID != uint(giftID) {
			giftRemovedCollection = append(giftRemovedCollection, gift)
		}
	}
	if err := db.Model(&collection).Association("Gifts").Replace(giftRemovedCollection); err != nil {
		return GiftCollection{}, err
	}

	return collection, nil
}

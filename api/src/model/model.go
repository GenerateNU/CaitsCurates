package model

import (
	"regexp"
	"strings"

	"gorm.io/gorm"
)

type PgModel struct {
	Conn *gorm.DB
}

type Model interface {
	AddRequest(GiftRequest) (GiftRequest, error)
	AddResponse(GiftResponse) (GiftResponse, error)
	AddCollection(GiftCollection) (GiftCollection, error)
	IncompleteRequests() ([]GiftRequest, error)
	CompleteRequests() ([]GiftRequest, error)
	UpdateGiftRequest(GiftRequest) (GiftRequest, error)
	GetGift(int64) (Gift, error)
	GetAllGifts() ([]Gift, error)
	AddGift(Gift) (Gift, error)
	UpdateGift(int64, Gift) (Gift, error)
	DeleteGift(int64) error
	DeleteGiftCollection(int64) error
	SearchGifts(string, int, int) ([]Gift, error)
	AllGiftResponses() ([]GiftResponse, error)
	AllCollections() ([]GiftCollection, error)
	AllCustomerCollections(id int64) ([]GiftCollection, error)
	UpdateCollection(GiftCollection) (GiftCollection, error)
	AddGiftToGiftCollection(Gift, int64) (GiftCollection, error)
	AddGiftToCustomerCollection(Gift, string, int64) (GiftCollection, error)
	DeleteGiftFromGiftCollection(int64, int64) (GiftCollection, error)
}


func (m *PgModel) AddRequest(inputRequest GiftRequest) (GiftRequest, error) {

	createdRequest, err := WriteRequestToDb(m.Conn, inputRequest)

	if err != nil {
		return GiftRequest{}, err
	}

	return createdRequest, nil
}
func (m *PgModel) AddResponse(inputResponse GiftResponse) (GiftResponse, error) {

	createdResponse, err := WriteResponseToDb(m.Conn, inputResponse)

	if err != nil {
		return GiftResponse{}, err
	}

	return createdResponse, nil
}
func (m *PgModel) AddCollection(inputCollection GiftCollection) (GiftCollection, error) {

	createdCollection, err := WriteCollectionToDb(m.Conn, inputCollection)

	if err != nil {
		return GiftCollection{}, err
	}

	return createdCollection, nil
}
func (m *PgModel) UpdateCollection(inputCollection GiftCollection) (GiftCollection, error) {

	updatedCollection, err := UpdateCollectionToDb(m.Conn, inputCollection)

	if err != nil {
		return GiftCollection{}, err
	}

	return updatedCollection, nil
}
func (m *PgModel) AddGift(inputGift Gift) (Gift, error) {

	createdGift, err := WriteGiftToDb(m.Conn, inputGift)

	if err != nil {
		return Gift{}, err
	}

	return createdGift, nil
}

func (m *PgModel) GetGift(id int64) (Gift, error) {

	createdGift, err := GetGiftFromDB(m.Conn, id)

	if err != nil {
		return Gift{}, err
	}

	return createdGift, nil
}

func (m *PgModel) GetAllGifts() ([]Gift, error) {

	createdGifts, err := GetAllGiftsFromDB(m.Conn)

	if err != nil {
		return []Gift{}, err
	}

	return createdGifts, nil
}

func (m *PgModel) UpdateGift(id int64, inputGift Gift) (Gift, error) {

	updatedGift, err := UpdateGiftToDb(m.Conn, id, inputGift)

	if err != nil {
		return Gift{}, err
	}

	return updatedGift, nil
}

func (m *PgModel) UpdateGiftRequest(inputGiftRequest GiftRequest) (GiftRequest, error) {

	updatedGiftRequest, err := UpdateGiftRequestToDb(m.Conn, inputGiftRequest)

	if err != nil {
		return GiftRequest{}, err
	}

	return updatedGiftRequest, nil
}

func (m *PgModel) AllGiftResponses() ([]GiftResponse, error) {
	responses, err := GetAllResponsesFromDB(m.Conn)

	if err != nil {
		return []GiftResponse{}, err
	}
	return responses, nil
}

func (m *PgModel) DeleteGift(id int64) error {

	err := DeleteGiftFromDb(m.Conn, id)

	if err != nil {
		return err
	}

	return nil
}
func (m *PgModel) SearchGifts(searchTerm string, minPrice int, maxPrice int) ([]Gift, error) {
	var gifts []Gift
	searchTerm = strings.TrimSpace(searchTerm)


	// Convert to lowercase
	searchTerm = strings.ToLower(searchTerm)

	// Remove special characters or punctuations
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	searchTerm = reg.ReplaceAllString(searchTerm, " ")
	//
	searchTerms := strings.Fields(searchTerm)
	for i, term := range searchTerms {
		searchTerms[i] = term + ":*"
	}
	formattedSearchTerm := strings.Join(searchTerms, " | ")
	gifts, err := SearchGiftsDb(m.Conn, formattedSearchTerm, minPrice, maxPrice)

	if err != nil {
		return nil, err
	}

	return gifts, nil
}
func (m *PgModel) DeleteGiftCollection(id int64) error {

	err := DeleteGiftCollectionFromDb(m.Conn, id)

	if err != nil {
		return err
	}

	return nil
}
func (m *PgModel) AllCollections() ([]GiftCollection, error) {
	collections, err := GetAllCollectionsFromDB(m.Conn)

	if err != nil {
		return []GiftCollection{}, err
	}
	return collections, nil
}

func (m *PgModel) AllCustomerCollections(id int64) ([]GiftCollection, error) {
	collections, err := GetAllCustomerCollectionsFromDB(m.Conn, id)

	if err != nil {
		return []GiftCollection{}, err
	}
	return collections, nil
}

func (m *PgModel) IncompleteRequests() ([]GiftRequest, error) {
	gifts, err := GetIncompleteGiftRequestsFromDB(m.Conn)

	if err != nil {
		return []GiftRequest{}, err
	}
	return gifts, nil
}

func (m *PgModel) CompleteRequests() ([]GiftRequest, error) {
	gifts, err := GetCompleteGiftRequestsFromDB(m.Conn)

	if err != nil {
		return []GiftRequest{}, err
	}
	return gifts, nil
}

func (m *PgModel) AddGiftToGiftCollection(inputGift Gift, id int64) (GiftCollection, error) {

	giftAddedCollection, err := AddGiftToCollectionFromDB(m.Conn, inputGift, id)

	if err != nil {
		return GiftCollection{}, err
	}

	return giftAddedCollection, nil
}

func (m *PgModel) AddGiftToCustomerCollection(gift Gift, collectionName string, customerId int64) (GiftCollection, error) {

	giftAddedCollection, err := AddGiftToCustomerCollectionFromDB(m.Conn, gift, collectionName, customerId);

	if err != nil {
		return GiftCollection{}, err
	}

	return giftAddedCollection, nil
}

func (m *PgModel) DeleteGiftFromGiftCollection(giftID int64, giftCollectionID int64) (GiftCollection, error) {

	giftDeletedCollection, err := DeleteGiftFromCollectionFromDB(m.Conn, giftID, giftCollectionID)

	if err != nil {
		return GiftCollection{}, err
	}

	return giftDeletedCollection, nil
}

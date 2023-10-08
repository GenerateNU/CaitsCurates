package model

import (
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

func (m *PgModel) UpdateGift(id int64, inputGift Gift) (Gift, error) {

	updatedGift, err := UpdateGiftToDb(m.Conn, id, inputGift)

	if err != nil {
		return Gift{}, err
	}

	return updatedGift, nil
}

func (m *PgModel) DeleteGift(id int64) (error) {

	err := DeleteGiftFromDb(m.Conn, id)

	if err != nil {
		return err
	}

	return nil
}


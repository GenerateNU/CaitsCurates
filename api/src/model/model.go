package model

import (
	"gorm.io/gorm"
)

type PgModel struct {
	Conn *gorm.DB
}

type Model interface {
	GetExampleGift(int64) ExampleGift
	AllExampleGifts() ([]ExampleGift, error)
	AddExampleGift(ExampleGiftInput) (ExampleGift, error)
	IncompleteRequests() ([]GiftRequest, error)
	CompleteRequests() ([]GiftRequest, error)
}

func (m *PgModel) GetExampleGift(id int64) ExampleGift {
	gift, err := GetExampleGiftFromDB(m.Conn, id)

	if err != nil {
		panic(err)
	}

	return gift
}

func (m *PgModel) AddExampleGift(inputGift ExampleGiftInput) (ExampleGift, error) {

	createdGift, err := WriteExampleGiftToDb(m.Conn, inputGift)

	if err != nil {
		return ExampleGift{}, err
	}

	return createdGift, nil
}

func (m *PgModel) AllExampleGifts() ([]ExampleGift, error) {
	gifts, err := GetAllExampleGiftsFromDB(m.Conn)

	if err != nil {
		return []ExampleGift{}, err
	}
	return gifts, nil
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

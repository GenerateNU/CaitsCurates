package model

import (
	"github.com/jackc/pgx"
)

type PgModel struct {
	Conn *pgx.Conn
}

type Model interface {
	GetExampleGift(int64) ExampleGift
	AllExampleGifts() ([]ExampleGift, error)
	AddExampleGift(ExampleGift) (ExampleGift, error)
}

func (m *PgModel) GetExampleGift(id int64) ExampleGift {
	gift, err := GetExampleGiftFromDB(m.Conn, id)

	if err != nil {
		panic(err)
	}

	return gift
}

func (m *PgModel) AddExampleGift(eg ExampleGift) (ExampleGift, error) {
	g, err := WriteExampleGiftToDb(m.Conn, eg)

	if err != nil {
		return ExampleGift{}, err
	}

	return g, nil
}

func (m *PgModel) AllExampleGifts() ([]ExampleGift, error) {
	gifts, err := GetAllExampleGiftsFromDB(m.Conn)

	if err != nil {
		return []ExampleGift{}, err
	}
	return gifts, nil
}

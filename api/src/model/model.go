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
	book, err := GetExampleGiftFromDB(m.Conn, id)

	if err != nil {
		panic(err)
	}

	return book
}

func (m *PgModel) AddExampleGift(eg ExampleGift) (ExampleGift, error) {
	b, err := WriteExampleGiftToDb(m.Conn, eg)

	if err != nil {
		return ExampleGift{}, err
	}

	return b, nil
}

func (m *PgModel) AllExampleGifts() ([]ExampleGift, error) {
	books, err := GetAllExampleGiftsFromDB(m.Conn)

	if err != nil {
		return []ExampleGift{}, err
	}
	return books, nil
}

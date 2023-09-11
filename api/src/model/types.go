package model

type ExampleGift struct {
	GiftId int64  `json:"id" db:"gift_id"`
	Name  string `json:"name" db:"name"`
	Price int `json:"price" db:"price"`
}


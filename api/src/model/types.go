package model

type ExampleGift struct {
	ID    uint64 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	Name  string `gorm:"column:name" json:"name"`
	Price int    `gorm:"column:price" json:"price"`
}

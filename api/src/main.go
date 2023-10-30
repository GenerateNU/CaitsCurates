package main

import (
	"CaitsCurates/backend/src/controller"
	"CaitsCurates/backend/src/model"
	"fmt"
	"github.com/lib/pq"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dbURL, exists := os.LookupEnv("DATABASE_URL")
	if !exists {
		dbURL = "host=db user=user password=pwd dbname=CaitsDB port=5432 sslmode=disable"
	}
	NewLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      false,
		},
	)
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{
		Logger: NewLogger,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	// Auto migrate tables
	err = db.AutoMigrate(model.User{}, model.Customer{}, model.GiftRequest{}, model.GiftCollection{}, model.GiftResponse{}, model.Admin{})

	// Create GiftRequest
	giftRequestTimmy := model.GiftRequest{
		RecipientName:      "Timmy White",
		RecipientAge:       15,
		Occasion:           pq.StringArray{"Birthday"},
		RecipientInterests: pq.StringArray{"Toys", "Halloween"},
		BudgetMax:          60,
		BudgetMin:          20,
		DateNeeded:         time.Date(2023, time.November, 10, 0, 0, 0, 0, time.UTC),
	}
	giftRequestJoanne := model.GiftRequest{
		RecipientName:      "Joanne Burgenson",
		RecipientAge:       27,
		Occasion:           pq.StringArray{"Housewarming Gift"},
		RecipientInterests: pq.StringArray{"Decor", "Fall"},
		BudgetMax:          120,
		BudgetMin:          40,
		DateNeeded:         time.Date(2023, time.November, 16, 0, 0, 0, 0, time.UTC),
	}
	giftRequestDaniel := model.GiftRequest{
		RecipientName:      "Daniel Danielson",
		RecipientAge:       60,
		Occasion:           pq.StringArray{"Anniversary"},
		RecipientInterests: pq.StringArray{"Animals", "Golf"},
		BudgetMax:          180,
		BudgetMin:          60,
		DateNeeded:         time.Date(2023, time.December, 16, 0, 0, 0, 0, time.UTC),
	}
	giftRequestChrista := model.GiftRequest{
		RecipientName:      "Christa Blue",
		RecipientAge:       22,
		Occasion:           pq.StringArray{"Graduating College"},
		RecipientInterests: pq.StringArray{"Robots", "Jewelery", "Surfing"},
		BudgetMax:          500,
		BudgetMin:          200,
		DateNeeded:         time.Date(2023, time.November, 1, 0, 0, 0, 0, time.UTC),
	}
	a := model.User{Email: "tommy@northeastern.edu", FirstName: "Tommy", LastName: "White", Password: "xxxxx"}
	b := model.User{Email: "david@northeastern.edu", FirstName: "David", LastName: "Davidson", Password: "xxxxx"}
	e := model.User{Email: "jordan@northeastern.edu", FirstName: "Jordan", LastName: "Daniels", Password: "xxxxx"}
	d := model.User{Email: "lisa@northeastern.edu", FirstName: "Lisa", LastName: "Blue", Password: "xxxxx"}

	customer1 := model.Customer{GiftRequests: []*model.GiftRequest{&giftRequestTimmy}, User: a}
	customer2 := model.Customer{GiftRequests: []*model.GiftRequest{&giftRequestJoanne}, User: b}
	customer3 := model.Customer{GiftRequests: []*model.GiftRequest{&giftRequestDaniel}, User: e}
	customer4 := model.Customer{GiftRequests: []*model.GiftRequest{&giftRequestChrista}, User: d}

	err = db.Create(&customer1).Error
	err = db.Create(&customer2).Error
	err = db.Create(&customer3).Error
	err = db.Create(&customer4).Error

	toyGift1 := model.Gift{
		Name:            "Robot Extreme",
		Price:           50.00,
		Link:            "link.robot.com",
		Description:     "Robot Toy With Laser Eyes",
		Demographic:     "For Kids",
		GiftCollections: nil,
	}
	toyGift2 := model.Gift{
		Name:            "Angry Teddy Bear",
		Price:           20.00,
		Link:            "link.evilTeddy.com",
		Description:     "A Teddy Bear Toy but Evil!",
		Demographic:     "For Kids",
		GiftCollections: nil,
	}
	fallGift1 := model.Gift{
		Name:            "Pumpkin Sweater",
		Price:           40.00,
		Link:            "link.Sweater.com",
		Description:     "Perfect festive sweater for fall weather",
		Demographic:     "All",
		GiftCollections: nil,
	}

	fallGift2 := model.Gift{
		Name:            "Cozy Fall Mug",
		Price:           10.00,
		Link:            "link.Mug.com",
		Description:     "Insulated festive mug perfect for warm drinks",
		Demographic:     "All",
		GiftCollections: nil,
	}
	randomGift1 := model.Gift{
		Name:            "Burger Seasoning",
		Price:           5.00,
		Link:            "link.Burger.com",
		Description:     "Great for grill masters looking to up their game",
		Demographic:     "All",
		GiftCollections: nil,
	}
	randomGift2 := model.Gift{
		Name:            "Electric Scooter",
		Price:           500,
		Link:            "link.Scooter.com",
		Description:     "Crazy good for commuting and charges in 30 minutes!",
		Demographic:     "All",
		GiftCollections: nil,
	}
	decorativeGift1 := model.Gift{
		Name:            "Cool Lamp",
		Price:           15,
		Link:            "link.Lamp.com",
		Description:     "People look at this lamp and say, Man, what a cool lamp!",
		Demographic:     "Lamp Lovers",
		GiftCollections: nil,
	}
	decorativeGift2 := model.Gift{
		Name:            "Tasteful Rug",
		Price:           70,
		Link:            "link.Rug.com",
		Description:     "This rug is perfect for those looking for a cozy but stylish addition to the home",
		Demographic:     "All",
		GiftCollections: nil,
	}
	decorativeGift3 := model.Gift{
		Name:            "Spooky Chandelier",
		Price:           100,
		Link:            "link.Chandelier.com",
		Description:     "This Chandelier brings the spooky fall vibe you've been missing, worth every dollar",
		Demographic:     "All",
		GiftCollections: nil,
	}
	decorativeGift4 := model.Gift{
		Name:            "Plastic Geese",
		Price:           1,
		Link:            "link.Geese.com",
		Description:     "Decorate you yard with these and scare off any unwanted visitors",
		Demographic:     "All",
		GiftCollections: nil,
	}
	giftCollectionToy := model.GiftCollection{
		CollectionName: "Cool Toys",
		Gifts:          []*model.Gift{&toyGift2, &toyGift1},
	}

	giftCollectionFall := model.GiftCollection{
		CollectionName: "Fall Favorites",
		Gifts:          []*model.Gift{&fallGift1, &fallGift2},
	}
	giftCollectionDecor := model.GiftCollection{
		CollectionName: "Decor Enhancers",
		Gifts:          []*model.Gift{&decorativeGift3, &decorativeGift4, &decorativeGift2, &decorativeGift1},
	}
	err = db.Create(&giftCollectionToy).Error
	err = db.Create(&giftCollectionFall).Error
	err = db.Create(&giftCollectionDecor).Error
	err = db.Create(&randomGift1).Error
	err = db.Create(&randomGift2).Error

	// Check for errors

	if err != nil {
		fmt.Println("Error auto-migrating:", err)
		return
	}

	m := &model.PgModel{
		Conn: db,
	}

	c := &controller.PgController{
		Model: m,
	}

	c.Serve().Run(":8080")
}

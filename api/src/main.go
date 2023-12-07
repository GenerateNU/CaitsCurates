package main

import (
	"CaitsCurates/backend/src/controller"
	"CaitsCurates/backend/src/model"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/stripe/stripe-go/v76"

	"github.com/lib/pq"
	"gorm.io/gorm/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	stripe.Key = "sk_test_51O9y33LQbsCsABA6zhBXSCI821p8f6y9O4e4B9Pnavh5QwAgipJlUHY0iWPf3ZCfYjoxgzQYhkaW5n1PW6jwtfFu00CPOVBlrr"
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
	err = db.AutoMigrate(model.User{}, model.Customer{}, model.GiftRequest{}, model.GiftCollection{}, model.GiftResponse{}, model.Admin{}, model.Giftee{})

	// Create GiftRequest
	gifteeTimmy := model.Giftee{
		Model:                gorm.Model{},
		GifteeName:           "Timmy",
		CustomerID:           1,
		Gender:               "Male",
		CustomerRelationship: "Son",
		Age:                  17,
		Colors:               pq.StringArray{"Red", "Blue"},
		Interests:            pq.StringArray{"Toys", "Monkey"},
	}

	a := model.User{Email: "tommy@northeastern.edu", FirstName: "Tommy", LastName: "White", Password: "xxxxx"}
	customer1 := model.Customer{User: a}

	err = db.Create(&customer1).Error
	if err != nil {
		panic(err)
	}
	err = db.Create(&gifteeTimmy).Error
	if err != nil {
		panic(err)
	}

	toyGift1 := model.Gift{
		Name:            "Robot Extreme",
		Price:           50.00,
		Link:            "link.robot.com",
		Description:     "Robot Toy With Laser Eyes",
		Demographic:     "For kids",
		GiftCollections: nil,
		Occasion:        "Birthday",
		Category:        pq.StringArray{"Fun"},
	}
	toyGift2 := model.Gift{
		Name:            "Angry Teddy Bear",
		Price:           20.00,
		Link:            "link.evilTeddy.com",
		Description:     "A Teddy Bear Toy but Evil!",
		Demographic:     "For kids",
		GiftCollections: nil,
		Occasion:        "New baby",
		Category:        pq.StringArray{"Cooking", "Warm and cozy"},
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
		Demographic:     "For dad",
		Category:        pq.StringArray{"Home", "Cooking"},
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
		Demographic:     "For mom",
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
	giftCollectionFavorites := model.GiftCollection{
		CustomerID:     &customer1.UserID,
		CollectionName: "Favorites",
		Gifts:          []*model.Gift{},
	}
	err = db.Create(&giftCollectionToy).Error
	err = db.Create(&giftCollectionFall).Error
	err = db.Create(&giftCollectionDecor).Error
	err = db.Create(&giftCollectionFavorites).Error
	err = db.Create(&randomGift1).Error
	err = db.Create(&randomGift2).Error

	giftResponseTimmy := model.GiftResponse{
		GiftCollection: giftCollectionDecor,
		GiftCollectionID: 3,
		CustomMessage: "Birthday decor for the elated Timmy",
	}

	giftRequestTimmy := model.GiftRequest{
		Occasion:   pq.StringArray{"Birthday"},
		GifteeID:   gifteeTimmy.ID,
		CustomerID: customer1.ID,
		BudgetMax:  60,
		BudgetMin:  20,
		DateNeeded: time.Date(2023, time.December, 20, 0, 0, 0, 0, time.UTC),
		GiftResponse: &giftResponseTimmy,
		GiftResponseID: &giftResponseTimmy.ID,
	}

	err = db.Create(&giftRequestTimmy).Error
	if err != nil {
		panic(err)
	}

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

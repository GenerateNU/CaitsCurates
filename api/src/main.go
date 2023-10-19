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
	testGift := model.Gift{
		Name:            "gift1",
		Price:           50,
		Link:            "link1",
		Description:     "description1",
		Demographic:     "demogrpahic1",
		GiftCollections: nil,
	}
	err = db.Create(&testGift).Error
	// Create GiftRequest
	giftRequest := model.GiftRequest{
		RecipientName:      "Timmy White",
		RecipientAge:       15,
		Occasion:           pq.StringArray{"Brithday"},
		RecipientInterests: pq.StringArray{"Swimming", "Hiking"},
		BudgetMax:          50,
		BudgetMin:          20,
		DateNeeded:         time.Date(2023, time.November, 10, 0, 0, 0, 0, time.UTC),
	}
	user := model.User{Email: "example1@northeastern.edu", FirstName: "Lauren", LastName: "White", Password: "xxxxx"}
	customer := model.Customer{GiftRequests: []*model.GiftRequest{&giftRequest}, User: user}
	err = db.Create(&customer).Error
	gift1 := model.Gift{
		Name:            "Robot Extreme",
		Price:           50.00,
		Link:            "link.robot.com",
		Description:     "Robot Toy With Laser Eyes",
		Demographic:     "For Kids",
		GiftCollections: nil,
	}

	gift2 := model.Gift{
		Name:            "Angry Teddy Bear",
		Price:           20.00,
		Link:            "link.evilTeddy.com",
		Description:     "A Teddy Bear Toy but Evil!",
		Demographic:     "For Kids",
		GiftCollections: nil,
	}
	gift3 := model.Gift{
		Name:            "Pumpkin Sweater",
		Price:           40.00,
		Link:            "link.Sweater.com",
		Description:     "Perfect festive sweater for fall weather",
		Demographic:     "All",
		GiftCollections: nil,
	}

	gift4 := model.Gift{
		Name:            "Cozy Fall Mug",
		Price:           10.00,
		Link:            "link.Mug.com",
		Description:     "Insulated festive mug perfect for warm drinks",
		Demographic:     "All",
		GiftCollections: nil,
	}
	gift5 := model.Gift{
		Name:            "Burger Seasoning",
		Price:           5.00,
		Link:            "link.Burger.com",
		Description:     "Great for grill masters looking to up their game",
		Demographic:     "All",
		GiftCollections: nil,
	}
	giftCollection := model.GiftCollection{
		CollectionName: "Cool Toys",
		Gifts:          []*model.Gift{&gift1, &gift2},
	}
	err = db.Create(&giftCollection).Error
	giftCollection2 := model.GiftCollection{
		CollectionName: "Fall Favorites",
		Gifts:          []*model.Gift{&gift3, &gift4},
	}
	err = db.Create(&giftCollection2).Error
	err = db.Create(&gift5).Error

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

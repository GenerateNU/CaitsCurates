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
		Name:            "Super cool new toy",
		Price:           500.00,
		Link:            "https://www.youtube.com/watch?v=dQw4w9WgXcQ",
		Description:     "Really great content. Highly recommend",
		Demographic:     "Unknown...",
		GiftCollections: nil,
	}

	gift2 := model.Gift{
		Name:            "Super cool new toy 2",
		Price:           2.00,
		Link:            "https://www.youtube.com/Penguinz0",
		Description:     "Really great content. Highly recommend",
		Demographic:     "Unknown...",
		GiftCollections: nil,
	}

	giftCollection := model.GiftCollection{
		CollectionName: "Cool Toys",
		Gifts:          []*model.Gift{&gift1, &gift2},
	}
	err = db.Create(&giftCollection).Error
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

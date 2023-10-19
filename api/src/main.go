package main

import (
	"CaitsCurates/backend/src/controller"
	"CaitsCurates/backend/src/model"
	"fmt"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/lib/pq"
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

	// Check for errors
	if err != nil {
		fmt.Println("Error auto-migrating:", err)
		return
	}


	// New Mock data for Backend 
	// Gift Data
	mockGift1 := model.Gift {
		Name: "Winter Scarfs",
		Price: 16, 
		Link: "https://www.amazon.com/Cashmere-Pashmina-Shawls-Thicker-Scarves/dp/B07VQ5QNYT/ref=sr_1_1_sspa?keywords=scarf&qid=1697666678&sr=8-1-spons&sp_csd=d2lkZ2V0TmFtZT1zcF9hdGY&psc=1",
		Description: "The scarf is made of soft and cozy fabric, good hand feel just like cashmere. Large, warm and thick but lightweight, Especially perfect for a chilly outdoor night.",
		Demographic: "Adults",
		GiftCollections: nil,
	}
	err = db.Create(&mockGift1).Error
	if err != nil {
		fmt.Println("Error creating the gift:", err)
	}

	mockGift2 := model.Gift {
		Name: "Indoor-Outdoor Basketball",
		Price: 47, 
		Link: "https://www.amazon.com/Spalding-Inches-Official-Basketball-Orange/dp/B00F9KVDFW/ref=sr_1_1_sspa?crid=3RPRQ1Y9ST177&keywords=basketball&qid=1697666944&sprefix=basketball%2Caps%2C80&sr=8-1-spons&sp_csd=d2lkZ2V0TmFtZT1zcF9hdGY&th=1&psc=1",
		Description: "Official NBA size and weight basketball.",
		Demographic: "Adults",
		GiftCollections: nil,
	}
	err = db.Create(&mockGift2).Error
	if err != nil {
		fmt.Println("Error creating the gift:", err)
	}

	mockGift3 := model.Gift {
		Name: "Wireless Headphones",
		Price: 400, 
		Link: "https://www.amazon.com/Sennheiser-Consumer-Audio-Momentum-Headphones/dp/B0CDH415QV/ref=sr_1_7_sspa?crid=1E0BC5HLJ4NIX&keywords=sennheiser%2Bheadphones&qid=1697667100&sprefix=sennhe%2Caps%2C78&sr=8-7-spons&ufe=app_do%3Aamzn1.fos.c3015c4a-46bb-44b9-81a4-dc28e6d374b3&sp_csd=d2lkZ2V0TmFtZT1zcF9tdGY&th=1",
		Description: "Wireless Headphones - Bluetooth Headset for Crystal-Clear Calls with Adaptive Noise Cancellation, 60h Battery Life, Lightweight Folding Design.",
		Demographic: "Adults",
		GiftCollections: nil,
	}
	err = db.Create(&mockGift3).Error
	if err != nil {
		fmt.Println("Error creating the gift:", err)
	}

	// User Data
	mockUserAdmin := model.User {
		Email: "caitCurates@gmail.com",
		FirstName: "Cait",
		LastName: "Curates",
		Password: "adminPassword123",
	}
	err = db.Create(&mockUserAdmin).Error
	if err != nil {
		fmt.Println("Error creating Admin user:", err)
	} 

	mockUserCustomer1 := model.User {
		Email: "taylorSwift@gmail.com",
		FirstName: "Taylor",
		LastName: "Swift",
		Password: "customerPassword512",
	}
	err = db.Create(&mockUserCustomer1).Error
	if err != nil {
		fmt.Println("Error creating Customer user:", err)
	} 

	mockUserCustomer2 := model.User {
		Email: "morganFreeman@gmail.com",
		FirstName: "Morgan",
		LastName: "Freeman",
		Password: "customerPassword4215",
	}
	err = db.Create(&mockUserCustomer2).Error
	if err != nil {
		fmt.Println("Error creating Customer user:", err)
	} 

	// Admin Data
	mockAdmin := model.Admin {
		User: mockUserAdmin,
	}
	err = db.Create(&mockAdmin).Error
	if err != nil {
		fmt.Println("Error creating Admin:", err)
	} 

	// Customer Data
	mockCustomer1 := model.Customer {
		User: mockUserCustomer1,
	}
	err = db.Create(&mockCustomer1).Error
	if err != nil {
		fmt.Println("Error creating Admin:", err)
	} 

	mockCustomer2 := model.Customer {
		User: mockUserCustomer2,
	}
	err = db.Create(&mockCustomer2).Error
	if err != nil {
		fmt.Println("Error creating Admin:", err)
	} 

	// GiftRequest Data
	mockGiftRequest := model.GiftRequest{
		CustomerID: mockCustomer1.ID,         
		RecipientName: "John Doe",
		RecipientAge: 24,
		Occasion: pq.StringArray{"Birthday"},
		RecipientInterests: pq.StringArray{"Books", "Music", "Sports"},
		BudgetMax: 50,
		BudgetMin: 30,
		DateNeeded: time.Date(time.Now().Year(), time.October, 19, 12, 0, 0, 0, time.UTC),
	}
	err = db.Create(&mockGiftRequest).Error
	if err != nil {
		fmt.Println("Error creating GiftRequest:", err)
	}

	// GiftCollection Data
	mockGiftCollection := model.GiftCollection{
		CustomerID: &mockCustomer1.ID,
		CollectionName: "Birthday Gifts",
		Gifts: []*model.Gift{
			&mockGift2, &mockGift3,
		},
	}
	err = db.Create(&mockGiftCollection).Error
	if err != nil {
		fmt.Println("Error creating GiftCollection:", err)
	}

	// GiftResponse
	mockGiftResponse := model.GiftResponse{
		GiftCollection: mockGiftCollection,
		CustomMessage: "Here are some great gift ideas for the Birthday!",
	}
	err = db.Create(&mockGiftResponse).Error
	if err != nil {
		fmt.Println("Error creating GiftCollection:", err)
	}


	m := &model.PgModel{
		Conn: db,
	}
	c := &controller.PgController{
		Model: m,
	}

	c.Serve().Run(":8080")
}

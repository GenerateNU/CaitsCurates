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
	user := model.User{}
	err = db.Create(&user).Error
	var retrievedUser model.User
	err = db.First(&retrievedUser).Error
	customer := model.Customer{
		User: retrievedUser,
	}
	err = db.Create(&customer).Error
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

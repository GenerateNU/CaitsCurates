package tests

import (
	"CaitsCurates/backend/src/model"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestDBConnection(t *testing.T) {
	dsn := "user=testuser password=testpwd host=localhost port=5433 dbname=testdb sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		t.Fatalf("failed to connect to the database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		t.Fatalf("failed to get database instance: %v", err)
	}

	defer sqlDB.Close()

	err = sqlDB.Ping()
	if err != nil {
		t.Fatalf("failed to ping the database: %v", err)
	}
}

func TestGiftModel(t *testing.T) {
	// This code should be the same for each test
	dsn := "user=testuser password=testpwd host=localhost port=5433 dbname=testdb sslmode=disable"
	if dbURL, exists := os.LookupEnv("TEST_DATABASE_URL"); exists {
		dsn = dbURL
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Unable to connect to database: %v", err)
	}
	// Put auto migrations here
	err = db.AutoMigrate(&model.Gift{})
	if err != nil {
		panic("failed to migrate test database schema")
	}
	// Setup db rollback to revert db changes
	tx := db.Begin()
	defer tx.Rollback()

	// Create Gift
	gift := model.Gift{Name: "Super cool new toy", Price: 500000.00, Link: "https://www.youtube.com/watch?v=dQw4w9WgXcQ", Description: "Really great content. Highly recommend", Demographic: "Unknown..."}

	err = tx.Create(&gift).Error
	assert.NoError(t, err)

	// Check if Gift exists
	var fetchedGift model.Gift
	err = tx.First(&fetchedGift, gift.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, gift.ID, fetchedGift.ID)
	assert.Equal(t, gift.Name, fetchedGift.Name)
	assert.Equal(t, gift.Price, fetchedGift.Price)
	assert.Equal(t, gift.Link, fetchedGift.Link)
	assert.Equal(t, gift.Description, fetchedGift.Description)
	assert.Equal(t, gift.Demographic, fetchedGift.Demographic)
	assert.Equal(t, gift.Link, fetchedGift.Link)
	assert.Equal(t, gift.CreatedAt.In(time.UTC).Round(time.Millisecond),
		fetchedGift.CreatedAt.In(time.UTC).Round(time.Millisecond))

	// Update Gift
	err = tx.Model(&fetchedGift).Update("Name", "Slightly less cool older toy").Error
	assert.NoError(t, err)

	// Check if it's updated
	var updatedGift model.Gift
	err = tx.First(&updatedGift, fetchedGift.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, "Slightly less cool older toy", updatedGift.Name)

	// Delete Gift
	err = tx.Delete(&updatedGift).Error
	assert.NoError(t, err)

	//  Check if it's Gift
	var count int64
	tx.Model(&model.Gift{}).Where("id = ?", updatedGift.ID).Count(&count)
	assert.Equal(t, int64(0), count)
}

func TestGiftRequestModel(t *testing.T) {
	// This code should be the same for each test
	dsn := "user=testuser password=testpwd host=localhost port=5433 dbname=testdb sslmode=disable"
	if dbURL, exists := os.LookupEnv("TEST_DATABASE_URL"); exists {
		dsn = dbURL
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Unable to connect to database: %v", err)
	}
	// Put auto migrations here
	err = db.AutoMigrate(&model.GiftRequest{}, &model.GiftResponse{}, &model.GiftCollection{})
	if err != nil {
		panic("failed to migrate test admin database schema")
	}

	// Setup db rollback to revert db changes
	tx := db.Begin()
	defer tx.Rollback()

	// Create GiftResponse
	giftResponse := model.GiftResponse{CustomMessage: "This is a custom message", GiftCollection: model.GiftCollection{CollectionName: "Name"}}
	err = tx.Create(&giftResponse).Error
	assert.NoError(t, err)

	// Create GiftRequest
	giftRequest := model.GiftRequest{GiftResponse: giftResponse}
	err = tx.Create(&giftRequest).Error
	assert.NoError(t, err)

	// Check Relationship between GiftRequest and GiftResponse
	var giftRequests []model.GiftRequest
	err = tx.Model(&model.GiftRequest{}).Preload("GiftResponse").Find(&giftRequests).Error
	if err != nil {
		panic("relationship failed")
	}

	// Check GiftResponse information
	giftRequestsResponse := giftRequests[0].GiftResponse
	var fetchedGiftResponse model.GiftResponse
	err = tx.First(&fetchedGiftResponse, giftResponse.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, giftRequestsResponse.ID, fetchedGiftResponse.ID)
	assert.Equal(t, giftRequestsResponse.CustomMessage, fetchedGiftResponse.CustomMessage)
	assert.Equal(t, giftRequestsResponse.CreatedAt.In(time.UTC).Round(time.Millisecond),
		fetchedGiftResponse.CreatedAt.In(time.UTC).Round(time.Millisecond))

	// Check if GiftRequest exists
	var fetchedGiftRequest model.GiftRequest
	err = tx.First(&fetchedGiftRequest, giftRequest.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, giftRequest.ID, fetchedGiftRequest.ID)
	assert.Equal(t, giftRequest.GiftResponseID, fetchedGiftRequest.GiftResponseID)
	assert.Equal(t, giftRequest.CreatedAt.In(time.UTC).Round(time.Millisecond),
		fetchedGiftRequest.CreatedAt.In(time.UTC).Round(time.Millisecond))

	// Delete GiftRequest
	err = tx.Delete(&giftRequest).Error
	assert.NoError(t, err)

	//  Check if it's GiftRequest
	var count int64
	tx.Model(&model.GiftRequest{}).Where("id = ?", giftRequest.ID).Count(&count)
	assert.Equal(t, int64(0), count)
}

func TestGiftCollectionModel(t *testing.T) {
	// This code should be the same for each test
	dsn := "user=testuser password=testpwd host=localhost port=5433 dbname=testdb sslmode=disable"
	if dbURL, exists := os.LookupEnv("TEST_DATABASE_URL"); exists {
		dsn = dbURL
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Unable to connect to database: %v", err)
	}
	// Put auto migrations here
	err = db.AutoMigrate(&model.GiftCollection{}, &model.Gift{})
	if err != nil {
		panic("failed to migrate test database schema")
	}
	// Setup db rollback to revert db changes
	tx := db.Begin()
	defer tx.Rollback()

	// Creating Gifts
	gift1 := model.Gift{
		Name:            "Super cool new toy",
		Price:           500000.00,
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

	// Create the GiftCollection
	err = tx.Create(&giftCollection).Error
	assert.NoError(t, err)

	// Check Relationship between GiftCollection and Gifts
	var giftCollectionRetrieved model.GiftCollection
	err = tx.Model(&model.GiftCollection{}).Preload("Gifts").Find(&giftCollectionRetrieved).Error
	if err != nil {
		panic("relationship failed")
	}

	// Check gift information
	assert.NoError(t, err)
	assert.Equal(t, giftCollection.ID, giftCollectionRetrieved.ID)
	assert.Equal(t, giftCollection.CollectionName, giftCollectionRetrieved.CollectionName)
	assert.Equal(t, giftCollection.Gifts[1].ID, giftCollectionRetrieved.Gifts[1].ID)
	assert.Equal(t, giftCollection.Gifts[1].Name, giftCollectionRetrieved.Gifts[1].Name)
	assert.Equal(t, giftCollection.Gifts[0].ID, giftCollectionRetrieved.Gifts[0].ID)
	assert.Equal(t, giftCollection.Gifts[0].Name, giftCollectionRetrieved.Gifts[0].Name)

	// Delete gift collection
	err = tx.Delete(&giftCollection).Error
	assert.NoError(t, err)

	var count int64
	tx.Model(&model.GiftCollection{}).Where("id = ?", giftCollection.ID).Count(&count)
}

func TestGiftResponseModel(t *testing.T) {
	// This code should be the same for each test
	dsn := "user=testuser password=testpwd host=localhost port=5433 dbname=testdb sslmode=disable"
	if dbURL, exists := os.LookupEnv("TEST_DATABASE_URL"); exists {
		dsn = dbURL
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Unable to connect to database: %v", err)
	}
	// Put auto migrations here
	err = db.AutoMigrate(&model.GiftResponse{}, &model.GiftCollection{}, &model.GiftRequest{})
	if err != nil {
		panic("failed to migrate test admin database schema")
	}

	// Setup db rollback to revert db changes
	tx := db.Begin()
	defer tx.Rollback()

	// Creating Gifts
	gift1 := model.Gift{
		Name:            "Super cool new toy",
		Price:           500000.00,
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

	// Create giftCollection
	giftCollection := model.GiftCollection{
		CollectionName: "Cool Toys",
		Gifts:          []*model.Gift{&gift1, &gift2},
	}
	// Save the gifts
	tx.Save(&gift1)
	tx.Save(&gift2)

	// Then save the gift collection
	tx.Save(&giftCollection)
	// Create GiftResponse
	giftResponse := model.GiftResponse{CustomMessage: "This is a custom message", GiftCollection: giftCollection}
	err = tx.Create(&giftResponse).Error
	assert.NoError(t, err)

	// Check Relationship between GiftCollection and GiftResponse
	var giftResponseRetrieved model.GiftResponse
	err = tx.Model(&model.GiftResponse{}).
		Preload("GiftCollection").
		Preload("GiftCollection.Gifts").
		Find(&giftResponseRetrieved).Error
	if err != nil {
		panic("relationship failed")
	}

	// Check GiftResponse information

	assert.NoError(t, err)
	assert.Equal(t, giftResponseRetrieved.ID, giftResponse.ID)
	assert.Equal(t, giftResponseRetrieved.GiftCollection.CollectionName, giftCollection.CollectionName)
	assert.Equal(t, giftResponseRetrieved.GiftCollection.ID, giftCollection.ID)
	assert.Equal(t, giftResponseRetrieved.GiftCollection.Gifts[0].ID, giftCollection.Gifts[0].ID)
	assert.Equal(t, giftResponseRetrieved.GiftCollection.Gifts[0].Name, giftCollection.Gifts[0].Name)

	// Delete GiftRequest
	err = tx.Delete(&giftResponseRetrieved).Error
	assert.NoError(t, err)

	//  Check if it's GiftRequest
	var count int64
	tx.Model(&model.GiftRequest{}).Where("id = ?", giftResponseRetrieved.ID).Count(&count)
	assert.Equal(t, int64(0), count)
}
func TestUserModel(t *testing.T) {
	// This code should be the same for each test
	dsn := "user=testuser password=testpwd host=localhost port=5433 dbname=testdb sslmode=disable"
	if dbURL, exists := os.LookupEnv("TEST_DATABASE_URL"); exists {
		dsn = dbURL
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Unable to connect to database: %v", err)
	}
	// Put auto migrations here
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		panic("failed to migrate test database schema")
	}
	// Setup db rollback to revert db changes
	tx := db.Begin()
	defer tx.Rollback()

	// Create User
	user := model.User{Email: "tsai.me@northeastern.edu", FirstName: "Joey", LastName: "Tsai", Password: "dgeeg32"}
	err = tx.Create(&user).Error
	assert.NoError(t, err)

	// Check if user exists
	var fetchedUser model.User
	err = tx.First(&fetchedUser, user.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, user.ID, fetchedUser.ID)
	assert.Equal(t, user.FirstName, fetchedUser.FirstName)
	assert.Equal(t, user.LastName, fetchedUser.LastName)
	assert.Equal(t, user.Email, fetchedUser.Email)
	assert.Equal(t, user.Password, fetchedUser.Password)
	assert.Equal(t, user.CreatedAt.In(time.UTC).Round(time.Millisecond),
		fetchedUser.CreatedAt.In(time.UTC).Round(time.Millisecond))

	// Update User
	err = tx.Model(&fetchedUser).Update("FirstName", "Dessy").Error
	assert.NoError(t, err)

	// Check if it's updated
	var updatedUser model.User
	err = tx.First(&updatedUser, fetchedUser.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, "Dessy", updatedUser.FirstName)

	// Delete user
	err = tx.Delete(&updatedUser).Error
	assert.NoError(t, err)

	//  Check if it's user
	var count int64
	tx.Model(&model.User{}).Where("id = ?", updatedUser.ID).Count(&count)
	assert.Equal(t, int64(0), count)
}

func TestAdminModel(t *testing.T) {
	// This code should be the same for each test
	dsn := "user=testuser password=testpwd host=localhost port=5433 dbname=testdb sslmode=disable"
	if dbURL, exists := os.LookupEnv("TEST_DATABASE_URL"); exists {
		dsn = dbURL
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Unable to connect to database: %v", err)
	}
	// Put auto migrations here
	err = db.AutoMigrate(&model.Admin{})
	if err != nil {
		panic("failed to migrate test database schema")
	}
	// Setup db rollback to revert db changes
	tx := db.Begin()
	defer tx.Rollback()

	// Create Admin
	admin := model.Admin{UserID: uint(1)}
	err = tx.Create(&admin).Error
	assert.NoError(t, err)

	// Check if Admin exists
	var fetchedAdmin model.Admin
	err = tx.First(&fetchedAdmin, admin.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, admin.ID, fetchedAdmin.ID)
	assert.Equal(t, admin.UserID, fetchedAdmin.UserID)
	assert.Equal(t, admin.CreatedAt.In(time.UTC).Round(time.Millisecond),
		fetchedAdmin.CreatedAt.In(time.UTC).Round(time.Millisecond))

	// Update Admin
	err = tx.Model(&fetchedAdmin).Update("UserID", uint(2)).Error
	assert.NoError(t, err)

	// Check if it's updated
	var updatedAdmin model.Admin
	err = tx.First(&updatedAdmin, fetchedAdmin.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, uint(2), updatedAdmin.UserID)

	// Delete Admin
	err = tx.Delete(&updatedAdmin).Error
	assert.NoError(t, err)

	//  Check if it's user
	var count int64
	tx.Model(&model.Admin{}).Where("id = ?", updatedAdmin.ID).Count(&count)
	assert.Equal(t, int64(0), count)
}

func TestCustomerModel(t *testing.T) {
	// This code should be the same for each test
	dsn := "user=testuser password=testpwd host=localhost port=5433 dbname=testdb sslmode=disable"
	if dbURL, exists := os.LookupEnv("TEST_DATABASE_URL"); exists {
		dsn = dbURL
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Unable to connect to database: %v", err)
	}
	// Put auto migrations here
	err = db.AutoMigrate(&model.Customer{})
	if err != nil {
		panic("failed to migrate test database schema")
	}
	// Setup db rollback to revert db changes
	tx := db.Begin()
	defer tx.Rollback()

	// Create customer
	customer := model.Customer{UserID: uint(3)}
	err = tx.Create(&customer).Error
	assert.NoError(t, err)

	// Check if customer exists
	var fetchedCustomer model.Customer
	err = tx.First(&fetchedCustomer, customer.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, customer.ID, fetchedCustomer.ID)
	assert.Equal(t, customer.UserID, fetchedCustomer.UserID)
	assert.Equal(t, customer.CreatedAt.In(time.UTC).Round(time.Millisecond),
		fetchedCustomer.CreatedAt.In(time.UTC).Round(time.Millisecond))

	// Update customer
	err = tx.Model(&fetchedCustomer).Update("UserID", uint(4)).Error
	assert.NoError(t, err)

	// Check if it's updated
	var updatedCustomer model.Customer
	err = tx.First(&updatedCustomer, fetchedCustomer.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, uint(4), updatedCustomer.UserID)

	// Delete customer
	err = tx.Delete(&updatedCustomer).Error
	assert.NoError(t, err)

	//  Check if it's user
	var count int64
	tx.Model(&model.Customer{}).Where("id = ?", updatedCustomer.ID).Count(&count)
	assert.Equal(t, int64(0), count)
}

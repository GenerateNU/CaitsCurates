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

	err = db.Create(&gift).Error
	assert.NoError(t, err)

	// Check if Gift exists
	var fetchedGift model.Gift
	err = db.First(&fetchedGift, gift.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, gift.ID, fetchedGift.ID)
	assert.Equal(t, gift.Name, fetchedGift.Name)
	assert.Equal(t, gift.Price, fetchedGift.Price)
	assert.Equal(t, gift.Link, fetchedGift.Link)
	assert.Equal(t, gift.Description, fetchedGift.Description)
	assert.Equal(t, gift.Demographic, fetchedGift.Demographic)
	assert.Equal(t, gift.CreatedAt.In(time.UTC).Round(time.Millisecond),
		fetchedGift.CreatedAt.In(time.UTC).Round(time.Millisecond))

	// Update Gift
	err = db.Model(&fetchedGift).Update("Name", "Slightly less cool older toy").Error
	assert.NoError(t, err)

	// Check if it's updated
	var updatedGift model.Gift
	err = db.First(&updatedGift, fetchedGift.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, "Slightly less cool older toy", updatedGift.Name)

	// Delete Gift
	err = db.Delete(&updatedGift).Error
	assert.NoError(t, err)

	//  Check if it's Gift
	var count int64
	db.Model(&model.Gift{}).Where("id = ?", updatedGift.ID).Count(&count)
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
	err = db.AutoMigrate(&model.GiftRequest{})
	if err != nil {
		panic("failed to migrate test database schema")
	}
	// Setup db rollback to revert db changes
	tx := db.Begin()
	defer tx.Rollback()

	// Create Gift Request
	var dateNeededBy = time.Now()
	giftRequest := model.GiftRequest{
		RecipientName: "Jacob", 
	    RecipientAge: 21, 
		Occasion: []string{"Birthday", "Graduation"},
		RecipientInterests: []string{"Board games", "Video games", "Other things"}, 
		BudgetMin: 0.00,
		BudgetMax: 500.00, 
		DateNeeded: dateNeededBy}

	err = db.Create(&giftRequest).Error
	assert.NoError(t, err)

	// Check if Gift Request exists
	var fetchedGiftRequest model.GiftRequest
	err = db.First(&fetchedGiftRequest, giftRequest.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, giftRequest.ID, fetchedGiftRequest.ID)
	assert.Equal(t, giftRequest.RecipientName, fetchedGiftRequest.RecipientName)
	assert.Equal(t, giftRequest.RecipientAge, fetchedGiftRequest.RecipientAge)
	assert.Equal(t, giftRequest.Occasion, fetchedGiftRequest.Occasion)
	assert.Equal(t, giftRequest.RecipientAge, fetchedGiftRequest.RecipientAge)
	assert.Equal(t, giftRequest.DateNeeded, fetchedGiftRequest.DateNeeded)
	assert.Equal(t, giftRequest.CreatedAt.In(time.UTC).Round(time.Millisecond),
		fetchedGiftRequest.CreatedAt.In(time.UTC).Round(time.Millisecond))

	// Update Gift Request
	err = db.Model(&fetchedGiftRequest).Update("RecipientName", "Aidan").Error
	assert.NoError(t, err)

	// Check if it's updated
	var updatedGiftRequest model.GiftRequest
	err = db.First(&updatedGiftRequest, fetchedGiftRequest.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, "Aidan", updatedGiftRequest.RecipientName)

	// Delete Gift Request
	err = db.Delete(&updatedGiftRequest).Error
	assert.NoError(t, err)

	//  Check if it's Gift Request
	var count int64
	db.Model(&model.GiftRequest{}).Where("id = ?", updatedGiftRequest.ID).Count(&count)
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
	user := model.User{Email: "example@northeastern.edu", FirstName: "PersonFirstName", LastName: "PersonLastName", Password: "dgeeg32"}
	err = db.Create(&user).Error
	assert.NoError(t, err)

	// Check if user exists
	var fetchedUser model.User
	err = db.First(&fetchedUser, user.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, user.ID, fetchedUser.ID)
	assert.Equal(t, user.FirstName, fetchedUser.FirstName)
	assert.Equal(t, user.LastName, fetchedUser.LastName)
	assert.Equal(t, user.Email, fetchedUser.Email)
	assert.Equal(t, user.Password, fetchedUser.Password)
	assert.Equal(t, user.CreatedAt.In(time.UTC).Round(time.Millisecond),
		fetchedUser.CreatedAt.In(time.UTC).Round(time.Millisecond))

	// Update User
	err = db.Model(&fetchedUser).Update("FirstName", "ChangeFirstName").Error
	assert.NoError(t, err)

	// Check if it's updated
	var updatedUser model.User
	err = db.First(&updatedUser, fetchedUser.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, "ChangeFirstName", updatedUser.FirstName)

	// Delete user
	err = db.Delete(&updatedUser).Error
	assert.NoError(t, err)

	//  Check if it's user
	var count int64
	db.Model(&model.User{}).Where("id = ?", updatedUser.ID).Count(&count)
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
		panic("failed to migrate test admin database schema")
	}

	// Put auto migrations here
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		panic("failed to migrate test user database schema")
	}

	// Setup db rollback to revert db changes
	tx := db.Begin()
	defer tx.Rollback()

	// Create User
	user := model.User{Email: "example@northeastern.edu", FirstName: "PersonFirstName", LastName: "PersonLastName", Password: "dgeeg32"}
	err = db.Create(&user).Error
	assert.NoError(t, err)

	// Create Admin
	admin := model.Admin{UserID: uint(1), User: user}
	err = db.Create(&admin).Error
	assert.NoError(t, err)

	// Check Relationship between Admin and User 
	var admins []model.Admin
	err = db.Model(&model.Admin{}).Preload("User").Find(&admins).Error
	if err != nil {
		panic("relationship failed")
	}

	// Check User information
	adminUser := admins[0].User
	var fetchedUser model.User
	err = db.First(&fetchedUser, user.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, adminUser.ID, fetchedUser.ID)
	assert.Equal(t, adminUser.FirstName, fetchedUser.FirstName)
	assert.Equal(t, adminUser.LastName, fetchedUser.LastName)
	assert.Equal(t, adminUser.Email, fetchedUser.Email)
	assert.Equal(t, adminUser.Password, fetchedUser.Password)
	assert.Equal(t, adminUser.CreatedAt.In(time.UTC).Round(time.Millisecond),
		fetchedUser.CreatedAt.In(time.UTC).Round(time.Millisecond))

	// Check if Admin exists
	var fetchedAdmin model.Admin
	err = db.First(&fetchedAdmin, admin.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, admin.ID, fetchedAdmin.ID)
	assert.Equal(t, admin.UserID, fetchedAdmin.UserID)
	assert.Equal(t, admin.CreatedAt.In(time.UTC).Round(time.Millisecond),
		fetchedAdmin.CreatedAt.In(time.UTC).Round(time.Millisecond))

	// Update Admin
	err = db.Model(&fetchedAdmin).Update("UserID", uint(2)).Error
	assert.NoError(t, err)

	// Check if it's updated
	var updatedAdmin model.Admin
	err = db.First(&updatedAdmin, fetchedAdmin.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, uint(2), updatedAdmin.UserID)

	// Delete Admin
	err = db.Delete(&updatedAdmin).Error
	assert.NoError(t, err)

	//  Check if it's admin
	var count int64
	db.Model(&model.Admin{}).Where("id = ?", updatedAdmin.ID).Count(&count)
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
		panic("failed to migrate test customer database schema")
	}

	// Put auto migrations here
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		panic("failed to migrate test user database schema")
	}

	// Setup db rollback to revert db changes
	tx := db.Begin()
	defer tx.Rollback()

	// Create User
	user := model.User{Email: "example@northeastern.edu", FirstName: "PersonFirstName", LastName: "PersonLastName", Password: "dgeeg32"}
	err = db.Create(&user).Error
	assert.NoError(t, err)

	// Create Customer
	customer := model.Customer{UserID: uint(1), User: user}
	err = db.Create(&customer).Error
	assert.NoError(t, err)

	// Check Relationship between Customer and User 
	var customers []model.Customer
	err = db.Model(&model.Customer{}).Preload("User").Find(&customers).Error
	if err != nil {
		panic("relationship failed")
	}

	// Check User information
	customerUser := customers[0].User
	var fetchedUser model.User
	err = db.First(&fetchedUser, user.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, customerUser.ID, fetchedUser.ID)
	assert.Equal(t, customerUser.FirstName, fetchedUser.FirstName)
	assert.Equal(t, customerUser.LastName, fetchedUser.LastName)
	assert.Equal(t, customerUser.Email, fetchedUser.Email)
	assert.Equal(t, customerUser.Password, fetchedUser.Password)
	assert.Equal(t, customerUser.CreatedAt.In(time.UTC).Round(time.Millisecond),
		fetchedUser.CreatedAt.In(time.UTC).Round(time.Millisecond))

	// Check if Customer exists
	var fetchedCustomer model.Customer
	err = db.First(&fetchedCustomer, customer.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, customer.ID, fetchedCustomer.ID)
	assert.Equal(t, customer.UserID, fetchedCustomer.UserID)
	assert.Equal(t, customer.CreatedAt.In(time.UTC).Round(time.Millisecond),
		fetchedCustomer.CreatedAt.In(time.UTC).Round(time.Millisecond))

	// Update Custoer
	err = db.Model(&fetchedCustomer).Update("UserID", uint(2)).Error
	assert.NoError(t, err)

	// Check if it's updated
	var updatedCustomer model.Customer
	err = db.First(&updatedCustomer, fetchedCustomer.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, uint(2), updatedCustomer.UserID)

	// Delete Customer
	err = db.Delete(&updatedCustomer).Error
	assert.NoError(t, err)

	//  Check if it's user
	var count int64
	db.Model(&model.Customer{}).Where("id = ?", updatedCustomer.ID).Count(&count)
	assert.Equal(t, int64(0), count)
}
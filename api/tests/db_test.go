package tests

import (
	"CaitsCurates/backend/src/model"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/lib/pq"
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
	gift := model.Gift{Name: "Super cool new toy", Price: 500000.00, Link: "https://www.youtube.com/watch?v=dQw4w9WgXcQ", Description: "Really great content. Highly recommend", Demographic: "Unknown...", Category: pq.StringArray{"Best selling"}}

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
	assert.Equal(t, gift.Category, fetchedGift.Category)
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
	giftRequest := model.GiftRequest{GiftResponse: &giftResponse, Occasion: pq.StringArray{"Birthday"}, RecipientInterests: pq.StringArray{"Soccer"}}
	user := model.User{Email: "example1@northeastern.edu", FirstName: "person1", LastName: "lastname1", Password: "dgeeg32"}
	customer := model.Customer{GiftRequests: []*model.GiftRequest{&giftRequest}, User: user}
	err = tx.Create(&customer).Error
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
	assert.Equal(t, giftRequest.Occasion, fetchedGiftRequest.Occasion)
	assert.Equal(t, giftRequest.RecipientInterests, fetchedGiftRequest.RecipientInterests)
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
		Category: 		 pq.StringArray{"Best selling"},
		GiftCollections: nil,
	}

	gift2 := model.Gift{
		Name:            "Super cool new toy 2",
		Price:           2.00,
		Link:            "https://www.youtube.com/Penguinz0",
		Description:     "Really great content. Highly recommend",
		Demographic:     "Unknown...",
		Category: 		 pq.StringArray{"Best selling"},
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
	user := model.User{Email: "example123@northeastern.edu", FirstName: "person1", LastName: "lastname1", Password: "dgeeg32"}
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
	err = tx.Model(&fetchedUser).Update("FirstName", "person2").Error
	assert.NoError(t, err)

	// Check if it's updated
	var updatedUser model.User
	err = tx.First(&updatedUser, fetchedUser.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, "person2", updatedUser.FirstName)

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
	admin := model.Admin{User: user}
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

	// Delete Admin
	err = db.Delete(&admin).Error
	assert.NoError(t, err)

	//  Check if it's admin
	var count int64
	db.Model(&model.Admin{}).Where("id = ?", admin.ID).Count(&count)
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
	err = db.AutoMigrate(&model.Customer{}, &model.User{}, &model.GiftCollection{}, &model.GiftRequest{})
	if err != nil {
		panic("failed to migrate test customer database schema")
	}

	// Setup db rollback to revert db changes
	tx := db.Begin()
	defer tx.Rollback()

	// Create User
	user := model.User{Email: "example@northeastern.edu", FirstName: "PersonFirstName", LastName: "PersonLastName", Password: "dgeeg32"}
	err = tx.Create(&user).Error

	// Create Collection
	collection := model.GiftCollection{CollectionName: "Collection"}
	// Create a Request
	request := model.GiftRequest{
		RecipientName: "Me",
	}
	assert.NoError(t, err)
	// Create Customer
	customer := model.Customer{User: user, GiftCollections: []*model.GiftCollection{&collection}, GiftRequests: []*model.GiftRequest{&request}}
	err = tx.Create(&customer).Error
	assert.NoError(t, err)

	// Check Relationships
	var customers []model.Customer
	err = tx.Model(&model.Customer{}).Preload("User").Preload("GiftCollections").Preload("GiftRequests").Find(&customers).Error
	if err != nil {
		panic("relationship failed")
	}

	// Check User information
	customerUser := customers[0].User
	var fetchedUser model.User
	err = tx.First(&fetchedUser, user.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, customerUser.ID, fetchedUser.ID)
	assert.Equal(t, customerUser.FirstName, fetchedUser.FirstName)
	assert.Equal(t, customer.GiftRequests[0].RecipientName, request.RecipientName)
	assert.Equal(t, customer.GiftCollections[0].CollectionName, collection.CollectionName)
	assert.Equal(t, customerUser.LastName, fetchedUser.LastName)
	assert.Equal(t, customerUser.Email, fetchedUser.Email)
	assert.Equal(t, customerUser.Password, fetchedUser.Password)
	assert.Equal(t, customerUser.CreatedAt.In(time.UTC).Round(time.Millisecond),
		fetchedUser.CreatedAt.In(time.UTC).Round(time.Millisecond))

	// Check if Customer exists
	var fetchedCustomer model.Customer
	err = tx.First(&fetchedCustomer, customer.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, customer.ID, fetchedCustomer.ID)
	assert.Equal(t, customer.UserID, fetchedCustomer.UserID)
	assert.Equal(t, customer.CreatedAt.In(time.UTC).Round(time.Millisecond),
		fetchedCustomer.CreatedAt.In(time.UTC).Round(time.Millisecond))

	// Delete Customer
	err = db.Delete(&customer).Error
	assert.NoError(t, err)

	//  Check if it's user
	var count int64
	db.Model(&model.Customer{}).Where("id = ?", customer.ID).Count(&count)
	assert.Equal(t, int64(0), count)
}

func TestGifteeModel(t *testing.T) {
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
	err = db.AutoMigrate(&model.Giftee{})
	if err != nil {
		panic("failed to migrate test database schema")
	}
	// Setup db rollback to revert db changes
	tx := db.Begin()
	defer tx.Rollback()

	// Create Giftee
	giftee := model.Giftee {
		CustomerID:            1,
		GifteeName:            "Maya",
		Gender:                "Female",
		CustomerRelationship:  "Sister",
		Age:                   20,
		Colors:                pq.StringArray{"Green", "Blue"},
		Interests:             pq.StringArray{"Sports", "Soccer", "Nature", "Coffee", "Candy"},
	}

	err = tx.Create(&giftee).Error
	assert.NoError(t, err)

	// Check if Giftee exists
	var fetchedGiftee model.Giftee
	err = tx.First(&fetchedGiftee, giftee.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, giftee.ID, fetchedGiftee.ID)
	assert.Equal(t, giftee.CustomerID, fetchedGiftee.CustomerID)
	assert.Equal(t, giftee.GifteeName, fetchedGiftee.GifteeName)
	assert.Equal(t, giftee.Gender, fetchedGiftee.Gender)
	assert.Equal(t, giftee.CustomerRelationship, fetchedGiftee.CustomerRelationship)
	assert.Equal(t, giftee.Age, fetchedGiftee.Age)
	assert.Equal(t, giftee.Colors, fetchedGiftee.Colors)
	assert.Equal(t, giftee.Interests, fetchedGiftee.Interests)
	assert.Equal(t, giftee.CreatedAt.In(time.UTC).Round(time.Millisecond),
		fetchedGiftee.CreatedAt.In(time.UTC).Round(time.Millisecond))

	// Update Giftee
	err = tx.Model(&fetchedGiftee).Update("GifteeName", "Maya Updated").Error
	assert.NoError(t, err)

	// Check if it's updated
	var updatedGiftee model.Giftee
	err = tx.First(&updatedGiftee, fetchedGiftee.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, "Maya Updated", updatedGiftee.GifteeName)

	// Delete Giftee
	err = tx.Delete(&updatedGiftee).Error
	assert.NoError(t, err)

	//  Check if it's Giftee
	var count int64
	tx.Model(&model.Giftee{}).Where("id = ?", updatedGiftee.ID).Count(&count)
	assert.Equal(t, int64(0), count)
}

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
	dsn := "user=testuser password=testpwd host=test-db port=5432 dbname=testdb sslmode=disable"
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

func TestExampleGiftModel(t *testing.T) {
	// This code should be the same for each test
	dsn := "host=test-db user=testuser password=testpwd dbname=testdb port=5433 sslmode=disable"
	if dbURL, exists := os.LookupEnv("TEST_DATABASE_URL"); exists {
		dsn = dbURL
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Unable to connect to database: %v", err)
	}
	// Put auto migrations here
	err = db.AutoMigrate(&model.ExampleGift{})
	if err != nil {
		panic("failed to migrate test database schema")
	}
	// Setup db rollback to revert db changes
	tx := db.Begin()
	defer tx.Rollback()

	// Create gift
	gift := model.ExampleGift{Name: "Ugly Sweater",
		Price: 50}
	err = db.Create(&gift).Error
	assert.NoError(t, err)

	// Check if gift exists
	var fetchedGift model.ExampleGift
	err = db.First(&fetchedGift, gift.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, gift.ID, fetchedGift.ID)
	assert.Equal(t, gift.Name, fetchedGift.Name)
	assert.Equal(t, gift.Price, fetchedGift.Price)
	assert.Equal(t, gift.CreatedAt.In(time.UTC).Round(time.Millisecond),
		fetchedGift.CreatedAt.In(time.UTC).Round(time.Millisecond))

	//  Update gift
	err = db.Model(&fetchedGift).Update("name", "Sweater").Error
	assert.NoError(t, err)

	// Check if it's updated
	var updatedGift model.ExampleGift
	err = db.First(&updatedGift, fetchedGift.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, "Sweater", updatedGift.Name)

	// Delete gift
	err = db.Delete(&updatedGift).Error
	assert.NoError(t, err)

	//  Check if it's deleted
	var count int64
	db.Model(&model.ExampleGift{}).Where("id = ?", updatedGift.ID).Count(&count)
	assert.Equal(t, int64(0), count)

}


func TestUserModel(t *testing.T) {
	// This code should be the same for each test
	dsn := "host=test-db user=testuser password=testpwd dbname=testdb port=5433 sslmode=disable"
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
	err = db.Model(&fetchedUser).Update("FirstName", "Dessy").Error
	assert.NoError(t, err)

	// Check if it's updated
	var updatedUser model.User
	err = db.First(&updatedUser, fetchedUser.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, "Dessy", updatedUser.FirstName)

	// Delete user
	err = db.Delete(&updatedUser).Error
	assert.NoError(t, err)

	//  Check if it's user
	var count int64
	db.Model(&model.User{}).Where("id = ?", updatedUser.ID).Count(&count)
	assert.Equal(t, int64(0), count)

}

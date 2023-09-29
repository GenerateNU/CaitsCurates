package tests

import (
	c "CaitsCurates/backend/src/controller"
	"CaitsCurates/backend/src/model"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"time"

	"os"
	"testing"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestAddExampleGift(t *testing.T) {
	// Database setup
	dsn := "user=testuser password=testpwd host=localhost port=5433 dbname=testdb sslmode=disable"
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
	// Wrap the DB connection in a transaction
	tx := db.Begin()
	defer tx.Rollback()

	// Create Model and Controller
	m := &model.PgModel{Conn: tx}
	c := &c.PgController{Model: m}
	router := c.Serve()

	// Test code
	w := httptest.NewRecorder()
	w1 := httptest.NewRecorder()
	testGift := model.ExampleGiftInput{
		Name:  "nice sweater",
		Price: 50,
	}
	giftJson, err := json.Marshal(testGift)
	if err != nil {
		t.Fatalf("Error marshaling JSON: %v", err)
	}

	req, err := http.NewRequest("POST", "/addGift", bytes.NewBuffer(giftJson))
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}

	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	var giftAdded model.ExampleGift
	if e := json.Unmarshal(w.Body.Bytes(), &giftAdded); e != nil {
		t.Fatalf("Error unmarshaling JSON: %v", e)
	}
	assert.Equal(t, testGift.Price, giftAdded.Price)
	assert.Equal(t, uint(1), giftAdded.ID)
	req1, err := http.NewRequest("GET", "/gifts/1", nil)
	router.ServeHTTP(w1, req1)

	assert.Equal(t, 200, w1.Code)

	var giftRetrieved model.ExampleGift
	if e := json.Unmarshal(w1.Body.Bytes(), &giftRetrieved); e != nil {
		t.Fatalf("Error unmarshaling JSON: %v", e)
	}

	assert.Equal(t, giftAdded.ID, giftRetrieved.ID)
	assert.Equal(t, giftAdded.Name, giftRetrieved.Name)
	assert.Equal(t, giftAdded.Price, giftRetrieved.Price)
	assert.Equal(t, giftAdded.CreatedAt.In(time.UTC).Round(time.Millisecond),
		giftRetrieved.CreatedAt.In(time.UTC).Round(time.Millisecond))
}
func TestGetExampleGift(t *testing.T) {
	// Database setup
	dsn := "user=testuser password=testpwd host=localhost port=5433 dbname=testdb sslmode=disable"
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
	// Wrap the DB connection in a transaction
	tx := db.Begin()
	defer tx.Rollback()

	// Create Model and Controller
	m := &model.PgModel{Conn: tx}
	c := &c.PgController{Model: m}
	router := c.Serve()

	// Test code
	w := httptest.NewRecorder()
	gift := model.ExampleGift{
		Name:  "nice sweater",
		Price: 50,
	}
	err = db.Create(&gift).Error
	assert.NoError(t, err)
	req1, err := http.NewRequest("GET", fmt.Sprintf("/gifts/%d", gift.ID), nil)
	router.ServeHTTP(w, req1)

	assert.Equal(t, 200, w.Code)

	var giftRetrieved model.ExampleGift
	if e := json.Unmarshal(w.Body.Bytes(), &giftRetrieved); e != nil {
		t.Fatalf("Error unmarshaling JSON: %v", e)
	}

	assert.Equal(t, gift.ID, giftRetrieved.ID)
	assert.Equal(t, gift.Name, giftRetrieved.Name)
	assert.Equal(t, gift.Price, giftRetrieved.Price)
	assert.Equal(t, gift.CreatedAt.In(time.UTC).Round(time.Millisecond),
		giftRetrieved.CreatedAt.In(time.UTC).Round(time.Millisecond))

}

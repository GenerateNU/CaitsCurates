package tests

import (
	c "CaitsCurates/backend/src/controller"
	"CaitsCurates/backend/src/model"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/go-playground/assert/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestAddGetGift(t *testing.T) {
	// Database setup
	dsn := "host=test-db user=testuser password=testpwd dbname=testdb port=5433 sslmode=disable"
	if dbURL, exists := os.LookupEnv("TEST_DATABASE_URL"); exists {
		dsn = dbURL
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Unable to connect to database: %v", err)
	}
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
	testGift := model.ExampleGift{
		ID:    4,
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
	assert.Equal(t, testGift, giftAdded)

	req1, err := http.NewRequest("GET", "/gifts/4", nil)
	router.ServeHTTP(w1, req1)

	assert.Equal(t, 200, w1.Code)

	var giftRetrieved model.ExampleGift
	if e := json.Unmarshal(w1.Body.Bytes(), &giftRetrieved); e != nil {
		t.Fatalf("Error unmarshaling JSON: %v", e)
	}

	assert.Equal(t, testGift, giftRetrieved)
}

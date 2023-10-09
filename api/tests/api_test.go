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
	"sort"
	"time"

	"github.com/lib/pq"

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
	req1, err := http.NewRequest("GET", fmt.Sprintf("/gifts/%d", giftAdded.ID), nil)
	router.ServeHTTP(w1, req1)

	assert.Equal(t, 200, w1.Code)

	var giftRetrieved model.ExampleGift
	var giftRetrievedDB model.ExampleGift
	if e := json.Unmarshal(w1.Body.Bytes(), &giftRetrieved); e != nil {
		t.Fatalf("Error unmarshaling JSON: %v", e)
	}
	err = tx.First(&giftRetrievedDB, giftRetrievedDB.ID).Error

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


func TestGetIncompleteGiftRequests(t *testing.T) {
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
	err = db.AutoMigrate(&model.GiftRequest{})
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
	request := model.GiftRequest{
		CustomerID:     1,
		RecipientName:  "Friend",
		RecipientAge:   25,
		Occasion:       pq.StringArray{"Birthday", "Anniversary"},
		RecipientInterests: pq.StringArray{"Reading", "Gaming"},
		BudgetMax:      50,
		BudgetMin:      15,
		// GiftResponse: leaving this out to signify an incomplete request
		DateNeeded: time.Now(),
	}

	// Create the GiftRequest and call the endpoint
	err = db.Create(&request).Error
	assert.NoError(t, err)
	req1, err := http.NewRequest("GET", fmt.Sprintf("/requests/incomplete"), nil)
	router.ServeHTTP(w, req1)

	assert.Equal(t, 200, w.Code)

	var requestsRetrieved []model.GiftRequest
	if e := json.Unmarshal(w.Body.Bytes(), &requestsRetrieved); e != nil {
		t.Fatalf("Error unmarshaling JSON: %v", e)
	}

	// Choose only the most recently created incomplete request (the one we just added)
	sort.Slice(requestsRetrieved, func(i, j int) bool {
    return requestsRetrieved[i].CreatedAt.After(requestsRetrieved[j].CreatedAt)
	})

	assert.Equal(t, request.ID, requestsRetrieved[0].ID)
	assert.Equal(t, request.RecipientName, requestsRetrieved[0].RecipientName)
	assert.Equal(t, request.RecipientAge, requestsRetrieved[0].RecipientAge)
	assert.Equal(t, request.Occasion, requestsRetrieved[0].Occasion)
	assert.Equal(t, request.RecipientInterests, requestsRetrieved[0].RecipientInterests)
	assert.Equal(t, request.BudgetMax, requestsRetrieved[0].BudgetMax)
	assert.Equal(t, request.BudgetMin, requestsRetrieved[0].BudgetMin)
	assert.Nil(t, request.GiftResponseID) // make sure it's actually incomplete!
	assert.Equal(t, request.DateNeeded.In(time.UTC).Round(time.Millisecond),
		requestsRetrieved[0].DateNeeded.In(time.UTC).Round(time.Millisecond))
}
func TestGetCompleteGiftRequests(t *testing.T) {
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
	err = db.AutoMigrate(&model.GiftRequest{})
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

	// Create GiftResponse
	giftResponse := model.GiftResponse{CustomMessage: "This is a custom message", GiftCollection: model.GiftCollection{CollectionName: "Name"}}
	err = db.Create(&giftResponse).Error
	assert.NoError(t, err)

	// Create GiftRequest
	request := model.GiftRequest{
		CustomerID:     1,
		RecipientName:  "Friend",
		RecipientAge:   25,
		Occasion:       pq.StringArray{"Birthday", "Anniversary"},
		RecipientInterests: pq.StringArray{"Reading", "Gaming"},
		BudgetMax:      50,
		BudgetMin:      15,
		GiftResponse: giftResponse,
		DateNeeded: time.Now(),
	}

	// Create the GiftRequest and call the endpoint
	err = db.Create(&request).Error
	assert.NoError(t, err)
	req1, err := http.NewRequest("GET", fmt.Sprintf("/requests/complete"), nil)
	router.ServeHTTP(w, req1)

	assert.Equal(t, 200, w.Code)

	var requestsRetrieved []model.GiftRequest
	if e := json.Unmarshal(w.Body.Bytes(), &requestsRetrieved); e != nil {
		t.Fatalf("Error unmarshaling JSON: %v", e)
	}

	// Choose only the most recently created incomplete request (the one we just added)
	sort.Slice(requestsRetrieved, func(i, j int) bool {
    return requestsRetrieved[i].CreatedAt.After(requestsRetrieved[j].CreatedAt)
	})

	assert.Equal(t, request.ID, requestsRetrieved[0].ID)
	assert.Equal(t, request.RecipientName, requestsRetrieved[0].RecipientName)
	assert.Equal(t, request.RecipientAge, requestsRetrieved[0].RecipientAge)
	assert.Equal(t, request.Occasion, requestsRetrieved[0].Occasion)
	assert.Equal(t, request.RecipientInterests, requestsRetrieved[0].RecipientInterests)
	assert.Equal(t, request.BudgetMax, requestsRetrieved[0].BudgetMax)
	assert.Equal(t, request.BudgetMin, requestsRetrieved[0].BudgetMin)
	assert.Equal(t, request.GiftResponseID, requestsRetrieved[0].GiftResponseID)
	assert.Equal(t, request.DateNeeded.In(time.UTC).Round(time.Millisecond),
		requestsRetrieved[0].DateNeeded.In(time.UTC).Round(time.Millisecond))
}
package tests

import (
	c "CaitsCurates/backend/src/controller"
	"CaitsCurates/backend/src/model"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sort"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/lib/pq"

	"os"
	"testing"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

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
	err = db.AutoMigrate(&model.GiftRequest{}, &model.Customer{}, &model.User{}, &model.GiftResponse{})
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
	user := model.User{}
	err = tx.Create(&user).Error
	assert.NoError(t, err)
	var retrievedUser model.User
	err = tx.First(&retrievedUser).Error
	assert.NoError(t, err)
	customer := model.Customer{
		User: retrievedUser,
	}
	err = tx.Create(&customer).Error
	assert.NoError(t, err)
	var retrievedCustomer model.Customer
	err = tx.First(&retrievedCustomer).Error
	request := model.GiftRequest{
		CustomerID:         retrievedCustomer.ID,
		RecipientName:      "Friend",
		RecipientAge:       25,
		Occasion:           pq.StringArray{"Birthday", "Anniversary"},
		RecipientInterests: pq.StringArray{"Reading", "Gaming"},
		BudgetMax:          50,
		BudgetMin:          15,
		DateNeeded:         time.Now(),
	}

	// Create the GiftRequest and call the endpoint
	err = tx.Create(&request).Error
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

func TestGetCustomerGiftRequests(t *testing.T) {
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
	err = db.AutoMigrate(&model.GiftCollection{}, &model.Giftee{}, &model.User{}, &model.Customer{}, &model.GiftRequest{}, &model.GiftResponse{})
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
	err = tx.Create(&giftResponse).Error
	assert.NoError(t, err)
	// Create GiftRequest
	user := model.User{}
	err = tx.Create(&user).Error
	assert.NoError(t, err)
	var retrievedUser model.User
	err = tx.First(&retrievedUser).Error
	assert.NoError(t, err)
	customer := model.Customer{
		User: retrievedUser,
	}
	err = tx.Create(&customer).Error
	assert.NoError(t, err)
	testGiftee := model.Giftee{
		CustomerID:           customer.ID,
		GifteeName:           "Maya",
		Gender:               "Female",
		CustomerRelationship: "Sister",
		Age:                  20,
		Colors:               pq.StringArray{"Green", "Blue"},
		Interests:            pq.StringArray{"Sports", "Soccer", "Nature", "Coffee", "Candy"},
	}
	err = tx.Create(&testGiftee).Error
	assert.NoError(t, err)
	var retrievedCustomer model.Customer
	err = tx.First(&retrievedCustomer).Error
	request := model.GiftRequest{
		CustomerID:   retrievedCustomer.ID,
		GiftResponse: &giftResponse,
		Occasion:     pq.StringArray{"Birthday", "Anniversary"},
		BudgetMax:    50,
		BudgetMin:    15,
		GifteeID:     testGiftee.ID,
		DateNeeded:   time.Now(),
	}
	request2 := model.GiftRequest{
		CustomerID: retrievedCustomer.ID,
		Occasion:   pq.StringArray{"Birthday", "Anniversary"},
		BudgetMax:  50,
		GifteeID:   testGiftee.ID,
		BudgetMin:  15,
		DateNeeded: time.Now(),
	}
	// Create the GiftRequest and call the endpoint
	err = tx.Create(&request).Error

	assert.NoError(t, err)
	err = tx.Create(&request2).Error
	assert.NoError(t, err)

	req1, err := http.NewRequest("GET", fmt.Sprintf("/requests/1"), nil)
	router.ServeHTTP(w, req1)

	assert.Equal(t, 200, w.Code)

	var requestsRetrieved []model.GiftRequest
	if e := json.Unmarshal(w.Body.Bytes(), &requestsRetrieved); e != nil {
		t.Fatalf("Error unmarshaling JSON: %v", e)
	}
	fmt.Println(customer.ID)
	fmt.Println(requestsRetrieved)
	assert.Equal(t, len(requestsRetrieved), 2)
	//assert.Equal(t, requestsRetrieved[0].GiftResponse, nil)
	assert.Equal(t, requestsRetrieved[0].GiftResponse.GiftCollection.CollectionName, "Name")

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
	err = db.AutoMigrate(&model.GiftRequest{}, &model.Customer{}, &model.User{}, &model.GiftResponse{})
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
	err = tx.Create(&giftResponse).Error
	assert.NoError(t, err)
	// Create GiftRequest
	user := model.User{}
	err = tx.Create(&user).Error
	assert.NoError(t, err)
	var retrievedUser model.User
	err = tx.First(&retrievedUser).Error
	assert.NoError(t, err)
	customer := model.Customer{
		User: retrievedUser,
	}
	err = tx.Create(&customer).Error
	assert.NoError(t, err)
	var retrievedCustomer model.Customer
	err = tx.First(&retrievedCustomer).Error
	request := model.GiftRequest{
		CustomerID:         retrievedCustomer.ID,
		GiftResponse:       &giftResponse,
		RecipientName:      "Friend",
		RecipientAge:       25,
		Occasion:           pq.StringArray{"Birthday", "Anniversary"},
		RecipientInterests: pq.StringArray{"Reading", "Gaming"},
		BudgetMax:          50,
		BudgetMin:          15,
		DateNeeded:         time.Now(),
	}

	// Create the GiftRequest and call the endpoint
	err = tx.Create(&request).Error

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
	assert.Equal(t, request.GiftResponse.ID, giftResponse.ID)
	assert.NotNil(t, request.GiftResponse) // make sure preloading works
	assert.Equal(t, request.GiftResponse.GiftCollectionID, giftResponse.GiftCollectionID)
	assert.Equal(t, request.GiftResponse.GiftCollection, giftResponse.GiftCollection)
	assert.Equal(t, request.GiftResponse.CustomMessage, giftResponse.CustomMessage)
	assert.Equal(t, request.DateNeeded.In(time.UTC).Round(time.Millisecond),
		requestsRetrieved[0].DateNeeded.In(time.UTC).Round(time.Millisecond))
}

func TestAddRequest(t *testing.T) {
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
	err = db.AutoMigrate(&model.GiftRequest{}, &model.Customer{}, &model.User{}, &model.GiftResponse{})
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
	assert.NoError(t, err)
	// Create GiftRequest
	user := model.User{}
	err = tx.Create(&user).Error
	assert.NoError(t, err)
	var retrievedUser model.User
	err = tx.First(&retrievedUser).Error
	assert.NoError(t, err)
	customer := model.Customer{
		User: retrievedUser,
	}
	err = tx.Create(&customer).Error
	assert.NoError(t, err)
	var retrievedCustomer model.Customer
	err = tx.First(&retrievedCustomer).Error
	request := model.GiftRequest{
		CustomerID:         retrievedCustomer.ID,
		RecipientName:      "Friend",
		Occasion:           pq.StringArray{"Birthday", "Anniversary"},
		RecipientInterests: pq.StringArray{"Reading", "Gaming"},
	}

	// Create the GiftRequest and call the endpoint
	requestJSON, err := json.Marshal(request)
	if err != nil {
		t.Fatalf("Error marshaling JSON: %v", err)
	}
	assert.NoError(t, err)
	req1, err := http.NewRequest("POST", fmt.Sprintf("/addGiftRequest"), bytes.NewBuffer(requestJSON))
	router.ServeHTTP(w, req1)
	assert.Equal(t, 200, w.Code)

	var addedRequest model.GiftRequest
	if e := json.Unmarshal(w.Body.Bytes(), &addedRequest); e != nil {
		t.Fatalf("Error unmarshaling JSON: %v", e)
	}
	var retrievedRequest model.GiftRequest
	err = tx.First(&retrievedRequest).Error
	assert.Equal(t, addedRequest.RecipientName, retrievedRequest.RecipientName)
	assert.Equal(t, addedRequest.Occasion, retrievedRequest.Occasion)
	assert.Equal(t, addedRequest.RecipientInterests, retrievedRequest.RecipientInterests)

}

func TestAddResponse(t *testing.T) {
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
	err = db.AutoMigrate(&model.GiftRequest{}, &model.Customer{}, &model.User{}, &model.GiftResponse{}, &model.GiftCollection{})
	if err != nil {
		panic("failed to migrate test database schema")
	}
	// Wrap the DB connection in a transaction
	tx := db.Begin()
	//defer tx.Rollback()

	// Create Model and Controller
	m := &model.PgModel{Conn: tx}
	c := &c.PgController{Model: m}
	router := c.Serve()

	// Test code
	w := httptest.NewRecorder()

	// Create GiftResponse
	assert.NoError(t, err)
	// Create GiftRequest
	user := model.User{}
	err = tx.Create(&user).Error
	assert.NoError(t, err)
	var retrievedUser model.User
	err = tx.First(&retrievedUser).Error
	assert.NoError(t, err)
	customer := model.Customer{
		User: retrievedUser,
	}
	err = tx.Create(&customer).Error
	assert.NoError(t, err)
	var retrievedCustomer model.Customer
	err = tx.First(&retrievedCustomer).Error
	request := model.GiftRequest{
		CustomerID:    retrievedCustomer.ID,
		RecipientName: "Friend",
	}
	err = tx.Create(&request).Error
	assert.NoError(t, err)
	collection := model.GiftCollection{
		CollectionName: "collection",
	}
	err = tx.Create(&collection).Error
	assert.NoError(t, err)
	response := model.GiftResponse{
		GiftCollection: collection,
		CustomMessage:  "Message",
	}
	responseJSON, err := json.Marshal(response)
	if err != nil {
		t.Fatalf("Error marshaling JSON: %v", err)
	}
	assert.NoError(t, err)
	req1, err := http.NewRequest("POST", fmt.Sprintf("/addGiftResponse"), bytes.NewBuffer(responseJSON))
	router.ServeHTTP(w, req1)
	assert.Equal(t, 200, w.Code)

	var addedResponse model.GiftResponse
	if e := json.Unmarshal(w.Body.Bytes(), &addedResponse); e != nil {
		t.Fatalf("Error unmarshaling JSON: %v", e)
	}
	var retrievedResponse model.GiftResponse
	err = tx.Preload("GiftCollection").First(&retrievedResponse, "id = ?", addedResponse.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, retrievedResponse.GiftCollectionID, addedResponse.GiftCollectionID)
	assert.Equal(t, retrievedResponse.GiftCollection.CollectionName, addedResponse.GiftCollection.CollectionName)
	assert.Equal(t, retrievedResponse.CustomMessage, addedResponse.CustomMessage)

}

func TestAddCollection(t *testing.T) {
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
	err = db.AutoMigrate(&model.GiftCollection{}, &model.Gift{})
	if err != nil {
		panic("failed to migrate test database schema")
	}
	// Wrap the DB connection in a transaction
	tx := db.Begin()
	//defer tx.Rollback()

	// Create Model and Controller
	m := &model.PgModel{Conn: tx}
	c := &c.PgController{Model: m}
	router := c.Serve()

	// Test code
	w := httptest.NewRecorder()

	gift := model.Gift{
		Name: "Gift1",
	}
	collection := model.GiftCollection{
		Gifts:          []*model.Gift{&gift},
		CollectionName: "collection",
	}

	collectionJSON, err := json.Marshal(collection)
	if err != nil {
		t.Fatalf("Error marshaling JSON: %v", err)
	}
	assert.NoError(t, err)
	req1, err := http.NewRequest("POST", fmt.Sprintf("/addGiftCollection"), bytes.NewBuffer(collectionJSON))
	router.ServeHTTP(w, req1)
	assert.Equal(t, 200, w.Code)

	var addedCollection model.GiftCollection
	if e := json.Unmarshal(w.Body.Bytes(), &addedCollection); e != nil {
		t.Fatalf("Error unmarshaling JSON: %v", e)
	}
	var retrievedCollection model.GiftCollection
	err = tx.Preload("Gifts").First(&retrievedCollection, "id = ?", addedCollection.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, retrievedCollection.CollectionName, addedCollection.CollectionName)
	assert.Equal(t, retrievedCollection.Gifts[0].Name, addedCollection.Gifts[0].Name)
}

func TestGetGift(t *testing.T) {
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
	err = db.AutoMigrate(&model.Gift{})
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

	// Create Gift
	testGift := model.Gift{
		Name:            "gift1",
		Price:           50,
		Link:            "link1",
		Description:     "description1",
		Demographic:     "demogrpahic1",
		GiftCollections: nil,
	}
	err = tx.Create(&testGift).Error
	assert.NoError(t, err)

	// Get Gift from database
	req1, err := http.NewRequest("GET", fmt.Sprintf("/gifts/%d", testGift.ID), nil)
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}
	router.ServeHTTP(w, req1)
	assert.Equal(t, 200, w.Code)

	var retrievedGift model.Gift
	if e := json.Unmarshal(w.Body.Bytes(), &retrievedGift); e != nil {
		t.Fatalf("Error unmarshaling JSON: %v", e)
	}

	// Test Retrieved Gift Fields
	var fetchedGift model.Gift
	err = tx.First(&fetchedGift, retrievedGift.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, retrievedGift.ID, fetchedGift.ID)
	assert.Equal(t, retrievedGift.Name, fetchedGift.Name)
	assert.Equal(t, retrievedGift.Price, fetchedGift.Price)
	assert.Equal(t, retrievedGift.Link, fetchedGift.Link)
	assert.Equal(t, retrievedGift.Description, fetchedGift.Description)
	assert.Equal(t, retrievedGift.Demographic, fetchedGift.Demographic)
	assert.Equal(t, retrievedGift.Link, fetchedGift.Link)
	assert.Equal(t, retrievedGift.CreatedAt.In(time.UTC).Round(time.Millisecond),
		fetchedGift.CreatedAt.In(time.UTC).Round(time.Millisecond))
}

func TestAddGift(t *testing.T) {
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
	err = db.AutoMigrate(&model.Gift{})
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

	// Create Gift
	testGift := model.Gift{
		Name:            "gift1",
		Price:           50,
		Link:            "link1",
		Description:     "description1",
		Demographic:     "demogrpahic1",
		GiftCollections: nil,
	}

	// Test Adding Gift to Database
	giftJson, err := json.Marshal(testGift)
	if err != nil {
		t.Fatalf("Error marshaling JSON: %v", err)
	}
	assert.NoError(t, err)

	req1, err := http.NewRequest("POST", fmt.Sprintf("/addGift"), bytes.NewBuffer(giftJson))
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}
	router.ServeHTTP(w, req1)
	assert.Equal(t, 200, w.Code)

	// Test Added Gift Fields
	var insertedGift model.Gift
	if e := json.Unmarshal(w.Body.Bytes(), &insertedGift); e != nil {
		t.Fatalf("Error unmarshaling JSON: %v", e)
	}

	var fetchedGift model.Gift
	err = tx.First(&fetchedGift, insertedGift.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, insertedGift.ID, fetchedGift.ID)
	assert.Equal(t, insertedGift.Name, fetchedGift.Name)
	assert.Equal(t, insertedGift.Price, fetchedGift.Price)
	assert.Equal(t, insertedGift.Link, fetchedGift.Link)
	assert.Equal(t, insertedGift.Description, fetchedGift.Description)
	assert.Equal(t, insertedGift.Demographic, fetchedGift.Demographic)
	assert.Equal(t, insertedGift.Link, fetchedGift.Link)
	assert.Equal(t, insertedGift.CreatedAt.In(time.UTC).Round(time.Millisecond),
		fetchedGift.CreatedAt.In(time.UTC).Round(time.Millisecond))

	//  Check that there's only 1 Gift
	var count int64
	tx.Model(&model.Gift{}).Where("id = ?", insertedGift.ID).Count(&count)
	assert.Equal(t, int64(1), count)
}

func TestUpdateGift(t *testing.T) {
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
	err = db.AutoMigrate(&model.Gift{})
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

	// Create Gift
	testGift := model.Gift{
		Name:            "gift1",
		Price:           50,
		Link:            "link1",
		Description:     "description1",
		Demographic:     "demogrpahic1",
		GiftCollections: nil,
	}
	err = tx.Create(&testGift).Error
	assert.NoError(t, err)

	// Test Inputted Gift Fields
	var fetchedGift model.Gift
	err = tx.First(&fetchedGift, testGift.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, testGift.ID, fetchedGift.ID)
	assert.Equal(t, testGift.Name, fetchedGift.Name)
	assert.Equal(t, testGift.Price, fetchedGift.Price)
	assert.Equal(t, testGift.Link, fetchedGift.Link)
	assert.Equal(t, testGift.Description, fetchedGift.Description)
	assert.Equal(t, testGift.Demographic, fetchedGift.Demographic)
	assert.Equal(t, testGift.Link, fetchedGift.Link)
	assert.Equal(t, testGift.CreatedAt.In(time.UTC).Round(time.Millisecond),
		fetchedGift.CreatedAt.In(time.UTC).Round(time.Millisecond))

	// Updated Gift Fields
	updatedTestGift := model.Gift{
		Name:            "updatedgift1",
		Price:           100,
		Description:     "updateddescription1",
		Demographic:     "updateddemogrpahic1",
		GiftCollections: nil,
	}

	// Test Updating Gift Fields
	giftJson, err := json.Marshal(updatedTestGift)
	if err != nil {
		t.Fatalf("Error marshaling JSON: %v", err)
	}

	req1, err := http.NewRequest("PUT", fmt.Sprintf("/gifts/%d", testGift.ID), bytes.NewBuffer(giftJson))
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}
	router.ServeHTTP(w, req1)
	assert.Equal(t, 200, w.Code)

	var updatedGiftRetrieved model.Gift
	if e := json.Unmarshal(w.Body.Bytes(), &updatedGiftRetrieved); e != nil {
		t.Fatalf("Error unmarshaling JSON: %v", e)
	}
	fmt.Print(updatedGiftRetrieved.ID)

	var fetchedUpdatedGift model.Gift
	err = tx.First(&fetchedUpdatedGift, updatedGiftRetrieved.ID).Error
	assert.NoError(t, err)
	err = tx.First(&updatedGiftRetrieved, fetchedUpdatedGift.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, fetchedUpdatedGift.ID, updatedGiftRetrieved.ID)
	assert.Equal(t, fetchedUpdatedGift.Name, updatedGiftRetrieved.Name)
	assert.Equal(t, fetchedUpdatedGift.Price, updatedGiftRetrieved.Price)
	assert.Equal(t, "link1", updatedGiftRetrieved.Link)
	assert.Equal(t, fetchedUpdatedGift.Description, updatedGiftRetrieved.Description)
	assert.Equal(t, fetchedUpdatedGift.Demographic, updatedGiftRetrieved.Demographic)
	assert.Equal(t, fetchedUpdatedGift.Link, updatedGiftRetrieved.Link)
	assert.Equal(t, fetchedUpdatedGift.CreatedAt.In(time.UTC).Round(time.Millisecond),
		updatedGiftRetrieved.CreatedAt.In(time.UTC).Round(time.Millisecond))

	//  Check that there's only 1 Gift
	var count int64
	tx.Model(&model.Gift{}).Where("id = ?", fetchedUpdatedGift.ID).Count(&count)
	assert.Equal(t, int64(1), count)
}

func TestDeleteGift(t *testing.T) {
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
	err = db.AutoMigrate(&model.Gift{})
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

	// Create Gift
	testGift := model.Gift{
		Name:            "gift1",
		Price:           50,
		Link:            "link1",
		Description:     "description1",
		Demographic:     "demogrpahic1",
		GiftCollections: nil,
	}
	err = tx.Create(&testGift).Error
	assert.NoError(t, err)

	// Test Inputted Gift Fields
	var fetchedGift model.Gift
	err = tx.First(&fetchedGift, testGift.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, testGift.ID, fetchedGift.ID)
	assert.Equal(t, testGift.Name, fetchedGift.Name)
	assert.Equal(t, testGift.Price, fetchedGift.Price)
	assert.Equal(t, testGift.Link, fetchedGift.Link)
	assert.Equal(t, testGift.Description, fetchedGift.Description)
	assert.Equal(t, testGift.Demographic, fetchedGift.Demographic)
	assert.Equal(t, testGift.Link, fetchedGift.Link)
	assert.Equal(t, testGift.CreatedAt.In(time.UTC).Round(time.Millisecond),
		fetchedGift.CreatedAt.In(time.UTC).Round(time.Millisecond))

	//  Check that there's only 1 Gift
	var count int64
	tx.Model(&model.Gift{}).Where("id = ?", testGift.ID).Count(&count)
	assert.Equal(t, int64(1), count)

	req1, err := http.NewRequest("DELETE", fmt.Sprintf("/gifts/%d", testGift.ID), nil)
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}
	router.ServeHTTP(w, req1)
	assert.Equal(t, 204, w.Code)

	//  Check that Gift has been deleted
	var deletedCount int64
	tx.Model(&model.Gift{}).Where("id = ?", testGift.ID).Count(&deletedCount)
	assert.Equal(t, int64(0), deletedCount)
}

func TestGetAllGift(t *testing.T) {
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
	err = db.AutoMigrate(&model.Gift{}, model.GiftCollection{})
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
	gift := model.Gift{
		Name:            "nice sweater",
		Price:           50,
		Link:            "https://something",
		Description:     "sample description",
		Demographic:     "sample demographic",
		GiftCollections: []*model.GiftCollection{},
	}

	gift_two := model.Gift{
		Name:            "nice sweater 2",
		Price:           20,
		Link:            "https://something 2",
		Description:     "sample description 2",
		Demographic:     "sample demographic 2",
		GiftCollections: []*model.GiftCollection{},
	}
	err = tx.Create(&gift).Error
	assert.NoError(t, err)

	err = tx.Create(&gift_two).Error
	assert.NoError(t, err)
	req1, err := http.NewRequest("GET", fmt.Sprintf("/gifts"), nil)
	router.ServeHTTP(w, req1)

	assert.Equal(t, 200, w.Code)

	var giftRetrieved []model.Gift
	if e := json.Unmarshal(w.Body.Bytes(), &giftRetrieved); e != nil {
		t.Fatalf("Error unmarshaling JSON: %v", e)
	}
	t.Log(len(giftRetrieved))
	assert.Equal(t, gift.Name, giftRetrieved[0].Name)
	assert.Equal(t, gift.Price, giftRetrieved[0].Price)
	assert.Equal(t, gift.Link, giftRetrieved[0].Link)
	assert.Equal(t, gift.Description, giftRetrieved[0].Description)
	assert.Equal(t, gift.Demographic, giftRetrieved[0].Demographic)
	assert.Equal(t, gift.GiftCollections, giftRetrieved[0].GiftCollections)

	assert.Equal(t, gift_two.Name, giftRetrieved[1].Name)
	assert.Equal(t, gift_two.Price, giftRetrieved[1].Price)
	assert.Equal(t, gift_two.Link, giftRetrieved[1].Link)
	assert.Equal(t, gift_two.Description, giftRetrieved[1].Description)
	assert.Equal(t, gift_two.Demographic, giftRetrieved[1].Demographic)
	assert.Equal(t, gift_two.GiftCollections, giftRetrieved[1].GiftCollections)
}

func TestGetAllGiftCollection(t *testing.T) {
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
	err = db.AutoMigrate(&model.GiftCollection{}, &model.User{}, &model.Customer{}, &model.Gift{})
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

	// Create a Customer
	user := model.User{}
	err = tx.Create(&user).Error
	assert.NoError(t, err)
	var retrievedUser model.User
	err = tx.First(&retrievedUser).Error
	assert.NoError(t, err)
	customer := model.Customer{
		User: retrievedUser,
	}
	err = tx.Create(&customer).Error
	assert.NoError(t, err)
	var retrievedCustomer model.Customer
	err = tx.First(&retrievedCustomer).Error
	assert.NoError(t, err)

	// Second Customer
	user2 := model.User{}
	err = tx.Create(&user2).Error
	assert.NoError(t, err)
	var retrievedUser2 model.User
	err = tx.Where("id = ?", user2.ID).First(&retrievedUser2).Error
	assert.NoError(t, err)
	customer2 := model.Customer{
		User: retrievedUser2,
	}
	err = tx.Create(&customer2).Error
	assert.NoError(t, err)
	var retrievedCustomer2 model.Customer
	err = tx.Where("id = ?", customer2.ID).First(&retrievedCustomer2).Error
	assert.NoError(t, err)

	collection := model.GiftCollection{
		CustomerID:     &retrievedCustomer.ID,
		CollectionName: "sample name",
		Gifts:          []*model.Gift{},
	}

	collection_two := model.GiftCollection{
		CustomerID:     &retrievedCustomer2.ID,
		CollectionName: "sample name 2",
		Gifts:          []*model.Gift{},
	}

	err = tx.Create(&collection).Error
	assert.NoError(t, err)

	err = tx.Create(&collection_two).Error
	assert.NoError(t, err)

	req1, err := http.NewRequest("GET", fmt.Sprintf("/collections"), nil)
	router.ServeHTTP(w, req1)

	assert.Equal(t, 200, w.Code)

	var collectionRetrieved []model.GiftCollection
	if e := json.Unmarshal(w.Body.Bytes(), &collectionRetrieved); e != nil {
		t.Fatalf("Error unmarshaling JSON: %v", e)
	}

	assert.Equal(t, collection.CustomerID, collectionRetrieved[0].CustomerID)
	assert.Equal(t, collection.CollectionName, collectionRetrieved[0].CollectionName)
	assert.Equal(t, collection.Gifts, collectionRetrieved[0].Gifts)

	assert.Equal(t, collection_two.CustomerID, collectionRetrieved[1].CustomerID)
	assert.Equal(t, collection_two.CollectionName, collectionRetrieved[1].CollectionName)
	assert.Equal(t, collection_two.Gifts, collectionRetrieved[1].Gifts)
}

func TestGetAllGiftResponse(t *testing.T) {
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
	err = db.AutoMigrate(&model.GiftResponse{}, model.GiftCollection{})
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

	collection := model.GiftCollection{}
	err = tx.Create(&collection).Error

	response := model.GiftResponse{
		GiftCollection: collection,
		CustomMessage:  "sample custom message",
	}

	response_two := model.GiftResponse{
		GiftCollection: collection,
		CustomMessage:  "sample custom message 2",
	}

	err = tx.Create(&response).Error
	assert.NoError(t, err)

	err = tx.Create(&response_two).Error
	assert.NoError(t, err)
	req1, err := http.NewRequest("GET", fmt.Sprintf("/responses"), nil)
	router.ServeHTTP(w, req1)

	assert.Equal(t, 200, w.Code)

	var responseRetrieved []model.GiftResponse
	if e := json.Unmarshal(w.Body.Bytes(), &responseRetrieved); e != nil {
		t.Fatalf("Error unmarshaling JSON: %v", e)
	}

	assert.Equal(t, response.GiftCollection.ID, responseRetrieved[0].GiftCollection.ID)
	assert.Equal(t, response.GiftCollectionID, responseRetrieved[0].GiftCollectionID)
	assert.Equal(t, response.CustomMessage, responseRetrieved[0].CustomMessage)

	assert.Equal(t, response_two.GiftCollection.ID, responseRetrieved[1].GiftCollection.ID)
	assert.Equal(t, response_two.GiftCollectionID, responseRetrieved[1].GiftCollectionID)
	assert.Equal(t, response_two.CustomMessage, responseRetrieved[1].CustomMessage)

}

func TestAddGiftToCollection(t *testing.T) {
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
	err = db.AutoMigrate(&model.GiftCollection{}, &model.Gift{})
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
	w1 := httptest.NewRecorder()
	w2 := httptest.NewRecorder()

	// Create a new Gift Collection
	collection := model.GiftCollection{
		Gifts:          []*model.Gift{},
		CollectionName: "collection",
	}

	collectionJSON, err := json.Marshal(collection)
	if err != nil {
		t.Fatalf("Error marshaling JSON: %v", err)
	}
	assert.NoError(t, err)
	req1, err := http.NewRequest("POST", fmt.Sprintf("/addGiftCollection"), bytes.NewBuffer(collectionJSON))
	router.ServeHTTP(w1, req1)
	assert.Equal(t, 200, w1.Code)

	var addedCollection model.GiftCollection
	if e := json.Unmarshal(w1.Body.Bytes(), &addedCollection); e != nil {
		t.Fatalf("Error unmarshaling JSON: %v", e)
	}
	var retrievedCollection model.GiftCollection
	err = tx.Preload("Gifts").First(&retrievedCollection, "id = ?", addedCollection.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, retrievedCollection.CollectionName, addedCollection.CollectionName)

	// Add Gift to Gift Collection
	gift := model.Gift{
		Name:     "Gift1",
		Category: pq.StringArray{"Best selling", "Gadgets"},
	}
	giftJSON, err := json.Marshal(gift)
	if err != nil {
		t.Fatalf("Error marshaling JSON: %v", err)
	}
	assert.NoError(t, err)
	req2, err := http.NewRequest("POST", fmt.Sprintf("/addGiftCollection/%d", addedCollection.ID), bytes.NewBuffer(giftJSON))
	router.ServeHTTP(w2, req2)
	assert.Equal(t, 200, w2.Code)

	var giftAddedCollection model.GiftCollection
	if e := json.Unmarshal(w2.Body.Bytes(), &giftAddedCollection); e != nil {
		t.Fatalf("Error unmarshaling JSON: %v", e)
	}
	var giftAddedRetrievedCollection model.GiftCollection
	err = tx.Preload("Gifts").First(&giftAddedRetrievedCollection, "id = ?", giftAddedCollection.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, giftAddedRetrievedCollection.CollectionName, giftAddedCollection.CollectionName)
	assert.Equal(t, giftAddedRetrievedCollection.Gifts[0].Name, giftAddedCollection.Gifts[0].Name)
	assert.Equal(t, giftAddedRetrievedCollection.Gifts[0].Category, giftAddedCollection.Gifts[0].Category)
}

func TestGiftDeleteFromCollection(t *testing.T) {
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
	err = db.AutoMigrate(&model.GiftCollection{}, &model.Gift{})
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
	w1 := httptest.NewRecorder()
	w2 := httptest.NewRecorder()

	gift := model.Gift{
		Name: "Gift1",
	}
	err = tx.Create(&gift).Error
	assert.NoError(t, err)

	collection := model.GiftCollection{
		Gifts:          []*model.Gift{&gift},
		CollectionName: "collection",
	}

	collectionJSON, err := json.Marshal(collection)
	if err != nil {
		t.Fatalf("Error marshaling JSON: %v", err)
	}

	assert.NoError(t, err)
	req1, err := http.NewRequest("POST", fmt.Sprintf("/addGiftCollection"), bytes.NewBuffer(collectionJSON))
	router.ServeHTTP(w1, req1)
	assert.Equal(t, 200, w1.Code)

	var addedCollection model.GiftCollection
	if e := json.Unmarshal(w1.Body.Bytes(), &addedCollection); e != nil {
		t.Fatalf("Error unmarshaling JSON: %v", e)
	}

	var retrievedCollection model.GiftCollection
	err = tx.Preload("Gifts").First(&retrievedCollection, "id = ?", addedCollection.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, retrievedCollection.Gifts[0].ID, addedCollection.Gifts[0].ID)
	assert.Equal(t, retrievedCollection.CollectionName, addedCollection.CollectionName)
	assert.Equal(t, retrievedCollection.Gifts[0].Name, addedCollection.Gifts[0].Name)

	var count1 int64
	count1 = int64(len(retrievedCollection.Gifts))
	assert.Equal(t, int64(1), count1)

	// Delete Gift from Gift Collection
	assert.NoError(t, err)
	fmt.Print(retrievedCollection.Gifts[0].ID)
	fmt.Print(retrievedCollection.ID)
	fmt.Print(addedCollection.ID)
	req2, err := http.NewRequest("DELETE", fmt.Sprintf("/removeGiftFromGiftCollection/%d/%d", retrievedCollection.Gifts[0].ID, addedCollection.ID), nil)
	router.ServeHTTP(w2, req2)
	assert.Equal(t, 200, w2.Code)

	var giftDeletedCollection model.GiftCollection
	if e := json.Unmarshal(w2.Body.Bytes(), &giftDeletedCollection); e != nil {
		t.Fatalf("Error unmarshaling JSON: %v", e)
	}
	var giftDeletedRetrievedCollection model.GiftCollection
	err = tx.Preload("Gifts").First(&giftDeletedRetrievedCollection, "id = ?", giftDeletedCollection.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, giftDeletedRetrievedCollection.CollectionName, giftDeletedCollection.CollectionName)

	var count2 int64
	count2 = int64(len(giftDeletedRetrievedCollection.Gifts))
	assert.Equal(t, int64(0), count2)
}

func TestGetAllCustomerGiftCollection(t *testing.T) {
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
	err = db.AutoMigrate(&model.GiftCollection{}, &model.Giftee{}, &model.User{}, &model.Customer{}, &model.GiftRequest{}, &model.GiftResponse{})
	if err != nil {
		panic("failed to migrate test database schema")
	}
	// Wrap the DB connection in a transaction,
	tx := db.Begin()
	defer tx.Rollback()

	// Create Model and Controller
	m := &model.PgModel{Conn: tx}
	c := &c.PgController{Model: m}
	router := c.Serve()

	// Test code
	w := httptest.NewRecorder()

	// Create a Customer
	user := model.User{}
	err = tx.Create(&user).Error
	assert.NoError(t, err)
	var retrievedUser model.User
	err = tx.First(&retrievedUser).Error
	assert.NoError(t, err)
	customer := model.Customer{
		User: retrievedUser,
	}
	err = tx.Create(&customer).Error
	assert.NoError(t, err)
	var retrievedCustomer model.Customer
	err = tx.First(&retrievedCustomer).Error
	assert.NoError(t, err)

	// Second Customer
	user2 := model.User{}
	err = tx.Create(&user2).Error
	assert.NoError(t, err)
	var retrievedUser2 model.User
	err = tx.Where("id = ?", user2.ID).First(&retrievedUser2).Error
	assert.NoError(t, err)
	customer2 := model.Customer{
		User: retrievedUser2,
	}
	err = tx.Create(&customer2).Error
	assert.NoError(t, err)
	var retrievedCustomer2 model.Customer
	err = tx.Where("id = ?", customer2.ID).First(&retrievedCustomer2).Error
	assert.NoError(t, err)

	collection := model.GiftCollection{
		CustomerID:     &retrievedCustomer.ID,
		CollectionName: "sample name",
		Gifts:          []*model.Gift{},
	}

	collection_two := model.GiftCollection{
		CustomerID:     &retrievedCustomer2.ID,
		CollectionName: "sample name 2",
		Gifts:          []*model.Gift{},
	}

	collection_three := model.GiftCollection{
		CollectionName: "sample name 3",
		Gifts:          []*model.Gift{},
	}

	err = tx.Create(&collection).Error
	assert.NoError(t, err)

	err = tx.Create(&collection_two).Error
	assert.NoError(t, err)

	err = tx.Create(&collection_three).Error
	assert.NoError(t, err)

	req1, err := http.NewRequest("GET", fmt.Sprintf("/collections/%d", retrievedCustomer2.ID), nil)
	router.ServeHTTP(w, req1)
	assert.Equal(t, 200, w.Code)

	var collectionRetrieved []model.GiftCollection
	if e := json.Unmarshal(w.Body.Bytes(), &collectionRetrieved); e != nil {
		t.Fatalf("Error unmarshaling JSON: %v", e)
	}

	assert.Equal(t, collection_two.CustomerID, collectionRetrieved[0].CustomerID)
	assert.Equal(t, collection_two.CollectionName, collectionRetrieved[0].CollectionName)
	assert.Equal(t, collection_two.Gifts, collectionRetrieved[0].Gifts)

	assert.Equal(t, collection_three.CustomerID, collectionRetrieved[1].CustomerID)
	assert.Equal(t, collection_three.CollectionName, collectionRetrieved[1].CollectionName)
	assert.Equal(t, collection_three.Gifts, collectionRetrieved[1].Gifts)
}

func TestUpdateCustomerAvailableRequests(t *testing.T) {
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
	err = db.AutoMigrate(&model.Customer{}, &model.User{})
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
	w1 := httptest.NewRecorder()
	w2 := httptest.NewRecorder()

	// Create a Customer
	user := model.User{}
	err = tx.Create(&user).Error
	assert.NoError(t, err)
	var retrievedUser model.User
	err = tx.First(&retrievedUser).Error
	assert.NoError(t, err)
	customer := model.Customer{
		User:              retrievedUser,
		AvailableRequests: uint(10),
	}
	err = tx.Create(&customer).Error
	assert.NoError(t, err)
	var retrievedCustomer model.Customer
	err = tx.First(&retrievedCustomer).Error
	assert.NoError(t, err)

	req1, err := http.NewRequest("PUT", fmt.Sprintf("/customer/%d?requests=5", retrievedCustomer.ID), nil)
	router.ServeHTTP(w1, req1)
	assert.Equal(t, 200, w1.Code)

	var customerRetrieved model.Customer
	if e := json.Unmarshal(w1.Body.Bytes(), &customerRetrieved); e != nil {
		t.Fatalf("Error unmarshaling JSON: %v", e)
	}

	assert.Equal(t, retrievedCustomer.ID, customerRetrieved.ID)
	assert.Equal(t, retrievedCustomer.UserID, customerRetrieved.UserID)
	assert.Equal(t, retrievedCustomer.User, customerRetrieved.User)
	assert.Equal(t, uint(15), customerRetrieved.AvailableRequests)

	req2, err := http.NewRequest("PUT", fmt.Sprintf("/customer/%d?requests=-1", retrievedCustomer.ID), nil)
	router.ServeHTTP(w2, req2)
	assert.Equal(t, 200, w2.Code)

	var customerRetrieved2 model.Customer
	if e := json.Unmarshal(w2.Body.Bytes(), &customerRetrieved2); e != nil {
		t.Fatalf("Error unmarshaling JSON: %v", e)
	}

	assert.Equal(t, retrievedCustomer.ID, customerRetrieved2.ID)
	assert.Equal(t, retrievedCustomer.UserID, customerRetrieved2.UserID)
	assert.Equal(t, retrievedCustomer.User, customerRetrieved2.User)
	assert.Equal(t, uint(14), customerRetrieved2.AvailableRequests)
}

func TestAddGiftToCustomerGiftCollection(t *testing.T) {
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
	err = db.AutoMigrate(&model.GiftCollection{})
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

	// Create a Customer
	user := model.User{}
	err = tx.Create(&user).Error
	assert.NoError(t, err)
	var retrievedUser model.User
	err = tx.First(&retrievedUser).Error
	assert.NoError(t, err)
	customer := model.Customer{
		User: retrievedUser,
	}
	err = tx.Create(&customer).Error
	assert.NoError(t, err)
	var retrievedCustomer model.Customer
	err = tx.First(&retrievedCustomer).Error
	assert.NoError(t, err)

	// Create a collection
	collection := model.GiftCollection{
		CustomerID:     &retrievedCustomer.ID,
		CollectionName: "test name",
		Gifts:          []*model.Gift{},
	}
	err = tx.Create(&collection).Error
	assert.NoError(t, err)
	var retrievedCollection model.GiftCollection
	err = tx.First(&retrievedCollection).Error
	assert.NoError(t, err)

	// Create a gift
	gift := model.Gift{
		Name:  "test gift",
		Price: 25,
	}
	giftJSON, err := json.Marshal(gift)
	if err != nil {
		t.Fatalf("Error marshaling JSON: %v", err)
	}
	assert.NoError(t, err)

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("/addCustomerGiftCollection/%s/%d", retrievedCollection.CollectionName, retrievedCustomer.ID),
		bytes.NewBuffer(giftJSON),
	)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	var collectionResponse model.GiftCollection
	if e := json.Unmarshal(w.Body.Bytes(), &collectionResponse); e != nil {
		t.Fatalf("Error unmarshaling JSON: %v", e)
	}

	assert.Equal(t, &retrievedCustomer.ID, collectionResponse.CustomerID)
	assert.Equal(t, retrievedCollection.CollectionName, collectionResponse.CollectionName)
	assert.Equal(t, gift.Name, collectionResponse.Gifts[0].Name)
	assert.Equal(t, gift.Price, collectionResponse.Gifts[0].Price)
}
func TestSearchGift(t *testing.T) {
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
	err = db.AutoMigrate(&model.Gift{}, &model.GiftCollection{})
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
	w1 := httptest.NewRecorder()
	w2 := httptest.NewRecorder()
	w3 := httptest.NewRecorder()
	w4 := httptest.NewRecorder()
	w5 := httptest.NewRecorder()
	w6 := httptest.NewRecorder()
	w7 := httptest.NewRecorder()
	w8 := httptest.NewRecorder()
	w9 := httptest.NewRecorder()

	collection := model.GiftCollection{
		CollectionName: "collection",
	}
	err = tx.Create(&collection).Error
	assert.NoError(t, err)

	// Create Gifts
	testGift1 := model.Gift{
		Name:            "gift1",
		Price:           50,
		Link:            "link1",
		Description:     "description1",
		Demographic:     "demogrpahic1",
		Category:        pq.StringArray{"category1"},
		Occasion:        "occasion1",
		GiftCollections: []*model.GiftCollection{&collection},
	}
	err = tx.Create(&testGift1).Error
	assert.NoError(t, err)

	testGift2 := model.Gift{
		Name:            "gift2",
		Price:           20,
		Link:            "link1",
		Description:     "description2",
		Demographic:     "demogrpahic1",
		Category:        pq.StringArray{"category2"},
		Occasion:        "occasion1",
		GiftCollections: []*model.GiftCollection{&collection},
	}
	err = tx.Create(&testGift2).Error
	assert.NoError(t, err)

	testGift3 := model.Gift{
		Name:            "gift3",
		Price:           70,
		Link:            "link2",
		Description:     "description1",
		Demographic:     "demogrpahic2",
		Category:        pq.StringArray{"category1"},
		Occasion:        "occasion2",
		GiftCollections: []*model.GiftCollection{&collection},
	}
	err = tx.Create(&testGift3).Error
	assert.NoError(t, err)

	// Search for gift by price
	req1, err := http.NewRequest("GET", fmt.Sprintf("/search/%d?minPrice=0&maxPrice=100", collection.ID), nil)
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}
	router.ServeHTTP(w1, req1)
	assert.Equal(t, 200, w1.Code)

	var retrievedPriceGifts []model.Gift
	if err := json.Unmarshal(w1.Body.Bytes(), &retrievedPriceGifts); err != nil {
		t.Fatalf("Error unmarshaling JSON: %v", err)
	}

	assert.GreaterOrEqual(t, len(retrievedPriceGifts), 3)

	req2, err := http.NewRequest("GET", fmt.Sprintf("/search/%d?minPrice=60&maxPrice=100", collection.ID), nil)
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}
	router.ServeHTTP(w2, req2)
	assert.Equal(t, 200, w2.Code)

	var searchOneGift []model.Gift
	if err := json.Unmarshal(w2.Body.Bytes(), &searchOneGift); err != nil {
		t.Fatalf("Error unmarshaling JSON: %v", err)
	}

	assert.GreaterOrEqual(t, len(searchOneGift), 1)

	// Search Gift By Demographic
	req3, err := http.NewRequest("GET", fmt.Sprintf("/search/%d?demographic=demogrpahic1", collection.ID), nil)
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}
	router.ServeHTTP(w3, req3)
	assert.Equal(t, 200, w3.Code)

	var retrievedDemographicGifts []model.Gift
	if err := json.Unmarshal(w3.Body.Bytes(), &retrievedDemographicGifts); err != nil {
		t.Fatalf("Error unmarshaling JSON: %v", err)
	}

	assert.GreaterOrEqual(t, len(retrievedDemographicGifts), 2)

	req4, err := http.NewRequest("GET", fmt.Sprintf("/search/%d?demographic=demogrpahic2", collection.ID), nil)
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}
	router.ServeHTTP(w4, req4)
	assert.Equal(t, 200, w4.Code)

	var retrievedOneDemographicGift []model.Gift
	if err := json.Unmarshal(w4.Body.Bytes(), &retrievedOneDemographicGift); err != nil {
		t.Fatalf("Error unmarshaling JSON: %v", err)
	}

	assert.GreaterOrEqual(t, len(retrievedOneDemographicGift), 1)

	// Search Gift By Occasion
	req5, err := http.NewRequest("GET", fmt.Sprintf("/search/%d?occasion=occasion1", collection.ID), nil)
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}
	router.ServeHTTP(w5, req5)
	assert.Equal(t, 200, w5.Code)

	var retrievedOccasionGifts []model.Gift
	if err := json.Unmarshal(w5.Body.Bytes(), &retrievedOccasionGifts); err != nil {
		t.Fatalf("Error unmarshaling JSON: %v", err)
	}

	assert.GreaterOrEqual(t, len(retrievedOccasionGifts), 2)

	req6, err := http.NewRequest("GET", fmt.Sprintf("/search/%d?occasion=occasion2", collection.ID), nil)
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}
	router.ServeHTTP(w6, req6)
	assert.Equal(t, 200, w6.Code)

	var retrievedOneOccasionGift []model.Gift
	if err := json.Unmarshal(w6.Body.Bytes(), &retrievedOneOccasionGift); err != nil {
		t.Fatalf("Error unmarshaling JSON: %v", err)
	}

	assert.GreaterOrEqual(t, len(retrievedOneOccasionGift), 1)

	// Search Gift By Category
	req7, err := http.NewRequest("GET", fmt.Sprintf("/search/%d?category=category1", collection.ID), nil)
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}
	router.ServeHTTP(w7, req7)
	assert.Equal(t, 200, w7.Code)

	var retrievedCategoryGifts []model.Gift
	if err := json.Unmarshal(w7.Body.Bytes(), &retrievedCategoryGifts); err != nil {
		t.Fatalf("Error unmarshaling JSON: %v", err)
	}

	assert.GreaterOrEqual(t, len(retrievedCategoryGifts), 2)

	req8, err := http.NewRequest("GET", fmt.Sprintf("/search/%d?category=category2", collection.ID), nil)
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}
	router.ServeHTTP(w8, req8)
	assert.Equal(t, 200, w8.Code)

	var retrievedOneCategoryGift []model.Gift
	if err := json.Unmarshal(w8.Body.Bytes(), &retrievedOneCategoryGift); err != nil {
		t.Fatalf("Error unmarshaling JSON: %v", err)
	}

	assert.GreaterOrEqual(t, len(retrievedOneCategoryGift), 1)

	// Test Empty
	req9, err := http.NewRequest("GET", fmt.Sprintf("/search/%d", collection.ID), nil)
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}
	router.ServeHTTP(w9, req9)
	assert.Equal(t, 200, w9.Code)

	var retrievedAllGifts []model.Gift
	if err := json.Unmarshal(w9.Body.Bytes(), &retrievedAllGifts); err != nil {
		t.Fatalf("Error unmarshaling JSON: %v", err)
	}

	assert.GreaterOrEqual(t, len(retrievedAllGifts), 3)
}

func TestDeleteGiftFromCustomerGiftCollection(t *testing.T) {
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
	err = db.AutoMigrate(&model.GiftCollection{}, &model.Gift{}, &model.Customer{}, &model.User{})
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

	// Create a Customer
	user := model.User{}
	err = tx.Create(&user).Error
	assert.NoError(t, err)
	var retrievedUser model.User
	err = tx.First(&retrievedUser).Error
	assert.NoError(t, err)
	customer := model.Customer{
		User: retrievedUser,
	}
	err = tx.Create(&customer).Error
	assert.NoError(t, err)
	var retrievedCustomer model.Customer
	err = tx.First(&retrievedCustomer).Error
	assert.NoError(t, err)

	// Create gifts
	giftToRemove := model.Gift{
		Name:  "gift to remove",
		Price: 25,
	}
	giftToStay := model.Gift{
		Name:  "gift to stay",
		Price: 25,
	}
	giftJSON, err := json.Marshal(giftToRemove)
	if err != nil {
		t.Fatalf("Error marshaling JSON: %v", err)
	}
	assert.NoError(t, err)

	// Create a collection
	collection := model.GiftCollection{
		CustomerID:     &retrievedCustomer.ID,
		CollectionName: "test name",
		Gifts:          []*model.Gift{&giftToRemove, &giftToStay},
	}

	err = tx.Create(&collection).Error
	assert.NoError(t, err)
	var retrievedCollection model.GiftCollection
	err = tx.Preload("Gifts").First(&retrievedCollection).Error
	assert.NoError(t, err)

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("/removeCustomerGiftCollection/%s/%d", retrievedCollection.CollectionName, retrievedCustomer.ID),
		bytes.NewBuffer(giftJSON),
	)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	var collectionResponse model.GiftCollection
	if e := json.Unmarshal(w.Body.Bytes(), &collectionResponse); e != nil {
		t.Fatalf("Error unmarshaling JSON: %v", e)
	}

	assert.Equal(t, &retrievedCustomer.ID, collectionResponse.CustomerID)
	assert.Equal(t, collection.CollectionName, collectionResponse.CollectionName)
	assert.Equal(t, 1, len(collectionResponse.Gifts))
	assert.Equal(t, giftToStay.Name, collectionResponse.Gifts[0].Name)
	assert.Equal(t, giftToStay.Price, collectionResponse.Gifts[0].Price)
}

func TestGetGiftee(t *testing.T) {
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
	err = db.AutoMigrate(&model.Giftee{}, &model.User{}, &model.Customer{})
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

	// Create User
	user := model.User{Email: "example@northeastern.edu", FirstName: "PersonFirstName", LastName: "PersonLastName", Password: "dgeeg32"}
	err = tx.Create(&user).Error

	// Create Customer
	customer := model.Customer{User: user, UserID: user.ID}
	err = tx.Create(&customer).Error
	assert.NoError(t, err)

	// Create Giftee
	testGiftee := model.Giftee{
		CustomerID:           customer.ID,
		GifteeName:           "Maya",
		Gender:               "Female",
		CustomerRelationship: "Sister",
		Age:                  20,
		Colors:               pq.StringArray{"Green", "Blue"},
		Interests:            pq.StringArray{"Sports", "Soccer", "Nature", "Coffee", "Candy"},
	}
	err = tx.Create(&testGiftee).Error
	assert.NoError(t, err)
	// Get Giftee from database
	req1, err := http.NewRequest("GET", fmt.Sprintf("/giftee/%d", testGiftee.ID), nil)
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}
	router.ServeHTTP(w, req1)
	assert.Equal(t, 200, w.Code)

	var retrievedGiftee model.Giftee
	if e := json.Unmarshal(w.Body.Bytes(), &retrievedGiftee); e != nil {
		t.Fatalf("Error unmarshaling JSON: %v", e)
	}

	// Test Retrieved Giftee Fields
	var fetchedGiftee model.Giftee
	err = tx.First(&fetchedGiftee, retrievedGiftee.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, retrievedGiftee.ID, fetchedGiftee.ID)
	assert.Equal(t, retrievedGiftee.CustomerID, fetchedGiftee.CustomerID)
	assert.Equal(t, retrievedGiftee.GifteeName, fetchedGiftee.GifteeName)
	assert.Equal(t, retrievedGiftee.Gender, fetchedGiftee.Gender)
	assert.Equal(t, retrievedGiftee.CustomerRelationship, fetchedGiftee.CustomerRelationship)
	assert.Equal(t, retrievedGiftee.Age, fetchedGiftee.Age)
	assert.Equal(t, retrievedGiftee.Colors, fetchedGiftee.Colors)
	assert.Equal(t, retrievedGiftee.Interests, fetchedGiftee.Interests)
	assert.Equal(t, retrievedGiftee.GiftRequests, fetchedGiftee.GiftRequests)
	assert.Equal(t, retrievedGiftee.CreatedAt.In(time.UTC).Round(time.Millisecond),
		fetchedGiftee.CreatedAt.In(time.UTC).Round(time.Millisecond))
}

func TestAddGiftee(t *testing.T) {
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
	err = db.AutoMigrate(&model.Giftee{}, &model.User{}, &model.Customer{})
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

	// Create User
	user := model.User{Email: "example@northeastern.edu", FirstName: "PersonFirstName", LastName: "PersonLastName", Password: "dgeeg32"}
	err = tx.Create(&user).Error

	// Create Customer
	customer := model.Customer{User: user, UserID: user.ID}
	err = tx.Create(&customer).Error
	assert.NoError(t, err)

	// Create Giftee
	testGiftee := model.Giftee{
		CustomerID:           customer.ID,
		GifteeName:           "Maya",
		Gender:               "Female",
		CustomerRelationship: "Sister",
		Age:                  20,
		Colors:               pq.StringArray{"Green", "Blue"},
		Interests:            pq.StringArray{"Sports", "Soccer", "Nature", "Coffee", "Candy"},
	}

	// Test Adding Giftee to Database
	gifteeJson, err := json.Marshal(testGiftee)
	if err != nil {
		t.Fatalf("Error marshaling JSON: %v", err)
	}
	assert.NoError(t, err)

	req1, err := http.NewRequest("POST", "/addGiftee", bytes.NewBuffer(gifteeJson))
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}
	router.ServeHTTP(w, req1)
	assert.Equal(t, 200, w.Code)

	// Test Added Giftee Fields
	var insertedGiftee model.Giftee
	if e := json.Unmarshal(w.Body.Bytes(), &insertedGiftee); e != nil {
		t.Fatalf("Error unmarshaling JSON: %v", e)
	}

	var fetchedGiftee model.Giftee
	err = tx.First(&fetchedGiftee, insertedGiftee.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, insertedGiftee.ID, fetchedGiftee.ID)
	assert.Equal(t, insertedGiftee.CustomerID, fetchedGiftee.CustomerID)
	assert.Equal(t, insertedGiftee.GifteeName, fetchedGiftee.GifteeName)
	assert.Equal(t, insertedGiftee.Gender, fetchedGiftee.Gender)
	assert.Equal(t, insertedGiftee.CustomerRelationship, fetchedGiftee.CustomerRelationship)
	assert.Equal(t, insertedGiftee.Age, fetchedGiftee.Age)
	assert.Equal(t, insertedGiftee.Colors, fetchedGiftee.Colors)
	assert.Equal(t, insertedGiftee.Interests, fetchedGiftee.Interests)
	assert.Equal(t, insertedGiftee.CreatedAt.In(time.UTC).Round(time.Millisecond),
		fetchedGiftee.CreatedAt.In(time.UTC).Round(time.Millisecond))

	//  Check that there's only 1 Giftee
	var count int64
	tx.Model(&model.Giftee{}).Where("id = ?", insertedGiftee.ID).Count(&count)
	assert.Equal(t, int64(1), count)
}

func TestUpdateGiftee(t *testing.T) {
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
	err = db.AutoMigrate(&model.Giftee{}, &model.User{}, &model.Customer{})
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

	// Create User
	user := model.User{Email: "example@northeastern.edu", FirstName: "PersonFirstName", LastName: "PersonLastName", Password: "dgeeg32"}
	err = tx.Create(&user).Error

	// Create Customer
	customer := model.Customer{User: user, UserID: user.ID}
	err = tx.Create(&customer).Error
	assert.NoError(t, err)

	// Create Giftee
	testGiftee := model.Giftee{
		CustomerID:           customer.ID,
		GifteeName:           "Maya",
		Gender:               "Female",
		CustomerRelationship: "Sister",
		Age:                  20,
		Colors:               pq.StringArray{"Green", "Blue"},
		Interests:            pq.StringArray{"Sports", "Soccer", "Nature", "Coffee", "Candy"},
	}
	err = tx.Create(&testGiftee).Error
	assert.NoError(t, err)

	// Test Inputted Giftee Fields
	var fetchedGiftee model.Giftee
	err = tx.First(&fetchedGiftee, testGiftee.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, testGiftee.ID, fetchedGiftee.ID)
	assert.Equal(t, testGiftee.CustomerID, fetchedGiftee.CustomerID)
	assert.Equal(t, testGiftee.GifteeName, fetchedGiftee.GifteeName)
	assert.Equal(t, testGiftee.Gender, fetchedGiftee.Gender)
	assert.Equal(t, testGiftee.CustomerRelationship, fetchedGiftee.CustomerRelationship)
	assert.Equal(t, testGiftee.Age, fetchedGiftee.Age)
	assert.Equal(t, testGiftee.Colors, fetchedGiftee.Colors)
	assert.Equal(t, testGiftee.Interests, fetchedGiftee.Interests)
	assert.Equal(t, testGiftee.GiftRequests, fetchedGiftee.GiftRequests)
	assert.Equal(t, testGiftee.CreatedAt.In(time.UTC).Round(time.Millisecond),
		fetchedGiftee.CreatedAt.In(time.UTC).Round(time.Millisecond))

	// Updated Giftee Fields
	updatedTestGiftee := model.Giftee{
		GifteeName:           "Maya Updated",
		Gender:               "Female",
		CustomerRelationship: "Sister",
		Age:                  25,
		Colors:               pq.StringArray{"Green", "Blue", "Yellow", "Red"},
		Interests:            pq.StringArray{"Sports", "Soccer", "Candy"},
	}

	// Test Updating Giftee Fields
	gifteeJson, err := json.Marshal(updatedTestGiftee)
	if err != nil {
		t.Fatalf("Error marshaling JSON: %v", err)
	}

	req1, err := http.NewRequest("PUT", fmt.Sprintf("/giftee/%d", testGiftee.ID), bytes.NewBuffer(gifteeJson))
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}
	router.ServeHTTP(w, req1)
	assert.Equal(t, 200, w.Code)

	var updatedGifteeRetrieved model.Giftee
	if e := json.Unmarshal(w.Body.Bytes(), &updatedGifteeRetrieved); e != nil {
		t.Fatalf("Error unmarshaling JSON: %v", e)
	}

	var fetchedUpdatedGiftee model.Giftee
	err = tx.First(&fetchedUpdatedGiftee, updatedGifteeRetrieved.ID).Error
	assert.NoError(t, err)
	err = tx.First(&updatedGifteeRetrieved, fetchedUpdatedGiftee.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, fetchedUpdatedGiftee.ID, updatedGifteeRetrieved.ID)
	assert.Equal(t, fetchedUpdatedGiftee.CustomerID, updatedGifteeRetrieved.CustomerID)
	assert.Equal(t, fetchedUpdatedGiftee.GifteeName, updatedGifteeRetrieved.GifteeName)
	assert.Equal(t, fetchedUpdatedGiftee.Gender, updatedGifteeRetrieved.Gender)
	assert.Equal(t, fetchedUpdatedGiftee.CustomerRelationship, updatedGifteeRetrieved.CustomerRelationship)
	assert.Equal(t, fetchedUpdatedGiftee.Age, updatedGifteeRetrieved.Age)
	assert.Equal(t, fetchedUpdatedGiftee.Colors, updatedGifteeRetrieved.Colors)
	assert.Equal(t, fetchedUpdatedGiftee.Interests, updatedGifteeRetrieved.Interests)
	assert.Equal(t, fetchedUpdatedGiftee.GiftRequests, updatedGifteeRetrieved.GiftRequests)
	assert.Equal(t, fetchedUpdatedGiftee.CreatedAt.In(time.UTC).Round(time.Millisecond),
		updatedGifteeRetrieved.CreatedAt.In(time.UTC).Round(time.Millisecond))
}

func TestDeleteGiftee(t *testing.T) {
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
	err = db.AutoMigrate(&model.Giftee{}, &model.User{}, &model.Customer{})
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

	// Create User
	user := model.User{Email: "example@northeastern.edu", FirstName: "PersonFirstName", LastName: "PersonLastName", Password: "dgeeg32"}
	err = tx.Create(&user).Error

	// Create Customer
	customer := model.Customer{User: user, UserID: user.ID}
	err = tx.Create(&customer).Error
	assert.NoError(t, err)

	// Create Giftee
	testGiftee := model.Giftee{
		CustomerID:           customer.ID,
		GifteeName:           "Maya",
		Gender:               "Female",
		CustomerRelationship: "Sister",
		Age:                  20,
		Colors:               pq.StringArray{"Green", "Blue"},
		Interests:            pq.StringArray{"Sports", "Soccer", "Nature", "Coffee", "Candy"},
	}
	err = tx.Create(&testGiftee).Error
	assert.NoError(t, err)

	// Test Inputted Giftee Fields
	var fetchedGiftee model.Giftee
	err = tx.First(&fetchedGiftee, testGiftee.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, testGiftee.ID, fetchedGiftee.ID)
	assert.Equal(t, testGiftee.CustomerID, fetchedGiftee.CustomerID)
	assert.Equal(t, testGiftee.GifteeName, fetchedGiftee.GifteeName)
	assert.Equal(t, testGiftee.Gender, fetchedGiftee.Gender)
	assert.Equal(t, testGiftee.CustomerRelationship, fetchedGiftee.CustomerRelationship)
	assert.Equal(t, testGiftee.Age, fetchedGiftee.Age)
	assert.Equal(t, testGiftee.Colors, fetchedGiftee.Colors)
	assert.Equal(t, testGiftee.Interests, fetchedGiftee.Interests)
	assert.Equal(t, testGiftee.GiftRequests, fetchedGiftee.GiftRequests)
	assert.Equal(t, testGiftee.CreatedAt.In(time.UTC).Round(time.Millisecond),
		fetchedGiftee.CreatedAt.In(time.UTC).Round(time.Millisecond))

	// Check that there's only 1 Giftee
	var count int64
	tx.Model(&model.Giftee{}).Where("id = ?", testGiftee.ID).Count(&count)
	assert.Equal(t, int64(1), count)

	req1, err := http.NewRequest("DELETE", fmt.Sprintf("/giftee/%d", testGiftee.ID), nil)
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}
	router.ServeHTTP(w, req1)
	assert.Equal(t, 204, w.Code)

	// Check that Giftee has been deleted
	var deletedCount int64
	tx.Model(&model.Giftee{}).Where("id = ?", testGiftee.ID).Count(&deletedCount)
	assert.Equal(t, int64(0), deletedCount)
}

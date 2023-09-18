package tests

import (
	"context"
	c "CaitsCurates/backend/src/controller"
	"CaitsCurates/backend/src/model"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-playground/assert/v2"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/jackc/pgx"
)

func TestAddGetGift(t *testing.T) {
	// This Code should be at the top of every test function written
	db_url, exists := os.LookupEnv("TEST_DATABASE_URL")
	fmt.Print(exists)
	cfg := pgx.ConnConfig{
		User:     "testuser",
		Database: "testdb",
		Password: "testpwd",
		Host:     "test-db",
		Port:     5433,
	}
	var err error
	if exists {
		cfg, err = pgx.ParseConnectionString(db_url)

		if err != nil {
			panic(err)
		}
	}

	conn, err := pgx.Connect(cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	defer conn.Close()
	ctx := context.Background()
	tx, err := conn.BeginEx(ctx, nil)

	if err != nil {
		t.Fatalf("Error starting transaction: %v", err)
	}
	// Always rollback the transaction when exiting this function
	defer tx.Rollback()

	m := &model.PgModel{
		Conn: conn,
	}
	c := &c.PgController{
		Model: m,
	}
	router := c.Serve()

	// This code is unique to each test you are writing
	w := httptest.NewRecorder()
	w1 := httptest.NewRecorder()
	testGift := model.ExampleGift{
		GiftId: 4,
		Name:  "nice sweater",
		Price: 50,
	}
	giftJson, err := json.Marshal(testGift)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", "/addGift", bytes.NewBuffer(giftJson))
	if err != nil {
		panic(err)
	}

	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	var giftAdded model.ExampleGift
	if e := json.Unmarshal(w.Body.Bytes(), &giftAdded); e != nil {
		panic(e)
	}
	assert.Equal(t, testGift, giftAdded)

	req1, err := http.NewRequest("GET", "/gifts/4", nil)
	router.ServeHTTP(w1, req1)

	assert.Equal(t, 200, w1.Code)
	var giftRetrieved model.ExampleGift

	if e := json.Unmarshal(w1.Body.Bytes(), &giftRetrieved); e != nil {
		panic(e)
	}


	assert.Equal(t, testGift, giftRetrieved)
}

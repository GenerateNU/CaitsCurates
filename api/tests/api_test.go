package tests

import (
	c "CaitsCurates/backend/src/controller"
	"CaitsCurates/backend/src/model"
	"encoding/json"
	"fmt"
	"github.com/go-playground/assert/v2"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/jackc/pgx"
)

func TestGetGifts(t *testing.T) {
	db_url, exists := os.LookupEnv("DATABASE_URL")

	cfg := pgx.ConnConfig{
		User:     "user",
		Database: "CaitsDB",
		Password: "pwd",
		Host:     "127.0.0.1",
		Port:     5432,
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

	m := &model.PgModel{
		Conn: conn,
	}
	c := &c.PgController{
		Model: m,
	}
	router := c.Serve()

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/gifts/1", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var gift model.ExampleGift

	if e := json.Unmarshal(w.Body.Bytes(), &gift); e != nil {
		panic(err)
	}

	testGift := model.ExampleGift{
		GiftId: 1,
		Name:  "nice sweater",
		Price: 50,
	}
	assert.Equal(t, testGift, gift)
}

package tests

import (
	"database/sql"
	"testing"

	_ "github.com/lib/pq"
)

func TestDBConnection(t *testing.T) {
	dbURL :=  "user=testuser password=testpwd host=test-db port=5432 dbname=testdb sslmode=disable"


	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		t.Fatalf("failed to connect to the database: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		t.Fatalf("failed to ping the database: %v", err)
	}
}

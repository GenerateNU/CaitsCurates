package main


import (
	"CaitsCurates/backend/src/controller"
	"CaitsCurates/backend/src/model"
	"fmt"
	"os"

	"github.com/jackc/pgx"
)

func main() {
	db_url, exists := os.LookupEnv("DATABASE_URL")

	cfg := pgx.ConnConfig{
		User:     "postgres",
		Database: db_url,
		Password: "pwd",
		Host:     "127.0.0.1",
		Port:     5432,
	}
	var err error
	if exists {
		cfg, err = pgx.ParseConnectionString(db_url)

		if err != nil {
			print(db_url)
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
	c := &controller.PgController{
		Model: m,
	}

	c.Serve().Run(":8080")
}


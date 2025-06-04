package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/marcolino/jukebox/gen/openapi"
	"github.com/marcolino/jukebox/internal/api"
	"github.com/marcolino/jukebox/internal/resources/database/postgres"
)

func main() {
	// db_url := os.Getenv("DATABASE_URL")
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	moduleHandler := postgres.New(db)
	handler := api.NewHandler(moduleHandler)

	server, err := openapi.NewServer(handler)
	if err != nil {
		log.Fatal(err)
	}

	if err = http.ListenAndServe(":9090", server); err != nil {
		log.Fatal(err)
	}

}

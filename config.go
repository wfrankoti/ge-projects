package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	// _ "github.com/lib/pq"
)

var (
	dbHost     = getEnv("DB_HOST", "localhost")
	dbPort     = getEnv("DB_PORT", "5432")
	dbUser     = getEnv("DB_USER", "user")
	dbPassword = getEnv("DB_PASSWORD", "password")
	dbName     = getEnv("DB_NAME", "dbname")
)

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

func NewPostgresClient() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	return db
}

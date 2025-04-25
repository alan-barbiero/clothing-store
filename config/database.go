package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var dbPool *pgxpool.Pool

func ConnectDB() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("error loading .env file:", err)
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	fmt.Println("Attempting to connect to database with connection string:", connString)

	dbPool, err = pgxpool.New(context.Background(), connString)
	if err != nil {
		log.Fatal("error connecting to database:", err)
	} else {
		fmt.Println("Successfully connected to database")
	}
}

func GetDB() *pgxpool.Pool {
	if dbPool == nil {
		log.Fatal("database connection is nil")
	}
	return dbPool
}

func CloseDB() {
	if dbPool != nil {
		dbPool.Close()
	}
}

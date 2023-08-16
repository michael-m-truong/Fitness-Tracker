package respository

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	db      *sql.DB
	dbMutex sync.Once

	dbUser string
	dbPass string
	dbHost string
	dbPort string
	dbName string
)

func init() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
		return
	}

	// Get database connection details from environment variables
	dbUser = os.Getenv("DB_USER")
	dbPass = os.Getenv("DB_PASS")
	dbHost = os.Getenv("DB_HOST")
	dbPort = os.Getenv("DB_PORT")
	dbName = os.Getenv("DB_NAME")

	initDB()
}

func initDB() error {

	// Construct the database connection string
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPass, dbHost, dbPort, dbName)

	// Open a database connection
	var errOpen error
	db, errOpen = sql.Open("postgres", connStr)
	if errOpen != nil {
		return errOpen
	}

	// Test the connection
	errPing := db.Ping()
	if errPing != nil {
		return errPing
	}

	log.Printf("Connected to PostgreSQL database!")
	return nil
}

// GetDB returns the singleton database connection instance
func GetDB() *sql.DB {
	dbMutex.Do(func() {
		if err := initDB(); err != nil {
			log.Fatal(err)
		}
	})
	return db
}

package databases

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq" // PostgreSQL driver
)

// InitializePostgres initializes the PostgreSQL database connection.
func InitializePostgres() *sql.DB {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	} else {
		log.Println("DATABASE_URL environment variable loaded successfully")
	}

	var db *sql.DB
	var err error
	maxRetries := 5
	for i := 0; i < maxRetries; i++ {
		db, err = sql.Open("postgres", dbURL)
		if err == nil {
			// Test the connection to the database
			if err = db.Ping(); err == nil {
				log.Println("Successfully connected to PostgreSQL database")
				break
			}
		}

		log.Printf("Failed to ping the database (attempt %d/%d): %v", i+1, maxRetries, err)
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatalf("Could not connect to the database after %d attempts: %v", maxRetries, err)
	}

	configureDBPooling(db)

	// Create the necessary tables
	if err := createTables(db); err != nil {
		log.Fatalf("Failed to create tables: %v", err)
	} else {
		log.Println("Tables created successfully in PostgreSQL database")
	}

	return db
}

// configureDBPooling configures the connection pool settings for the database.
func configureDBPooling(db *sql.DB) {
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)
	log.Println("Database connection pooling configured")
}

// createTables creates the necessary tables in the PostgreSQL database.
func createTables(db *sql.DB) error {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100),
		email VARCHAR(100) UNIQUE NOT NULL,
		is_active BOOLEAN DEFAULT TRUE,
		user_level INTEGER DEFAULT 0
	);`

	if _, err := db.Exec(createUsersTable); err != nil {
		return err
	}

	return nil
}

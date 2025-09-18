package db

import (
	"context"
	"database/sql"
	"fmt"
	"os"
)

var dbInstance *sql.DB

type DatabaseConfig struct {
	host     string
	port     string
	user     string
	password string
	dbname   string
	DB       *sql.DB
}

func NewDatabase(user, host, port, dbname, password string) *DatabaseConfig {
	db := &DatabaseConfig{
		user:     user,
		host:     host,
		port:     port,
		dbname:   dbname,
		password: password,
	}

	return db
}

func (db *DatabaseConfig) Connect() error {
	connStr := fmt.Sprintf("user=%s host=%s port=%s dbname=%s password=%s sslmode=disable", db.user, db.host, db.port, db.dbname, db.password)
	dbProxy, err := sql.Open("postgres", connStr)

	if err != nil {
		if dbProxy != nil {
			dbProxy.Close()
		}
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	if err := dbProxy.PingContext(context.Background()); err != nil {
		dbProxy.Close()
		return fmt.Errorf("pinging database: %w", err)
	}

	db.DB = dbProxy
	return nil
}

func (db *DatabaseConfig) Close() {
	if db.DB != nil {
		db.DB.Close()
	}
}

// getDatabase returns a singleton instance of the database connection
func getDatabase() (*sql.DB, error) {
	if dbInstance == nil {
		user := os.Getenv("DB_USER")
		host := os.Getenv("DB_HOST")
		port := os.Getenv("DB_PORT")
		dbname := os.Getenv("DB_NAME")
		password := os.Getenv("DB_PASSWORD")

		if user == "" || host == "" || port == "" || dbname == "" || password == "" {
			return nil, fmt.Errorf("database configuration is not set properly in environment variables")
		}

		db := NewDatabase(user, host, port, dbname, password)

		if err := db.Connect(); err != nil {
			return nil, fmt.Errorf("database connection failed: %w", err)
		}

		dbInstance = db.DB
	}

	return dbInstance, nil
}

// CloseDatabase closes the global database connection
func CloseDatabase() {
	if dbInstance != nil {
		dbInstance.Close()
		dbInstance = nil
	}
}

// ConnectionAvailability checks if database connection is available
func ConnectionAvailability() bool {
	_, err := getDatabase()
	return err == nil
}

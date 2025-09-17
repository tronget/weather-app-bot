package db

import (
	"context"
	"database/sql"
	"fmt"
)

var dbInstance *sql.DB

type DatabaseConfig struct {
	host     string
	port     int
	user     string
	password string
	dbname   string
	DB       *sql.DB
}

func NewDatabase(user, host string, port int, dbname, password string) *DatabaseConfig {
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
	connStr := fmt.Sprintf("user=%s host=%s port=%d dbname=%s password=%s sslmode=disable", db.user, db.host, db.port, db.dbname, db.password)
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

func getDatabase() (*sql.DB, error) {
	if dbInstance == nil {
		db := NewDatabase(
			"tronget",
			"localhost",
			5432,
			"weather_app_bot",
			"postgres",
		)

		if err := db.Connect(); err != nil {
			return nil, fmt.Errorf("database connection failed: %w", err)
		}

		dbInstance = db.DB
	}

	return dbInstance, nil
}

// ConnectionAvailability checks if database connection is available
func ConnectionAvailability() bool {
	_, err := getDatabase()
	return err == nil
}

package db

import (
	"fmt"
	"github.com/tronget/weather-app-bot/internal/locales"
	"github.com/tronget/weather-app-bot/internal/network/db/models"
	"log"
	"slices"
)

func IsUserExist(username string) (*models.User, bool) {
	db, err := getDatabase()
	if err != nil {
		log.Printf("Database not available: %v", err)
		return nil, false
	}

	row := db.QueryRow("SELECT username, lang_code FROM users WHERE username = $1", username)

	var name, langCode string
	if err := row.Scan(&name, &langCode); err != nil {
		return nil, false
	}
	return models.NewUser(username, langCode), true
}

func CreateUser(username, langCode string) (*models.User, error) {
	if !slices.Contains(locales.AvailableLanguages, langCode) {
		langCode = locales.LANG_EN
	}

	db, err := getDatabase()
	if err != nil {
		log.Printf("Database not available, cannot create user: %v", err)
		return models.NewUser(username, langCode), nil // Return user object even without DB
	}

	_, err = db.Exec("INSERT INTO users (username, lang_code) VALUES ($1, $2)", username, langCode)
	if err != nil {
		return nil, err
	}

	return models.NewUser(username, langCode), nil
}

func GetUserLanguage(username string) string {
	if !ConnectionAvailability() {
		log.Printf("Database not available, using default language")
		return locales.LANG_EN
	}

	user, err := createUserIfNotExists(username, locales.LANG_EN)
	if err != nil {
		log.Printf("error creating user before getting user language: %v", err)
		return locales.LANG_EN
	}

	return user.LangCode
}

// SetUserLanguage sets user's language in the database.
// If user does not exist, it creates a new user with the given language code.
// If database is not available, it logs the error but doesn't fail.
func SetUserLanguage(username, langCode string) {
	db, err := getDatabase()
	if err != nil {
		log.Printf("Database not available, cannot save user language: %v", err)
		return
	}

	_, err = createUserIfNotExists(username, langCode)
	if err != nil {
		log.Printf("error creating user before setting user language: %v", err)
		return
	}

	_, err = db.Exec("UPDATE users SET lang_code = $1 WHERE username = $2", langCode, username)
	if err != nil {
		log.Printf("error updating user language: %v", err)
	}
}

// Inserting user into a database if they don't exist
// and return *models.User as in database
func createUserIfNotExists(username, langCode string) (*models.User, error) {
	if user, ok := IsUserExist(username); ok {
		return user, nil
	}

	user, err := CreateUser(username, langCode)

	if err != nil {
		return nil, fmt.Errorf("error creating user in database: %w", err)
	}

	return user, nil
}

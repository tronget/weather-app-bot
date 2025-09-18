package locales

import (
	"encoding/json"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"log"
	"os"
	"path/filepath"
	"slices"
)

var bundle *i18n.Bundle

func InitI18n() {
	// Try multiple possible paths for the locales directory
	possiblePaths := []string{
		"internal/locales/json",       // For local development
		"./internal/locales/json",     // Alternative local path
		"/app/internal/locales/json",  // Docker build context path
		"/root/internal/locales/json", // Docker runtime path
	}

	var localesPath string
	var files []os.DirEntry
	var err error

	// Find the correct path that exists
	for _, path := range possiblePaths {
		files, err = os.ReadDir(path)
		if err == nil {
			localesPath = path
			break
		}
	}

	if err != nil {
		log.Panicf("Could not find locales directory in any of the expected paths: %v. Error: %v", possiblePaths, err)
	}

	b := i18n.NewBundle(language.English)
	b.RegisterUnmarshalFunc("json", json.Unmarshal)

	// Load translation files
	for _, file := range files {
		if file.IsDir() || !isJSON(file.Name()) {
			continue
		}

		filePath := filepath.Join(localesPath, file.Name())
		_, err := b.LoadMessageFile(filePath)
		if err != nil {
			log.Panicf("Failed to load translation file %s: %v", filePath, err)
		}
	}

	bundle = b
	log.Printf("Successfully loaded %d translation files", len(files))
}

func isJSON(fileName string) bool {
	return len(fileName) > 5 && filepath.Ext(fileName) == ".json"
}

func newLocalizer(lang string) *i18n.Localizer {
	if bundle == nil {
		InitI18n()
	}
	if !slices.Contains(AvailableLanguages, lang) {
		lang = "en"
	}
	return i18n.NewLocalizer(bundle, lang)
}

func Translate(messageID, lang string) string {
	localizer := newLocalizer(lang)
	localizeConfig := &i18n.LocalizeConfig{
		MessageID: messageID,
	}

	s, err := localizer.Localize(localizeConfig)

	if err == nil {
		return s
	}

	localizer = newLocalizer(LANG_EN)
	return localizer.MustLocalize(localizeConfig)
}

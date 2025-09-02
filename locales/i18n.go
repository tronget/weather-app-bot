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

var availableLanguages = []string{
	LANG_EN,
	LANG_RU,
	LANG_ES,
	LANG_DE,
	LANG_FR,
	LANG_IT,
	LANG_PT,
	LANG_ZH,
	LANG_SV,
	LANG_FI,
}

var bundle *i18n.Bundle

func InitI18n() {
	files, err := os.ReadDir("locales/json")
	if err != nil {
		log.Fatal(err)
	}

	b := i18n.NewBundle(language.English)
	b.RegisterUnmarshalFunc("json", json.Unmarshal)

	// Load translation files
	for _, file := range files {
		if file.IsDir() || !isJSON(file.Name()) {
			continue
		}

		_, err := b.LoadMessageFile("locales/json/" + file.Name())
		if err != nil {
			log.Fatal(err)
		}
	}

	bundle = b
}

func isJSON(fileName string) bool {
	return len(fileName) > 5 && filepath.Ext(fileName) == ".json"
}

func newLocalizer(lang string) *i18n.Localizer {
	if bundle == nil {
		InitI18n()
	}
	if !slices.Contains(availableLanguages, lang) {
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

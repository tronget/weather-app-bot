package locales

import (
	"encoding/json"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"slices"
)

var availableLanguages = []string{LANG_EN, LANG_RU, LANG_ES}
var files = []string{"locales/json/en.json", "locales/json/ru.json", "locales/json/es.json"}
var bundle *i18n.Bundle

func InitI18n() {
	b := i18n.NewBundle(language.English)
	b.RegisterUnmarshalFunc("json", json.Unmarshal)

	// Load translation files
	for _, file := range files {
		_, err := b.LoadMessageFile(file)
		if err != nil {
			panic(err)
		}
	}

	bundle = b
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
	return localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: messageID,
	})
}

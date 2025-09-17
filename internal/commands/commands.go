package commands

import (
	"errors"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/tronget/weather-app-bot/internal/config"
	"github.com/tronget/weather-app-bot/internal/ierrors"
	"github.com/tronget/weather-app-bot/internal/locales"
	"github.com/tronget/weather-app-bot/internal/weather/service"
	"log"
)

func Handle(commandName string, msgConfig *tgbotapi.MessageConfig, lang string) string {
	var replyMessageText string

	switch commandName {
	case "start":
		replyMessageText = start(lang)
	case "help":
		replyMessageText = help(lang)
	case "language":
		replyMessageText = language(msgConfig, lang)
	default:
		replyMessageText = locales.Translate(locales.UNKNOWN_CMD, lang)
	}

	return replyMessageText
}

func start(lang string) string {
	return locales.Translate(locales.START_MESSAGE, lang)
}

func help(lang string) string {
	return locales.Translate(locales.HELP_MESSAGE, lang)
}

func language(msgConfig *tgbotapi.MessageConfig, lang string) string {
	msgConfig.ReplyMarkup = locales.CreateLanguageKeyboard()
	return locales.Translate(locales.CHOOSE_LANG, lang)
}

func HandleDefault(update *tgbotapi.Update, cfg *config.Config, lang string) string {
	// Name of the city passed by user
	userMessageText := update.Message.Text

	username := update.Message.From.UserName
	cityName, err := service.GetCorrectCityName(userMessageText, cfg)

	var replyMessageText string
	var cityNotFoundError *ierrors.CityNotFoundError

	switch {
	case errors.As(err, &cityNotFoundError):
		formatString := locales.Translate(locales.CITY_NOT_FOUND, lang)
		replyMessageText = fmt.Sprintf(formatString, userMessageText)
	case err != nil:
		replyMessageText = locales.Translate(locales.ERROR_MESSAGE, lang)
		log.Printf("Error occurred during user request: %v", err)
	default:
		userLang := cfg.UserLanguage(username)
		weather, err := service.GetWeatherInfo(cityName, cfg, userLang)
		if err != nil {
			replyMessageText = locales.Translate(locales.ERROR_MESSAGE, lang)
			log.Printf("Error occurred during user request: %v", err)
			break
		}
		replyMessageText = weather.BuildMessage(lang)
	}

	// Trim message if it's too long
	if len(replyMessageText) > 4096 {
		replyMessageText = replyMessageText[:4093] + "..."
	}

	return replyMessageText
}

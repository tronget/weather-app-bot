package commands

import (
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/tronget/weather-app-bot/config"
	"github.com/tronget/weather-app-bot/ierrors"
	"github.com/tronget/weather-app-bot/locales"
	"github.com/tronget/weather-app-bot/weather/service"
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
		replyMessageText = language(msgConfig)
	default:
		replyMessageText = "I don't know that command :("
	}

	return replyMessageText
}

func start(lang string) string {
	return locales.Translate(locales.START_MESSAGE, lang)
}

func help(lang string) string {
	return locales.Translate(locales.HELP_MESSAGE, lang)
}

func language(msgConfig *tgbotapi.MessageConfig) string {
	msgConfig.ReplyMarkup = createLanguageKeyboard()
	return "Choose language:"
}

func HandleDefault(update *tgbotapi.Update, cfg *config.Config) string {
	userMessageText := update.Message.Text
	userID := update.Message.From.ID
	cityName, err := service.GetCorrectCityName(userMessageText, cfg)

	var replyMessageText string
	var cityNotFoundError *ierrors.CityNotFoundError

	switch {
	case errors.As(err, &cityNotFoundError):
		replyMessageText = err.Error()
	case err != nil:
		replyMessageText = "Error occurred during request. Be sure you passed correct city name."
		log.Printf("Error occurred during user request: %v", err)
	default:
		userLang := cfg.UserLanguage(userID)
		weather, err := service.GetWeatherInfo(cityName, cfg, userLang)
		if err != nil {
			replyMessageText = "Error occurred during request. Be sure you passed correct city name."
			log.Printf("Error occurred during user request: %v", err)
			break
		}
		replyMessageText = weather.BuildMessage()
	}

	// Trim message if it's too long
	if len(replyMessageText) > 4096 {
		replyMessageText = replyMessageText[:4093] + "..."
	}

	return replyMessageText
}

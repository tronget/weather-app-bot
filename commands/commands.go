package commands

import (
	"errors"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/tronget/weather-app-bot/config"
	"github.com/tronget/weather-app-bot/ierrors"
	"github.com/tronget/weather-app-bot/weather/models"
	"github.com/tronget/weather-app-bot/weather/service"
	"log"
	"time"
)

func Handle(msgConfig *tgbotapi.MessageConfig, commandName string) string {
	var replyMessageText string

	switch commandName {
	case "start", "help":
		replyMessageText = start()
	case "language":
		replyMessageText = language(msgConfig)
	default:
		replyMessageText = "I don't know that command :("
	}

	return replyMessageText
}

func start() string {
	return `This is simple weather bot!
You can rapidly get information about weather all around the world!
Simply enter the name of the place where you want to know the weather.
For example, New York`
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
		replyMessageText = BuildWeatherMessage(weather) // TODO: implement better
	}

	return replyMessageText
}

func BuildWeatherMessage(weather *models.Weather) string {

	weatherDesc := "Нет данных"
	weatherEmoji := "🌤️"
	if len(weather.DescriptionList) > 0 {
		weatherDesc = weather.DescriptionList[0].Description
		weatherEmoji = iconCodeToEmoji(weather.DescriptionList[0].IconID)
	}

	loc := time.FixedZone("local", weather.TimeZone)
	sunrise := weather.Sys.Sunrise.In(loc).Format("15:04")
	sunset := weather.Sys.Sunset.In(loc).Format("15:04")

	msg := fmt.Sprintf(
		"🌍 %s, %s\n"+
			"%s %s\n"+
			"🌡️ Температура: %.1f°C (ощущается как %.1f°C)\n"+
			"💨 Ветер: %.1f м/с\n"+
			"🌅 Восход: %s\n"+
			"🌇 Закат: %s",
		weather.CityName, weather.Sys.Country,
		weatherEmoji, weatherDesc,
		weather.Temperature.Temp, weather.Temperature.FeelsLike,
		weather.Wind.Speed,
		sunrise, sunset,
	)

	return msg
}

func iconCodeToEmoji(code string) string {
	switch code {
	case "01d":
		return "☀️" // clear sky day
	case "01n":
		return "🌑" // clear sky night
	case "02d":
		return "🌤️" // few clouds day
	case "02n":
		return "☁️🌙" // few clouds night
	case "03d", "03n":
		return "☁️" // scattered clouds
	case "04d", "04n":
		return "☁️☁️" // broken clouds
	case "09d", "09n":
		return "🌧️" // shower rain
	case "10d":
		return "🌦️" // rain day
	case "10n":
		return "🌧️🌙" // rain night
	case "11d", "11n":
		return "⛈️" // thunderstorm
	case "13d", "13n":
		return "❄️" // snow
	case "50d", "50n":
		return "🌫️" // mist
	default:
		return "❔" // unknown
	}
}

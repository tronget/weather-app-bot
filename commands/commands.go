package commands

import (
	"errors"
	"github.com/tronget/weather-app-bot/config"
	"github.com/tronget/weather-app-bot/ierrors"
	"github.com/tronget/weather-app-bot/weather/service"
)

func Handle(command string) string {
	var replyMessageText string

	switch command {
	case "start", "help":
		replyMessageText = `This is simple weather bot!
You can rapidly get information about weather all around the world!
Simply enter the name of the place where you want to know the weather.
For example, New York`

	default:
		replyMessageText = "I don't know that command :("
	}

	return replyMessageText
}

func HandleDefault(messageText string, cfg *config.Config) string {
	cityName, err := service.GetCorrectCityName(messageText, cfg)

	var replyMessageText string
	var cityNotFoundError *ierrors.CityNotFoundError

	switch {
	case errors.As(err, &cityNotFoundError):
		replyMessageText = err.Error()
	case err != nil:
		replyMessageText = "Error occurred during request. Sorry, maybe we have some problems.\n" + err.Error()
	default:
		replyMessageText = cityName
	}

	return replyMessageText
}

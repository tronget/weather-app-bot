package commands

import (
	"fmt"
	"github.com/tronget/weather-app-bot/weather"
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

func HandleDefault(messageText string) string {
	cities, err := weather.GetCity(messageText)

	var replyMessageText string

	switch {
	case err != nil:
		replyMessageText = "Error occurred during request. Sorry, maybe we have some problems.\n" + err.Error()
	case len(cities) == 0:
		replyMessageText = fmt.Sprintf("City \"%s\" not found.", messageText)
	default:
		city := cities[0]
		replyMessageText = fmt.Sprintf("%s\n%f : %f, ", city.Name, city.Lon, city.Lat)
	}

	return replyMessageText
}

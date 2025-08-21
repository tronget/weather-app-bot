package internal

const (
	TELEGRAM_APIKEY         = "TELEGRAM_APIKEY"
	OPENWEATHER_APIKEY      = "OPENWEATHER_APIKEY"
	COORDINATES_REQUEST_URL = "http://api.openweathermap.org/geo/1.0/direct?q=%s&appid=%s"
	WEATHER_REQUEST_URL     = "https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric"
)

type WeatherEmoji int

const (
	ClearDay WeatherEmoji = iota
	ClearNight
	FewCloudsDay
	FewCloudsNight
	ScatteredClouds
	BrokenClouds
	ShowerRain
	Rain
)

var weatherEmojiName = map[WeatherEmoji]string{
	ClearDay:     "☀️",
	ClearNight:   "🌕",
	FewCloudsDay: "⛅",
}

func (w WeatherEmoji) String() string {
	return weatherEmojiName[w]
}

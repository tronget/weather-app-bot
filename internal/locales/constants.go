package locales

const (
	LANG_EN = "en"
	LANG_RU = "ru"
	LANG_ES = "es"
	LANG_DE = "de"
	LANG_FR = "fr"
	LANG_IT = "it"
	LANG_PT = "pt"
	LANG_ZH = "zh"
	LANG_SV = "sv"
	LANG_FI = "fi"
)

var AvailableLanguages = []string{
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

const (
	EMPTY_MESSAGE  = "empty_message"
	UNKNOWN_CMD    = "unknown_command"
	START_MESSAGE  = "start_message"
	HELP_MESSAGE   = "help_message"
	CHOOSE_LANG    = "choose_language"
	LANG_CHOSEN    = "language_chosen"
	LANG_SAVED     = "language_saved"
	CITY_NOT_FOUND = "city_not_found"
	ERROR_MESSAGE  = "error_message"
	NO_DATA        = "no_data"
	TEMPERATURE    = "temperature"
	FEELS_LIKE     = "feels_like"
	WIND           = "wind"
	SUNRISE        = "sunrise"
	SUNSET         = "sunset"

	WEATHER_MSG_FORMAT = "🌍 %s, %s\n" +
		"%s %s\n" +
		"🌡️ %s: %.1f°C (%s %.1f°C)\n" +
		"💨 %s: %.1f\n" +
		"🌅 %s: %s\n" +
		"🌇 %s: %s"
)

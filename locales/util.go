package locales

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func CreateLanguageKeyboard() tgbotapi.InlineKeyboardMarkup {
	row1 := tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("🇺🇸 English", "lang_en"),
		tgbotapi.NewInlineKeyboardButtonData("🇷🇺 Русский", "lang_ru"),
		tgbotapi.NewInlineKeyboardButtonData("🇪🇸 Español", "lang_es"),
	)
	row2 := tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("🇩🇪 Deutsch", "lang_de"),
		tgbotapi.NewInlineKeyboardButtonData("🇫🇷 Français", "lang_fr"),
		tgbotapi.NewInlineKeyboardButtonData("🇮🇹 Italiano", "lang_it"),
	)
	row3 := tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("🇵🇹 Português", "lang_pt"),
		tgbotapi.NewInlineKeyboardButtonData("🇨🇳 中文", "lang_zh"),
	)
	row4 := tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("🇸🇪 Svenska", "lang_sv"),
		tgbotapi.NewInlineKeyboardButtonData("🇫🇮 Suomi", "lang_fi"),
	)

	return tgbotapi.NewInlineKeyboardMarkup(row1, row2, row3, row4)
}

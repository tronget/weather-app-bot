package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func createLanguageKeyboard() tgbotapi.InlineKeyboardMarkup {
	row1 := tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("English", "lang_en"),
		tgbotapi.NewInlineKeyboardButtonData("Русский", "lang_ru"),
	)
	row2 := tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Español", "lang_es"),
	)

	return tgbotapi.NewInlineKeyboardMarkup(row1, row2)
}

package telegramBot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramBot struct {
	ApiToken string "env: TELEGRAM_API_TOKEN"
}

type Commands struct {
	weather string
	chatgpt string
}

type TelegramBotImpl struct {
	Bot      *tgbotapi.BotAPI
	commands *Commands
}

package telegramBot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

const (
	defaultTimeout = 60
)

func initTelegram(cfg *TelegramBot) *TelegramBotImpl {
	bot, err := tgbotapi.NewBotAPI(cfg.ApiToken)

	if err != nil {
		log.Panic(err)
	}

	return &TelegramBotImpl{Bot: bot}
}

func (t *TelegramBotImpl) ListenForMessage(fn func(message tgbotapi.Update)) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = defaultTimeout

	updates := t.Bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		fn(update)
	}
}

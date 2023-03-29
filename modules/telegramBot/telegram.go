package telegramBot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func initTelegram(cfg *TelegramBot) *TelegramBotImpl {
	bot, err := tgbotapi.NewBotAPI(cfg.ApiToken)

	if err != nil {
		log.Panic(err)
	}

	return &TelegramBotImpl{Bot: bot}
}

func (b *TelegramBotImpl) ListenForMessage(fn func(message tgbotapi.Update)) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.Bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		fn(update)
	}
}

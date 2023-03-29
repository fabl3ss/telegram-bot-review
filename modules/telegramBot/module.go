package telegramBot

import (
	"go.uber.org/fx"
	"os"
)

var Module = fx.Options(
	fx.Provide(
		func() *TelegramBot {
			return &TelegramBot{ApiToken: os.Getenv("TELEGRAM_API_TOKEN")}
		},
		initTelegram,
	),
)

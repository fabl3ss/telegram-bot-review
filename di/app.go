package di

import (
	"go.uber.org/fx"
	"telegram_bot/modules/chatGPT"
	"telegram_bot/modules/telegramBot"
)

func AppProviders() []fx.Option {
	modules := []fx.Option{
		telegramBot.Module,
		chatGPT.Module,
	}

	return modules
}

package main

import (
	"go.uber.org/fx"
	"telegram_bot/di"
	"telegram_bot/pkg/telegram_server"
)

// WTF ???
func getAllOptions() []fx.Option {
	fxOptions := di.AppProviders()

	return fxOptions
}

func main() {
	fxOptions := getAllOptions()

	telegram_server.StartTelegram(fxOptions)
}

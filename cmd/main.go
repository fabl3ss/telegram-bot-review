package main

import (
	"telegram_bot/di"
	"telegram_bot/pkg/telegram_server"
)

func main() {
	fxOptions := di.GetAllOptions()

	telegram_server.StartTelegram(fxOptions)
}

package config

type ChatGpt struct {
	Token string "env: CHAT_TOKEN"
}

type Telegram struct {
	Token string "env: TELEGRAM_TOKEN"
}

type Config struct {
	ChatGpt  ChatGpt
	Telegram Telegram
}

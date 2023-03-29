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
	//chatGPT.StartChatGPT(fxOptions)
	//fx.New(
	//	chatGPT.Module,
	//	telegramBot.Module,
	//	fx.Invoke(
	//		func(lifecycle fx.Lifecycle, bot *telegramBot.TelegramBotImpl, chatGpt *chatGPT.ChatGptImpl) {
	//			lifecycle.Append(
	//				fx.Hook{
	//					OnStart: func(context.Context) error {
	//						go bot.ListenForMessage(func(update tgbotapi.Update) {
	//
	//							fmt.Printf("Received message: %s\n", update.Message.Text)
	//
	//							update.Message.Text = chatGpt.SendChatGPT(update.Message.Text)
	//
	//							fmt.Println("update.Message.Text", update.Message.Text)
	//							msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	//							msg.ReplyToMessageID = update.Message.MessageID
	//
	//							bot.Bot.Send(msg)
	//						})
	//						return nil
	//					},
	//					OnStop: func(context.Context) error {
	//						return errors.New("ERROR")
	//					},
	//				},
	//			)
	//		},
	//	),
	//).Run()
}

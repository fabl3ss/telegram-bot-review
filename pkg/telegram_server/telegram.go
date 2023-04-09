package telegram_server

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sashabaranov/go-openai"
	"go.uber.org/fx"
	"telegram_bot/modules/chatGPT"
	"telegram_bot/modules/telegramBot"
)

func StartTelegram(fxOptions []fx.Option) {

	fxOptions = append(
		fxOptions,
		fx.Invoke(
			func(lc fx.Lifecycle, bot *telegramBot.TelegramBotImpl, chatGPT *chatGPT.ChatGptImpl) {
				fmt.Println("INIT_TELEGRAM", chatGPT)
				messages := make([]openai.ChatCompletionMessage, 0)

				bot.ListenForMessage(func(update tgbotapi.Update) {
					messages = append(messages, openai.ChatCompletionMessage{
						Role:    openai.ChatMessageRoleUser,
						Content: update.Message.Text,
					})

					update.Message.Text = chatGPT.SendChatGPT(messages)

					messages = append(messages, openai.ChatCompletionMessage{
						Role:    openai.ChatMessageRoleAssistant,
						Content: update.Message.Text,
					})

					msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
					msg.ReplyToMessageID = update.Message.MessageID
					bot.Bot.Send(msg)
				})
			},
		),
	)

	app := fx.New(fxOptions...)

	app.Run()
}

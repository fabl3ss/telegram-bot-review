package chatGPT

import (
	"go.uber.org/fx"
	"os"
)

var Module = fx.Options(
	fx.Provide(
		func() *ChatGpt {
			return &ChatGpt{ApiToken: os.Getenv("CHAT_GPT_API_TOKEN")}
		},
		initChatGpt,
	),
)

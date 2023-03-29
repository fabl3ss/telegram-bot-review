package chatGPT

import (
	openai "github.com/sashabaranov/go-openai"
)

type ChatGpt struct {
	ApiToken string "env: CHAT_GPT_API_TOKEN"
}

type ChatGptImpl struct {
	ChatGpt *openai.Client
}

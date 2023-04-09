package chatGPT

import (
	"context"
	"github.com/sashabaranov/go-openai"
)

func initChatGpt(cfg *ChatGpt) *ChatGptImpl {
	chatGPT := openai.NewClient(cfg.ApiToken)

	return &ChatGptImpl{ChatGpt: chatGPT}
}

func (c *ChatGptImpl) SendChatGPT(messages []openai.ChatCompletionMessage) string {
	resp, err := c.ChatGpt.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    openai.GPT3Dot5Turbo,
			Messages: messages,
		},
	)

	if err != nil {
		return "ChatGPT API error"
	} else {
		return resp.Choices[0].Message.Content
	}
}

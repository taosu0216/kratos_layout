package utils

import (
	"context"
	"github.com/sashabaranov/go-openai"
	"log"
)

func Gpt(contextInfo, key, model, url string) string {
	openaiConfig := openai.DefaultConfig(key)
	openaiConfig.BaseURL = url
	client := openai.NewClientWithConfig(openaiConfig)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: model,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: contextInfo,
				},
			},
		},
	)

	if err != nil {
		log.Panicf("ChatCompletion error: %v\n", err)
		return ""
	}

	return resp.Choices[0].Message.Content
}

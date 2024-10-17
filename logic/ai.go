package logic

import (
	"context"
	"fmt"

	"github.com/sashabaranov/go-openai"
)

func Chat(key string, input string) string {
	client := openai.NewClient(key)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4oMini,
			Messages: []openai.ChatCompletionMessage{
				{
					Role: openai.ChatMessageRoleSystem,
					Content: "You are an expert at summarizing documents. You will be given a title and paragraphs from the document. Your task is to summarize the document into a concise, logically coherent summary. You should use headers to organize the summary.",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: input,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return ""
	}

	return resp.Choices[0].Message.Content
}
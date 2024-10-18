package logic

import (
	"context"
	"fmt"
	"os"

	"github.com/sashabaranov/go-openai"
)

func GetSummaryStream(input string) (*openai.ChatCompletionStream, error) {
	c := openai.NewClient(os.Getenv("OPENAI_API_KEY"))
	ctx := context.Background()
	req := openai.ChatCompletionRequest{
		Model:     openai.GPT4oMini,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: "You are an expert at summarizing documents. You will be given a title and paragraphs from the document. Your task is to summarize the document into a concise, logically coherent summary. You should use headers to organize the summary.",
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: input,
			},
		},
		Stream: true,
	}
	stream, err := c.CreateChatCompletionStream(ctx, req)
	if err != nil {
		fmt.Printf("ChatCompletionStream error: %v\n", err)
		return nil, err
	}
	return stream, nil
	
}
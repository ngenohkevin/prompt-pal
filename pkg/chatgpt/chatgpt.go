package chatgpt

import (
	"context"
	"errors"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"github.com/spf13/viper"
)

// ChatGPT is a struct representing an instance of the OpenAI GPT chatbot.
type ChatGPT struct {
	client *openai.Client
}

func GetChatResponse(prompt string) (*ChatGPT, error) {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	apiKey := viper.GetString("API_KEY")
	if apiKey == "" {
		return nil, errors.New("missing OpenAI API Key")
	}

	client := openai.NewClient(apiKey)

	return &ChatGPT{client: client}, nil
}

// GenerateResponse generates a response from the OpenAI GPT chatbot given a prompt.
func (cg *ChatGPT) GenerateResponse(ctx context.Context, prompt string, length int, temperature float64) (string, error) {
	completionReq := &openai.CompletionRequest{
		Model:       "gpt-3.5-turbo-0613",
		Prompt:      prompt,
		MaxTokens:   length,
		Temperature: float32(temperature),
	}
	//completionReq := &openai.CompletionRequest{
	resp, err := cg.client.CreateCompletion(ctx, *completionReq)
	if err != nil {
		return "", fmt.Errorf("failed to generate repsonse: %v", err)
	}
	if len(resp.Choices) == 0 {
		return "", errors.New("no response generated")
	}

	return resp.Choices[0].Text, nil
}

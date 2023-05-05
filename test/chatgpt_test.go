package test

import (
	"errors"
	"github.com/ngenohkevin/prompt-pal.git/pkg/chatgpt"
	"github.com/sashabaranov/go-openai"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetChatResponse(t *testing.T) {
	type test struct {
		name        string
		apiKey      string
		expectedErr error
	}

	tests := []test{
		{
			name:        "success",
			apiKey:      "test-api-key",
			expectedErr: nil,
		},
		{
			name:        "missing api key",
			apiKey:      "",
			expectedErr: errors.New("missing OpenAPI key"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			viper.SetConfigFile(".env")
			err := viper.ReadInConfig()
			if err != nil {
				return
			}
			viper.Set("API_KEY", tt.apiKey)

			client := &openai.Client{}

			chatGPT, err := chatgpt.GetChatResponse("test-prompt")
			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, client, chatGPT)
			}
		})
	}

}

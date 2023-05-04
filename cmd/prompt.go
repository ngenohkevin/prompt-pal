package cmd

import (
	"context"
	"fmt"
	"github.com/ngenohkevin/prompt-pal.git/pkg/chatgpt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

var promptCmd = &cobra.Command{
	Use:   "prompt",
	Short: "Interact with OpenAI's ChatGPT API",
	Long: `The prompt command allows you to interact with OpenAI's ChatGPT API. 
With the prompt command, you can provide a prompt to the API and get a response back.`,
	Run: func(cmd *cobra.Command, args []string) {
		executePrompt()
	},
}

func executePrompt() {
	prompt := viper.GetString("prompt")
	length := viper.GetInt("length")
	temperature := viper.GetFloat64("temperature")

	chatbot, err := chatgpt.GetChatResponse(prompt)
	if err != nil {
		log.Fatal(err)
	}

	response, err := chatbot.GenerateResponse(context.Background(), prompt, length, temperature)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(response)
}

func init() {
	rootCmd.AddCommand(promptCmd)
	promptCmd.Flags().StringP("prompt", "p", "", "The prompt to send to ChatGPT")
	err := viper.BindPFlag("prompt", promptCmd.Flags().Lookup("prompt"))
	if err != nil {
		return
	}
	err = promptCmd.MarkFlagRequired("prompt")
	if err != nil {
		return
	}

	promptCmd.Flags().IntP("length", "l", 128, "The length of the response to generate")
	err = viper.BindPFlag("length", promptCmd.Flags().Lookup("length"))
	if err != nil {
		return
	}

	promptCmd.Flags().Float64P("temperature", "t", 0.5, "The temperature to use for response generation")
	err = viper.BindPFlag("temperature", promptCmd.Flags().Lookup("temperature"))
	if err != nil {
		return
	}
}

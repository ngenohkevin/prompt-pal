package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "prompt-pal",
	Short: "A CLI for OpenAI's GPT-3 language model",
	Long: `PromptPal is a powerful tool for handling user prompts in a simple and intuitive way. 
	Its name reflects its core functionality and emphasizes its ability to provide personalized and helpful responses.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to PromptPal! Use the --help flag to see the available commands.")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "prompt-pal",
	Short: "A command-line tool for integrating OpenAI's ChatGPT API",
	Long: `PromptPal is a command-line tool that allows you to interact with OpenAI's 
ChatGPT API from your terminal. With PromptPal, you can ask technical questions 
and get personalized feedbacks instantly.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to PromptPal. Use 'prompt-pal help' to see available commands.")
		err := cmd.Help()
		if err != nil {
			return
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

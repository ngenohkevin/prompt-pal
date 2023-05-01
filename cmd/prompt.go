package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var promptCmd = &cobra.Command{
	Use:   "prompt",
	Short: "Interact with OpenAI's ChatGPT API",
	Long: `The prompt command allows you to interact with OpenAI's ChatGPT API. 
With the prompt command, you can provide a prompt to the API and get a response back.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("This is the prompt subcommand.")
	},
}

func init() {
	rootCmd.AddCommand(promptCmd)
}

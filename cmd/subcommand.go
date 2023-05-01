package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate response to a prompt using GPT-3",
	Run: func(cmd *cobra.Command, args []string) {
		prompt, _ := cmd.Flags().GetString("prompt")
		length, _ := cmd.Flags().GetInt("length")
		temperature, _ := cmd.Flags().GetFloat64("temperature")

		//	TODO: Implement generating response using GPT-3
		fmt.Printf("Generating response to prompt: %s\n", prompt)
		fmt.Printf("Length: %d,  Temperature: %.2f\n", length, temperature)
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().StringP("prompt", "p", "", "prompt for generating usage")
	err := generateCmd.MarkFlagRequired("prompt")
	if err != nil {
		return
	}
	generateCmd.Flags().IntP("length", "l", 1024, "Length of response")
	generateCmd.Flags().Float64P("temperature", "t", 0.7, "Temperature for generating response")
}

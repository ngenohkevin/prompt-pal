package cmd

//import (
//	"context"
//	"fmt"
//	"github.com/ngenohkevin/prompt-pal.git/pkg/chatgpt"
//	"github.com/spf13/cobra"
//	"github.com/spf13/viper"
//	"log"
//)
//
//var subCommandCmd = &cobra.Command{
//	Use:   "subcommand",
//	Short: "Interact with the OpenAI Chatbot via a subcommand",
//	Run: func(cmd *cobra.Command, args []string) {
//		prompt, _ := cmd.Flags().GetString("prompt")
//		length, _ := cmd.Flags().GetInt("length")
//		temperature, _ := cmd.Flags().GetFloat64("temperature")
//
//		chatbot, err := chatgpt.GetChatResponse(prompt)
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		response, err := chatbot.GenerateResponse(context.Background(), prompt, length, temperature)
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		fmt.Printf("Response: %s\n", response)
//	},
//}
//
//func init() {
//	rootCmd.AddCommand(subCommandCmd)
//
//	subCommandCmd.Flags().StringP("prompt", "p", "", "prompt for generating usage")
//	err := subCommandCmd.MarkFlagRequired("prompt")
//	if err != nil {
//		return
//	}
//	subCommandCmd.Flags().IntP("length", "l", 128, "Length of response")
//	subCommandCmd.Flags().Float64P("temperature", "t", 0.7, "Temperature for generating response")
//	err = viper.BindPFlag("prompt", subCommandCmd.Flags().Lookup("prompt"))
//	if err != nil {
//		return
//	}
//	err = subCommandCmd.MarkFlagRequired("prompt")
//	if err != nil {
//		return
//	}
//}

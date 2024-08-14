/*
Copyright © 2024 jhonatan <jhonatanbds97@gmail.com>
*/
package cmd

import (
	"github.com/spf13/cobra"

	"github.com/jhonatanbds/LlamaShell/internal/chat"
)

// chatCmd represents the chat command
var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Chat with LlamaShell",
	Long: `This iterative chat system utilizes LangChain Go with Ollama to create 
an interactive conversation between you and an the Llama model.`,
	Run: func(cmd *cobra.Command, args []string) {
		chat.Chat()
	},
}

func init() {
	rootCmd.AddCommand(chatCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// chatCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// chatCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

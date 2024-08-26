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
}

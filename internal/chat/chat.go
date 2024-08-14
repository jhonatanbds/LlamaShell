package chat

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

var initialMessage = `You are a CLI assistent who should suggest commands to help Linux users.
	You should only provide the commands for the desired task when required with no formatting.`

func Chat() {
	start := time.Now()
	llm, err := ollama.New(ollama.WithModel("llama3.1"))

	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()

	// Initial system message
	messages := []llms.MessageContent{
		llms.TextParts(llms.ChatMessageTypeSystem, initialMessage),
	}
	elapsed := time.Since(start)
	log.Printf("Ollama loading took %s", elapsed)

	fmt.Println("Chat with 🦙 LlamaShell. Type 'exit' to end the conversation.")

	for {
		// Get user input
		fmt.Print(">>> ")
		reader := bufio.NewReader(os.Stdin)
		userInput, _ := reader.ReadString('\n')
		userInput = strings.TrimSpace(userInput)

		if userInput == "exit" {
			fmt.Println("Ending chat.")
			break
		}

		// Add user message to the conversation
		messages = append(messages, llms.TextParts(llms.ChatMessageTypeHuman, userInput))

		s := spinner.New(spinner.CharSets[21], 100*time.Millisecond)
		s.Suffix = " 🦙"
		s.Start()
		stopSpinner := make(chan struct{})

		completion, err := llm.GenerateContent(ctx, messages, llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
			select {
			case <-stopSpinner:
			default:
				close(stopSpinner)
				s.Stop()
			}
			fmt.Print(string(chunk))
			return nil
		}))
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println() // New line after AI response

		// Add AI response to the conversation
		messages = append(messages, llms.TextParts(llms.ChatMessageTypeAI, completion.Choices[0].Content))
	}
}

package run

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/eiannone/keyboard"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

var initialMessage = `You are a CLI assistent who should suggest commands to help Linux users.
	You should only provide the commands for the desired task when required with no formatting.
	You do not need to explain the meaning of the commands. Assume your answer will be served directly as an input to the terminal.`

type Option int

const (
	run = iota
	regen
	quit
)

func chooseOption(cmd string) (Option, error) {
	options := []string{cmd, "Regenerate"}
	selectedIndex := 0

	err := keyboard.Open()
	if err != nil {
		panic(err)
	}
	defer keyboard.Close()

	for {
		menuString := ""
		for i, option := range options {
			if i == selectedIndex {
				menuString += fmt.Sprintf("> %s\n", option)
			} else {
				menuString += fmt.Sprintf("  %s\n", option)
			}
		}

		// Move cursor up and print the menu
		fmt.Printf("\033[%dA\r%s", len(options), menuString)

		// Wait for keypress
		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		switch key {
		case keyboard.KeyArrowUp:
			if selectedIndex > 0 {
				selectedIndex--
			}
		case keyboard.KeyArrowDown:
			if selectedIndex < len(options)-1 {
				selectedIndex++
			}
		case keyboard.KeyEnter:
			if selectedIndex == 1 {
				return regen, nil
			} else {
				return run, nil
			}
		case keyboard.KeyEsc:
			return quit, nil
		}

		if char == 'q' {
			return quit, nil
		}
	}
}

func Run(prompt string) {
	llm, err := ollama.New(ollama.WithModel("llama3.1"))
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()

	// Initial system message
	messages := []llms.MessageContent{
		llms.TextParts(llms.ChatMessageTypeSystem, initialMessage),
	}

	// Add user message to the conversation
	messages = append(messages, llms.TextParts(llms.ChatMessageTypeHuman, prompt))

	fmt.Println()
	fmt.Println()

	for {
		generatedCmd := ""
		// Generate AI response
		completion, err := llm.GenerateContent(ctx, messages, llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
			generatedCmd += string(chunk)
			return nil
		}))
		if err != nil {
			log.Fatal(err)
		}
		messages = append(messages, llms.TextParts(llms.ChatMessageTypeAI, completion.Choices[0].Content))

		option, err := chooseOption(generatedCmd)
		if err != nil {
			panic(err)
		}

		switch option {
		case run:
			cmd := exec.Command("sh", "-c", generatedCmd)

			currentDir, err := os.Getwd()
			if err != nil {
				panic(err)
			}
			cmd.Dir = currentDir

			out, err := cmd.CombinedOutput()
			if err != nil {
				fmt.Printf("Command failed with error: %v\nOutput: %s\n", err, out)
				return
			}

			fmt.Println(string(out))
		case regen:
			regenMessage := "Provide another option"
			messages = append(messages, llms.TextParts(llms.ChatMessageTypeHuman, regenMessage))
		}
		if option != regen {
			break
		}
	}
}

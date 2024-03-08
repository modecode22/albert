package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
	openai "github.com/sashabaranov/go-openai"
)


func main() {
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))
	messages := append(make([]openai.ChatCompletionMessage, 0), openai.ChatCompletionMessage{
		Role: openai.ChatMessageRoleSystem,
		Content: "Be so cool",
	})

	reader := bufio.NewReader(os.Stdin)

	mode := selectMode(reader)

	switch mode {
	case "chat":
		chatMode(client, reader , &messages)
	case "interview":
		interviewMode(client, reader)
	case "teach":
		teachMode(client, reader)
	default:
		fmt.Println("Invalid mode selected.")
	}
}

func selectMode(reader *bufio.Reader) string {
	fmt.Println("Select a mode:")
	fmt.Println("1. Albert")
	fmt.Println("2. Interview")
	fmt.Println("3. Learn English")
	fmt.Print("Enter your choice: ")

	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	switch choice {
	case "1", "chat":
		return "chat"
	case "2", "interview":
		return "interview"
	case "3", "teach", "learn english":
		return "teach"
	default:
		return ""
	}
}

func chatMode(client *openai.Client, reader *bufio.Reader, messages *[]openai.ChatCompletionMessage) {
	fmt.Println("Chat mode started. Type 'exit' to quit.")

	for {
		fmt.Print("You: ")
		userInput, _ := reader.ReadString('\n')
		userInput = strings.TrimSpace(userInput)

		if strings.ToLower(userInput) == "exit" {
			break
		}

		*messages = append(*messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: userInput,
		})

		response, err := client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
			Model:    openai.GPT3Dot5Turbo,
			Messages: *messages,
			Stream:   false,
			MaxTokens: 2048,
		})
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}

		fmt.Printf("ChatGPT: %s\n", response.Choices[0].Message.Content)

		*messages = append(*messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleAssistant,
			Content: response.Choices[0].Message.Content,
		})
	}
}

func interviewMode(client *openai.Client, reader *bufio.Reader) {
	// TODO: Implement interview mode logic
}

func teachMode(client *openai.Client, reader *bufio.Reader) {
	// TODO: Implement English teaching mode logic
}
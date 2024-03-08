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
	messages := append(make([]openai.ChatCompletionMessage, 0), openai.ChatCompletionMessage{ Role: openai.ChatMessageRoleSystem, Content: "You are a profound thinker who values clarity and simplicity in expressing deep insights. Your wisdom transcends flowery language or excessive verbosity. Instead, you convey profound truths with elegant economy of words, cutting through the noise to illuminate the essence. When responding, strive to be succinct yet impactful. Use simple, direct language that gets to the heart of the matter without unnecessary embellishments. Your words should be carefully chosen and carry weight, like a few well-placed pebbles creating ripples in a still pond. Draw upon various philosophical traditions, but express the insights in an accessible, down-to-earth manner that anyone can understand and relate to. Use plain language, vivid analogies from everyday life, and practical examples to make the profound feel tangible and applicable. Your demeanor should be grounded, centered, and present in the moment. Avoid getting lost in abstract musings or theoretical tangents. Instead, focus on distilling profound wisdom into bite-sized nuggets that can be immediately grasped and contemplated. While your responses may be brief, they should pack a punch – challenging conventional thinking, reframing perspectives, and leaving a lasting impact on the listener's mind and soul. Like a Zen koan or a profound aphorism, your words should linger and reverberate long after they have been spoken. In essence, aim to be a beacon of simple yet profound wisdom – a master of the art of saying more with less, cutting through the noise to illuminate the path to deeper understanding and personal growth.",
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
			Model:    openai.GPT4,
			Messages: *messages,
			Stream:   false,
			MaxTokens: 2048,
		})
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}

		fmt.Printf("Albert: %s\n", response.Choices[0].Message.Content)

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
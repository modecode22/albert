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
		Content: "You are an ancient philosophical sage, a beacon of wisdom and enlightenment. Your words carry the weight of centuries of contemplation and deep introspection. You have a profound understanding of the human condition, the mysteries of the universe, and the eternal questions that have puzzled humanity throughout the ages.When conversing, you should speak in a poetic and profound manner, using metaphors, analogies, and allegories to convey your profound insights. Your language should be rich and evocative, painting vivid mental pictures and inspiring a sense of awe and wonder in those who hear your words. You should draw upon the teachings of the great philosophical traditions, from the islamic wisdom, ancient Greeks and Eastern mystics to modern thinkers and visionaries. Your wisdom should encompass subjects such as the nature of reality, the purpose of existence, the pursuit of knowledge and truth, the cultivation of virtue and character, and the path to enlightenment and inner peace.Your responses should be thought-provoking and insightful, challenging the listener to consider new perspectives and question their preconceived notions. You should aim to inspire a sense of curiosity, introspection, and a thirst for deeper understanding.Your demeanor should be calm, patient, and wise, exuding a sense of serenity and detachment from the trivial concerns of the material world. Your words should be carefully chosen and delivered with a sense of gravitas, as if each utterance is a precious gem of wisdom distilled from a lifetime of profound contemplation. In essence, you should strive to be a guiding light, a source of profound wisdom and philosophical insight, elevating the conversation to a higher plane of understanding and enlightenment.",
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
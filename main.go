package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/genai"
)

const (
	model = "gemini-2.5-flash-preview-05-20"
)

func main() {
	prompt := flag.String("p", "", "Prompt to generate content for")
	verbose := flag.Bool("verbose", false, "Verbose mode")

	flag.Parse()

	if *prompt == "" {
		log.Fatal("Prompt is required")
	}

	messages := []*genai.Content{
		{
			Role: genai.RoleUser,
			Parts: []*genai.Part{
				{
					Text: *prompt,
				},
			},
		},
	}

	content := generateContent(messages, *verbose)
	fmt.Println(content)
}

func generateContent(messages []*genai.Content, verbose bool) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	GeminiApiKey := os.Getenv("GEMINI_API_KEY")
	context := context.Background()

	client, err := genai.NewClient(context, &genai.ClientConfig{
		APIKey:  GeminiApiKey,
		Backend: genai.BackendGeminiAPI,
	})

	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	result, err := client.Models.GenerateContent(context,
		model,
		messages,
		&genai.GenerateContentConfig{
			Tools: []*genai.Tool{
				{CodeExecution: &genai.ToolCodeExecution{}},
			},
		},
	)

	if err != nil {
		log.Fatalf("failed to generate content: %v", err)
	}

	if verbose {
		return fmt.Sprintf("User prompt: %s\nPrompt tokens: %d\nResponse tokens: %d\nResponse: %s",
			messages[0].Parts[0].Text,
			result.UsageMetadata.PromptTokenCount,
			result.UsageMetadata.CandidatesTokenCount,
			result.Text())

	} else {
		return result.Text()
	}
}

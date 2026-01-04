package gemini

import (
	"context"
	"log"
	"os"

	"google.golang.org/genai"
)

var aiClient *genai.Client

func InitAi() {
	key := os.Getenv("GEMINI_API_KEY")
	ctx := context.Background()

	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  key,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}

	aiClient = client
}

func Ask(prompt string) (string, error) {
	ctx := context.Background()

	result, err := aiClient.Models.GenerateContent(
		ctx,
		"gemini-2.5-flash",
		genai.Text(prompt),
		nil,
	)

	if err != nil {
		return "", err
	}

	return result.Text(), nil
}

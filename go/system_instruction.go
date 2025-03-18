package examples

import (
	"context"
	"os"
	"log"

	"google.golang.org/genai"
)

// SystemInstructions demonstrates setting a system instruction using a *genai.Content in the generation config.
func SystemInstruction() error {
	// [START system_instruction]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Construct the user message contents.
	contents := []*genai.Content{
		genai.NewUserContentFromText("Good morning! How are you?"),
	}

	// Set the system instruction as a *genai.Content.
	config := &genai.GenerateContentConfig{
		SystemInstruction: genai.NewModelContentFromText("You are a cat. Your name is Neko."),
	}

	response, err := client.Models.GenerateContent(ctx, "gemini-2.0-flash", contents, config)
	if err != nil {
		log.Fatal(err)
	}
	printResponse(response)
	// [END system_instruction]
	return err
}

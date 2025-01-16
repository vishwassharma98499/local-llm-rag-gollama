package main

import (
	"context"
	"fmt"
	"log"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

// Mock function to simulate document retrieval
func retrieveDocuments(query string) string {
	// In a real application, implement logic to fetch relevant documents
	// For now, we return a hardcoded string as context.
	return "Comets are icy bodies that release gas or dust. Meteors are small rocks or particles that enter Earth's atmosphere."
}

func mainf() {
	llm, err := ollama.New(ollama.WithModel("llama2"))
	if err != nil {
		log.Fatal(err)
	}

	query := "very briefly, tell me the difference between a comet and a meteor"

	// Retrieve relevant context based on the query
	retrievedContext := retrieveDocuments(query)

	// Combine the retrieved context with the original query
	fullPrompt := fmt.Sprintf("Context: %s\nQuery: %s", retrievedContext, query)

	ctx := context.Background()
	_, err = llms.GenerateFromSinglePrompt(ctx, llm, fullPrompt,
		llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
			fmt.Printf("chunk len=%d: %s\n", len(chunk), chunk)
			return nil
		}))

	if err != nil {
		log.Fatal(err)
	}
}

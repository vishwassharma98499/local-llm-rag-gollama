package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

const ollamaAPI = "http://127.0.0.1:11434/api/chat"

type ChatRequest struct {
    Model  string `json:"model"`
    Prompt string `json:"prompt"`
}

type ChatResponse struct {
    Model      string `json:"model"`
    CreatedAt  string `json:"created_at"`
    Message    Message `json:"message"`
    DoneReason string `json:"done_reason"`
    Done       bool   `json:"done"`
}

type Message struct {
    Role    string `json:"role"`
    Content string `json:"content"`
}

func main() {
    model := "llama2"

    // List of prompts to test
    prompts := []string{
        "Can you explain the main differences between a comet and a meteor in simple terms?",
        "What are the main characteristics of comets and meteors?",
        "Can you describe what a comet is and how it differs from a meteor?",
        "Explain comets and meteors in simple terms.",
    }

    for _, prompt := range prompts {
        // Create the request payload
        requestPayload := ChatRequest{Model: model, Prompt: prompt}
        jsonData, err := json.Marshal(requestPayload)
        if err != nil {
            fmt.Println("Error marshalling JSON:", err)
            return
        }

        // Log the request payload
        fmt.Println("Request Payload:", string(jsonData))

        // Send the request to the Ollama API.
        resp, err := http.Post(ollamaAPI, "application/json", bytes.NewBuffer(jsonData))
        if err != nil {
            fmt.Println("Error making POST request:", err)
            return
        }
        defer resp.Body.Close()

        // Check if response status is not OK
        if resp.StatusCode != http.StatusOK {
            body, _ := ioutil.ReadAll(resp.Body)
            fmt.Printf("Error: received status %s with body: %s\n", resp.Status, body)
            return
        }

        // Read and parse the response.
        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            fmt.Println("Error reading response body:", err)
            return
        }

        var chatResponse ChatResponse
        if err := json.Unmarshal(body, &chatResponse); err != nil {
            fmt.Println("Error unmarshalling response:", err)
            return
        }

        // Output detailed information about the response.
        fmt.Printf("Ollama Response:\nModel: %s\nCreated At: %s\nRole: %s\nContent: %s\nDone Reason: %s\nDone: %t\n",
            chatResponse.Model,
            chatResponse.CreatedAt,
            chatResponse.Message.Role,
            chatResponse.Message.Content,
            chatResponse.DoneReason,
            chatResponse.Done,
        )
    }
}

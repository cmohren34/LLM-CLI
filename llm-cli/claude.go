package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	claude "github.com/potproject/claude-sdk-go"
)

// claude-3-opus-20240229
// claude-3-haiku-20240307
// claude-3-sonnet-20240229
// TODO: should probably pull this out into model class or something
const CLAUDE_MODEL_STRING = "claude-3-opus-20240229"
const MAX_TOKENS = 4096
const TEMP = 0.3

// PrepareSystemMessage prepends context to the system message and returns the formatted system message
func GetSystemMessage() string {
	var sb strings.Builder

	// Retrieve the system_message
	system_message := os.Getenv("SYSTEM_MESSAGE")
	if system_message == "" {
		system_message = ""
	}

	// Append the static system message
	// double check we are getting a clean output from env variable
	sb.WriteString(system_message)

	return sb.String()
}

// InitClaude initializes the Claude client
func InitClaude() {
	// start timer
	startTime := time.Now()
	apiKey := os.Getenv("CLAUDE_API_KEY")
	if apiKey == "" {
		log.Fatalln("API key not set in environment variables")
	}
	client := claude.NewClient(apiKey)
	ctx := context.Background()

	Println()
	Print("Session started with ")
	Println(Cyan(CLAUDE_MODEL_STRING))
	Println("---------------------------------------------------------------")
	Print("> ")

	s := bufio.NewScanner(os.Stdin)
	var inputLines []string
	var conversationHistory []claude.RequestBodyMessagesMessages

	for {
		for s.Scan() {
			line := s.Text()
			if line == "done" { // Signal to end the conversation
				Println()
				log.Println(Yellow("End of request signal received."))
				break
			}
			if line == "end" { // Kill program
				duration := formatDuration(time.Since(startTime))
				Println()
				Println(Blue(fmt.Sprintf("Session duration: %s", duration)))
				log.Fatalln(Red("Received end chat signal..."))
				os.Exit(0) // exit
			}
			if line == "refresh" { // Refresh the client
				Println()
				log.Println(Yellow("Refreshing the client..."))
				client = claude.NewClient(apiKey)
				conversationHistory = nil // Clear conversation history
				Println(Cyan("Client refreshed."))
				Print("> ")
				continue
			}
			inputLines = append(inputLines, line)
		}

		userInput := strings.Join(inputLines, "\n")
		inputLines = nil // Clear inputLines for the next request

		// Add the user's message to the conversation history
		conversationHistory = append(conversationHistory, claude.RequestBodyMessagesMessages{
			Role:    claude.MessagesRoleUser,
			Content: userInput,
		})

		// Concatenate all messages into a single string
		var conversationString string
		for _, message := range conversationHistory {
			conversationString += message.Content + "\n"
		}

		// Prepare the request with the user's message
		body_request := claude.RequestBodyMessages{
			Model:       CLAUDE_MODEL_STRING,
			System:      GetSystemMessage(),
			MaxTokens:   MAX_TOKENS,
			Temperature: TEMP,
			Messages: []claude.RequestBodyMessagesMessages{
				{
					Role:    claude.MessagesRoleUser,
					Content: conversationString,
				},
			},
		}

		Println("\nawaiting response...")
		// Send the message and receive a response
		response, err := client.CreateMessages(ctx, body_request)
		if err != nil {
			log.Print(Red("ChatCompletion error: %v\n", err))
			continue
		}

		Println()
		log.Println(fmt.Sprintf(Cyan("*%s Response*"), CLAUDE_MODEL_STRING))
		Println()
		if len(response.Content) > 0 {
			Printf("%s\n\n", response.Content[0].Text)
		}
		Print("> ")
	}
}

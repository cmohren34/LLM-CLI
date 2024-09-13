package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	go_openai "github.com/sashabaranov/go-openai"
)

// const OPENAI_MODEL = go_openai.GPT4Turbo1106
const OPENAI_MODEL = "gpt-4o"

// InitGPT initializes the GPT client
func InitGPT() {
	// start timer
	startTime := time.Now()
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatalln("API key not set in environment variables")
	}
	client := go_openai.NewClient(apiKey)

	request := go_openai.ChatCompletionRequest{
		Model:       OPENAI_MODEL,
		MaxTokens:   4096,
		Temperature: 0.3,
		Messages: []go_openai.ChatCompletionMessage{
			{
				Role:    go_openai.ChatMessageRoleSystem,
				Content: GetSystemMessage(),
			},
		},
	}

	Println()
	Print("Session started with ")
	Println(Cyan(OPENAI_MODEL))
	Println("---------------------------------------------------------------")
	Print("> ")

	s := bufio.NewScanner(os.Stdin)
	var inputLines []string

	for {
		for s.Scan() {
			line := s.Text()
			if line == "done" { // lets make this cmd+enter to end the conversation
				Println()
				log.Println(Yellow("*end of request signal received*"))
				break
			}
			if line == "end" { // kill program
				duration := formatDuration(time.Since(startTime))
				Println()
				Println(Blue(fmt.Sprintf("Session duration: %s", duration)))
				log.Fatalln(Red("Received [end] chat signal..."))
				os.Exit(0)
			}
			if line == "clear" { // TODO: clear the chat context
				inputLines = nil
				Print("> ")
				continue
			}
			inputLines = append(inputLines, line)
		}

		userInput := strings.Join(inputLines, "\n")
		inputLines = nil // Clear inputLines for the next request

		request.Messages = append(request.Messages, go_openai.ChatCompletionMessage{
			Role:    go_openai.ChatMessageRoleUser,
			Content: userInput,
		})

		Println("\nawaiting response...")
		// makes a new call to CreateChatCompletion with the full request history
		response, err := client.CreateChatCompletion(context.Background(), request)
		if err != nil {
			log.Print(Red("ChatCompletion error: %v\n", err))
			continue
		}

		Println()
		log.Println(fmt.Sprintf(Cyan("*%s Response*"), OPENAI_MODEL))
		Println()
		Printf("%s\n\n", response.Choices[0].Message.Content)
		request.Messages = append(request.Messages, response.Choices[0].Message)
		Print("> ")
	}
}

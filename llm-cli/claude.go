package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

// claude-3-opus-20240229
// claude-3-haiku-20240307
// claude-3-sonnet-20240229
const (
	CLAUDE_MODEL_STRING = "claude-3-5-sonnet-20240620"
	MAX_TOKENS          = 8192
	TEMPERATURE         = 0.3
	API_URL             = "https://api.anthropic.com/v1/messages"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type APIRequest struct {
	Model       string    `json:"model"`
	MaxTokens   int       `json:"max_tokens"`
	Temperature float64   `json:"temperature"`
	System      string    `json:"system"`
	Messages    []Message `json:"messages"`
}

type APIResponse struct {
	Content []struct {
		Text string `json:"text"`
	} `json:"content"`
}

// this is called in openai.go as well
func GetSystemMessage() string {
	systemMessage := os.Getenv("SYSTEM_MESSAGE")
	if systemMessage == "" {
		systemMessage = "You are a helpful AI assistant."
	}
	return systemMessage
}

func sendRequest(client *http.Client, apiKey string, request APIRequest) (*APIResponse, error) {
	jsonData, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request: %v", err)
	}

	req, err := http.NewRequest("POST", API_URL, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", apiKey)
	req.Header.Set("anthropic-version", "2023-06-01")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status code %d: %s", resp.StatusCode, string(body))
	}

	var apiResponse APIResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %v", err)
	}

	return &apiResponse, nil
}

func InitClaude() {
	startTime := time.Now()
	apiKey := os.Getenv("CLAUDE_API_KEY")
	if apiKey == "" {
		log.Fatalln("API key not set in environment variables")
	}
	client := &http.Client{}

	Println()
	Print("Session started with ")
	Println(Cyan(CLAUDE_MODEL_STRING))
	Println("---------------------------------------------------------------")
	Print("> ")

	scanner := bufio.NewScanner(os.Stdin)
	var conversationHistory []Message

	for {
		var userInput strings.Builder
		for scanner.Scan() {
			line := scanner.Text()
			switch line {
			case "done":
				Println()
				log.Println(Yellow("End of request signal received."))
				break
			case "end":
				duration := formatDuration(time.Since(startTime))
				Println()
				Println(Blue(fmt.Sprintf("Session duration: %s", duration)))
				log.Fatalln(Red("Received end chat signal..."))
				return
			case "refresh":
				log.Println(Yellow("\nRefreshing the conversation..."))
				conversationHistory = nil
				Println(Cyan("Conversation refreshed."))
				Print("> ")
				continue
			default:
				userInput.WriteString(line + "\n")
				continue
			}
			break
		}

		if userInput.Len() > 0 {
			conversationHistory = append(conversationHistory, Message{
				Role:    "user",
				Content: strings.TrimSpace(userInput.String()),
			})

			fmt.Println("\nAwaiting response...")
			request := APIRequest{
				Model:       CLAUDE_MODEL_STRING,
				MaxTokens:   MAX_TOKENS,
				Temperature: TEMPERATURE,
				System:      GetSystemMessage(),
				Messages:    conversationHistory,
			}

			resp, err := sendRequest(client, apiKey, request)
			if err != nil {
				log.Print(Red("ChatCompletion error: %v\n", err))
				Print("> ")
				continue
			}

			Println()
			log.Printf(fmt.Sprintf(Cyan("*%s Response*"), CLAUDE_MODEL_STRING))
			if len(resp.Content) > 0 {
				Println()
				Printf(resp.Content[0].Text)
				Println()
				conversationHistory = append(conversationHistory, Message{
					Role:    "assistant",
					Content: resp.Content[0].Text,
				})
			} else {
				Println("Received an empty response from the API.")
			}
			Println()
			Print("> ")
		}
	}
}

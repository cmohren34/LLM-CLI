package main

import (
	"flag"
)

var llmClientType string

func init() {
	flag.StringVar(&llmClientType, "client", "claude", "Specify the LLM client to use (gpt, claude, mistral)")
}

func main() {

	flag.Parse()

	switch llmClientType {
	case "gpt":
		InitGPT()
	case "claude":
		InitClaude()
	case "mistral":
		//InitMistral()
	default:
		panic("Invalid client type")
	}
}

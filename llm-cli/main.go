package main

import (
	"flag"
)

//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

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
	case "groq":
		InitGroq()
	default:
		panic("Invalid client type")
	}
}

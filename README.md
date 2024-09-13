# LLM-CLI

command-line interface designed to use as few resources as possible because using in a browser feels very slow. Uses the APIs of OpenAI, Anthropic, and plans for future integration with Mistral and Meta's Llama-3. minimal-resource tool to communicate with these powerful language models directly from the terminal.

## Features

- **Multiple API Integrations**: switch between different LLM APIs like OpenAI and Claude by Anthropic.
- **Efficiency**: Optimized for speed and minimal system resource usage.
- **Terminal Readability**: Enhanced formatting for improved readability directly in the terminal (lol).

### Built-in Commands

when using the client, you can use the following built-in commands by typing them on a new line:

- **`done`**: Marks the end of a prompt to an LLM. This tells the client that you have finished your input and are ready for the LLM to process it.
- **`end`**: Completely closes out the client session. Use this command when you are finished with your session.

### Environment Variables

To securely manage your API keys and system configurations, it is recommended to use environment variables. Set the following environment variables before running the program:

- **`OPENAI_API_KEY`**: Your API key for OpenAI. This key is necessary to authenticate requests to OpenAI's services.
- **`CLAUDE_API_KEY`**: Your API key for Anthropic's Claude. This key is used to access Claude's language model.
- **`SYSTEM_MESSAGE`**: A custom system message that can be used within the client to provide additional instructions and context.

## Running the Program

You can run the program using the following commands:

```bash
$ go run . -client=claude 
$ go run . -client=gpt
$ go run main.go # defaults to openai
```





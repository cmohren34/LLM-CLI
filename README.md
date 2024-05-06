# LLM-CLI

LLM-CLI is a streamlined command-line interface designed to facilitate rapid and resource-efficient interaction with leading language models from various providers. This tool leverages the APIs of OpenAI, Anthropic's Claude, and plans for future integration with Mistral and Meta's Llama-3. The primary goal of LLM-CLI is to provide developers and researchers with a fast, minimal-resource tool to communicate with these powerful language models directly from the terminal.

## Features

- **Multiple API Integrations**: Seamlessly switch between different LLM APIs like OpenAI and Claude by Anthropic.
- **Efficiency**: Optimized for speed and minimal system resource usage.
- **Terminal Readability**: Enhanced formatting for improved readability directly in the terminal.

### Built-in Commands

While interacting with the client, you can use the following built-in commands by typing them on a new line:

- **`done`**: Marks the end of a prompt to an LLM. This tells the client that you have finished your input and are ready for the LLM to process it.
- **`end`**: Completely closes out the client session. Use this command when you are finished with your session and wish to exit the application.

### Environment Variables

To securely manage your API keys and system configurations, it is recommended to use environment variables. Set the following environment variables before running LLM-CLI:

- **`OPENAI_API_KEY`**: Your API key for OpenAI. This key is necessary to authenticate requests to OpenAI's services.
- **`CLAUDE_API_KEY`**: Your API key for Anthropic's Claude. This key is used to access Claude's language model.
- **`SYSTEM_MESSAGE`**: A custom system message that can be used within the client to provide additional instructions and context.



package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"strings"
)

var client *openai.Client

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	openAIAPIKey := os.Getenv("OPENAI_API_KEY")
	if openAIAPIKey == "" {
		log.Fatal("OPENAI_API_KEY not set in .env file")
	}

	cli.AppHelpTemplate = `NAME:
   {{.HelpName}} - {{.Usage}}

USAGE:
   {{if .VisibleFlags}}{{.HelpName}} [options] prompt{{end}}
   {{if .VisibleCommands}}{{.HelpName}} prompt{{end}}

OPTIONS:
   {{range .VisibleFlags}}{{.}}
   {{end}}
`
	app := &cli.App{
		Name:  "regexai",
		Usage: "Generate regex from a prompt",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "explain",
				Aliases: []string{"e"},
				Value:   false,
				Usage:   "Provide a short explanation about the generated regex",
			},
		},
		Action: func(c *cli.Context) error {
			if c.NArg() != 1 {
				err := cli.ShowAppHelp(c)
				if err != nil {
					return err
				}
				return cli.Exit("Error: No prompt provided", 1)
			}

			prompt := c.Args().Get(0)
			explain := c.Bool("explain")
			regex, err := getRegexFromPrompt(prompt, explain)

			if err != nil {
				return cli.Exit(fmt.Sprintf("Error: %v", err), 1)
			}

			fmt.Println(regex)
			return nil
		},
	}

	// Initialize the OpenAI client
	client = openai.NewClient(openAIAPIKey)

	err = app.Run(os.Args)
	if err != nil {
		_, err := fmt.Fprintln(os.Stderr, err)
		if err != nil {
			return
		}
		os.Exit(1)
	}
}

func getRegexFromPrompt(prompt string, explain bool) (string, error) {
	var maxTokens int
	if explain {
		prompt = fmt.Sprintf("You are a regex generator CLI.\n"+
			"Generate a regex pattern for the request being provided.\n"+
			"Answer with the regex pattern as the first line, then briefly explain the generated regex"+
			"If the request is irrelevant, answer 'Invalid request'\n\n"+
			"Request:\n%s", prompt)
		maxTokens = 250
	} else {
		prompt = fmt.Sprintf("You are a regex generator CLI.\n"+
			"Generate a regex pattern for the request being provided.\n"+
			"Only answer with the regex pattern. "+
			"If the request is irrelevant, answer 'Invalid request'\n\n"+
			"Request:\n%s", prompt)
		maxTokens = 150
	}

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
			MaxTokens:   maxTokens,
			N:           1,
			Stop:        nil,
			Temperature: 0.7,
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return "", err
	}

	return strings.TrimSpace(resp.Choices[0].Message.Content), nil
}

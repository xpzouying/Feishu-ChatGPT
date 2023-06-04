package main

import (
	"os"

	"github.com/sashabaranov/go-openai"
	"github.com/xpzouying/Feishu-ChatGPT/internal/domain"
	"github.com/xpzouying/Feishu-ChatGPT/internal/repo/chatgpt"
)

func newLLM() domain.LLMer {
	var (
		openAIToken = os.Getenv("OPENAI_TOKEN")

		// 如果不为空，则设置 openai 的代理模式
		openAIURL = os.Getenv("OPENAI_URL")
	)

	config := openai.DefaultConfig(openAIToken)
	if openAIURL != "" {
		config.BaseURL = openAIURL
	}

	openaiClient := openai.NewClientWithConfig(config)
	llm := chatgpt.NewChatGPT(openaiClient)

	return llm
}

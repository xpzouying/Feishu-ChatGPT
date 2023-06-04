package main

import (
	"os"

	"github.com/agi-cn/llmplugin"
	"github.com/agi-cn/llmplugin/llm/openai"
	"github.com/agi-cn/llmplugin/plugins/agicn_search"
	"github.com/agi-cn/llmplugin/plugins/calculator"
	"github.com/agi-cn/llmplugin/plugins/google"
	"github.com/agi-cn/llmplugin/plugins/stablediffusion"
	"github.com/sirupsen/logrus"
)

func newLLMPluginManager() *llmplugin.PluginManager {
	var (
		openAIToken       = os.Getenv("OPENAI_TOKEN")
		pluginOpenaiModel = os.Getenv("PLUGIN_OPENAI_MODEL")
	)

	var chatgpt *openai.ChatGPT
	if pluginOpenaiModel == "" {
		chatgpt = openai.NewChatGPT(openAIToken)
	} else {
		chatgpt = openai.NewChatGPT(openAIToken, openai.WithModel(pluginOpenaiModel))
	}

	plugins := makePlugins(chatgpt)

	loggingPlugin(plugins)

	return llmplugin.NewPluginManager(
		chatgpt,
		llmplugin.WithPlugins(plugins),
	)
}

func makePlugins(chatgpt *openai.ChatGPT) []llmplugin.Plugin {

	plugins := []llmplugin.Plugin{
		calculator.NewCalculator(),
	}

	{ // Google Search Engine
		var (
			googleEngineID = os.Getenv("GOOGLE_ENGINE_ID")
			googleToken    = os.Getenv("GOOGLE_TOKEN")
		)

		if googleEngineID != "" && googleToken != "" {
			plugins = append(plugins,
				google.NewGoogle(googleEngineID, googleToken, google.WithSummarizer(chatgpt)))
		}
	}

	{ // Customize Search Engine: agi.cn search
		var (
			enableAgicnSearch = os.Getenv("ENABLE_AGICN_SEARCH")
		)
		if enableAgicnSearch == "true" {
			plugins = append(plugins, agicn_search.NewAgicnSearch())
		}
	}

	{ // Stable Diffusion
		if sdAddr := os.Getenv("SD_ADDR"); sdAddr != "" {
			plugins = append(plugins, stablediffusion.NewStableDiffusion(sdAddr))
		}
	}

	return plugins
}

func loggingPlugin(plugins []llmplugin.Plugin) {
	for _, p := range plugins {
		logrus.Infof("load plugin: %v", p.GetName())
	}
}

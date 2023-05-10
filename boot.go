package main

import (
	agents "github.com/scottraio/app-maker/agents"
	plum "github.com/scottraio/plum"
)

func Boot() plum.AppConfig {
	boot := plum.Boot(plum.Initialize{
		// Embedding function
		Embedding: plum.InitEmbeddings("openai"),
		// Which LLMS to use
		LLM: "openai",
	})

	// register the agents
	boot.RegisterAgent("code", agents.AppMakerAgent())

	return boot
}

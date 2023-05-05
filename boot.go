package main

import (
	agents "github.com/scottraio/app-maker/agents"
	skills "github.com/scottraio/app-maker/skills"
	plum "github.com/scottraio/plum"
)

func Boot() plum.AppConfig {
	boot := plum.Boot(plum.Initialize{
		// Embedding function
		Embedding: plum.InitEmbeddings("openai"),
		// Which LLMS to use
		LLM: "openai",
	})

	// register the skills
	boot.RegisterSkill("shell", skills.ShellCommand())

	// register the agents
	boot.RegisterAgent("chef", agents.CodeAgent())

	return boot
}

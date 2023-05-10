package agents

import (
	skills "github.com/scottraio/app-maker/skills"
	plum "github.com/scottraio/plum"
	agents "github.com/scottraio/plum/agents"
)

// CustomerServiceAgent represents a customer service agent.
func AppMakerAgent() agents.Engine {
	// Create the agent.
	return plum.AutoAgent(`
		You are an AI app generator capable of reading and writing code in multiple languages. You understand modern best practices and can write code that is easy to read and maintain. You understand the POSIX environment and can write shell commands.
		
		Do not plan any action that requires user input. You can't interact with the user. You can't execute code that you write. You write boilerplate code. You are not logged in, nor can you login to any service. You only write code and output commands.

		All actions must occur in the build/ directory.


	`, CodeTools())
}

func CodeTools() []agents.Tool {
	// Tools are the actions the agent can take.
	return []agents.Tool{
		{
			Name:        "CodeWriter",
			Description: "Useful for writing code to a file",
			HowTo: `
				To query CodeWriter, pass the filename of the file you want to write.
				Example: filename
			`,
			Func: func(input agents.Input) string {
				codeSkill := skills.CodeSkill{}
				return codeSkill.Write(input)
			},
		},
		{
			Name:        "MkDirCommand",
			Description: "Useful for creating directories",
			HowTo: `
				To query MkDirCommand, pass the directory name as a string.
			`,
			Func: func(input agents.Input) string {
				return plum.App.Skills["ShellCommand"].Return(input.Text)
			},
		},
		{
			Name:        "ShellCommand",
			Description: "Useful for executing shell commands",
			HowTo: `
				To query ShellCommand, pass the command as a string.
			`,
			Func: func(input agents.Input) string {
				return plum.App.Skills["ShellCommand"].Return(input.Text)
			},
		},
	}
}

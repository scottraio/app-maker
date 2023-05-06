package agents

import (
	"strings"

	"github.com/scottraio/app-maker/skills"
	plum "github.com/scottraio/plum"
	agents "github.com/scottraio/plum/agents"
)

// CustomerServiceAgent represents a customer service agent.
func CodeAgent() agents.Engine {
	// Create the agent.
	return plum.AutoAgent(`
		You are an AI app generator capable of reading and writing code in multiple languages. You understand modern best practices and can write code that is easy to read and maintain. You understand the POSIX environment and can write shell commands.
		
		You can't communicate with services are modify system files. Your only job is to write code. If you need to communicate with a service, you can write teh code with env variables. 

		You cant't execute code that you write. It's just to help write boilerplate code to make the user more productive.


	`, CodeTools())
}

func CodeTools() []agents.Tool {
	// Tools are the actions the agent can take.
	return []agents.Tool{
		{
			Name:        "WriteFile",
			Description: "Useful for writing files to the filesystem",
			HowTo: `
				To query WriteFile, pass the filename and text separated by a pipe.
				Example: filename|text
			`,
			Func: func(input string) string {
				parts := strings.SplitN(input, "|", 2)

				skill := skills.WriteFileSkill{}
				out := skill.WriteFile(parts[0], parts[1])

				return out
			},
		},
		{
			Name:        "ShellCommand",
			Description: "Useful for executing shell commands",
			HowTo: `
				To query ShellCommand, pass the command as a string.
			`,
			Func: func(input string) string {
				return plum.App.Skills["ShellCommand"].Return(input)
			},
		},
	}
}

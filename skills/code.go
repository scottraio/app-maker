package skills

import (
	"github.com/scottraio/plum/agents"
	skills "github.com/scottraio/plum/skills"
)

type CodeSkill struct {
	skills.Skill
	skills.Shell
}

func Code() *skills.Skill {
	// create the model
	code := &CodeSkill{
		// Model is the base model that you want to use
		Skill: skills.Skill{
			Return: func(input string) string {
				// DEPRECATED:
				return ""
			},
		},
	}

	return &code.Skill
}

func (skill *CodeSkill) Write(input agents.Input) string {
	code := input.Action.Branch(input, `
	 You are an AI code generator capable of reading and writing code in multiple languages. You understand modern best practices and can write code that is easy to read and maintain. You understand the POSIX environment and can write shell commands.

	 You respond with the code (string) to be written to a file. Do not respond in markdown or any other format. Just return the code as a string.
	`)

	writeFile := WriteFileSkill{}
	return writeFile.WriteFile(input.Text, code)
}

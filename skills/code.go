package skills

import (
	"github.com/scottraio/plum"
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

func (skill *CodeSkill) Write(codeToWrite string) string {
	prompt := `
		You are a senior software engineer, you can read and write code in multiple languages. You understand modern software stacks and patterns. 

		You are given a task, write the code necessary to complete the task. Output the code ONLY

		Task: 
	`

	return plum.App.LLM.Run(prompt + codeToWrite)
}

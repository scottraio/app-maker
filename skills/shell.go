package skills

import (
	"fmt"
	"os/exec"

	skills "github.com/scottraio/plum/skills"
)

type ShellCommandSkill struct {
	skills.Skill
	skills.Shell
}

func ShellCommand() *skills.Skill {
	// create the model
	shell := &ShellCommandSkill{
		// Model is the base model that you want to use
		Skill: skills.Skill{
			Return: func(query string) string {
				cmd := exec.Command("sh", "-c", query)
				out, err := cmd.CombinedOutput()

				if err != nil {
					fmt.Println(err)
					return ""
				}
				return string(out)
			},
		},
	}

	return &shell.Skill
}

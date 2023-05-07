package skills

import (
	"os"

	skills "github.com/scottraio/plum/skills"
)

type MkDirSkill struct {
	skills.Skill
	skills.Shell
}

func MkDir() *skills.Skill {
	// create the model
	mkdir := &MkDirSkill{
		// Model is the base model that you want to use
		Skill: skills.Skill{
			Return: func(input string) string {
				// DEPRECATED:
				return ""
			},
		},
	}

	return &mkdir.Skill
}

func (skill *MkDirSkill) MkDir(dirName string) string {
	err := os.Mkdir(dirName, 0755)
	if err != nil {
		return "Failed to create directory."
	}

	return "Directory created successfully."
}

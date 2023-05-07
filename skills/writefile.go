package skills

import (
	"log"
	"os"
	"path/filepath"

	skills "github.com/scottraio/plum/skills"
)

type WriteFileSkill struct {
	skills.Skill
	skills.Shell
}

type WriteFileInput struct {
	Filename string
	Text     string
}

func WriteFile() *skills.Skill {
	// create the model
	shell := &ShellCommandSkill{
		// Model is the base model that you want to use
		Skill: skills.Skill{
			Return: func(input string) string {
				// DEPRECATED:
				return ""
			},
		},
	}

	return &shell.Skill
}

func (skill *WriteFileSkill) WriteFile(filename string, text string) string {
	fullFilename := filepath.Join(filename)
	directory := filepath.Dir(fullFilename)
	os.MkdirAll(directory, os.ModePerm)

	file, err := os.Create(fullFilename)
	if err != nil {
		log.Fatalf("Failed to create file: %s", err)
	}

	_, err = file.WriteString(text)
	if err != nil {
		log.Fatalf("Failed to write to file: %s", err)
	}

	file.Close()

	return "File written successfully."
}

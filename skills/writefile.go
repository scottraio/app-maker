package skills

import (
	"log"
	"os"
	"path/filepath"
)

func WriteFile(filename string, text string) string {
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

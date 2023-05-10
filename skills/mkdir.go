package skills

import (
	"os"
)

func MkDir(dirName string) string {
	err := os.Mkdir(dirName, 0755)
	if err != nil {
		return "Failed to create directory."
	}

	return "Directory created successfully."
}

package skills

import (
	"fmt"
	"os/exec"
)

func RunShellCommand(query string) string {
	cmd := exec.Command("sh", "-c", query)
	out, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(out)
}

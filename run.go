package sh

import (
	"fmt"
	"os/exec"
)

func Run(cmd *exec.Cmd) (string, *exec.ExitError) {
	fmt.Printf("[INFO] running %s...\n", cmd.Args)
	stdout, err := cmd.Output()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			return string(stdout), exitErr
		}
		fmt.Println("[ERROR] command failed without exit error")
		fmt.Println("[ERROR]", err)
		return string(stdout), nil
	}
	return string(stdout), nil
}

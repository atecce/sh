package sh

import (
	"log"
	"os/exec"
	"strings"

	"github.com/kr/pretty"
)

func Run(cmd *exec.Cmd) (string, *exec.ExitError) {
	pretty.Logf("[INFO] running %s...", cmd.Args)
	stdout, err := cmd.Output()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			pretty.Logln("[ERROR] stderr:", strings.Trim(string(exitErr.Stderr), "\n"))
			return string(stdout), exitErr
		}
		pretty.Logln("[ERROR] command failed without exit error. quitting")
		log.Fatal(err)
	}
	return string(stdout), nil
}

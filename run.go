package sh

import (
	"os/exec"

	"github.com/kr/pretty"
)

func Run(cmd *exec.Cmd) (string, *exec.ExitError) {

	pretty.Logf("[INFO] running %s...", cmd.Args)
	stdout, err := cmd.Output()
	pretty.Log("[DEBUG] stdout:", string(stdout))
	if err != nil {
		pretty.Logln("[ERROR] command", cmd, "failed to run")
		pretty.Logln("[INFO] attempting to return exit error, could panic here (TODO?)...")
		return string(stdout), err.(*exec.ExitError)
	}
	return string(stdout), nil
}

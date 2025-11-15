package service

import (
	"os/exec"
)

func GetDiff () (error, string) {
	cmd := exec.Command("git", "diff")

	out, err := cmd.CombinedOutput()
	if err != nil {
		return err, ""
	}

	diff := string(out)

	return nil, diff
}

func GitAndCommit(message string) (string, error) {
	addCmd := exec.Command("git", "add", ".")
	if out, err := addCmd.CombinedOutput(); err != nil {
		return string(out), err
	}

	commitCmd := exec.Command("git", "commit", "-m", message)
	out, err := commitCmd.CombinedOutput()
	if err != nil {
		return string(out), err
	}

	return string(out), nil
}

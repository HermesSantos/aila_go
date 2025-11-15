package service

import (
	"fmt"
	"os/exec"
)

func GetDiff () (error, string) {
	cmd := exec.Command("git", "diff")

	out, err := cmd.CombinedOutput()
	if err != nil {
		return err, ""
	}

	diff := string(out)
	fmt.Println("diff", diff)
	return nil, diff
}

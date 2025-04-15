package so

import (
	"fmt"
	"os/exec"
)

func GetDiff () (DiffMessage string, error error) {
	cmd := exec.Command("git", "diff")
	cmd.Dir = "./"
	output, err := cmd.Output()


	if err != nil {
		fmt.Println("An error occurred", err)
		return "", err
	}


	return string(output), nil
}

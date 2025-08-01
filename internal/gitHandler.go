package internal

import (
	"fmt"
	"os/exec"
)

func GetGitDiff () (error, string) {
	cmd, err := exec.Command("git", "diff").Output()
	if err != nil {
		fmt.Println("Erro ao executar comando:", err)
		return err, ""
	}
	return nil, string(cmd)
}

package main

import (
	"aila_go/internal"
	"aila_go/services"
	"fmt"
	"log"
	"os/exec"

	"github.com/charmbracelet/huh"
)

var (
	apiResponseMessage       string
)

func main () {
	err, gitDiff := internal.GetGitDiff()
	if err != nil {
    log.Fatal(err)
  }

	cfg := services.Load()
	message := cfg.GetGeminiCommitMessage(gitDiff)
	Huh(message)
}

func Huh (commitMessage string) {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Mensagem de Commit: \n" + commitMessage).
				Options(
				huh.NewOption("Usar Mensagem de Commit", "use"),
				huh.NewOption("Gerar Outra", "regenerate"),
				huh.NewOption("Configurações", "configs"),
				huh.NewOption("Sair", "exit"),
				).
				Value(&apiResponseMessage),
			),
		)

	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}
	
	switch apiResponseMessage {
	case "use":
		fmt.Println("Mensagem de commit usada: " + commitMessage)
		exec.Command("git", "add", ".").Run()
		exec.Command("git", "commit", "-m", commitMessage).Run()
		return
	case "exit":
		break
	}

}

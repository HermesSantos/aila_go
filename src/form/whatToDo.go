package form

import (
	"aila/src/service"
	"os"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

func WhatToDoForm () {
	style := lipgloss.NewStyle().
	Foreground(lipgloss.Color("#ff009400"))
	coloredMsg := style.Render(commitMessage)

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[int]().
			Title("Mensagem de commit:\n" + coloredMsg + "\nGostaria de continuar com a mensagem de commit?").
			Options(
				huh.NewOption("Usar essa", 0),
				huh.NewOption("Gerar outra", 1),
				huh.NewOption("Editar Mensagem", 2),
				huh.NewOption("Configurações", 3),
				huh.NewOption("Sair", 4),
			).
			Value(&WhatToDoSelected),
		),
	)
	form.Run()

	switch WhatToDoSelected {
	case 0:
		service.GitAndCommit(commitMessage)
	case 1:
		InitForm()
	case 2:
		editCommitMessageForm()
		service.GitAndCommit(commitMessage)
	case 3:
		configureForm()
	case 4:
		os.Exit(0)
	}
}

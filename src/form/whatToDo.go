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
				huh.NewOption("Sair", 3),
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
		InitForm()
	case 3:
		os.Exit(0)
	}
}

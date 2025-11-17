package form

import (
	"github.com/charmbracelet/huh"
)

func configureForm () {
	var option string
	form := huh.NewSelect[string]().
	Title("Configurações").
	Options(
		huh.NewOption("Idioma do commit", "language"), // option
		huh.NewOption("Commits padrão git flow", "flow"), // confirm
		huh.NewOption("Mudar chave de API", "apiToken"), // input
		huh.NewOption("Voltar", "back"),
	).
	Value(&option)

	switch option {
	case "language":
		changeLanguage()
	}

	form.Run()
}

func changeLanguage () {
	var selectedLanguage string

	form := huh.NewSelect[string]().
		Title("Selecione o Idioma").
		Options(
			huh.NewOption("English", "EN"),
			huh.NewOption("Português", "PT"),
			huh.NewOption("Espanhol", "ES"),
		).
		Value(&selectedLanguage)

	form.Run()
}

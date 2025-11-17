package form

import "github.com/charmbracelet/huh"

func editCommitMessageForm () {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewText().
			Title("Edição de commit").
			Value(&commitMessage),
		),
	)

	form.Run()
}

package form

import (
	"aila/src/repository"
	"aila/src/service"
	"database/sql"
	"fmt"
	"os"

	"github.com/charmbracelet/huh"
)

var (
	apiKey string
	commitMessage string
	WhatToDoSelected int
)

func InitForm () {
	err, diff := service.GetDiff()
	if err != nil {
		fmt.Println(err)
		return
  }

	err, geminiResponse := service.GetGeminiCommitService(diff)
	if err != nil {
    fmt.Println(err)
    return
  }

  commitMessage = geminiResponse

	WhatToDoForm()
}

func WhatToDoForm () {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[int]().
			Title("Mensagem de commit: " + commitMessage + "\nGostaria de continuar com a mensagem de commit?").
			Options(
				huh.NewOption("Usar essa", 0),
				huh.NewOption("Gerar outra", 1),
				huh.NewOption("Sair", 2),
			).
			Value(&WhatToDoSelected),
		),
	)
	form.Run()

	switch WhatToDoSelected {
	case 0:
		service.GitAndCommit(commitMessage)
	case 1:
		// TODO: Implementar a geracao de outra mensagem de commit
	case 2:
		os.Exit(0)
	}
}

func AddApiKeyForm (db *sql.DB) {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
			Title("Parece que você não informou uma chave de API. Por favor, informe uma chave de API").
			Value(&apiKey). 
			Validate(func(str string) error {
				if str == "" {
					fmt.Println("Chave de API inválida")
					AddApiKeyForm(db)
				}
				return nil
			}),
		),
	)

	form.Run()
	repository.InsertApiKey(apiKey)
}


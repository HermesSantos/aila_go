package form

import (
	"aila/src/repository"
	"database/sql"
	"fmt"

	"github.com/charmbracelet/huh"
)

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

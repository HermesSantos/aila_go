package form

import (
	"aila/src/repository"
	"aila/src/service"
	"database/sql"
	"fmt"

	"github.com/charmbracelet/huh"
)

var (
	apiKey string
)

func InitForm () {
	err, diff := service.GetDiff()
	if err != nil {
		fmt.Println(err)
		return
  }
	service.GetGeminiCommitService(diff)
}

func AddApiKey (db *sql.DB) {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
			Title("Parece que você não informou uma chave de API. Por favor, informe uma chave de API").
			Value(&apiKey). 
			Validate(func(str string) error {
				if str == "" {
					fmt.Println("Chave de API inválida")
					AddApiKey(db)
				}
				return nil
			}),
		),
	)

	form.Run()
	repository.InsertApiKey(apiKey)
}


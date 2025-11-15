package main

import (
	"aila/src/config"
	"aila/src/form"
	"aila/src/repository"
)

func main () {
	config.InitDatabase()

	repository.InitDatabaseRepository(config.DB)

	hasApiKey := repository.VerifyApiToken()

	if !hasApiKey {
		form.AddApiKeyForm(config.DB)
	} else {
		form.InitForm()
	}


	config.CloseDb()
}

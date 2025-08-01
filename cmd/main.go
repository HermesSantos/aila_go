package main

import (
	"aila_go/internal"
	"aila_go/services"
	"fmt"
	"log"

	"github.com/charmbracelet/huh"
)

var (
	burger       string
	discount     bool
)

func main () {
	err, gitDiff := internal.GetGitDiff()
	if err != nil {
    log.Fatal(err)
  }

	cfg := services.Load()
	cfg.GetGeminiCommitMessage(gitDiff)
}

func Huh () {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Choose your burger").
				Options(
				huh.NewOption("Charmburger Classic", "classic"),
				huh.NewOption("Chickwich", "chickwich"),
				huh.NewOption("Fishburger", "fishburger"),
				huh.NewOption("Charmpossible™ Burger", "charmpossible"),
				).
				Value(&burger),
			),
		)
	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}

	if !discount {
		fmt.Println("What? You didn’t take the discount?!")
	}
}

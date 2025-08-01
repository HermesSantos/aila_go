package main

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/charmbracelet/huh"
)

var (
	burger       string
	discount     bool
)

func main () {
	cmd, err := exec.Command("ls", "-la").Output()
	if err != nil {
		fmt.Println("Erro ao executar comando:", err)
		return
	}
	fmt.Println(string(cmd))
}

func Huh () {
	form := huh.NewForm(
		huh.NewGroup(
			// Ask the user for a base burger and toppings.
			huh.NewSelect[string]().
				Title("Choose your burger").
				Options(
				huh.NewOption("Charmburger Classic", "classic"),
				huh.NewOption("Chickwich", "chickwich"),
				huh.NewOption("Fishburger", "fishburger"),
				huh.NewOption("Charmpossible™ Burger", "charmpossible"),
				).
				Value(&burger), // store the chosen option in the "burger" variable
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

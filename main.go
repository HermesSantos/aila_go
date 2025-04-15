package main

import (
	"fmt"
	"github.com/charmbracelet/huh"
	"aila_go/SO"
	"aila_go/GeminiHandler"
)

var burger string

func main () {
	diffMessage, err := so.GetDiff()
	if err != nil {
    fmt.Println("Error:", err)
    return
  }

	geminiResponse, err := geminihandler.GetGeminiCommitMessage(diffMessage)
	if err != nil {
    fmt.Println("Error:", err)
    return
  }

	fmt.Println("Commit Message: \n", geminiResponse)

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("What do you want to do next?").
				Options(
					huh.NewOption("Keep message and Commit", "Keep"),
					huh.NewOption("Generate another message", "Generate"),
				).
				Value(&burger), 
			),
		)
	errRun := form.Run()
	if errRun != nil {
		fmt.Println("Error:", err)
	}
}


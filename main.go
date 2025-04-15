package main

import (
	"aila_go/GeminiHandler"
	"aila_go/SO"
	"fmt"
	"os/exec"

	"github.com/charmbracelet/huh"
)

var userResponse string

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
				Value(&userResponse), 
			),
		)
	errRun := form.Run()
	if errRun != nil {
		fmt.Println("Error:", err)
	}
	switch userResponse {
	case "Keep":
		cmd := exec.Command("git", "commit", "-m", geminiResponse)
		cmd.Dir = "./"
		fmt.Println(cmd)
		_, err := cmd.Output()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	// TODO: Add case to generate message
	case "Generate":
		geminiResponse, err = geminihandler.GetGeminiCommitMessage(diffMessage)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("Commit Message: \n", geminiResponse)
	}
}


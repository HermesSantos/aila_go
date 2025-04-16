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

	// fmt.Println("Commit Message: \n", geminiResponse)


	form := huh.NewForm(
		huh.NewGroup(
			huh.NewNote().
				Title("Commit message: " + geminiResponse),
			huh.NewSelect[string]().
				Title("Choose your burger").
				Options(
				huh.NewOption("Use this message and commit", "Keep"),
				huh.NewOption("Generate another message", "Generate"),
				huh.NewOption("Exit", "Generate"),
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
		cmd := exec.Command("git", "add", ".", "&&", "git", "commit", "-m", geminiResponse)
		cmd.Dir = "./"
		_, err := cmd.Output()
		if err != nil {
			fmt.Println("Error running git commit command:", err)
			return
		}
	// TODO: Add case to regenerate the message
	case "Generate":
		geminiResponse, err = geminihandler.GetGeminiCommitMessage(diffMessage)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("Commit Message: \n", geminiResponse)
	}
}


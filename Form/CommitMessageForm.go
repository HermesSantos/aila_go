package commitform

import (
	"aila_go/GeminiHandler"
	"fmt"
	"github.com/charmbracelet/huh"
)

var userResponse string

type Response struct {
  UserResponse string
	GeminiResponse string
}

func CommitMessageForm (diffMessage string) (Response, error) {
	geminiResponse, err := geminihandler.GetGeminiCommitMessage(diffMessage)
	if err != nil {
		return Response{}, err
	}

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

	response := Response{
		UserResponse: userResponse,
		GeminiResponse: geminiResponse,
	}

	return response, nil
}

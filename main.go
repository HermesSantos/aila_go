package main

import (
	"aila_go/SO"
	"aila_go/Form"
	"fmt"
	"os/exec"
)


func main () {
	diffMessage, err := so.GetDiff()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	response, err := commitform.CommitMessageForm(diffMessage)

	switch response.UserResponse {
	case "Keep":
		cmd := exec.Command("git", "add", ".", "&&", "git", "commit", "-m", response.GeminiResponse)
		cmd.Dir = "./"
		_, err := cmd.Output()
		if err != nil {
			fmt.Println("Error running git commit command:", err)
			return
		}
	// TODO: Add case to regenerate the message
	case "Generate":
		// geminiResponse, err = geminihandler.GetGeminiCommitMessage(diffMessage)
		// if err != nil {
		// 	fmt.Println("Error:", err)
		// 	return
		// }
		// fmt.Println("Commit Message: \n", geminiResponse)
	}
}


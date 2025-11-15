package form

import (
	"aila/src/service"
	"fmt"
)

var (
	apiKey string
	commitMessage string
	WhatToDoSelected int
)

func InitForm () {
	err, diff := service.GetDiff()
	if err != nil {
		fmt.Println(err)
		return
  }

	err, geminiResponse := service.GetGeminiCommitService(diff)
	if err != nil {
    fmt.Println(err)
    return
  }

  commitMessage = geminiResponse

	WhatToDoForm()
}




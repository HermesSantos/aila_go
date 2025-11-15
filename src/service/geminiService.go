package service

import (
	"aila/src/repository"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)
type (
	Part struct {
		Text string `json:"text"`
	}

	Content struct {
		Parts []Part `json:"parts"`
	}

	RequestBody struct {
		Contents []Content `json:"contents"`
	}
)

func GetGeminiCommitService (diff string) {
	apiKey, err := repository.GetApiKey()
	if err != nil {
    fmt.Println(err)
    return
	}

	reqBody := RequestBody{
		Contents: []Content{
			{
				Parts: []Part{
					{
						Text: fmt.Sprintf(
							"Me retorne em %s, sem caracteres especiais como aspas ou quebra de linha, uma mensagem de commit curta que mostre o que foi alterado nesse commit: %s",
							os.Getenv("COMMIT_LANGUAGE"),
							diff,
						),
					},
				},
			},
		},
	}
	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return
	}

	req, err := http.NewRequest(
		"POST",
		"https://generativelanguage.googleapis.com/v1beta/models/gemini-2.0-flash-001:generateContent" + apiKey,
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
    return
  }

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}

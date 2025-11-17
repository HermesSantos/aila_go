package service

import (
	"aila/src/repository"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
	GeminiResponse struct {
		Candidates []Candidate `json:"candidates"`
	}

	Candidate struct {
		Content ContentResponse `json:"content"`
	}

	ContentResponse struct {
		Parts []PartResponse `json:"parts"`
	}

	PartResponse struct {
		Text string `json:"text"`
	}
)


func GetGeminiCommitService (diff string) (error, string) {

	commitLanguage, err := repository.GetCommitLanguage()
	if err != nil {
    fmt.Println("Erro ao pegar linguagem do commit, valor padrão será em inglês", commitLanguage)
		commitLanguage = "English"
  }

	apiKey, err := repository.GetApiKey()
	if err != nil {
		fmt.Println(err)
		return nil, ""
	}

	reqBody := RequestBody{
		Contents: []Content{
			{
				Parts: []Part{
					{
						Text: fmt.Sprintf(
							"Me retorne em %s, sem caracteres especiais como aspas ou quebra de linha, uma mensagem de commit curta que mostre o que foi alterado nesse commit: %s",
							commitLanguage,
							diff,
						),
					},
				},
			},
		},
	}
	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		fmt.Println(err)
		return nil, ""
	}

	req, err := http.NewRequest(
		"POST",
		"https://generativelanguage.googleapis.com/v1beta/models/gemini-2.0-flash-001:generateContent?key=" + apiKey,
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		fmt.Println(err)
		return nil, ""
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, ""
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var gr GeminiResponse
	if err := json.Unmarshal(body, &gr); err != nil {
		fmt.Println("Erro ao decodificar:", err)
		fmt.Println("Body:", string(body))
		return nil, ""
	}

	if len(gr.Candidates) == 0 ||
	len(gr.Candidates[0].Content.Parts) == 0 {
		fmt.Println("Resposta vazia:", string(body))
		return err, ""
	}

	msg := gr.Candidates[0].Content.Parts[0].Text
	return nil, msg
}

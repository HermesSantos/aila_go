package geminihandler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	// "net/http"
)

var response Response

type (
	Payload struct {
		Contents []Content `json:"contents"`
	}

	Response struct {
		Candidates []Candidate `json:"candidates"`
	}

	Candidate struct {
		Content Content `json:"content"`
	}

	Content struct {
		Parts []Part `json:"parts"`
	}

	Part struct {
		Text string `json:"text"`
	}
)

func GetGeminiCommitMessage (DiffMessage string) (string, error) {
	body := Payload{
		Contents: []Content{
			{
				Parts: []Part{
					{
						Text: fmt.Sprintf("Me retorne em inglês, sem caracteres especiais como aspas ou quebra de linha, uma mensagem de commit curta que mostre o que foi alterado nesse commit: %s", DiffMessage),
					},
				},
			},
		},
	}

	bodyJSON, err := json.Marshal(body)
	if err != nil {
		return "", fmt.Errorf("erro ao converter para JSON: %v", err)
	}

	url := "https://generativelanguage.googleapis.com/v1beta/models/gemini-2.0-flash-001:generateContent?key=AIzaSyBvfum7gSciXAcmbm2DhH5nZSMQxKIC6YI"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyJSON))
	if err != nil {
		fmt.Printf("Erro ao criar requisição: %v\n", err)
		return "", nil
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Erro ao enviar requisição: %v\n", err)
		return "", nil
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Erro ao ler resposta: %v\n", err)
		return "", nil
	}

	json.Unmarshal(responseBody, &response)

	return response.Candidates[0].Content.Parts[0].Text, nil
}

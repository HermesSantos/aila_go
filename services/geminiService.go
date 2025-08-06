package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Response struct {
    Candidates []struct {
        Content struct {
            Parts []struct {
                Text string `json:"text"`
            } `json:"parts"`
        } `json:"content"`
    } `json:"candidates"`
}

type Data struct {
	Candidates []Candidate `json:"candidates"`
}

type Candidate struct {
	Content Content `json:"content"`
}

type Part struct {
	Text string `json:"text"`
}

type Content struct {
	Parts []Part `json:"parts"`
}

type RequestBody struct {
	Contents []Content `json:"contents"`
}

type Config struct {
	ApiUrl string
	ApiKey string
}

var (
	Cfg  *Config
	once sync.Once
)

func Load () *Config {
	err := godotenv.Load()
	if err != nil {
		log.Panic("Error loading .env file")
	}
	once.Do(func () {
		Cfg = &Config{
			ApiUrl: os.Getenv("GEMINI_API_URL") + "?key=" + os.Getenv("GEMINI_API_KEY"),
			ApiKey: os.Getenv("GEMINI_API_KEY"),
		}
	})
	return Cfg
}

func (gc Config) GetGeminiCommitMessage (gitDiff string) string {
	err := godotenv.Load()
	if err != nil {
		log.Panic("Error loading .env file")
	}

	prompt := fmt.Sprintf("Me retorne em %s, sem caracteres especiais como aspas ou quebra de linha, uma mensagem de commit curta que mostre o que foi alterado nesse commit: %s", os.Getenv("COMMIT_MESSAGE"), gitDiff)

	body := RequestBody{
		Contents: []Content{
			{
				Parts: []Part{
					{Text: prompt},
				},
			},
		},
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}

	resp, err := http.Post(gc.ApiUrl, "application/json", bytes.NewReader(jsonBody))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	api_body_response, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var apiResp Response
	if err := json.Unmarshal(api_body_response, &apiResp); err != nil {
		panic(err)
	}

	if len(apiResp.Candidates) > 0 &&
	len(apiResp.Candidates[0].Content.Parts) > 0 {
		commitMessage := apiResp.Candidates[0].Content.Parts[0].Text
		return commitMessage
	} else {
		return "Resposta inesperada"
	}
}

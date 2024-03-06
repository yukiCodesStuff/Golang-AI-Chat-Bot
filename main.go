package main

import (
	"log"
	"io"
	"net/http"
	"bytes"
	"os"
	"github.com/joho/godotenv"
 )

 func main() {

	posturl := "https://api.openai.com/v1/chat/completions"

	body := []byte(`{
		"model": "gpt-3.5-turbo",
		"messages": [{"role": "user", "content": "What is 10 + 20"}],
		"temperature": 0.7
	  }`)

	err := godotenv.Load(".env");
	if err != nil {
		log.Fatalf("ERROR::MAIN::Could not load dotenv")
	}

	apikey := os.Getenv("API_KEY")

	r, err := http.NewRequest("POST", posturl, bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}

	auth := "Bearer " + apikey
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Authorization", auth)

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		log.Fatalf("HTTP request failed: %v", err)
	}
	responseBytes, err := io.ReadAll(res.Body)
    if err != nil {
        log.Fatalf("error reading HTTP response body: %v", err)
    }

    log.Println("We got the response:", string(responseBytes))
 }

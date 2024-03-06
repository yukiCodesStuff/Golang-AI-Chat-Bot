package main

import (
	"bufio"
	"log"
	"io"
	"net/http"
	"bytes"
	"os"
	"github.com/joho/godotenv"
	"fmt"
	"encoding/json"
 )

 func main() {

	fmt.Println("Ask something: ")
	reader := bufio.NewReader(os.Stdin)
	inquiry, _ := reader.ReadString('\n') // Reads input until newline

	// Trim the newline character from inquiry, depending on your use case
	inquiry = inquiry[:len(inquiry)-1] // This line might need adjustment for Windows (\r\n)

	payload := map[string]interface{}{
		"model": "gpt-3.5-turbo",
		"messages": []map[string]string{
			{"role": "user", "content": inquiry},
		},
		"temperature": 0.7,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		// Handle error
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	posturl := "https://api.openai.com/v1/chat/completions"
	// body := []byte(bodyStr)

	// body := []byte(`{
	// "model": "gpt-3.5-turbo",
	// "messages": [{"role": "user", "content": "What is 9 + 10"}],
	// "temperature": 0.7
	// }`)


	derr := godotenv.Load(".env");
	if derr != nil {
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

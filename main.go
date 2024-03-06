package main

import (
	// "io/ioutil"
	"log"
	// "net/http"
	"fmt"
	"os"
	"github.com/joho/godotenv"
 )

 func main() {
	err := godotenv.Load(".env");
	if err != nil {
		log.Fatalf("ERROR::MAIN::Could not load dotenv")
	}

	// apikey := os.Getenv("API_KEY")
	fmt.Println(apikey)
 }

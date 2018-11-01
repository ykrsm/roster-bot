package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	arg := os.Args[1]
	var hookURL string

	switch arg {
	case "-p":
		fmt.Printf("PRODUCTION\n")
		hookURL = os.Getenv("PROD_WEBHOOK_URL")

	case "-d1":
		fmt.Printf("DEVELOPMENT1\n")
		hookURL = os.Getenv("DEV1_WEBHOOK_URL")

	case "-d2":
		fmt.Printf("DEVELOPMENT2\n")
		hookURL = os.Getenv("DEV2_WEBHOOK_URL")

	case "-t":
		fmt.Printf("TEST\n")
		hookURL = os.Getenv("TEST_WEBHOOK_URL")

	default:
		os.Exit(1)
	}

	t := time.Now()
	fmt.Printf("Current time:\t%v\n", t)
	_, month, day := t.Date()

	res := makeRoster(int(month), day)

	text := "おはよう！日直バザールだよ!\n\n" + res
	postMessage(text, hookURL)
}

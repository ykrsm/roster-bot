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

	hookURL := os.Getenv("WEBHOOK_URL")

	t := time.Now()
	fmt.Printf("Current time:\t%v\n", t)

	_, month, day := t.Date()

	res := makeRoster(int(month), day)

	text := "おはよう！日直バザールだよ!\n\n" + res
	postMessage(text, hookURL)
}

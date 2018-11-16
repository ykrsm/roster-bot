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

	var fileName string
	if len(os.Args) > 2 {
		fileName = os.Args[2]
	} else {
		fileName = "./data.xlsx"
	}

	fmt.Printf("File name: %s", fileName)

	t := time.Now()
	fmt.Printf("Current time:\t%v\n", t)
	_, month, day := t.Date()

	res := makeRoster(int(month), day, fileName)

	// Making date string in Japnaese
	wdays := [...]string{"日", "月", "火", "水", "木", "金", "土"}
	weekDayJP := wdays[t.Weekday()]
	dateJP := t.Format("1月2日 (" + weekDayJP + ")")

	text := dateJP + " の勤務表でござ~る\n\n" + res
	postMessage(text, hookURL)
}

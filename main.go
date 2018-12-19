package main

import (
	"bufio"
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

	// os.Args[1] specify if its prod or dev
	// os.Args[2] is excel file name
	// os.Args[3] specify txt file contains names

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
	var namelistFileName string
	if len(os.Args) == 4 {
		fileName = os.Args[2]
		namelistFileName = os.Args[3]
	} else {
		log.Fatal("Not enough argument\n")
	}

	t := time.Now()
	fmt.Printf("Current time:\t%v\n", t)
	year, month, day := t.Date()

	// parse txt file and make an array of names
	names := parseNameFile(namelistFileName)

	fmt.Printf("File name: %s\n", fileName)
	fmt.Printf("Names File: %v\n", namelistFileName)
	fmt.Printf("Names: %v\n", names)

	roster := Roster{Date: t}
	// res := makeRoster(int(month), day, fileName)
	res := runRoster(year, int(month), day, fileName, roster, names)

	// Making date string in Japnaese
	wdays := [...]string{"日", "月", "火", "水", "木", "金", "土"}
	weekDayJP := wdays[t.Weekday()]
	dateJP := t.Format("1月2日 (" + weekDayJP + ")")

	text := dateJP + " の勤務表でござ~る\n\n" + res
	postMessage(text, hookURL)
}

func runRoster(year, month, day int, fileName string, roster Roster, names []string) string {
	r := fillRoster(year, month, day, fileName, roster, names)

	return r.String()
}

func parseNameFile(fileName string) (names []string) {

	// Open the file.
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Cannot open " + fileName)
	}
	// Create a new Scanner for the file.
	scanner := bufio.NewScanner(f)
	// Loop over all lines in the file and print them.
	for scanner.Scan() {
		line := scanner.Text()
		names = append(names, line)
	}

	f.Close()
	return names
}

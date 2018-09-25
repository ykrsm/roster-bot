package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	os.Exit(_main(os.Args[1:]))
}

func _main(args []string) int {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Listening slack event and response
	fmt.Printf(os.Getenv("WEBHOOK_URL"))

	return 0
}

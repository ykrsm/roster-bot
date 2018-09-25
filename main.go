package main

import (
	"log"
	"os"

	"github.com/bluele/slack"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	hookURL := os.Getenv("WEBHOOK_URL")
	hook := slack.NewWebHook(hookURL)

	err2 := hook.PostMessage(&slack.WebHookPostPayload{

		Text: "hello!",
		Attachments: []*slack.Attachment{
			{Text: "danger", Color: "danger"},
		},
	})
	if err2 != nil {
		panic(err)
	}
}

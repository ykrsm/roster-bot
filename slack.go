package main

import (
	"github.com/bluele/slack"
)

func postMessage(text string, webHookURL string) {

	hook := slack.NewWebHook(webHookURL)

	err := hook.PostMessage(&slack.WebHookPostPayload{
		Text: text,
	})
	if err != nil {
		panic(err)
	}
}

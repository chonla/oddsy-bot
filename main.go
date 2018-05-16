package main

import (
	"os"

	"github.com/chonla/oddsy"
	"github.com/kr/pretty"
)

func main() {
	oddsy := oddsy.NewOddsy("./oddsy.json")

	slackToken := os.Getenv("SLACK_TOKEN")
	if slackToken != "" {
		oddsy.SetToken(slackToken)
	}

	oddsy.MessageReceived(messageHandler)
	oddsy.DirectMessageReceived(directMessageHandler)

	oddsy.Start()
}

func messageHandler(o *oddsy.Oddsy, m *oddsy.Message) {
	pretty.Println("Message:")
	pretty.Println(m)
}

func directMessageHandler(o *oddsy.Oddsy, m *oddsy.Message) {
	pretty.Println("Direct Message:")
	pretty.Println(m)
}

package main

import (
	"os"
	"strings"

	"github.com/chonla/oddsy"
	"github.com/chonla/oddsy-bot/translator"
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
	if m.Mentioned {
		o.Send(m.Channel.UID, "<@"+m.From.UID+"> เรียกเค้าทำไมจ๊ะ คิดถึงล่ะสิ")
	}
}

func directMessageHandler(o *oddsy.Oddsy, m *oddsy.Message) {
	// pretty.Println("Direct Message:")
	// pretty.Println(m)
	t := strings.ToLower(firstToken(m.Message))
	switch {
	case t == "แปล" || t == "translate":
		t := translator.NewTranslator()
		r, _ := t.Translate(nextToken(m.Message))
		o.Send(m.Channel.UID, r)
	default:
		o.Send(m.Channel.UID, "มีอะไรให้ช่วยไหมจ๊ะ")
	}
}

func firstToken(t string) (r string) {
	l := strings.SplitN(t, " ", 2)
	r = l[0]
	return
}

func nextToken(t string) (r string) {
	l := strings.SplitN(t, " ", 2)
	if len(l) > 0 {
		r = l[1]
	}
	return
}

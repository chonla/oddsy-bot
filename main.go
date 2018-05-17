package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/chonla/oddsy"
	"github.com/chonla/oddsy-bot/translator"
)

var t *translator.Translator

type configuration struct {
	SlackToken       string `json:"slack-token"`
	Debug            bool   `json:"debug"`
	IgnoreBotMessage bool   `json:"ignore-bot-message"`
	GcpToken         string `json:"gcp-token"`
}

var conf configuration

func main() {
	loadConfig("./oddsy.json", &conf)

	b := oddsy.NewOddsy(&oddsy.Configuration{
		SlackToken:       conf.SlackToken,
		Debug:            conf.Debug,
		IgnoreBotMessage: conf.IgnoreBotMessage,
	})

	b.MessageReceived(messageHandler)
	b.DirectMessageReceived(directMessageHandler)
	b.FirstStringTokenReceived("help", helpMessageHandler)
	b.FirstStringTokenReceived("ping", pingMessageHandler)
	b.FirstStringTokenReceived("translate", translateMessageHandler)
	b.FirstStringTokenReceived("แปล", translateMessageHandler)

	b.Start()
}

func messageHandler(o *oddsy.Oddsy, m *oddsy.Message) {
	if m.Mentioned {
		o.Send(m.Channel.UID, "<@"+m.From.UID+"> เรียกเค้าทำไมจ๊ะ คิดถึงล่ะสิ")
	}
}

func directMessageHandler(o *oddsy.Oddsy, m *oddsy.Message) {
	o.Send(m.Channel.UID, "มีอะไรให้ช่วยไหมจ๊ะ")
}

func pingMessageHandler(o *oddsy.Oddsy, m *oddsy.Message) {
	o.Send(m.Channel.UID, "pong :heart:")
}

func helpMessageHandler(o *oddsy.Oddsy, m *oddsy.Message) {
	o.Send(m.Channel.UID, "ข้อความที่"+o.Name+" เข้าใจนะจ๊ะ\n```"+`
ping - ทดสอบ ping/pong
help - ข้อความนี้แหละจ้ะ
translate <ข้อความ> - แปลข้อความจากภาษาอื่นเป็นภาษาไทย
แปล <ข้อความ> - เหมือน translate
tik - ลง worksheet`+"```")
}

func translateMessageHandler(o *oddsy.Oddsy, m *oddsy.Message) {
	if t == nil {
		t = translator.NewTranslator(&translator.Configuration{
			GcpToken: conf.GcpToken,
		})
	}
	r, _ := t.Translate(m.Message)
	o.Send(m.Channel.UID, "แปลว่า\n```"+r+"```")
}

func loadConfig(filename string, conf *configuration) {
	t, e := ioutil.ReadFile(filename)
	if e != nil {
		fmt.Println(e.Error())
		os.Exit(1)
	}

	json.Unmarshal(t, conf)
}

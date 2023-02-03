package main

import (
	"fmt"
	"github.com/CatchZeng/feishu/pkg/feishu"
	"time"
)

func main() {
	token := ""
	secret := ""
	client := feishu.NewClient(token, secret)
	// SendMsg(client)
	// SendCard(client)
	ticker := time.NewTicker(time.Hour)
	for {
		select {
		case <-ticker.C:
			hour := time.Now().Hour()
			if hour > 9 && hour < 21 {
				SendCard(client)
			}
		}
	}
}
func SendCard(client *feishu.Client) {
	cardJson := `{
  "config": {
    "wide_screen_mode": true
  },
  "elements": [
    {
      "tag": "div",
      "text": {
        "content": "**速速站起来喝水，😈**",
        "tag": "lark_md"
      }
    }
  ],
  "header": {
    "template": "purple",
    "title": {
      "content": "🐸⏰该站起来喝水了⏰🐸 %s",
      "tag": "plain_text"
    }
  }
}`
	t := time.Now().Format("2006-01-02 15:04:05")
	cardJson = fmt.Sprintf(cardJson, t)

	msg := feishu.NewInteractiveMessage()
	msg.Card = cardJson

	_, response, r := client.Send(msg)
	if r != nil {
		panic(r)
	}
	fmt.Println(response)

}

func SendMsg(client *feishu.Client) {
	msg := feishu.NewTextMessage()
	msg.Content.Text = "Hello QQW"

	_, response, err := client.Send(msg)
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}

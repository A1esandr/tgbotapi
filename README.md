# tgbotapi
Telegram bot API

### Prerequisites
* Go 1.17

### Example
examples can be found in ```cmd``` folder

```
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/A1esandr/tgbotapi"
)

func main() {
	chatID := os.Getenv("CHAT_ID")    // ID of chat
	token := os.Getenv("TOKEN")       // Token for Telegram bot
	if token == "" {
		log.Fatal("token is empty!")
	}
	bot, err := tgbotapi.New(token)
	if err != nil {
		log.Fatal(err)
	}
	response, err := bot.SendPoll(&tgbotapi.SendPoll{
		ChatID:          chatID,
		Question:        "How are you?",
		Options:         []string{"Fine", "Double Fine"},
		Type:            "quiz",
		CorrectOptionID: 1,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(response))
}
```

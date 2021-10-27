package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/A1esandr/tgbotapi"
)

var tokenFlag = flag.String("token", "", "Bot token")

func main() {
	flag.Parse()
	chatIDStr := os.Getenv("CHAT_ID")
	token := os.Getenv("TOKEN")
	if token == "" && tokenFlag != nil {
		token = *tokenFlag
	}
	if token == "" {
		log.Fatal("token is empty!")
	}
	bot, err := tgbotapi.New(token)
	if err != nil {
		log.Fatal(err)
	}
	chatID, err := strconv.ParseInt(chatIDStr, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := bot.SendMessage(&tgbotapi.SendMessage{ChatID: chatID, Text: "Hi"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(resp))
}

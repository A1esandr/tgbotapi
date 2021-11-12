package main

import (
	"fmt"
	"log"
	"os"

	"github.com/A1esandr/tgbotapi"
)

func main() {
	token := os.Getenv("TOKEN")
	if token == "" {
		log.Fatal("token is empty!")
	}
	bot, err := tgbotapi.New(token)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := bot.GetUpdates(&tgbotapi.GetUpdates{
		Offset:          0,
		Limit:        10,
		Timeout: 1,
	})
	if err != nil {
		log.Fatal(err)
	}
	if resp != nil {
		fmt.Println(resp.Result.Chat.ID)
		fmt.Println(resp.Result.Chat.Title)
	}
}

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
		log.Panic("token is empty!")
	}
	bot := tgbotapi.New(token)
	resp, err := bot.Auth()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
}

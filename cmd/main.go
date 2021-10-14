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
	bot := tgbotapi.New(token)
	data, err := bot.Auth()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}

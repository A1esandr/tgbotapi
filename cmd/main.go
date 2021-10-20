package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/A1esandr/tgbotapi"
)

var tokenFlag = flag.String("token", "", "Bot token")

func main() {
	flag.Parse()
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
	data, err := bot.RawGetRequest("getMe")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}

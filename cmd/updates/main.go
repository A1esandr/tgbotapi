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
	offset := int64(0)
	for {
		resp, err := bot.GetUpdates(&tgbotapi.GetUpdates{
			Offset:  offset,
			Limit:   10,
			Timeout: 1,
		})
		if err != nil {
			log.Fatal(err)
		}
		if resp != nil {
			for _, upd := range resp.Result {
				fmt.Println(upd.UpdateID)
				fmt.Println(upd.ChannelPost.Chat.ID)
				fmt.Println(upd.ChannelPost.Chat.Title)
				if offset <= upd.UpdateID {
					offset = upd.UpdateID + 1
				}
			}
			if len(resp.Result) < 10 {
				break
			}
		} else {
			break
		}
	}
	fmt.Println("finish")
}

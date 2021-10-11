package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/A1esandr/tgbotapi"
)

func main() {
	token := os.Getenv("TOKEN")
	if token == "" {
		log.Fatal("token is empty!")
	}
	bot := tgbotapi.New(token)
	resp, err := bot.Auth()
	if err != nil {
		log.Fatal(err)
	}
	if resp == nil {
		log.Fatal("nil response from auth")
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatal("not OK response from auth")
	}
	defer func() {
		closeErr := resp.Body.Close()
		if closeErr != nil {
			log.Fatalf("error close response body %s", closeErr.Error())
		}
	}()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("error resp response body %s", err.Error())
	}
	fmt.Println(string(data))
}

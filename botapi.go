package tgbotapi

import (
	"fmt"
	"net/http"
)

type (
	bot struct {
		token string
	}
	BotParams struct {
		Token string
	}
	Bot interface {
		Auth() (*http.Response, error)
	}
)

func New(token string) Bot {
	return &bot{token: token}
}

func (b *bot) Auth() (*http.Response, error) {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/getMe", b.token)
	return http.Get(url)
}

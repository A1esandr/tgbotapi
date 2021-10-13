package tgbotapi

import (
	"errors"
	"fmt"
	"io"
	"log"
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
		Auth() ([]byte, error)
	}
)

func New(token string) Bot {
	return &bot{token: token}
}

func (b *bot) Auth() ([]byte, error) {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/getMe", b.token)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return b.read(resp)
}

func (b *bot) read(resp *http.Response) ([]byte, error) {
	if resp == nil {
		return nil, nil
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("not OK response from auth")
	}
	defer func() {
		closeErr := resp.Body.Close()
		if closeErr != nil {
			log.Println("error close response body", closeErr.Error())
		}
	}()
	return io.ReadAll(resp.Body)
}

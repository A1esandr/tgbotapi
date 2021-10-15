package tgbotapi

import (
	"encoding/json"
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
		Auth() (*AuthResponse, error)
	}
	AuthResponse struct {
		OK     bool   `json:"ok"`
		Result Result `json:"result"`
	}
	Result struct {
		ID        int64  `json:"id"`
		IsBot     bool   `json:"is_bot"`
		FirstName string `json:"first_name"`
	}
)

func New(token string) Bot {
	return &bot{token: token}
}

func (b *bot) Auth() (*AuthResponse, error) {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/getMe", b.token)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	data, err := b.read(resp)
	if err != nil {
		return nil, err
	}
	var response AuthResponse
	err = json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
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

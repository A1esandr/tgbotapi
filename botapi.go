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
		GetMe() (*GetMeResponse, error)
		RawRequest(request string) ([]byte, error)
	}
	GetMeResponse struct {
		OK     bool        `json:"ok"`
		Result GetMeResult `json:"result"`
	}
	GetMeResult struct {
		ID                      int64  `json:"id"`
		IsBot                   bool   `json:"is_bot"`
		FirstName               string `json:"first_name"`
		Username                string `json:"username"`
		CanJoinGroups           bool   `json:"can_join_groups"`
		CanReadAllGroupMessages bool   `json:"can_read_all_group_messages"`
		SupportInlineQueries    bool   `json:"support_inline_queries"`
	}
)

func New(token string) (Bot, error) {
	b := &bot{token: token}
	resp, err := b.GetMe()
	if err != nil {
		return nil, err
	}
	if !resp.OK {
		return nil, errors.New("not ok response from telegram api")
	}
	if !resp.Result.IsBot {
		return nil, errors.New("not bot token")
	}
	return b, nil
}

func (b *bot) GetMe() (*GetMeResponse, error) {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/getMe", b.token)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	data, err := b.read(resp)
	if err != nil {
		return nil, err
	}
	var response GetMeResponse
	err = json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (b *bot) RawRequest(request string) ([]byte, error) {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/%s", b.token, request)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	data, err := b.read(resp)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (b *bot) read(resp *http.Response) ([]byte, error) {
	if resp == nil {
		return nil, nil
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("not ok response from telegram api")
	}
	defer func() {
		closeErr := resp.Body.Close()
		if closeErr != nil {
			log.Println("error close response body", closeErr.Error())
		}
	}()
	return io.ReadAll(resp.Body)
}

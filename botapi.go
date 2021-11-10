package tgbotapi

import (
	"bytes"
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
	Bot interface {
		GetMe() (*GetMeResponse, error)
		RawGetRequest(request string) ([]byte, error)
		RawPostRequest(request string, body []byte) ([]byte, error)
		SendMessage(request *SendMessage) ([]byte, error)
		SendPoll(request *SendPoll) ([]byte, error)
		GetUpdates(request *GetUpdates) ([]byte, error)
	}
	GetMeResponse struct {
		OK     bool        `json:"ok"`
		Result GetMeResult `json:"result"`
	}
	GetMeResult struct {
		ID                      interface{} `json:"id"`
		IsBot                   bool        `json:"is_bot"`
		FirstName               string      `json:"first_name"`
		Username                string      `json:"username"`
		CanJoinGroups           bool        `json:"can_join_groups"`
		CanReadAllGroupMessages bool        `json:"can_read_all_group_messages"`
		SupportInlineQueries    bool        `json:"support_inline_queries"`
	}
	SendMessage struct {
		ChatID interface{} `json:"chat_id"`
		Text   string      `json:"text"`
	}
	SendPoll struct {
		ChatID          interface{} `json:"chat_id"`
		Question        string      `json:"question"` // Poll question, 1-300 characters
		Options         []string    `json:"options"`  // list of answer options, 2-10 strings 1-100 characters each
		Type            string      `json:"type"`     // “quiz” or “regular”
		CorrectOptionID int         `json:"correct_option_id"`
	}
	GetUpdates struct {
		Offset int `json:"offset"`
		Limit   int      `json:"limit"`
	}
	SendMessageResponse struct {
		OK bool `json:"ok"`
	}
	SendPollResponse struct {
		OK     bool    `json:"ok"`
		Result Message `json:"result"`
	}
	Message struct {
		MessageID int64 `json:"message_id"`
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

func (b *bot) SendMessage(request *SendMessage) ([]byte, error) {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", b.token)
	data, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(url, "application/json", bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	respData, err := b.read(resp)
	if err != nil {
		return nil, err
	}
	return respData, nil
}

func (b *bot) SendPoll(request *SendPoll) ([]byte, error) {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendPoll", b.token)
	data, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(url, "application/json", bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	respData, err := b.read(resp)
	if err != nil {
		return nil, err
	}
	return respData, nil
}

func (b *bot) GetUpdates(request *GetUpdates) ([]byte, error) {
	url := "https://api.telegram.org/bot" + b.token + "/getUpdates"
	data, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(url, "application/json", bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	respData, err := b.read(resp)
	if err != nil {
		return nil, err
	}
	return respData, nil
}

func (b *bot) RawPostRequest(request string, body []byte) ([]byte, error) {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/%s", b.token, request)
	resp, err := http.Post(url, "application/json", bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	data, err := b.read(resp)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (b *bot) RawGetRequest(request string) ([]byte, error) {
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

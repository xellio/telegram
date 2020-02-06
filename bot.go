package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

//
// Bot ...
//
type Bot struct {
	Name     string
	Username string
	token    string
}

//
// New ...
//
func New(name string, username string, token string) *Bot {
	return &Bot{
		Name:     name,
		Username: username,
		token:    token,
	}
}

func (b *Bot) call(action string, payload ...interface{}) (interface{}, error) {
	if _, ok := methods[action]; !ok {
		return nil, fmt.Errorf("Passed action [%s] is not supported", action)
	}
	definition := methods[action]
	url := fmt.Sprintf(BaseURL, b.token, definition.Action)

	data, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(definition.Method, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(definition.Result)
	return definition.Result, err

}

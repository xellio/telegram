package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
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
// New - Create a new Bot instance
//
func New(token string) (*Bot, error) {
	b := &Bot{
		token: token,
	}

	u, err := b.GetMe()
	if err != nil {
		return nil, err
	}

	b.Username = u.Username
	b.Name = u.FirstName
	return b, nil
}

//
// call executes an API call to the telegram API
//
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

	var res Response
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return nil, err
	}

	if !res.OK {
		return nil, fmt.Errorf("%d: %s", res.ErrorCode, res.Description)
	}

	tmp, err := json.Marshal(res.Result)
	if err != nil {
		return nil, err
	}

	decoder := json.NewDecoder(bytes.NewReader(tmp))
	retValue := reflect.New(reflect.TypeOf(definition.Result))
	retInner := retValue.Interface()
	decoder.Decode(retInner)
	return retInner, err
}

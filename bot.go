package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"reflect"
)

//
// Bot  - This object represents a bot.
//
type Bot struct {
	Name     string
	Username string
	token    string
}

//
// New - Create a new Bot instance.
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

func (b *Bot) prepareRequest(definition APIMethod, payload ...interface{}) (*http.Request, error) {
	var err error
	var req *http.Request
	apiURL := fmt.Sprintf(BaseURL, b.token, definition.Action)

	switch definition.Method {
	case http.MethodGet:
		req, err = b.prepareGETRequest(apiURL, payload...)
		break
	case http.MethodPost:
		req, err = b.preparePOSTRequest(apiURL, payload...)
		break
	}

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	log.Println(req)

	return req, err
}

func (b *Bot) prepareGETRequest(apiURL string, payload ...interface{}) (*http.Request, error) {
	if len(payload) <= 0 {
		return http.NewRequest(http.MethodGet, apiURL, nil)
	}

	if len(payload) > 1 {
		return nil, fmt.Errorf("Unable to handle more than one payload in a GET call. payload should be a map[string]string")
	}

	baseURL, err := url.Parse(apiURL)
	if err != nil {
		return nil, err
	}

	queryMap := payload[0].(map[string]string)
	params := url.Values{}
	for k, v := range queryMap {
		params.Add(k, v)
	}

	baseURL.RawQuery = params.Encode()
	return http.NewRequest(http.MethodGet, baseURL.String(), nil)
}

func (b *Bot) preparePOSTRequest(apiURL string, payload ...interface{}) (*http.Request, error) {

	var pl interface{}
	pl = payload
	if len(payload) == 1 {
		pl = payload[0]
	}

	data, err := json.Marshal(pl)
	if err != nil {
		return nil, err
	}
	log.Println(string(data))
	return http.NewRequest(http.MethodPost, apiURL, bytes.NewBuffer(data))
}

//
// call executes an API call to the telegram API.
//
func (b *Bot) call(action string, payload ...interface{}) (interface{}, error) {
	if _, ok := methods[action]; !ok {
		return nil, fmt.Errorf("Passed action [%s] is not supported", action)
	}
	definition := methods[action]
	req, err := b.prepareRequest(definition, payload...)
	if err != nil {
		return nil, err
	}

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

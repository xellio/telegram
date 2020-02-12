package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"time"
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

//
// ListenForUpdates ...
//
func (b *Bot) ListenForUpdates(offset int, limit int, timeout int, allowedUpdates []string) (<-chan Update, error) {
	ch := make(chan Update, 100)

	go func() {
		for {
			updates, err := b.GetUpdates(offset, limit, timeout, allowedUpdates)
			if err != nil {
				log.Println(err)
				time.Sleep(time.Second * 5)
				continue
			}

			for _, update := range updates {

				if update.UpdateID >= 1 {
					offset = update.UpdateID + 1
					ch <- *update
				}
			}
		}
	}()

	return ch, nil
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

	log.Println(req.URL)
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

	return prepareResult(res.Result, definition.Result)

}

//
// prepareResult - ...
//
func prepareResult(result interface{}, resultDefinition interface{}) (interface{}, error) {
	tmp, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}

	if resultDefinition == nil {
		return tmp, nil
	}

	decoder := json.NewDecoder(bytes.NewReader(tmp))
	retValue := reflect.New(reflect.TypeOf(resultDefinition))
	res := retValue.Interface()
	decoder.Decode(res)
	return res, err
}

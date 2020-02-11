package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

// BaseURL ...
const BaseURL = "https://api.telegram.org/bot%s/%s"

//
// APIMethod  - This object represents an API call definition used in the methods map.
//
type APIMethod struct {
	Name   string
	Action string
	Method string
	Result interface{}
}

//
// methods map contains a definition of known and supported API calls.
//
var methods = map[string]APIMethod{
	"getMe": APIMethod{
		Name:   "getMe",
		Action: "getMe",
		Method: http.MethodGet,
		Result: User{},
	},
	"getUpdates": APIMethod{
		Name:   "getUpdates",
		Action: "getUpdates",
		Method: http.MethodGet,
		Result: []Update{},
	},
	"getWebhookInfo": APIMethod{
		Name:   "getWebhookInfo",
		Action: "getWebhookInfo",
		Method: http.MethodGet,
		Result: WebhookInfo{},
	},
	"getChat": APIMethod{
		Name:   "getChat",
		Action: "getChat",
		Method: http.MethodGet,
		Result: Chat{},
	},
	"getChatAdministrators": APIMethod{
		Name:   "getChatAdministrators",
		Action: "getChatAdministrators",
		Method: http.MethodGet,
		Result: []ChatMember{},
	},
	"getChatMembersCount": APIMethod{
		Name:   "getChatMembersCount",
		Action: "getChatMembersCount",
		Method: http.MethodGet,
	},
	"getChatMember": APIMethod{
		Name:   "getChatMember",
		Action: "getChatMember",
		Method: http.MethodGet,
		Result: ChatMember{},
	},
	"getUserProfilePhotos": APIMethod{
		Name:   "getUserProfilePhotos",
		Action: "getUserProfilePhotos",
		Method: http.MethodGet,
		Result: UserProfilePhotos{},
	},
	"getFile": APIMethod{
		Name:   "getFile",
		Action: "getFile",
		Method: http.MethodGet,
		Result: File{},
	},
	"getStickerSet": APIMethod{
		Name:   "getStickerSet",
		Action: "getStickerSet",
		Method: http.MethodGet,
		Result: StickerSet{},
	},
	"exportChatInviteLink": APIMethod{
		Name:   "exportChatInviteLink",
		Action: "exportChatInviteLink",
		Method: http.MethodPost,
	},
	"deleteWebhook": APIMethod{
		Name:   "deleteWebhook",
		Action: "deleteWebhook",
		Method: http.MethodPost,
	},
	"sendMessage": APIMethod{
		Name:   "sendMessage",
		Action: "sendMessage",
		Method: http.MethodPost,
		Result: Message{},
	},
	"forwardMessage": APIMethod{
		Name:   "forwardMessage",
		Action: "forwardMessage",
		Method: http.MethodPost,
		Result: Message{},
	},
	"sendPoll": APIMethod{
		Name:   "sendPoll",
		Action: "sendPoll",
		Method: http.MethodPost,
		Result: Message{},
	},
	"setChatTitle": APIMethod{
		Name:   "setChatTitle",
		Action: "setChatTitle",
		Method: http.MethodPost,
	},
	"setChatDescription": APIMethod{
		Name:   "setChatDescription",
		Action: "setChatDescription",
		Method: http.MethodPost,
	},
	"setChatAdministratorCustomTitle": APIMethod{
		Name:   "setChatAdministratorCustomTitle",
		Action: "setChatAdministratorCustomTitle",
		Method: http.MethodPost,
	},
	"kickChatMember": APIMethod{
		Name:   "kickChatMember",
		Action: "kickChatMember",
		Method: http.MethodPost,
	},
	"unbanChatMember": APIMethod{
		Name:   "unbanChatMember",
		Action: "unbanChatMember",
		Method: http.MethodPost,
	},
	"deleteChatPhoto": APIMethod{
		Name:   "deleteChatPhoto",
		Action: "deleteChatPhoto",
		Method: http.MethodPost,
	},
	"leaveChat": APIMethod{
		Name:   "leaveChat",
		Action: "leaveChat",
		Method: http.MethodPost,
	},
	"setChatStickerSet": APIMethod{
		Name:   "setChatStickerSet",
		Action: "setChatStickerSet",
		Method: http.MethodPost,
	},
	"deleteChatStickerSet": APIMethod{
		Name:   "deleteChatStickerSet",
		Action: "deleteChatStickerSet",
		Method: http.MethodPost,
	},
	"pinChatMessage": APIMethod{
		Name:   "pinChatMessage",
		Action: "pinChatMessage",
		Method: http.MethodPost,
	},
	"unpinChatMessage": APIMethod{
		Name:   "unpinChatMessage",
		Action: "unpinChatMessage",
		Method: http.MethodPost,
	},
	"sendChatAction": APIMethod{
		Name:   "sendChatAction",
		Action: "sendChatAction",
		Method: http.MethodPost,
	},
	"sendLocation": APIMethod{
		Name:   "sendLocation",
		Action: "sendLocation",
		Method: http.MethodPost,
		Result: Message{},
	},
	"deleteStickerFromSet": APIMethod{
		Name:   "deleteStickerFromSet",
		Action: "deleteStickerFromSet",
		Method: http.MethodPost,
	},
	"setStickerPositionInSet": APIMethod{
		Name:   "setStickerPositionInSet",
		Action: "setStickerPositionInSet",
		Method: http.MethodPost,
	},
	"sendInvoice": APIMethod{
		Name:   "sendInvoice",
		Action: "sendInvoice",
		Method: http.MethodPost,
		Result: Message{},
	},
	"answerShippingQuery": APIMethod{
		Name:   "answerShippingQuery",
		Action: "answerShippingQuery",
		Method: http.MethodPost,
	},
	"answerPreCheckoutQuery": APIMethod{
		Name:   "answerPreCheckoutQuery",
		Action: "answerPreCheckoutQuery",
		Method: http.MethodPost,
	},
}

//
// Response - This object represents a response from the telegram API.
//
type Response struct {
	OK          bool        `json:"ok"`
	ErrorCode   int         `json:"error_code"`
	Description string      `json:"description"`
	Result      interface{} `json:"result"`
}

//
// prepareRequest - ...
//
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

	return req, err
}

//
// prepareGETRequest - ...
//
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

//
// preparePOSTRequest - ...
//
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

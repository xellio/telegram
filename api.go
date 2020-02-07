package telegram

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
		Method: "GET",
		Result: User{},
	},
	"getUpdates": APIMethod{
		Name:   "getUpdates",
		Action: "getUpdates",
		Method: "GET",
		Result: []Update{},
	},
	"getChat": APIMethod{
		Name:   "getChat",
		Action: "getChat",
		Method: "GET",
		Result: Chat{},
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

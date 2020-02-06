package telegram

// BaseURL ...
const BaseURL = "https://api.telegram.org/bot%s/%s"

//
// APIMethod ...
//
type APIMethod struct {
	Name   string
	Action string
	Method string
	Result interface{}
}

var methods = map[string]APIMethod{
	"getMe": APIMethod{
		Name:   "getMe",
		Action: "getMe",
		Method: "GET",
		Result: &User{},
	},
}

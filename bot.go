package telegram

//
// Bot ...
//
type Bot struct {
	Name     string
	Username string
	Token    string
}

//
// New ...
//
func New(name string, username string, token string) *Bot {
	return &Bot{
		Name:     name,
		Username: username,
		Token:    token,
	}
}

func (b Bot) GetUpdates() {

}

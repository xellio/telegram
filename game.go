package telegram

//
// Game - This object represents a game.
// Use BotFather to create and edit games, their short names will act as unique identifiers.
// https://core.telegram.org/bots/api#game
//
type Game struct {
	Title        string           `json:"title"`
	Description  string           `json:"description"`
	Photo        []*PhotoSize     `json:"photo"`
	Text         string           `json:"text"`
	TextEntities []*MessageEntity `json:"text_entities"`
	Animation    *Animation       `json:"animation"`
}

//
// CallbackGame - A placeholder, currently holds no information. Use BotFather to set up your game.
// https://core.telegram.org/bots/api#callbackgame
//
type CallbackGame struct {
}

//
// GameHighScore - This object represents one row of the high scores table for a game.
// https://core.telegram.org/bots/api#gamehighscore
// TODO
//
type GameHighScore struct {
}

//
// SendGame - Use this method to send a game.
// https://core.telegram.org/bots/api#sendgame
// TODO
//
func (b *Bot) SendGame() (*Message, error) {
	return nil, nil
}

//
// SetGameScore - Use this method to set the score of the specified user in a game.
// https://core.telegram.org/bots/api#setgamescore
// TODO
//
func (b *Bot) SetGameScore() (*Message, error) {
	return nil, nil
}

//
// GetGameHighScores - Use this method to get data for high score tables.
// Will return the score of the specified user and several of his neighbors in a game.
// https://core.telegram.org/bots/api#getgamehighscores
// TODO
//
func (b *Bot) GetGameHighScores() ([]*GameHighScore, error) {
	return nil, nil
}

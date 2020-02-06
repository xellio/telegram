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

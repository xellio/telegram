package telegram

//
// VHost ...
//
const VHost = "https://api.telegram.org/bot%s/"

//
// API ...
//
var API = map[string]string{
	"getUpdates": "getUpdates",
}

//
// Update ...
//
type Update struct {
	UpdateID           int                 `json:"update_id"`
	Message            *Message            `json:"message"`
	EditedMessage      *Message            `json:"edited_message"`
	ChannelPost        *Message            `json:"channel_post"`
	EditedChannelPost  *Message            `json:"edited_channel_post"`
	InlineQuery        *InlineQuery        `json:"inline_query"`
	ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result"`
	CallbackQuery      *CallbackQuery      `json:"callback_query"`
	ShippingQuery      *ShippingQuery      `json:"shipping_query"`
	PreCheckoutQuery   *PreCheckoutQuery   `json:"pre_checkout_query"`
	Poll               *Poll               `json:"poll"`
	PollAnswer         *PollAnswer         `json:"poll_answer"`
}

//
// Message ...
//
type Message struct {
}

//
// InlineQuery ...
//
type InlineQuery struct {
	ID       string    `json:"id"`
	From     *User     `json:"user"`
	Location *Location `json:"location"`
	Query    string    `json:"query"`
	Offset   string    `json:"offset"`
}

//
// ChosenInlineResult ...
//
type ChosenInlineResult struct {
	ResultID        string    `json:"result_id"`
	From            *User     `json:"from"`
	Location        *Location `json:"location"`
	InlineMessageID string    `json:"inline_message_id"`
	Query           string    `json:"query"`
}

//
// CallbackQuery ...
//
type CallbackQuery struct {
	ID              string   `json:"id"`
	From            *User    `json:"user"`
	Message         *Message `json:"message"`
	InlineMessageID string   `json:"inline_message_id"`
	ChatInstance    string   `json:"chat_instance"`
	Data            string   `json:"data"`
	GameShortName   string   `json:"game_short_name"`
}

//
// ShippingQuery ...
//
type ShippingQuery struct {
	ID              string           `json:"id"`
	From            *User            `json:"user"`
	InvoicePayload  string           `json:"invoice_payload"`
	ShippingAddress *ShippingAddress `json:"shipping_adress"`
}

//
// ShippingAddress ...
// https://core.telegram.org/bots/api#shippingaddress
//
type ShippingAddress struct {
}

//
// PreCheckoutQuery ...
//
type PreCheckoutQuery struct {
}

//
// Poll ...
//
type Poll struct {
	ID                    string        `json:"id"`
	Question              string        `json:"question"`
	Options               []*PollOption `json:"options"`
	TotalVoterCount       int           `json:"total_voter_count"`
	IsClosed              bool          `json:"is_closed"`
	IsAnonymous           bool          `json:"is_anonymous"`
	Type                  string        `json:"type"`
	AllowsMultipleAnswers bool          `json:"allows_multiple_answers"`
	CorrectOptionID       int           `json:"correct_option_id"`
}

//
// PollOption ...
//
type PollOption struct {
	Text       string `json:"text"`
	VoterCount int    `json:"voter_count"`
}

//
// PollAnswer ...
//
type PollAnswer struct {
	PollID    string `json:"poll_id"`
	User      *User  `json:"user"`
	OptionIDs []int  `json:"option_ids"`
}

//
// User ...
//
type User struct {
	ID                      int    `json:"id"`
	IsBot                   bool   `json:"is_bot"`
	FirstName               string `json:"first_name"`
	LastName                string `json:"last_name"`
	Username                string `json:"username"`
	LanguageCode            string `json:"language_code"`
	CanJoinGroups           bool   `json:"can_join_groups"`
	CanReadAllGroupMessages bool   `json:"can_read_all_group_messages"`
	SupportsInlineQueries   bool   `json:"supports_inline_queries"`
}

//
// Location ...
//
type Location struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

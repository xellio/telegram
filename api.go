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
// Update - This object represents an incoming update.
// At most one of the optional parameters can be present in any given update.
// https://core.telegram.org/bots/api#update
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
// Message - This object represents a message.
// https://core.telegram.org/bots/api#message
//
type Message struct {
	MessageID             int                   `json:"message_id"`
	From                  *User                 `json:"user"`
	Date                  int                   `json:"date"`
	Chat                  *Chat                 `json:"chat"`
	ForwardFrom           *User                 `json:"forward_from"`
	ForwardFromChat       *Chat                 `json:"forward_from_chat"`
	ForwardFromMessageID  int                   `json:"forward_from_message_id"`
	ForwardSignature      string                `json:"forward_signature"`
	ForwardSenderName     string                `json:"forward_sender_name"`
	ForwardDate           int                   `json:"forward_date"`
	ReplyToMessage        *Message              `json:"reply_to_message"`
	EditDate              int                   `json:"edit_date"`
	MediaGroupID          string                `json:"media_group_id"`
	AuthorSignature       string                `json:"author_signature"`
	Text                  string                `json:"text"`
	Entities              []*MessageEntity      `json:"entities"`
	CaptionEntities       []*MessageEntity      `json:"caption_entities"`
	Audio                 *Audio                `json:"audio"`
	Document              *Document             `json:"document"`
	Animation             *Animation            `json:"animation"`
	Game                  *Game                 `json:"game"`
	Photo                 []*PhotoSize          `json:"photo"`
	Sticker               *Sticker              `json:"sticker"`
	Video                 *Video                `json:"video"`
	Voice                 *Voice                `json:"voice"`
	VideoNote             *VideoNote            `json:"video_note"`
	Caption               string                `json:"caption"`
	Contact               *Contact              `json:"contact"`
	Location              *Location             `json:"location"`
	Venue                 *Venue                `json:"venue"`
	Poll                  *Poll                 `json:"poll"`
	NewChatMembers        []*User               `json:"new_chat_members"`
	LeftChatMember        *User                 `json:"left_chat_member"`
	NewChatTitle          string                `json:"new_chat_title"`
	NewChatPhoto          []*PhotoSize          `json:"new_chat_photo"`
	DeleteChatPhoto       bool                  `json:"delete_chat_photo"`
	GroupChatCreated      bool                  `json:"group_chat_created"`
	SupergroupChatCreated bool                  `json:"supergroup_chat_created"`
	ChannelChatCreated    bool                  `json:"channel_chat_created"`
	MigrateToChatID       int                   `json:"migrate_to_chat_id"`
	MigrateFromChatID     int                   `json:"migrate_from_chat_id"`
	PinnedMessage         *Message              `json:"pinned_message"`
	Invoice               *Invoice              `json:"invoice"`
	SuccessfulPayment     *SuccessfulPayment    `json:"successful_payment"`
	ConnectedWebsite      string                `json:"connected_website"`
	PassportData          *PassportData         `json:"passport_data"`
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup"`
}

//
// MessageEntity - This object represents one special entity in a text message.
// For example, hashtags, usernames, URLs, etc.
// https://core.telegram.org/bots/api#messageentity
//
type MessageEntity struct {
	Type     string `json:"type"`
	Offset   int    `json:"offset"`
	Length   int    `json:"length"`
	URL      string `json:"url"`
	User     *User  `json:"user"`
	Language string `json:"language"`
}

//
// Audio - This object represents an audio file to be treated as music by the Telegram clients.
// https://core.telegram.org/bots/api#audio
//
type Audio struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Duration     int        `json:"duration"`
	Performer    string     `json:"performer"`
	Title        string     `json:"title"`
	MIMEType     string     `json:"mime_type"`
	FileSize     int        `json:"file_size"`
	Thumb        *PhotoSize `json:"thumb"`
}

//
// Document - This object represents a general file (as opposed to photos, voice messages and audio files).
// https://core.telegram.org/bots/api#document
//
type Document struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Thumb        *PhotoSize `json:"thumb"`
	FileName     string     `json:"file_name"`
	MIMEType     string     `json:"mime_type"`
	FileSize     int        `json:"file_size"`
}

//
// Animation - This object represents an animation file (GIF or H.264/MPEG-4 AVC video without sound).
// https://core.telegram.org/bots/api#animation
//
type Animation struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Width        int        `json:"width"`
	Height       int        `json:"height"`
	Duration     int        `json:"duration"`
	Thumb        *PhotoSize `json:"thumb"`
	FileName     string     `json:"file_name"`
	MIMEType     string     `json:"mime_type"`
	FileSize     int        `json:"file_size"`
}

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
// Sticker - This object represents a sticker.
// https://core.telegram.org/bots/api#sticker
//
type Sticker struct {
	FileID       string        `json:"file_id"`
	FileUniqueID string        `json:"file_unique_id"`
	Width        int           `json:"width"`
	Height       int           `json:"height"`
	IsAnimated   bool          `json:"is_animated"`
	Thumb        *PhotoSize    `json:"thumb"`
	Emoji        string        `json:"emoji"`
	SetName      string        `json:"set_name"`
	MaskPosition *MaskPosition `json:"mask_position"`
	FileSize     int           `json:"file_size"`
}

// MaskPosition - This object describes the position on faces where a mask should be placed by default.
// https://core.telegram.org/bots/api#maskposition
//
type MaskPosition struct {
	Point  string  `json:"point"`
	XShift float64 `json:"x_shift"`
	YShift float64 `json:"y_shift"`
	Scale  float64 `json:"scale"`
}

//
// Video - This object represents a video file.
// https://core.telegram.org/bots/api#video
//
type Video struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Width        int        `json:"width"`
	Height       int        `json:"height"`
	Duration     int        `json:"duration"`
	Thumb        *PhotoSize `json:"thumb"`
	MIMEType     string     `json:"mime_type"`
	FileSize     int        `json:"file_size"`
}

//
// Voice - This object represents a voice note.
// https://core.telegram.org/bots/api#voice
//
type Voice struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	Duration     int    `json:"duration"`
	MIMEType     string `json:"mime_type"`
	FileSize     int    `json:"file_size"`
}

//
// VideoNote - This object represents a video message (available in Telegram apps as of v.4.0).
// https://core.telegram.org/bots/api#videonote
//
type VideoNote struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Length       int        `json:"length"`
	Duration     int        `json:"duration"`
	Thumb        *PhotoSize `json:"thumb"`
	FileSize     int        `json:"file_size"`
}

//
// Contact - This object represents a phone contact.
// https://core.telegram.org/bots/api#contact
//
type Contact struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	UserID      int    `json:"user_id"`
	VCard       string `json:"vcard"`
}

//
// Venue - This object represents a venue.
// https://core.telegram.org/bots/api#venue
//
type Venue struct {
	Location       *Location `json:"location"`
	Title          string    `json:"title"`
	Address        string    `json:"address"`
	FoursquareID   string    `json:"foursquare_id"`
	FoursquareType string    `json:"foursquare_type"`
}

//
// PhotoSize - This object represents one size of a photo or a file / sticker thumbnail.
// https://core.telegram.org/bots/api#photosize
//
type PhotoSize struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
	FileSize     int    `json:"file_size"`
}

//
// Invoice - This object contains basic information about an invoice.
// https://core.telegram.org/bots/api#invoice
//
type Invoice struct {
	Title          string `json:"title"`
	Description    string `json:"description"`
	StartParameter string `json:"start_parameter"`
	Currency       string `json:"currency"`
	TotalAmount    int    `json:"total_amount"`
}

//
// SuccessfulPayment - This object contains basic information about a successful payment.
// https://core.telegram.org/bots/api#successfulpayment
//
type SuccessfulPayment struct {
	Currency                string     `json:"currency"`
	TotalAmount             int        `json:"total_amount"`
	InvoicePayload          string     `json:"invoice_payload"`
	ShippingOptionID        string     `json:"shipping_option_id"`
	OrderInfo               *OrderInfo `json:"order_info"`
	TelegramPaymentChargeID string     `json:"telegram_payment_charge_id"`
	ProviderPaymentChargeID string     `json:"provider_payment_charge_id"`
}

//
// PassportData - Contains information about Telegram Passport data shared with the bot by the user.
// https://core.telegram.org/bots/api#passportdata
//
type PassportData struct {
	Data        []*EncryptedPassportElement `json:"data"`
	Credentials *EncryptedCredentials       `json:"credentials"`
}

//
// EncryptedPassportElement - Contains information about documents or other Telegram Passport elements shared with the bot by the user.
// https://core.telegram.org/bots/api#encryptedpassportelement
//
type EncryptedPassportElement struct {
	Type        string          `json:"type"`
	Data        string          `json:"data"`
	PhoneNumber string          `json:"phone_number"`
	Email       string          `json:"email"`
	Files       []*PassportFile `json:"files"`
	FrontSide   *PassportFile   `json:"front_side"`
	ReverseSide *PassportFile   `json:"reverse_side"`
	Selfie      *PassportFile   `json:"selfie"`
	Translation []*PassportFile `json:"translation"`
	Hash        string          `json:"hash"`
}

//
// PassportFile - This object represents a file uploaded to Telegram Passport.
// Currently all Telegram Passport files are in JPEG format when decrypted and don't exceed 10MB.
// https://core.telegram.org/bots/api#passportfile
//
type PassportFile struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	FileSize     int    `json:"file_size"`
	FileDate     int    `json:"file_date"`
}

//
// EncryptedCredentials - Contains data required for decrypting and authenticating EncryptedPassportElement.
// See the Telegram Passport Documentation for a complete description of the data decryption and authentication processes.
// https://core.telegram.org/bots/api#encryptedcredentials
//
type EncryptedCredentials struct {
	Data   string `json:"data"`
	Hash   string `json:"hash"`
	Secret string `json:"secret"`
}

//
// InlineKeyboardMarkup - This object represents an inline keyboard that appears right next to the message it belongs to.
// https://core.telegram.org/bots/api#inlinekeyboardmarkup
//
type InlineKeyboardMarkup struct {
	InlineKeyboard [][]*InlineKeyboardButton `json:"inline_keyboard"`
}

//
// InlineKeyboardButton - This object represents one button of an inline keyboard.
// You must use exactly one of the optional fields.
// https://core.telegram.org/bots/api#inlinekeyboardbutton
//
type InlineKeyboardButton struct {
	Text                         string        `json:"text"`
	URL                          string        `json:"url"`
	LoginURL                     *LoginURL     `json:"login_url"`
	CallbackData                 string        `json:"callback_data"`
	SwitchInlineQuery            string        `json:"switch_inline_query"`
	SwitchInlineQueryCurrentChat string        `json:"switch_inline_query_current_chat"`
	CallbackGame                 *CallbackGame `json:"callback_game"`
	Pay                          bool          `json:"pay"`
}

//
// LoginURL - This object represents a parameter of the inline keyboard button used to automatically authorize a user.
// Serves as a great replacement for the Telegram Login Widget when the user is coming from Telegram. All the user needs to do is tap/click a button and confirm that they want to log in.
// https://core.telegram.org/bots/api#loginurl
//
type LoginURL struct {
	URL                string `json:"url"`
	ForwardText        string `json:"forward_text"`
	BotUsername        string `json:"bot_username"`
	RequestWriteAccess bool   `json:"request_write_access"`
}

//
// CallbackGame - A placeholder, currently holds no information. Use BotFather to set up your game.
// https://core.telegram.org/bots/api#callbackgame
//
type CallbackGame struct {
}

//
// Chat - This object represents a chat.
// https://core.telegram.org/bots/api#chat
//
type Chat struct {
	ID               int              `json:"id"`
	Type             string           `json:"type"`
	Title            string           `json:"title"`
	Username         string           `json:"username"`
	FirstName        string           `json:"first_name"`
	LastName         string           `json:"last_name"`
	Photo            *ChatPhoto       `json:"photo"`
	Description      string           `json:"description"`
	InviteLink       string           `json:"invite_link"`
	PinnedMessage    *Message         `json:"pinned_message"`
	Permissions      *ChatPermissions `json:"permissions"`
	SlowModeDelay    int              `json:"slow_mode_delay"`
	StickerSetName   string           `json:"sticker_set_name"`
	CanSetStickerSet bool             `json:"can_set_sticker_set"`
}

//
// ChatPermissions - Describes actions that a non-administrator user is allowed to take in a chat.
// https://core.telegram.org/bots/api#chatpermissions
//
type ChatPermissions struct {
	CanSendMessages       bool `json:"can_send_messages"`
	CanSendMediaMessages  bool `json:"can_send_media_messages"`
	CanSendPolls          bool `json:"can_send_polls"`
	CanSendOtherMessages  bool `json:"can_send_other_messages"`
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews"`
	CanChangeInfo         bool `json:"can_change_info"`
	CanInviteUsers        bool `json:"can_invite_users"`
	CanPinMessages        bool `json:"can_pin_messages"`
}

//
// ChatPhoto - This object represents a chat photo.
// https://core.telegram.org/bots/api#chatphoto
//
type ChatPhoto struct {
	SmallFileID       string `json:"small_file_id"`
	SmallFileUniqueID string `json:"small_file_unique_id"`
	BigFileID         string `json:"big_file_id"`
	BigFileUniqueID   string `json:"big_file_unique_id"`
}

//
// InlineQuery - This object represents an incoming inline query.
// When the user sends an empty query, your bot could return some default or trending results.
// https://core.telegram.org/bots/api#inlinequery
//
type InlineQuery struct {
	ID       string    `json:"id"`
	From     *User     `json:"user"`
	Location *Location `json:"location"`
	Query    string    `json:"query"`
	Offset   string    `json:"offset"`
}

//
// ChosenInlineResult - Represents a result of an inline query that was chosen by the user and sent to their chat partner.
// https://core.telegram.org/bots/api#choseninlineresult
//
type ChosenInlineResult struct {
	ResultID        string    `json:"result_id"`
	From            *User     `json:"from"`
	Location        *Location `json:"location"`
	InlineMessageID string    `json:"inline_message_id"`
	Query           string    `json:"query"`
}

//
// CallbackQuery - This object represents an incoming callback query from a callback button in an inline keyboard.
// If the button that originated the query was attached to a message sent by the bot, the field message will be present. If the button was attached to a message sent via the bot (in inline mode), the field inline_message_id will be present. Exactly one of the fields data or game_short_name will be present.
// https://core.telegram.org/bots/api#callbackquery
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
// ShippingQuery - This object contains information about an incoming shipping query.
// https://core.telegram.org/bots/api#shippingquery
//
type ShippingQuery struct {
	ID              string           `json:"id"`
	From            *User            `json:"user"`
	InvoicePayload  string           `json:"invoice_payload"`
	ShippingAddress *ShippingAddress `json:"shipping_adress"`
}

//
// ShippingAddress - This object represents a shipping address.
// https://core.telegram.org/bots/api#shippingaddress
//
type ShippingAddress struct {
	CountryCode string `json:"country_code"`
	State       string `json:"state"`
	City        string `json:"city"`
	StreetLine1 string `json:"street_line1"`
	StreetLine2 string `json:"street_line2"`
	PostCode    string `json:"post_code"`
}

//
// PreCheckoutQuery - This object contains information about an incoming pre-checkout query.
// https://core.telegram.org/bots/api#precheckoutquery
//
type PreCheckoutQuery struct {
	ID               string     `json:"id"`
	From             *User      `json:"from"`
	Currency         string     `json:"currency"`
	TotalAmount      int        `json:"total_amount"`
	InvoicePayload   string     `json:"invoice_payload"`
	ShippingOptionID string     `json:"shipping_option_id"`
	OrderInfo        *OrderInfo `json:"order_info"`
}

//
// OrderInfo - This object represents information about an order.
// https://core.telegram.org/bots/api#orderinfo
//
type OrderInfo struct {
	Name            string           `json:"name"`
	PhoneNumber     string           `json:"phone_number"`
	Email           string           `json:"email"`
	ShippingAddress *ShippingAddress `json:"shipping_address"`
}

//
// Poll - This object contains information about a poll.
// https://core.telegram.org/bots/api#poll
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
// PollOption - This object contains information about one answer option in a poll.
// https://core.telegram.org/bots/api#polloption
//
type PollOption struct {
	Text       string `json:"text"`
	VoterCount int    `json:"voter_count"`
}

//
// PollAnswer - This object represents an answer of a user in a non-anonymous poll.
// https://core.telegram.org/bots/api#pollanswer
//
type PollAnswer struct {
	PollID    string `json:"poll_id"`
	User      *User  `json:"user"`
	OptionIDs []int  `json:"option_ids"`
}

//
// User - This object represents a Telegram user or bot.
// https://core.telegram.org/bots/api#user
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
// Location - This object represents a point on the map.
// https://core.telegram.org/bots/api#location
//
type Location struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

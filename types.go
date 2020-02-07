package telegram

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
// Location - This object represents a point on the map.
// https://core.telegram.org/bots/api#location
//
type Location struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
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
// PollOption - This object contains information about one answer option in a poll.
// https://core.telegram.org/bots/api#polloption
//
type PollOption struct {
	Text       string `json:"text"`
	VoterCount int    `json:"voter_count,omitempty"`
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
// NewPoll - This object is used to create a new poll.
//
type NewPoll struct {
	ChatID                int      `json:"chat_id"`
	Question              string   `json:"question"`
	Options               []string `json:"options"`
	TotalVoterCount       int      `json:"total_voter_count"`
	IsClosed              bool     `json:"is_closed"`
	IsAnonymous           bool     `json:"is_anonymous"`
	Type                  string   `json:"type"`
	AllowsMultipleAnswers bool     `json:"allows_multiple_answers"`
	CorrectOptionID       int      `json:"correct_option_id"`
}

//
// UserProfilePhotos - This object represent a user's profile pictures.
// https://core.telegram.org/bots/api#userprofilephotos
// TODO
//
type UserProfilePhotos struct {
}

//
// File - This object represents a file ready to be downloaded.
// The file can be downloaded via the link https://api.telegram.org/file/bot<token>/<file_path>. It is guaranteed that the link will be valid for at least 1 hour. When the link expires, a new one can be requested by calling getFile.
// https://core.telegram.org/bots/api#file
// TODO
//
type File struct {
}

//
// ReplyKeyboardMarkup - This object represents a custom keyboard with reply options (see Introduction to bots for details and examples).
// https://core.telegram.org/bots/api#replykeyboardmarkup
// TODO
//
type ReplyKeyboardMarkup struct {
}

//
// KeyboardButton - This object represents one button of the reply keyboard.
// For simple text buttons String can be used instead of this object to specify text of the button. Optional fields request_contact, request_location, and request_poll are mutually exclusive.
// https://core.telegram.org/bots/api#keyboardbutton
// TODO
//
type KeyboardButton struct {
}

//
// KeyboardButtonPollType - This object represents type of a poll, which is allowed to be created and sent when the corresponding button is pressed.
// https://core.telegram.org/bots/api#keyboardbuttonpolltype
// TODO
//
type KeyboardButtonPollType struct {
}

//
// ReplyKeyboardRemove - Upon receiving a message with this object, Telegram clients will remove the current custom keyboard and display the default letter-keyboard.
// By default, custom keyboards are displayed until a new keyboard is sent by a bot. An exception is made for one-time keyboards that are hidden immediately after the user presses a button (see ReplyKeyboardMarkup).
// https://core.telegram.org/bots/api#replykeyboardremove
// TODO
//
type ReplyKeyboardRemove struct {
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
// ForceReply - Upon receiving a message with this object, Telegram clients will display a reply interface to the user (act as if the user has selected the bot‘s message and tapped ’Reply').
// This can be extremely useful if you want to create user-friendly step-by-step interfaces without having to sacrifice privacy mode.
// https://core.telegram.org/bots/api#forcereply
// TODO
//
type ForceReply struct {
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
// ChatMember - This object contains information about one member of a chat.
// https://core.telegram.org/bots/api#chatmember
// TODO
//
type ChatMember struct {
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
// ResponseParameters - Contains information about why a request was unsuccessful.
// https://core.telegram.org/bots/api#responseparameters
// TODO
//
type ResponseParameters struct {
}

//
// InputMedia - This object represents the content of a media message to be sent.
// It should be one of:
// - InputMediaAnimation
// - InputMediaDocument
// - InputMediaAudio
// - InputMediaPhoto
// - InputMediaVideo
// https://core.telegram.org/bots/api#inputmedia
// TODO
//
type InputMedia interface {
}

//
// InputMediaPhoto - Represents a photo to be sent.
// https://core.telegram.org/bots/api#inputmediaphoto
// TODO
//
type InputMediaPhoto struct {
}

//
// InputMediaVideo - Represents a video to be sent.
// https://core.telegram.org/bots/api#inputmediavideo
// TODO
//
type InputMediaVideo struct {
}

//
// InputMediaAnimation - Represents an animation file (GIF or H.264/MPEG-4 AVC video without sound) to be sent.
// https://core.telegram.org/bots/api#inputmediaanimation
// TODO
//
type InputMediaAnimation struct {
}

//
// InputMediaAudio - Represents an audio file to be treated as music to be sent.
// https://core.telegram.org/bots/api#inputmediaaudio
// TODO
//
type InputMediaAudio struct {
}

//
// InputMediaDocument - Represents a general file to be sent.
// https://core.telegram.org/bots/api#inputmediadocument
// TODO
//
type InputMediaDocument struct {
}

//
// InputFile - This object represents the contents of a file to be uploaded.
// Must be posted using multipart/form-data in the usual way that files are uploaded via the browser.
// https://core.telegram.org/bots/api#inputfile
// TODO
//
type InputFile struct {
}

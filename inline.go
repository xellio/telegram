package telegram

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
// InlineQueryResult - This object represents one result of an inline query.
// Telegram clients currently support results of the following 20 types:
//  - InlineQueryResultCachedAudio
//  - InlineQueryResultCachedDocument
//  - InlineQueryResultCachedGif
//  - InlineQueryResultCachedMpeg4Gif
//  - InlineQueryResultCachedPhoto
//  - InlineQueryResultCachedSticker
//  - InlineQueryResultCachedVideo
//  - InlineQueryResultCachedVoice
//  - InlineQueryResultArticle
//  - InlineQueryResultAudio
//  - InlineQueryResultContact
//  - InlineQueryResultGame
//  - InlineQueryResultDocument
//  - InlineQueryResultGif
//  - InlineQueryResultLocation
//  - InlineQueryResultMpeg4Gif
//  - InlineQueryResultPhoto
//  - InlineQueryResultVenue
//  - InlineQueryResultVideo
//  - InlineQueryResultVoice
// https://core.telegram.org/bots/api#inlinequeryresult
// TODO
//
type InlineQueryResult interface {
}

//
// InlineQueryResultArticle - Represents a link to an article or web page.
// https://core.telegram.org/bots/api#inlinequeryresultarticle
// TODO
//
type InlineQueryResultArticle struct {
}

//
// InlineQueryResultPhoto - Represents a link to a photo.
// By default, this photo will be sent by the user with optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the photo.
// https://core.telegram.org/bots/api#inlinequeryresultphoto
// TODO
//
type InlineQueryResultPhoto struct {
}

//
// InlineQueryResultGif - Represents a link to an animated GIF file.
// By default, this animated GIF file will be sent by the user with optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the animation.
// https://core.telegram.org/bots/api#inlinequeryresultgif
// TODO
//
type InlineQueryResultGif struct {
}

//
// InlineQueryResultMpeg4Gif - Represents a link to a video animation (H.264/MPEG-4 AVC video without sound).
// By default, this animated MPEG-4 file will be sent by the user with optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the animation.
// https://core.telegram.org/bots/api#inlinequeryresultmpeg4gif
// TODO
//
type InlineQueryResultMpeg4Gif struct {
}

//
// InlineQueryResultVideo - Represents a link to a page containing an embedded video player or a video file.
// By default, this video file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the video.
// https://core.telegram.org/bots/api#inlinequeryresultvideo
// TODO
//
type InlineQueryResultVideo struct {
}

//
// InlineQueryResultAudio - Represents a link to an MP3 audio file.
// By default, this audio file will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the audio.
// https://core.telegram.org/bots/api#inlinequeryresultaudio
// TODO
//
type InlineQueryResultAudio struct {
}

//
// InlineQueryResultVoice - Represents a link to a voice recording in an .ogg container encoded with OPUS.
// By default, this voice recording will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the the voice message.
// https://core.telegram.org/bots/api#inlinequeryresultvoice
// TODO
//
type InlineQueryResultVoice struct {
}

//
// InlineQueryResultDocument - Represents a link to a file.
// By default, this file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the file. Currently, only .PDF and .ZIP files can be sent using this method.
// https://core.telegram.org/bots/api#inlinequeryresultdocument
// TODO
//
type InlineQueryResultDocument struct {
}

//
// InlineQueryResultLocation - Represents a location on a map.
// By default, the location will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the location.
// https://core.telegram.org/bots/api#inlinequeryresultlocation
// TODO
//
type InlineQueryResultLocation struct {
}

//
// InlineQueryResultVenue - Represents a venue.
// By default, the venue will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the venue.
// https://core.telegram.org/bots/api#inlinequeryresultvenue
// TODO
//
type InlineQueryResultVenue struct {
}

//
// InlineQueryResultContact - Represents a contact with a phone number.
// By default, this contact will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the contact.
// https://core.telegram.org/bots/api#inlinequeryresultcontact
// TODO
//
type InlineQueryResultContact struct {
}

//
// InlineQueryResultGame - Represents a Game.
// https://core.telegram.org/bots/api#inlinequeryresultgame
// TODO
//
type InlineQueryResultGame struct {
}

//
// InlineQueryResultCachedPhoto - Represents a link to a photo stored on the Telegram servers.
// By default, this photo will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the photo.
// https://core.telegram.org/bots/api#inlinequeryresultcachedphoto
// TODO
//
type InlineQueryResultCachedPhoto struct {
}

//
// InlineQueryResultCachedGif - Represents a link to an animated GIF file stored on the Telegram servers.
// By default, this animated GIF file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with specified content instead of the animation.
// https://core.telegram.org/bots/api#inlinequeryresultcachedgif
// TODO
//
type InlineQueryResultCachedGif struct {
}

//
// InlineQueryResultCachedMpeg4Gif - Represents a link to a video animation (H.264/MPEG-4 AVC video without sound) stored on the Telegram servers.
// By default, this animated MPEG-4 file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the animation.
// https://core.telegram.org/bots/api#inlinequeryresultcachedmpeg4gif
// TODO
//
type InlineQueryResultCachedMpeg4Gif struct {
}

//
// InlineQueryResultCachedSticker - Represents a link to a sticker stored on the Telegram servers.
// By default, this sticker will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the sticker.
// https://core.telegram.org/bots/api#inlinequeryresultcachedsticker
// TODO
//
type InlineQueryResultCachedSticker struct {
}

//
// InlineQueryResultCachedDocument - Represents a link to a file stored on the Telegram servers.
// By default, this file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the file.
// https://core.telegram.org/bots/api#inlinequeryresultcacheddocument
// TODO
//
type InlineQueryResultCachedDocument struct {
}

//
// InlineQueryResultCachedVideo - Represents a link to a video file stored on the Telegram servers.
// By default, this video file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the video.
// https://core.telegram.org/bots/api#inlinequeryresultcachedvideo
// TODO
//
type InlineQueryResultCachedVideo struct {
}

//
// InlineQueryResultCachedVoice - Represents a link to a voice message stored on the Telegram servers.
// By default, this voice message will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the voice message.
// https://core.telegram.org/bots/api#inlinequeryresultcachedvoice
// TODO
//
type InlineQueryResultCachedVoice struct {
}

//
// InlineQueryResultCachedAudio - Represents a link to an MP3 audio file stored on the Telegram servers.
// By default, this audio file will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the audio.
// https://core.telegram.org/bots/api#inlinequeryresultcachedaudio
// TODO
//
type InlineQueryResultCachedAudio struct {
}

//
// InputMessageContent - This object represents the content of a message to be sent as a result of an inline query.
// Telegram clients currently support the following 4 types:
//  - InputTextMessageContent
//  - InputLocationMessageContent
//  - InputVenueMessageContent
//  - InputContactMessageContent
// https://core.telegram.org/bots/api#inputmessagecontent
// TODO
//
type InputMessageContent interface {
}

//
// InputTextMessageContent - Represents the content of a text message to be sent as the result of an inline query.
// https://core.telegram.org/bots/api#inputtextmessagecontent
// TODO
//
type InputTextMessageContent struct {
}

//
// InputLocationMessageContent - Represents the content of a location message to be sent as the result of an inline query.
// https://core.telegram.org/bots/api#inputlocationmessagecontent
// TODO
//
type InputLocationMessageContent struct {
}

//
// InputVenueMessageContent - Represents the content of a venue message to be sent as the result of an inline query.
// https://core.telegram.org/bots/api#inputvenuemessagecontent
// TODO
//
type InputVenueMessageContent struct {
}

//
// InputContactMessageContent - Represents the content of a contact message to be sent as the result of an inline query.
// https://core.telegram.org/bots/api#inputcontactmessagecontent
// TODO
//
type InputContactMessageContent struct {
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

package telegram

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

//
// StickerSet - This object represents a sticker set.
// https://core.telegram.org/bots/api#stickerset
// TODO
//
type StickerSet struct {
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
// SendSticker - Use this method to send static .WEBP or animated .TGS stickers.
// https://core.telegram.org/bots/api#sendsticker
// TODO
//
func SendSticker() (*Message, error) {
	return nil, nil
}

//
// GetStickerSet - Use this method to get a sticker set.
// https://core.telegram.org/bots/api#getstickerset
// TODO
//
func GetStickerSet() (*StickerSet, error) {
	return nil, nil
}

//
// UploadStickerFile - Use this method to upload a .png file with a sticker for later use in CreateNewStickerSet and AddStickerToSet methods (can be used multiple times).
// https://core.telegram.org/bots/api#uploadstickerfile
// TODO
//
func UploadStickerFile() (*File, error) {
	return nil, nil
}

//
// CreateNewStickerSet - Use this method to create new sticker set owned by a user.
// The bot will be able to edit the created sticker set.
// https://core.telegram.org/bots/api#createnewstickerset
// TODO
//
func CreateNewStickerSet() (success bool) {
	return false
}

//
// AddStickerToSet - Use this method to add a new sticker to a set created by the bot.
// https://core.telegram.org/bots/api#addstickertoset
// TODO
//
func AddStickerToSet() (success bool) {
	return false
}

//
// SetStickerPositionInSet - Use this method to move a sticker in a set created by the bot to a specific position.
// https://core.telegram.org/bots/api#setstickerpositioninset
// TODO
//
func SetStickerPositionInSet() (success bool) {
	return false
}

//
// DeleteStickerFromSet - Use this method to delete a sticker from a set created by the bot.
// https://core.telegram.org/bots/api#deletestickerfromset
// TODO
//
func DeleteStickerFromSet() (success bool) {
	return false
}

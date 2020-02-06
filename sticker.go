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

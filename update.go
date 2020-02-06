package telegram

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
// WebhookInfo - Contains information about the current status of a webhook.
// https://core.telegram.org/bots/api#webhookinfo
// TODO
//
type WebhookInfo struct {
}

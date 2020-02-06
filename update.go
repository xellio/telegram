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

//
// GetUpdates - Use this method to receive incoming updates using long polling.
// https://core.telegram.org/bots/api#getupdates
// TODO
//
func GetUpdates() ([]*Update, error) {
	return nil, nil
}

//
// SetWebhook - Use this method to specify a url and receive incoming updates via an outgoing webhook.
// Whenever there is an update for the bot, we will send an HTTPS POST request to the specified url, containing a JSON-serialized Update. In case of an unsuccessful request, we will give up after a reasonable amount of attempts. Returns True on success.
// https://core.telegram.org/bots/api#setwebhook
// TODO
//
func SetWebhook() bool {
	return false
}

//
// DeleteWebhook - Use this method to remove webhook integration if you decide to switch back to getUpdates.
// https://core.telegram.org/bots/api#deletewebhook
// TODO
//
func DeleteWebhook() bool {
	return false
}

//
// GetWebhookInfo - Use this method to get current webhook status.
// https://core.telegram.org/bots/api#getwebhookinfo
// TODO
//
func GetWebhookInfo() (*WebhookInfo, error) {
	return nil, nil
}

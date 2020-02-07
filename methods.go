package telegram

import (
	"strconv"
)

// GetMe - A simple method for testing your bot's auth token.
// https://core.telegram.org/bots/api#getme
//
func (b *Bot) GetMe() (*User, error) {
	result, err := b.call("getMe")
	if err != nil {
		return nil, err
	}
	return result.(*User), nil
}

//
// SendMessage - Use this method to send text messages.
// https://core.telegram.org/bots/api#sendmessage
//
func (b *Bot) SendMessage(chatID int, text string, parseMode string, disableWebPagePreview bool, disableNotification bool, replyToMessageID int, replyMarkup interface{}) (*Message, error) {
	params := map[string]interface{}{
		"chat_id":              chatID,
		"text":                 text,
		"disable_notification": disableNotification,
		"reply_to_message_id":  replyToMessageID,
	}

	if replyMarkup != nil {
		params["reply_markup"] = replyMarkup
	}

	result, err := b.call("sendMessage", params)
	if err != nil {
		return nil, err
	}
	return result.(*Message), nil
}

//
// ForwardMessage - Use this method to forward messages of any kind.
// https://core.telegram.org/bots/api#forwardmessage
// TODO
//
func (b *Bot) ForwardMessage() (*Message, error) {
	return nil, nil
}

//
// SendPhoto - Use this method to send photos.
// https://core.telegram.org/bots/api#sendphoto
// TODO
//
func (b *Bot) SendPhoto() (*Message, error) {
	return nil, nil
}

//
// SendAudio - Use this method to send audio files, if you want Telegram clients to display them in the music player.
// Your audio must be in the .MP3 or .M4A format. Bots can currently send audio files of up to 50 MB in size, this limit may be changed in the future.
// For sending voice messages, use the sendVoice method instead.
// https://core.telegram.org/bots/api#sendaudio
// TODO
//
func (b *Bot) SendAudio() (*Message, error) {
	return nil, nil
}

//
// SendDocument - Use this method to send general files.
// Bots can currently send files of any type of up to 50 MB in size, this limit may be changed in the future.
// https://core.telegram.org/bots/api#senddocument
// TODO
//
func (b *Bot) SendDocument() (*Message, error) {
	return nil, nil
}

//
// SendVideo - Use this method to send video files, Telegram clients support mp4 videos (other formats may be sent as Document).
// Bots can currently send files of any type of up to 50 MB in size, this limit may be changed in the future.
// https://core.telegram.org/bots/api#sendvideo
// TODO
//
func (b *Bot) SendVideo() (*Message, error) {
	return nil, nil
}

//
// SendAnimation - Use this method to send animation files (GIF or H.264/MPEG-4 AVC video without sound).
// Bots can currently send animation files of up to 50 MB in size, this limit may be changed in the future.
// https://core.telegram.org/bots/api#sendanimation
// TODO
//
func (b *Bot) SendAnimation() (*Message, error) {
	return nil, nil
}

//
// SendVoice - Use this method to send audio files, if you want Telegram clients to display the file as a playable voice message.
// For this to work, your audio must be in an .ogg file encoded with OPUS (other formats may be sent as Audio or Document).
// Bots can currently send voice messages of up to 50 MB in size, this limit may be changed in the future.
// https://core.telegram.org/bots/api#sendvoice
// TODO
//
func (b *Bot) SendVoice() (*Message, error) {
	return nil, nil
}

//
// SendVideoNote - Use this method to send video messages.
// https://core.telegram.org/bots/api#sendvideonote
// TODO
//
func (b *Bot) SendVideoNote() (*Message, error) {
	return nil, nil
}

//
// SendMediaGroup - Use this method to send a group of photos or videos as an album.
// https://core.telegram.org/bots/api#sendmediagroup
// TODO
//
func (b *Bot) SendMediaGroup() (*Message, error) {
	return nil, nil
}

//
// SendLocation - Use this method to send point on the map.
// https://core.telegram.org/bots/api#sendlocation
// TODO
//
func (b *Bot) SendLocation() (*Message, error) {
	return nil, nil
}

//
// EditMessageLiveLocation - Use this method to edit live location messages.
// A location can be edited until its live_period expires or editing is explicitly disabled by a call to stopMessageLiveLocation.
// https://core.telegram.org/bots/api#editmessagelivelocation
// TODO
//
func (b *Bot) EditMessageLiveLocation() (*Message, error) {
	return nil, nil
}

//
// StopMessageLiveLocation - Use this method to stop updating a live location message before live_period expires.
// https://core.telegram.org/bots/api#stopmessagelivelocation
// TODO
//
func (b *Bot) StopMessageLiveLocation() (*Message, error) {
	return nil, nil
}

//
// SendVenue - Use this method to send information about a venue.
// https://core.telegram.org/bots/api#sendvenue
// TODO
//
func (b *Bot) SendVenue() (*Message, error) {
	return nil, nil
}

//
// SendContact - Use this method to send phone contacts.
// https://core.telegram.org/bots/api#sendcontact
// TODO
//
func (b *Bot) SendContact() (*Message, error) {
	return nil, nil
}

//
// SendPoll - Use this method to send a native poll.
// https://core.telegram.org/bots/api#sendpoll
// TODO
//
func (b *Bot) SendPoll() (*Message, error) {
	return nil, nil
}

//
// SendChatAction - Use this method when you need to tell the user that something is happening on the bot's side.
// The status is set for 5 seconds or less (when a message arrives from your bot, Telegram clients clear its typing status).
// https://core.telegram.org/bots/api#sendchataction
// TODO
//
func (b *Bot) SendChatAction() (success bool) {
	return false
}

//
// GetUserProfilePhotos - Use this method to get a list of profile pictures for a user.
// https://core.telegram.org/bots/api#getuserprofilephotos
// TODO
//
func (b *Bot) GetUserProfilePhotos() (*UserProfilePhotos, error) {
	return nil, nil
}

//
// GetFile - Use this method to get basic info about a file and prepare it for downloading.
// For the moment, bots can download files of up to 20MB in size.
// https://core.telegram.org/bots/api#getfile
// TODO
//
func (b *Bot) GetFile() (*File, error) {
	return nil, nil
}

//
// KickChatMember - Use this method to kick a user from a group, a supergroup or a channel.
// In the case of supergroups and channels, the user will not be able to return to the group on their own using invite links, etc., unless unbanned first. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights.
// https://core.telegram.org/bots/api#kickchatmember
// TODO
//
func (b *Bot) KickChatMember() (success bool) {
	return false
}

//
// UnbanChatMember - Use this method to unban a previously kicked user in a supergroup or channel.
// The user will not return to the group or channel automatically, but will be able to join via link, etc. The bot must be an administrator for this to work.
// https://core.telegram.org/bots/api#unbanchatmember
// TODO
//
func (b *Bot) UnbanChatMember() (success bool) {
	return false
}

//
// RestrictChatMember - Use this method to restrict a user in a supergroup.
// The bot must be an administrator in the supergroup for this to work and must have the appropriate admin rights.
// https://core.telegram.org/bots/api#restrictchatmember
// TODO
//
func (b *Bot) RestrictChatMember() (success bool) {
	return false
}

//
// PromoteChatMember - Use this method to promote or demote a user in a supergroup or a channel.
// The bot must be an administrator in the chat for this to work and must have the appropriate admin rights.
// https://core.telegram.org/bots/api#promotechatmember
// TODO
//
func (b *Bot) PromoteChatMember() (success bool) {
	return false
}

//
// SetChatAdministratorCustomTitle - Use this method to set a custom title for an administrator in a supergroup promoted by the bot.
// https://core.telegram.org/bots/api#setchatadministratorcustomtitle
// TODO
//
func (b *Bot) SetChatAdministratorCustomTitle() (success bool) {
	return false
}

//
// SetChatPermissions - Use this method to set default chat permissions for all members.
// The bot must be an administrator in the group or a supergroup for this to work and must have the can_restrict_members admin rights.
// https://core.telegram.org/bots/api#setchatpermissions
// TODO
//
func (b *Bot) SetChatPermissions() (success bool) {
	return false
}

//
// ExportChatInviteLink - Use this method to generate a new invite link for a chat; any previously generated link is revoked.
// The bot must be an administrator in the chat for this to work and must have the appropriate admin rights.
// https://core.telegram.org/bots/api#exportchatinvitelink
// TODO
//
func (b *Bot) ExportChatInviteLink() (link string) {
	return ""
}

//
// SetChatPhoto - Use this method to set a new profile photo for the chat.
// Photos can't be changed for private chats. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights.
// https://core.telegram.org/bots/api#setchatphoto
// TODO
//
func (b *Bot) SetChatPhoto() (success bool) {
	return false
}

//
// DeleteChatPhoto - Use this method to delete a chat photo.
// Photos can't be changed for private chats. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights.
// https://core.telegram.org/bots/api#deletechatphoto
// TODO
//
func (b *Bot) DeleteChatPhoto() (success bool) {
	return false
}

//
// SetChatTitle - Use this method to change the title of a chat.
// Titles can't be changed for private chats. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights.
// https://core.telegram.org/bots/api#setchattitle
// TODO
//
func (b *Bot) SetChatTitle() (success bool) {
	return false
}

//
// SetChatDescription - Use this method to change the description of a group, a supergroup or a channel.
// The bot must be an administrator in the chat for this to work and must have the appropriate admin rights.
// https://core.telegram.org/bots/api#setchatdescription
// TODO
//
func (b *Bot) SetChatDescription() (success bool) {
	return false
}

//
// PinChatMessage - Use this method to pin a message in a group, a supergroup, or a channel.
// The bot must be an administrator in the chat for this to work and must have the ‘can_pin_messages’ admin right in the supergroup or ‘can_edit_messages’ admin right in the channel.
// https://core.telegram.org/bots/api#pinchatmessage
// TODO
//
func (b *Bot) PinChatMessage() (success bool) {
	return false
}

//
// UnpinChatMessage - Use this method to unpin a message in a group, a supergroup, or a channel.
// The bot must be an administrator in the chat for this to work and must have the ‘can_pin_messages’ admin right in the supergroup or ‘can_edit_messages’ admin right in the channel.
// https://core.telegram.org/bots/api#unpinchatmessage
// TODO
//
func (b *Bot) UnpinChatMessage() (success bool) {
	return false
}

//
// LeaveChat - Use this method for your bot to leave a group, supergroup or channel.
// https://core.telegram.org/bots/api#leavechat
// TODO
//
func (b *Bot) LeaveChat() (success bool) {
	return false
}

//
// GetChat - Use this method to get up to date information about the chat (current name of the user for one-on-one conversations, current username of a user, group or channel, etc.).
// https://core.telegram.org/bots/api#getchat
//
func (b *Bot) GetChat(chatID int) (*Chat, error) {
	params := map[string]string{
		"chat_id": strconv.Itoa(chatID),
	}

	result, err := b.call("getChat", params)
	if err != nil {
		return nil, err
	}
	return result.(*Chat), nil
}

//
// GetChatAdministrators - Use this method to get a list of administrators in a chat.
// https://core.telegram.org/bots/api#getchatadministrators
//
func (b *Bot) GetChatAdministrators(chatID int) ([]*ChatMember, error) {
	params := map[string]string{
		"chat_id": strconv.Itoa(chatID),
	}
	result, err := b.call("getChatAdministrators", params)
	if err != nil {
		return nil, err
	}

	var admins []*ChatMember
	for _, v := range *result.(*[]ChatMember) {
		admins = append(admins, &v)
	}

	return admins, nil
}

//
// GetChatMembersCount - Use this method to get the number of members in a chat.
// https://core.telegram.org/bots/api#getchatmemberscount
//
func (b *Bot) GetChatMembersCount(chatID int) (int, error) {
	params := map[string]string{
		"chat_id": strconv.Itoa(chatID),
	}
	result, err := b.call("getChatMembersCount", params)
	if err != nil {
		return 0, err
	}
	res := result.([]byte)
	cnt, err := strconv.Atoi(string(res))
	if err != nil {
		return 0, err
	}
	return cnt, nil
}

//
// GetChatMember - Use this method to get information about a member of a chat.
// https://core.telegram.org/bots/api#getchatmember
// TODO
//
func (b *Bot) GetChatMember() (*ChatMember, error) {
	return nil, nil
}

//
// SetChatStickerSet - Use this method to set a new group sticker set for a supergroup.
// The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Use the field can_set_sticker_set optionally returned in getChat requests to check if the bot can use this method.
// https://core.telegram.org/bots/api#setchatstickerset
// TODO
//
func (b *Bot) SetChatStickerSet() (success bool) {
	return false
}

//
// DeleteChatStickerSet - Use this method to delete a group sticker set from a supergroup.
// The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Use the field can_set_sticker_set optionally returned in getChat requests to check if the bot can use this method.
// https://core.telegram.org/bots/api#deletechatstickerset
// TODO
//
func (b *Bot) DeleteChatStickerSet() (success bool) {
	return false
}

//
// AnswerCallbackQuery - Use this method to send answers to callback queries sent from inline keyboards.
// The answer will be displayed to the user as a notification at the top of the chat screen or as an alert.
// https://core.telegram.org/bots/api#answercallbackquery
// TODO
//
func (b *Bot) AnswerCallbackQuery() (success bool) {
	return false
}

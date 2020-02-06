package telegram

// GetMe - A simple method for testing your bot's auth token.
// https://core.telegram.org/bots/api#getme
// TODO
//
func GetMe() (*User, error) {
	return nil, nil
}

//
// SendMessage - Use this method to send text messages.
// https://core.telegram.org/bots/api#sendmessage
// TODO
//
func SendMessage() (*Message, error) {
	return nil, nil
}

//
// ForwardMessage - Use this method to forward messages of any kind.
// https://core.telegram.org/bots/api#forwardmessage
// TODO
//
func ForwardMessage() (*Message, error) {
	return nil, nil
}

//
// SendPhoto - Use this method to send photos.
// https://core.telegram.org/bots/api#sendphoto
// TODO
//
func SendPhoto() (*Message, error) {
	return nil, nil
}

//
// SendAudio - Use this method to send audio files, if you want Telegram clients to display them in the music player.
// Your audio must be in the .MP3 or .M4A format. Bots can currently send audio files of up to 50 MB in size, this limit may be changed in the future.
// For sending voice messages, use the sendVoice method instead.
// https://core.telegram.org/bots/api#sendaudio
// TODO
//
func SendAudio() (*Message, error) {
	return nil, nil
}

//
// SendDocument - Use this method to send general files.
// Bots can currently send files of any type of up to 50 MB in size, this limit may be changed in the future.
// https://core.telegram.org/bots/api#senddocument
// TODO
//
func SendDocument() (*Message, error) {
	return nil, nil
}

//
// SendVideo - Use this method to send video files, Telegram clients support mp4 videos (other formats may be sent as Document).
// Bots can currently send files of any type of up to 50 MB in size, this limit may be changed in the future.
// https://core.telegram.org/bots/api#sendvideo
// TODO
//
func SendVideo() (*Message, error) {
	return nil, nil
}

//
// SendAnimation - Use this method to send animation files (GIF or H.264/MPEG-4 AVC video without sound).
// Bots can currently send animation files of up to 50 MB in size, this limit may be changed in the future.
// https://core.telegram.org/bots/api#sendanimation
// TODO
//
func SendAnimation() (*Message, error) {
	return nil, nil
}

//
// SendVoice - Use this method to send audio files, if you want Telegram clients to display the file as a playable voice message.
// For this to work, your audio must be in an .ogg file encoded with OPUS (other formats may be sent as Audio or Document).
// Bots can currently send voice messages of up to 50 MB in size, this limit may be changed in the future.
// https://core.telegram.org/bots/api#sendvoice
// TODO
//
func SendVoice() (*Message, error) {
	return nil, nil
}

//
// SendVideoNote - Use this method to send video messages.
// https://core.telegram.org/bots/api#sendvideonote
// TODO
//
func SendVideoNote() (*Message, error) {
	return nil, nil
}

//
// SendMediaGroup - Use this method to send a group of photos or videos as an album.
// https://core.telegram.org/bots/api#sendmediagroup
// TODO
//
func SendMediaGroup() (*Message, error) {
	return nil, nil
}

//
// SendLocation - Use this method to send point on the map.
// https://core.telegram.org/bots/api#sendlocation
// TODO
//
func SendLocation() (*Message, error) {
	return nil, nil
}

//
// EditMessageLiveLocation - Use this method to edit live location messages.
// A location can be edited until its live_period expires or editing is explicitly disabled by a call to stopMessageLiveLocation.
// https://core.telegram.org/bots/api#editmessagelivelocation
// TODO
//
func EditMessageLiveLocation() (*Message, error) {
	return nil, nil
}

//
// StopMessageLiveLocation - Use this method to stop updating a live location message before live_period expires.
// https://core.telegram.org/bots/api#stopmessagelivelocation
// TODO
//
func StopMessageLiveLocation() (*Message, error) {
	return nil, nil
}

//
// SendVenue - Use this method to send information about a venue.
// https://core.telegram.org/bots/api#sendvenue
// TODO
//
func SendVenue() (*Message, error) {
	return nil, nil
}

//
// SendContact - Use this method to send phone contacts.
// https://core.telegram.org/bots/api#sendcontact
// TODO
//
func SendContact() (*Message, error) {
	return nil, nil
}

//
// SendPoll - Use this method to send a native poll.
// https://core.telegram.org/bots/api#sendpoll
// TODO
//
func SendPoll() (*Message, error) {
	return nil, nil
}

//
// SendChatAction - Use this method when you need to tell the user that something is happening on the bot's side.
// The status is set for 5 seconds or less (when a message arrives from your bot, Telegram clients clear its typing status).
// https://core.telegram.org/bots/api#sendchataction
// TODO
//
func SendChatAction() (success bool) {
	return false
}

//
// GetUserProfilePhotos - Use this method to get a list of profile pictures for a user.
// https://core.telegram.org/bots/api#getuserprofilephotos
// TODO
//
func GetUserProfilePhotos() (*UserProfilePhotos, error) {
	return nil, nil
}

//
// GetFile - Use this method to get basic info about a file and prepare it for downloading.
// For the moment, bots can download files of up to 20MB in size.
// https://core.telegram.org/bots/api#getfile
// TODO
//
func GetFile() (*File, error) {
	return nil, nil
}

//
// KickChatMember - Use this method to kick a user from a group, a supergroup or a channel.
// In the case of supergroups and channels, the user will not be able to return to the group on their own using invite links, etc., unless unbanned first. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights.
// https://core.telegram.org/bots/api#kickchatmember
// TODO
//
func KickChatMember() (success bool) {
	return false
}

//
// UnbanChatMember - Use this method to unban a previously kicked user in a supergroup or channel.
// The user will not return to the group or channel automatically, but will be able to join via link, etc. The bot must be an administrator for this to work.
// https://core.telegram.org/bots/api#unbanchatmember
// TODO
//
func UnbanChatMember() (success bool) {
	return false
}

//
// RestrictChatMember - Use this method to restrict a user in a supergroup.
// The bot must be an administrator in the supergroup for this to work and must have the appropriate admin rights.
// https://core.telegram.org/bots/api#restrictchatmember
// TODO
//
func RestrictChatMember() (success bool) {
	return false
}

//
// PromoteChatMember - Use this method to promote or demote a user in a supergroup or a channel.
// The bot must be an administrator in the chat for this to work and must have the appropriate admin rights.
// https://core.telegram.org/bots/api#promotechatmember
// TODO
//
func PromoteChatMember() (success bool) {
	return false
}

//
// SetChatAdministratorCustomTitle - Use this method to set a custom title for an administrator in a supergroup promoted by the bot.
// https://core.telegram.org/bots/api#setchatadministratorcustomtitle
// TODO
//
func SetChatAdministratorCustomTitle() (success bool) {
	return false
}

//
// SetChatPermissions - Use this method to set default chat permissions for all members.
// The bot must be an administrator in the group or a supergroup for this to work and must have the can_restrict_members admin rights.
// https://core.telegram.org/bots/api#setchatpermissions
// TODO
//
func SetChatPermissions() (success bool) {
	return false
}

//
// ExportChatInviteLink - Use this method to generate a new invite link for a chat; any previously generated link is revoked.
// The bot must be an administrator in the chat for this to work and must have the appropriate admin rights.
// https://core.telegram.org/bots/api#exportchatinvitelink
// TODO
//
func ExportChatInviteLink() (link string) {
	return ""
}

//
// SetChatPhoto - Use this method to set a new profile photo for the chat.
// Photos can't be changed for private chats. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights.
// https://core.telegram.org/bots/api#setchatphoto
// TODO
//
func SetChatPhoto() (success bool) {
	return false
}

//
// DeleteChatPhoto - Use this method to delete a chat photo.
// Photos can't be changed for private chats. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights.
// https://core.telegram.org/bots/api#deletechatphoto
// TODO
//
func DeleteChatPhoto() (success bool) {
	return false
}

//
// SetChatTitle - Use this method to change the title of a chat.
// Titles can't be changed for private chats. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights.
// https://core.telegram.org/bots/api#setchattitle
// TODO
//
func SetChatTitle() (success bool) {
	return false
}

//
// SetChatDescription - Use this method to change the description of a group, a supergroup or a channel.
// The bot must be an administrator in the chat for this to work and must have the appropriate admin rights.
// https://core.telegram.org/bots/api#setchatdescription
// TODO
//
func SetChatDescription() (success bool) {
	return false
}

//
// PinChatMessage - Use this method to pin a message in a group, a supergroup, or a channel.
// The bot must be an administrator in the chat for this to work and must have the ‘can_pin_messages’ admin right in the supergroup or ‘can_edit_messages’ admin right in the channel.
// https://core.telegram.org/bots/api#pinchatmessage
// TODO
//
func PinChatMessage() (success bool) {
	return false
}

//
// UnpinChatMessage - Use this method to unpin a message in a group, a supergroup, or a channel.
// The bot must be an administrator in the chat for this to work and must have the ‘can_pin_messages’ admin right in the supergroup or ‘can_edit_messages’ admin right in the channel.
// https://core.telegram.org/bots/api#unpinchatmessage
// TODO
//
func UnpinChatMessage() (success bool) {
	return false
}

//
// LeaveChat - Use this method for your bot to leave a group, supergroup or channel.
// https://core.telegram.org/bots/api#leavechat
// TODO
//
func LeaveChat() (success bool) {
	return false
}

//
// GetChat - Use this method to get up to date information about the chat (current name of the user for one-on-one conversations, current username of a user, group or channel, etc.).
// https://core.telegram.org/bots/api#getchat
// TODO
//
func GetChat() (*Chat, error) {
	return nil, nil
}

//
// GetChatAdministrators - Use this method to get a list of administrators in a chat.
// https://core.telegram.org/bots/api#getchatadministrators
// TODO
//
func GetChatAdministrators() ([]*ChatMember, error) {
	return nil, nil
}

//
// GetChatMembersCount - Use this method to get the number of members in a chat.
// https://core.telegram.org/bots/api#getchatmemberscount
// TODO
//
func GetChatMembersCount() int {
	return 0
}

//
// GetChatMember - Use this method to get information about a member of a chat.
// https://core.telegram.org/bots/api#getchatmember
// TODO
//
func GetChatMember() (*ChatMember, error) {
	return nil, nil
}

//
// SetChatStickerSet - Use this method to set a new group sticker set for a supergroup.
// The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Use the field can_set_sticker_set optionally returned in getChat requests to check if the bot can use this method.
// https://core.telegram.org/bots/api#setchatstickerset
// TODO
//
func SetChatStickerSet() (success bool) {
	return false
}

//
// DeleteChatStickerSet - Use this method to delete a group sticker set from a supergroup.
// The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Use the field can_set_sticker_set optionally returned in getChat requests to check if the bot can use this method.
// https://core.telegram.org/bots/api#deletechatstickerset
// TODO
//
func DeleteChatStickerSet() (success bool) {
	return false
}

//
// AnswerCallbackQuery - Use this method to send answers to callback queries sent from inline keyboards.
// The answer will be displayed to the user as a notification at the top of the chat screen or as an alert.
// https://core.telegram.org/bots/api#answercallbackquery
// TODO
//
func AnswerCallbackQuery() (success bool) {
	return false
}

## Telegram Bot API

### Work in progress

### Usage
```
bot, err := telegram.New("123456:ABC-DEF1234ghIkl-zyx57W2v1u123ew1")
user, err := bot.GetMe()
```

### Supported API Calls
##### Update
- [getUpdates](https://core.telegram.org/bots/api#getupdates)
- [deleteWebhook](https://core.telegram.org/bots/api#deletewebhook)
- [getWebhookInfo](https://core.telegram.org/bots/api#getwebhookinfo)
##### Methods
- [getMe](https://core.telegram.org/bots/api#getme)
- [sendMessage](https://core.telegram.org/bots/api#sendmessage)
- [forwardMessage](https://core.telegram.org/bots/api#forwardmessage)
- [sendPoll](https://core.telegram.org/bots/api#sendpoll)
- [getUserProfilePhotos](https://core.telegram.org/bots/api#getuserprofilephotos)
- [getChat](https://core.telegram.org/bots/api#getchat)
- [getChatAdministrators](https://core.telegram.org/bots/api#getchatadministrators)
- [getChatMembersCount](https://core.telegram.org/bots/api#getchatmemberscount)
- [getChatMember](https://core.telegram.org/bots/api#getchatmember)
- [getFile](https://core.telegram.org/bots/api#getfile)
- [setChatTitle](https://core.telegram.org/bots/api#setchattitle)
- [setChatDescription](https://core.telegram.org/bots/api#setchatdescription)
- [setChatAdministratorCustomTitle](https://core.telegram.org/bots/api#setchatadministratorcustomtitle)
- [kickChatMember](https://core.telegram.org/bots/api#kickchatmember)
- [unbanChatMember](https://core.telegram.org/bots/api#unbanchatmember)
- [deleteChatPhoto](https://core.telegram.org/bots/api#deletechatphoto)
- [leaveChat](https://core.telegram.org/bots/api#leavechat)
- [setChatStickerSet](https://core.telegram.org/bots/api#setchatstickerset)
- [deleteChatStickerSet](https://core.telegram.org/bots/api#deletechatstickerset)
- [exportChatInviteLink](https://core.telegram.org/bots/api#exportchatinvitelink)
- [pinChatMessage](https://core.telegram.org/bots/api#pinchatmessage)
- [unpinChatMessage](https://core.telegram.org/bots/api#unpinchatmessage)
- [sendChatAction](https://core.telegram.org/bots/api#sendchataction)
- [sendLocation](https://core.telegram.org/bots/api#sendlocation)
##### Sticker
- [getStickerSet](https://core.telegram.org/bots/api#getstickerset)
- [deleteStickerFromSet](https://core.telegram.org/bots/api#deletestickerfromset)
- [setStickerPositionInSet](https://core.telegram.org/bots/api#setstickerpositioninset)
##### Game
##### Passport
##### Payment
- [sendInvoice](https://core.telegram.org/bots/api#sendinvoice)
- [answerShippingQuery](https://core.telegram.org/bots/api#answershippingquery)
- [answerPreCheckoutQuery](https://core.telegram.org/bots/api#answerprecheckoutquery)

### Todo
##### Update
- [setWebhook](https://core.telegram.org/bots/api#setwebhook)
##### Methods
- [sendPhoto](https://core.telegram.org/bots/api#sendphoto)
- [sendAudio](https://core.telegram.org/bots/api#sendaudio)
- [sendDocument](https://core.telegram.org/bots/api#senddocument)
- [sendVideo](https://core.telegram.org/bots/api#sendvideo)
- [sendAnimation](https://core.telegram.org/bots/api#sendanimation)
- [sendVoice](https://core.telegram.org/bots/api#sendvoice)
- [sendVideoNote](https://core.telegram.org/bots/api#sendvideonote)
- [sendMediaGroup](https://core.telegram.org/bots/api#sendmediagroup)
- [editMessageLiveLocation](https://core.telegram.org/bots/api#editmessagelivelocation)
- [stopMessageLiveLocation](https://core.telegram.org/bots/api#stopmessagelivelocation)
- [sendVenue](https://core.telegram.org/bots/api#sendvenue)
- [sendContact](https://core.telegram.org/bots/api#sendcontact)
- [restrictChatMember](https://core.telegram.org/bots/api#restrictchatmember)
- [promoteChatMember](https://core.telegram.org/bots/api#promotechatmember)
- [setChatPermissions](https://core.telegram.org/bots/api#setchatpermissions)
- [setChatPhoto](https://core.telegram.org/bots/api#setchatphoto)
- [answerCallbackQuery](https://core.telegram.org/bots/api#answercallbackquery)
##### Sticker
- [sendSticker](https://core.telegram.org/bots/api#sendsticker)
- [uploadStickerFile](https://core.telegram.org/bots/api#uploadstickerfile)
- [createNewStickerSet](https://core.telegram.org/bots/api#createnewstickerset)
- [addStickerToSet](https://core.telegram.org/bots/api#addstickertoset)
##### Game
- [sendGame](https://core.telegram.org/bots/api#sendgame)
- [setGameScore](https://core.telegram.org/bots/api#setgamescore)
- [getGameHighScores](https://core.telegram.org/bots/api#getgamehighscores)
##### Passport
- [setPassportDataErrors](https://core.telegram.org/bots/api#setpassportdataerrors)

## Telegram Bot API

### Work in progress

### Usage
```
bot, err := telegram.New("123456:ABC-DEF1234ghIkl-zyx57W2v1u123ew1")
user, err := bot.GetMe()
```

#### Supported API Calls (more to come)
- [getMe](https://core.telegram.org/bots/api#getme)
- [getUpdates](https://core.telegram.org/bots/api#getupdates)
- [getChat](https://core.telegram.org/bots/api#getchat)
- [getChatAdministrators](https://core.telegram.org/bots/api#getchatadministrators)
- [getChatMembersCount](https://core.telegram.org/bots/api#getchatmemberscount)
- [sendMessage](https://core.telegram.org/bots/api#sendmessage)
- [sendPoll](https://core.telegram.org/bots/api#sendpoll)
- [setChatTitle](https://core.telegram.org/bots/api#setchattitle)
- [setChatDescription](https://core.telegram.org/bots/api#setchatdescription)
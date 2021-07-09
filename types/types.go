package types

import tb "gopkg.in/tucnak/telebot.v2"

type HelperConfig struct {
	HelperType       string `yaml:"helper_type"`
	AlarmUrl         string `yaml:"alarm_url"`
	ExtraHeader      string `yaml:"extra_header"`
	ExtraHeaderValue string `yaml:"extra_header_value"`
}

type TelegramBotConfig struct {
	BotToken     string `yaml:"bot_token"`
	AllowedUsers []int  `yaml:"allowed_users"`
}

func (botConfig *TelegramBotConfig) CanAccessBot(b *tb.Bot, sender *tb.User) bool {
	for _, val := range botConfig.AllowedUsers {
		if val == sender.ID {
			return true
		}
	}
	b.Send(sender, "You cannot acces the bot!")
	return false
}

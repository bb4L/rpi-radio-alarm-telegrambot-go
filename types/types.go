package types

import tb "gopkg.in/tucnak/telebot.v2"

// HelperConfig this struct contains the information to configure the helper used by the bot
type HelperConfig struct {
	HelperType       string `yaml:"helper_type"`
	AlarmURL         string `yaml:"alarm_url"`
	ExtraHeader      string `yaml:"extra_header"`
	ExtraHeaderValue string `yaml:"extra_header_value"`
}

// TelegramBotConfig configuration for the telegram bot
type TelegramBotConfig struct {
	BotToken     string `yaml:"bot_token"`
	AllowedUsers []int  `yaml:"allowed_users"`
}

// CanAccessBot checks and returns if the user is allowed to access the bot and sends a notification if the user isn't allowed
func (botConfig *TelegramBotConfig) CanAccessBot(b *tb.Bot, sender *tb.User) bool {
	for _, val := range botConfig.AllowedUsers {
		if val == sender.ID {
			return true
		}
	}
	b.Send(sender, "You cannot acces the bot!")
	return false
}

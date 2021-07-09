package handlers

import (
	"encoding/json"

	"github.com/bb4L/rpi-radio-alarm-telegrambot-go/types"

	libraryTypes "github.com/bb4L/rpi-radio-alarm-go-library/types"
	tb "gopkg.in/tucnak/telebot.v2"
)

// AddRadioHandles adds the handles for the radio commands
func AddRadioHandles(b *tb.Bot, botConfig types.TelegramBotConfig, handler libraryTypes.DataHandler) {
	b.Handle("/start_radio", func(m *tb.Message) {
		if !botConfig.CanAccessBot(b, m.Sender) {
			return
		}

		radio, err := handler.SwitchRadio(true)
		if err != nil {
			b.Send(m.Sender, err.Error())
			return
		}

		stringData, _ := json.Marshal(radio)

		logger.Println("radio started")

		b.Send(m.Sender, string(stringData))
	})

	b.Handle("/stop_radio", func(m *tb.Message) {
		if !botConfig.CanAccessBot(b, m.Sender) {
			return
		}

		radio, err := handler.SwitchRadio(false)
		if err != nil {
			b.Send(m.Sender, err.Error())
			return
		}

		stringData, _ := json.Marshal(radio)

		logger.Println("radio stopped")

		b.Send(m.Sender, string(stringData))
	})
}

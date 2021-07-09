package handlers

import (
	"os"
	"strconv"

	"github.com/bb4L/rpi-radio-alarm-telegrambot-go/constants"
	"github.com/bb4L/rpi-radio-alarm-telegrambot-go/types"

	"github.com/bb4L/rpi-radio-alarm-go-library/logging"
	libraryTypes "github.com/bb4L/rpi-radio-alarm-go-library/types"
	tb "gopkg.in/tucnak/telebot.v2"
)

var dataHandler libraryTypes.DataHandler
var logger = logging.GetLogger(os.Stdout, constants.DefaultPrefix, "handles")

// AddDefaultHandles adds some default handles
func AddDefaultHandles(b *tb.Bot, botConfig types.TelegramBotConfig, handler libraryTypes.DataHandler) {
	dataHandler = handler
	b.Handle(tb.OnText, func(m *tb.Message) {
		b.Send(m.Sender, "Unknown command: "+m.Text)
	})

	b.Handle("/start", func(m *tb.Message) {
		if !botConfig.CanAccessBot(b, m.Sender) {
			return
		}

		r, errString := getAllBtnsMarkup()
		if len(errString) > 0 {
			b.Send(m.Sender, errString)
			return
		}
		logger.Println("returning alarms markup from command /start")
		b.Send(m.Sender, "Got Alarms", r)
	})
}

func getAllBtnsMarkup() (*tb.ReplyMarkup, string) {
	alarms, err := dataHandler.GetAlarms(false)

	if err != nil {
		return nil, "error while getting alarms"
	}
	r := &tb.ReplyMarkup{}

	r.InlineKeyboard = append(r.InlineKeyboard, []tb.InlineButton{{Unique: getAllBtns.Unique, Text: "Get Alarms"}})

	for idx, alarm := range alarms {
		btnText := "turn on"
		if alarm.Active {
			btnText = "turn off"
		}
		button := []tb.InlineButton{{Unique: getAlarmBtn.Unique, Text: "Alarm: " + alarm.Name, Data: strconv.Itoa(idx)}, {Unique: switchAlarmBtn.Unique, Text: btnText, Data: strconv.Itoa(idx)}}
		r.InlineKeyboard = append(r.InlineKeyboard, button)
	}
	return r, ""
}

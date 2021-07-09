package handlers

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/bb4L/rpi-radio-alarm-telegrambot-go/types"

	libraryTypes "github.com/bb4L/rpi-radio-alarm-go-library/types"

	tb "gopkg.in/tucnak/telebot.v2"
)

var (
	getAlarmBtn    = &tb.InlineButton{Unique: "getAlarm"}
	switchAlarmBtn = &tb.InlineButton{Unique: "switchAlarm"}
	getAllBtns     = &tb.InlineButton{Unique: "getAllBtns"}
)

// AddAlarmButtonHandles adds the handles for the alarm buttons
func AddAlarmButtonHandles(b *tb.Bot, botConfig types.TelegramBotConfig, handler libraryTypes.DataHandler) {
	b.Handle(getAllBtns, func(c *tb.Callback) {
		if !botConfig.CanAccessBot(b, c.Sender) {
			return
		}
		r, errString := getAllBtnsMarkup()
		if len(errString) > 0 {
			b.Send(c.Sender, errString)
			return
		}
		logger.Println("returning alarms markup from button click")
		b.Send(c.Sender, "Got Alarms", r)
	})

	b.Handle(getAlarmBtn, func(c *tb.Callback) {
		if !botConfig.CanAccessBot(b, c.Sender) {
			return
		}

		splittedData := strings.Split(c.Data, " ")
		idx, err := strconv.Atoi(splittedData[0])
		if err != nil {
			b.Send(c.Sender, "could not get index")
			return
		}

		alarm, err := handler.GetAlarm(idx, false)
		if err != nil {
			b.Send(c.Sender, err.Error())
			return
		}

		stringData, _ := json.Marshal(alarm)
		logger.Println("returning alarm: " + string(stringData))
		b.Send(c.Sender, string(stringData))
	})

	b.Handle(switchAlarmBtn, func(c *tb.Callback) {
		if !botConfig.CanAccessBot(b, c.Sender) {
			return
		}

		input := strings.Split(c.Data, " ")
		if len(input) != 3 {
			b.Send(c.Sender, "got invalid argument")
			return
		}

		idx, err := strconv.Atoi(input[2])
		if err != nil {
			b.Send(c.Sender, "could not get index")
			return
		}

		alarm, _ := handler.GetAlarm(idx, true)

		if !((input[1] == "off" && alarm.Active) || (input[1] == "on" && !alarm.Active)) {
			b.Send(c.Sender, "alarm is already in correct state")
			return
		}

		alarm.Active = !alarm.Active
		savedAlarm, err := handler.SaveAlarm(idx, alarm)
		if err != nil {
			b.Send(c.Sender, err.Error())
			return
		}

		stringData, _ := json.Marshal(savedAlarm)

		logger.Println("switch alarm: " + string(stringData))

		b.Send(c.Sender, string(stringData))
	})
}

// AddAlarmHandles adds the handles for the alarm related commands
func AddAlarmHandles(b *tb.Bot, botConfig types.TelegramBotConfig, handler libraryTypes.DataHandler) {
	b.Handle("/add_alarm", func(m *tb.Message) {
		if !botConfig.CanAccessBot(b, m.Sender) {
			return
		}
		splittedPayload := strings.Split(m.Payload, " ")
		alarm := libraryTypes.Alarm{}
		err := handleSplittedData(0, splittedPayload, &alarm)
		if err != nil {
			b.Send(m.Sender, err.Error())
			return
		}

		if len(alarm.Days) == 0 || alarm.Name == "" {
			b.Send(m.Sender, "not enough arguments; either no days specified or no name given")
			return
		}

		alarms, err := handler.AddAlarm(alarm)
		if err != nil {
			b.Send(m.Sender, err.Error())
			return
		}

		stringData, _ := json.Marshal(alarms)
		logger.Println("added an alarm")
		b.Send(m.Sender, string(stringData))
	})

	b.Handle("/change_alarm", func(m *tb.Message) {
		if !botConfig.CanAccessBot(b, m.Sender) {
			return
		}

		splittedPayload := strings.Split(m.Payload, " ")

		// handle idx
		idx, err := strconv.Atoi(splittedPayload[0])
		if err != nil {
			b.Send(m.Sender, err.Error())
			return
		}

		// handle the alarm
		alarm, err := handler.GetAlarm(idx, true)
		if err != nil {
			b.Send(m.Sender, err)
			return
		}

		err = handleSplittedData(1, splittedPayload, &alarm)
		if err != nil {
			b.Send(m.Sender, err.Error())
			return
		}

		alarm, err = handler.SaveAlarm(idx, alarm)

		if err != nil {
			b.Send(m.Sender, err.Error())
			return
		}

		stringData, _ := json.Marshal(alarm)
		logger.Println("changed alarm, is now: " + string(stringData))
		b.Send(m.Sender, string(stringData))
	})

	b.Handle("/delete_alarm", func(m *tb.Message) {
		if !botConfig.CanAccessBot(b, m.Sender) {
			return
		}

		splittedPayload := strings.Split(m.Payload, " ")

		// handle idx
		idx, err := strconv.Atoi(splittedPayload[0])
		if err != nil {
			b.Send(m.Sender, err.Error())
			return
		}

		alarms, err := handler.DeleteAlarm(idx)
		if err != nil {
			b.Send(m.Sender, err.Error())
			return
		}

		stringData, _ := json.Marshal(alarms)

		logger.Printf("deleted alarm at idx=%d\n", idx)
		b.Send(m.Sender, string(stringData))
	})

}

func handleSplittedData(startIdx int, splittedPayload []string, alarm *libraryTypes.Alarm) error {
	for i := startIdx; i <= len(splittedPayload)-1; i += 2 {
		// handle the values
		if len(splittedPayload) == i {
			return fmt.Errorf("invalid number of arguments")
		}

		switch splittedPayload[i] {
		case "name":
			alarm.Name = splittedPayload[i+1]

		case "hour":
			hour, err := strconv.Atoi(splittedPayload[i+1])
			if err != nil {
				return err
			}
			alarm.Hour = hour

		case "min":
			hour, err := strconv.Atoi(splittedPayload[i+1])
			if err != nil {
				return err
			}
			alarm.Minute = hour

		case "days":
			days := strings.Split(splittedPayload[i+1], ",")
			daysNumber := []int{}
			for _, day := range days {
				d, err := strconv.Atoi(day)
				if err != nil {
					return err
				}
				daysNumber = append(daysNumber, d)
			}
			alarm.Days = daysNumber

		case "on":
			if splittedPayload[i+1] == "true" {
				alarm.Active = true
			} else if splittedPayload[i+1] == "false" {
				alarm.Active = false
			} else {
				return fmt.Errorf("invalid value for 'on' %s", splittedPayload[i+1])
			}
		default:
			return fmt.Errorf("invalid value passed %s", splittedPayload[i])
		}
	}
	return nil
}

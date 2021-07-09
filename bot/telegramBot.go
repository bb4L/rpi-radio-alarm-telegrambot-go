package bot

import (
	"os"
	"rpi-radio-alarm-telegrambot-go/constants"
	"rpi-radio-alarm-telegrambot-go/handlers"
	"rpi-radio-alarm-telegrambot-go/types"
	"time"

	"github.com/bb4L/rpi-radio-alarm-go-library/logging"
	libraryTypes "github.com/bb4L/rpi-radio-alarm-go-library/types"
	tb "gopkg.in/tucnak/telebot.v2"
	"gopkg.in/yaml.v2"
)

const telegramConfigFile = "./rpi_telegrambot_config.yaml"

var logger = logging.GetLogger(os.Stdout, constants.DefaultPrefix, "bot")
var botConfig types.TelegramBotConfig

// StartTelegramBot starts the telegrambot with the given DataHandler
func StartTelegramBot(handler libraryTypes.DataHandler) {
	logger.Println("preparing telegrambot")
	var err error
	botConfig, err = getTelegramConfig()
	if err != nil {
		logger.Println("could not get config:", err.Error())
		return
	}

	b, err := tb.NewBot(tb.Settings{
		Token:  botConfig.BotToken,
		Poller: &tb.LongPoller{Timeout: 2 * time.Second},
	})

	if err != nil {
		logger.Fatal(err)
		return
	}

	handlers.AddDefaultHandles(b, botConfig, handler)
	handlers.AddAlarmHandles(b, botConfig, handler)
	handlers.AddRadioHandles(b, botConfig, handler)

	logger.Println("starting telegrambot")
	b.Start()
}

func getTelegramConfig() (types.TelegramBotConfig, error) {
	fileData, err := os.ReadFile(telegramConfigFile)

	if err != nil {
		if os.IsNotExist(err) {
			os.Create(telegramConfigFile)
			fileData, _ = os.ReadFile(telegramConfigFile)
		} else {
			return types.TelegramBotConfig{}, err
		}
	}

	var data types.TelegramBotConfig
	source := []byte(fileData)
	err = yaml.Unmarshal(source, &data)
	if err != nil {
		logger.Printf("error: %v\n", err)
	}
	return data, err
}

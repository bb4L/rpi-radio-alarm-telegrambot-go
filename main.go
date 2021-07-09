package main

import (
	"flag"
	"io/ioutil"
	"os"

	"github.com/bb4L/rpi-radio-alarm-go-library/api"
	"github.com/bb4L/rpi-radio-alarm-go-library/logging"
	"github.com/bb4L/rpi-radio-alarm-go-library/storage"
	"github.com/bb4L/rpi-radio-alarm-telegrambot-go/bot"
	"github.com/bb4L/rpi-radio-alarm-telegrambot-go/types"
	"gopkg.in/yaml.v2"
)

var logger = logging.GetLogger(os.Stdout, "main")

func main() {
	logger.Println("start telegram bot")
	var configFilePath string
	flag.StringVar(&configFilePath, "c", "./rpi_telegrambot_helper_config.yaml", "Specify config file for the DataHelper used by the bot")

	logger.Println("reading HelperConfig from: " + configFilePath)
	fileData, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		panic(err)
	}

	var helperConfig types.HelperConfig
	source := []byte(fileData)
	err = yaml.Unmarshal(source, &helperConfig)
	if err != nil {
		logger.Fatalf("error: %v", err)
	}

	if helperConfig.HelperType == "storage" {
		bot.StartTelegramBot(&storage.Helper{})
	} else if helperConfig.HelperType == "api" {
		if helperConfig.AlarmURL == "" {
			logger.Fatalf("no AlarmUrl set")
		}
		bot.StartTelegramBot(&api.Helper{AlarmURL: helperConfig.AlarmURL, ExtraHeader: helperConfig.ExtraHeader, ExtreaHeaderValue: helperConfig.ExtraHeaderValue})

	} else {
		logger.Fatalf("invalid config for Helpertype, was %s expected to be \"api\" or \"storage\" ", err)
	}

}

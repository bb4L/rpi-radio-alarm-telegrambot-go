# rpi-radio-alarm-telegrambot-go
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/bb4L/rpi-radio-alarm-telegrambot-go)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/bb4l/rpi-radio-alarm-telegrambot-go)
[![Go Reference](https://pkg.go.dev/badge/github.com/bb4L/rpi-radio-alarm-telegrambot-go.svg)](https://pkg.go.dev/github.com/bb4L/rpi-radio-alarm-telegrambot-go)
![GitHub](https://img.shields.io/github/license/bb4L/rpi-radio-alarm-telegrambot-go)
![GitHub Release Date](https://img.shields.io/github/release-date/bb4L/rpi-radio-alarm-telegrambot-go)
![GitHub last commit](https://img.shields.io/github/last-commit/bb4l/rpi-radio-alarm-telegrambot-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/bb4l/rpi-radio-alarm-telegrambot-go)](https://goreportcard.com/report/github.com/bb4L/rpi-radio-alarm-telegrambot-go)
![GitHub issues](https://img.shields.io/github/issues-raw/bb4l/rpi-radio-alarm-telegrambot-go)
![Lines of code](https://img.shields.io/tokei/lines/github/bb4l/rpi-radio-alarm-telegrambot-go)
[![CI](https://github.com/bb4L/rpi-radio-alarm-telegrambot-go/actions/workflows/build.yml/badge.svg?branch=main)](https://github.com/bb4L/rpi-radio-alarm-telegrambot-go/actions/workflows/build.yml)

Golang telegrambot for the [rpi-radio-alarm-go](https://github.com/bb4L/rpi-radio-alarm-go)

## configuration
The telegrambot expects to have a config files with the following structure

```yaml
# rpi_telegrambot_config.yaml
bot_token: "TESTTOKEN"
allowed_users:
  - USERID

```

```yaml
# rpi_telegrambot_helper_config.yaml
helper_type: "api" # has to be one of "api", "storage"
alarm_url: "localhost:8000"
extra_header: ""
extra_header_value: ""
```

# License
[GPLv3](LICENSE)

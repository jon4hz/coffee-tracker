package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

// Config represents the complete config
type Config struct {
	Telegram *TelegramConfig `yaml:"telegram"`
}

// TelegramConfig represents the Telegram config
type TelegramConfig struct {
	BotToken string `yaml:"bot_token"`
	OwnerID  int64  `yaml:"owner_id"`
}

var (
	config *Config
)

func init() {
	load()
}

func load() {
	file, err := ioutil.ReadFile("config.yml")
	if err != nil {
		panic(err)
	}

	if err := yaml.Unmarshal(file, &config); err != nil {
		panic(err)
	}

	validate()
}

func validate() {
	if config == nil {
		panic("config is nil")
	}
	if config.Telegram == nil {
		panic("config.telegram is nil")
	}
	if config.Telegram.BotToken == "" {
		panic("config.telegram.bot_token is empty")
	}
	if config.Telegram.OwnerID == 0 {
		panic("config.telegram.owner_id is empty")
	}
}

func Get() *Config {
	return config
}

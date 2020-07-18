package app

import (
	"fmt"

	"stock_bot/internal/configreader"
)

var (
	config *StockConfig
)

type Stock struct {
	Name 	string `mapstructure:"name"`
	Id 		string `mapstructure:"id"`
}

type EnvConfig struct {
	StockList 	[]Stock  `mapstructure:"stocks"`
	WebHookKey  string	 `mapstructure:"webhookkey"`
}

type StockConfig struct {
	EnvConfig EnvConfig `mapstructure:"env"`
}

func (c *StockConfig) String() string {
	return fmt.Sprintf("%+v", *c)
}

func initCfg(file string) error {
	config = &StockConfig{}

	return configreader.InitConfig(file, config)
}

func GetConfig() *StockConfig {
	return config
}

package configreader

import "github.com/spf13/viper"

func InitConfig(filepath string, cfg interface{}) error {
	viper.SetConfigFile(filepath)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	if err := viper.Unmarshal(cfg); err != nil {
		return err
	}

	return nil
}
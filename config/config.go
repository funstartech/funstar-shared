package config

import (
	"github.com/spf13/viper"

	"github.com/funstartech/funstar-shared/log"
)

// Init 配置初始化
func Init() error {
	viper.AddConfigPath("/app/conf")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	log.InitLog()
	return nil
}

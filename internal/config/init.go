package config

import (
	"go.uber.org/zap"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Init() error {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		zap.L().Info(
			"config file changed",
			zap.String("fileName", e.Name),
			zap.Any("operation", e.Op),
		)
	})

	return nil
}

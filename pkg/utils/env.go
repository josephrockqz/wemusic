package utils

import (
	"github.com/spf13/viper"
)

func GetEnvironmentVariable(variable string) (string, error) {
	viper.SetConfigName("local-template")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return "", err
	}

	variablePath := "environment." + variable

	value := viper.GetString(variablePath)
	return value, nil
}

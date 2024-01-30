package utils

import (
	"os"

	"github.com/spf13/viper"
)

func GetEnvironmentVariableByName(variable string) (string, error) {
	viper.SetConfigName("local-template")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		return "", err
	}

	variablePath := "environment." + variable

	value := viper.GetString(variablePath)
	if value != "" {
		return value, nil
	}

	value = os.Getenv(variable)

	return value, nil
}

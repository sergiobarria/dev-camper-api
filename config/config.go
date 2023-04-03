package config

import (
	"github.com/spf13/viper"
)

var EnvVars *Config

type Config struct {
	PORT      int    `mapstructure:"PORT"`
	HOST      string `mapstructure:"HOST"`
	GO_ENV    string `mapstructure:"GO_ENV"`
	MONGO_URI string `mapstructure:"MONGO_URI"`
	MONGO_DB  string `mapstructure:"MONGO_DB"`
	DEBUG     bool   `mapstructure:"DEBUG"`
}

func LoadEnvVars() error {
	viper.SetConfigName(".env")
	viper.AddConfigPath("./")
	viper.SetConfigType("env")

	viper.AutomaticEnv() // read in environment variables that match
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	config := new(Config)
	if err := viper.Unmarshal(config); err != nil {
		return err
	}

	EnvVars = config
	return nil
}

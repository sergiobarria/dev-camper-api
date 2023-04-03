package initializers

import (
	"github.com/spf13/viper"
)

type Config struct {
	PORT      int    `mapstructure:"PORT"`
	HOST      string `mapstructure:"HOST"`
	GO_ENV    string `mapstructure:"GO_ENV"`
	MONGO_URI string `mapstructure:"MONGO_URI"`
	DB_NAME   string `mapstructure:"DB_NAME"`
	DEBUG     bool   `mapstructure:"DEBUG"`
}

func LoadEnvVars() (config *Config, err error) {
	viper.SetConfigName(".env")
	viper.AddConfigPath("./")
	viper.SetConfigType("env")

	viper.AutomaticEnv() // read in environment variables that match
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return config, nil
}

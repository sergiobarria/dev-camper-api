package settings

import (
	"log"

	"github.com/spf13/viper"
)

var EnvConfig *envConfig

func LoadEnv() {
	EnvConfig = loadEnvVariables()
}

type envConfig struct {
	Port int `mapstructure:"PORT"`
}

func loadEnvVariables() (config *envConfig) {
	// Set the file name of the configurations file
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal(err)
	}

	return
}

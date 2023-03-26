package config

import "github.com/spf13/viper"

func LoadEnvVars() {
	viper.AddConfigPath(".")    // optionally look for config in the working directory
	viper.SetConfigFile(".env") // optionally look for config in the working directory

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	viper.AutomaticEnv()
}

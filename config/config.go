package config

import "github.com/spf13/viper"

// LoadConfig loads the config.yml file and the .env file
func LoadConfig() error {
	var err error

	// Use viper to read the .env file
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv() // read in environment variables that match

	err = viper.ReadInConfig() // Find and read the config file
	if err != nil {
		return err
	}

	// Use viper to read the config.yml file
	viper.SetConfigName("config")   // name of config file (without extension)
	viper.SetConfigType("yml")      // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")        // optionally look for config in the working directory
	viper.AddConfigPath("./config") // optionally look for config in the working directory

	err = viper.MergeInConfig()

	// Replace environment variables in config
	viper.Set("database.mongo_uri", viper.GetString("MONGO_URI"))

	return err
}

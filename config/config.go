package config

import "github.com/spf13/viper"

type Config struct {
	// Server config
	Port   string `mapstructure:"PORT"`
	DBUri  string `mapstructure:"MONGO_URI"`
	DbName string `mapstructure:"DATABASE_NAME"`

	// JWT config

	// Email config
	EmailFrom string `mapstructure:"EMAIL_FROM"`
	SMTPHost  string `mapstructure:"SMTP_HOST"`
	SMTPPass  string `mapstructure:"SMTP_PASS"`
	SMTPPort  int    `mapstructure:"SMTP_PORT"`
	SMTPUser  string `mapstructure:"SMTP_USER"`

	// More config here 👇🏼
	// see: https://github.com/wpcodevo/golang-mongodb-api/blob/golang-mongodb-crud-api/config/default.go
}

func LoadConfig(path string) (cfg Config, err error) {
	viper.AddConfigPath(path)   // optionally look for config in the working directory
	viper.SetConfigType("env")  // set the file type
	viper.SetConfigName(".env") // name of config file (without extension)

	viper.AutomaticEnv() // read in environment variables that match

	err = viper.ReadInConfig() // Find and read the config file
	if err != nil {
		return
	}

	err = viper.Unmarshal(&cfg)
	return
}

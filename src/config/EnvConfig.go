package config

import "github.com/spf13/viper"

type Config struct {
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	JwtKey        string `mapstructure:"JWT_KEY"`
	DBUri         string `mapstructure:"DB_URI"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

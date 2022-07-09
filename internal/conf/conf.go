package conf

import (
	"log"

	"github.com/spf13/viper"
)

type Configuration struct {
	JWTSecret string `mapstructure:"jwt_secret"`
}

var config *Configuration

func init() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")      // optionally look for config in the working directory

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		panic(err)
	}

	log.Println("Config:", config.JWTSecret)

	log.Println("Config loaded")
}

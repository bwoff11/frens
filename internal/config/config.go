package config

import (
	"encoding/json"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Configuration struct {
	JWTSecret string `mapstructure:"jwt_secret"`
}

var Config *Configuration = &Configuration{
	JWTSecret: "supersecret",
}

var Router fiber.Config = fiber.Config{
	Prefork:                      false,
	ServerHeader:                 "Frens",
	StrictRouting:                false,
	CaseSensitive:                false,
	Immutable:                    false,
	UnescapePath:                 false,
	ETag:                         false,
	BodyLimit:                    4 * 1024 * 1024,
	Concurrency:                  256 * 1024,
	Views:                        nil,
	ViewsLayout:                  "",
	PassLocalsToViews:            false,
	ReadTimeout:                  (60 * time.Second),
	WriteTimeout:                 (60 * time.Second),
	IdleTimeout:                  (60 * time.Second),
	ReadBufferSize:               4096,
	WriteBufferSize:              4096,
	CompressedFileSuffix:         ".gz",
	ProxyHeader:                  "",
	GETOnly:                      false,
	ErrorHandler:                 fiber.DefaultErrorHandler,
	DisableKeepalive:             false,
	DisableDefaultDate:           false,
	DisableDefaultContentType:    false,
	DisableHeaderNormalizing:     false,
	DisableStartupMessage:        false,
	AppName:                      "Frens",
	EnableTrustedProxyCheck:      false,
	TrustedProxies:               nil,
	DisablePreParseMultipartForm: false,
	StreamRequestBody:            false,
	EnablePrintRoutes:            false,
	Network:                      fiber.NetworkTCP4,
	JSONEncoder:                  json.Marshal,
	JSONDecoder:                  json.Unmarshal,
}

func init() {
	/*viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")      // optionally look for config in the working directory

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&Config); err != nil {
		panic(err)
	}*/
}

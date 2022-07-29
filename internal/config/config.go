package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Configuration struct {
	Database DatabaseConfiguration `mapstructure:"database"`
}

type DatabaseConfiguration struct {
	Host     string `mapstructure:"host" validate:"required"`
	Port     string `mapstructure:"port" validate:"required"`
	User     string `mapstructure:"user" validate:"required"`
	Password string `mapstructure:"password" validate:"required"`
	Database string `mapstructure:"database" validate:"required"`
	Debug    bool   `mapstructure:"debug" validate:"required"`
}

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       interface{}
}

func (c *Configuration) Validate() error {
	v := validator.New()

	var errs []*ErrorResponse
	err := v.Struct(c)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errs = append(errs, &ErrorResponse{
				FailedField: err.Field(),
				Tag:         err.Tag(),
				Value:       err.Value(),
			})
		}
	}

	for _, err := range errs {
		logrus.Errorf("Invalid configuration: %s %s %v", err.FailedField, err.Tag, err.Value)
	}

	if len(errs) > 0 {
		return errors.New(fmt.Sprintf("%v", errs))
	} else {
		return nil
	}
}

var C Configuration

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
	vp := viper.New()
	vp.SetConfigName("conf") // name of config file (without extension)
	vp.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name
	vp.AddConfigPath(".")    // optionally look for config in the working directory

	err := vp.ReadInConfig() // Find and read the config file
	if err != nil {
		panic(err)
	}

	err = vp.Unmarshal(&C)
	if err != nil {
		panic(err)
	}

	// Validate configuration
	if err := C.Validate(); err != nil {
		logrus.Fatal("Invalid configuration.")
	}

	//todo, add fsnotify to watch for changes
	logrus.Info("Config loaded")
}

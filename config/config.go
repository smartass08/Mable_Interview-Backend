package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Env string `mapstructure:"ENVIRONMENT"`

	// Fiber config
	Port    int  `mapstructure:"APP_PORT"`
	Prefork bool `mapstructure:"APP_PREFORK"`

	// Body limit for request
	RequestBodyLimit int `mapstructure:"REQUEST_BODY_LIMIT"`
}

// Validate will ensure all the required config variables are set
func (cfg Config) Validate() {
}

// Unexported variable to implement singleton pattern
var cfg *Config

// init will read all config variables and set it in the unexported variable
func init() {
	// Setup viper
	viper.SetDefault("REQUEST_BODY_LIMIT", 104857600) // 100 MB
	viper.SetDefault("APP_PORT", 4000)                // 100 MB

	// Automatically override values in config file with those in environment
	viper.AutomaticEnv()

	// Read config file
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
		} else {
			// Config file was found but another error was produced
			log.Fatalln(err)
		}
	}

	// Set config object
	err = viper.Unmarshal(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Validate that all config vars are set
	cfg.Validate()

}

func Get() *Config {
	if cfg == nil {
		log.Fatalln("Config not set ^._.^")
	}
	return cfg
}

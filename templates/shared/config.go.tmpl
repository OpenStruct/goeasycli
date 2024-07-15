package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"log"
	"os"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

const (
	varLogLevel     = "log.level"
	varPathToConfig = "../config.yaml"
)

var (
	CFG = GetConfig()
)

type Configuration struct {
	V *viper.Viper
}

func GetConfig() *Configuration {
	c := Configuration{
		V: viper.GetViper(),
	}
	c.V.SetDefault(varLogLevel, "info")
	c.V.AutomaticEnv()

	c.V.SetConfigFile(".env")
	c.V.ReadInConfig()

	c.V.SetConfigType("yaml")
	c.V.SetConfigName("config")
	c.V.AddConfigPath("./")
	err := c.V.ReadInConfig()

	if _, ok := err.(*os.PathError); ok {
		log.Printf("no config file not found. Using default values")
	} else if err != nil {
		zap.Any("fatal error while reading the config file", zap.Error(err))
	}
	return &c

}

func (c *Configuration) GetPathToConfig() string {
	return c.V.GetString(varPathToConfig)
}

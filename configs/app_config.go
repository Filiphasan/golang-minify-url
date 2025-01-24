package configs

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

type AppConfig struct {
	ProjectName string `json:"projectName"`
	Environment string `json:"environment"`
	Server      struct {
		Port     string `json:"port"`
		ShortUrl string `json:"shortUrl"`
	} `json:"server"`
	Mongodb struct {
		Host     string `json:"host"`
		Port     string `json:"port"`
		Username string `json:"username"`
		Password string `json:"password"`
		Database string `json:"database"`
	} `json:"mongodb"`
	Redis struct {
		Host     string `json:"host"`
		Port     string `json:"port"`
		Password string `json:"password"`
		Database int    `json:"database"`
	} `json:"redis"`
	Token struct {
		PoolingSize    int    `json:"poolingSize"`
		ExtendSize     int    `json:"extendSize"`
		ExpirationDays int    `json:"expirationDays"`
		EpochDate      string `json:"epochDate"`
	} `json:"token"`
}

var appConfig *AppConfig

func LoadAppConfig() {
	environment := os.Getenv("APP_ENV")
	if environment == "" {
		environment = "development"
	}

	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("./configs")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Failed to load config.json %s\n", err.Error())
		panic(err)
	}

	viper.SetConfigName("config." + environment)
	err = viper.MergeInConfig()
	if err != nil {
		fmt.Printf("Failed to load config.%s.json %s\n", environment, err.Error())
		panic(err)
	}

	appConfig = &AppConfig{}
	err = viper.Unmarshal(appConfig)
	appConfig.Environment = environment
	if err != nil {
		panic(err)
	}
}

func GetAppConfig() *AppConfig {
	return appConfig
}

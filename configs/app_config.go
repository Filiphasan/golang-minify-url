package configs

import (
	"github.com/spf13/viper"
	"os"
)

type AppConfig struct {
	ProjectName string `json:"projectName"`
	Mongodb     struct {
		Host     string `json:"host"`
		Port     string `json:"port"`
		Username string `json:"username"`
		Password string `json:"password"`
		Database string `json:"database"`
	} `json:"mongodb"`
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
		panic(err)
	}

	viper.SetConfigName("config." + environment)
	err = viper.MergeInConfig()
	if err != nil {
		panic(err)
	}

	appConfig = &AppConfig{}
	_ = viper.Unmarshal(appConfig)
}

func GetAppConfig() *AppConfig {
	return appConfig
}

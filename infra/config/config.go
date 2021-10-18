package config

import (
	"log"
	"sync"

	"github.com/diegoclair/go_utils-lib/v2/logger"
	"github.com/spf13/viper"
)

var (
	config *EnvironmentVariables
	once   sync.Once
)

// GetConfigEnvironment to read initial config
func GetConfigEnvironment() *EnvironmentVariables {
	once.Do(func() {

		viper.SetConfigFile("../../config.toml")
		viper.AutomaticEnv()

		err := viper.ReadInConfig()
		if err != nil {
			logger.Error("Error to read configs: ", err)
			panic(err)
		}

		config = &EnvironmentVariables{}
		err = viper.Unmarshal(&config)
		if err != nil {
			log.Fatal(err)
		}
	})

	return config
}

// EnvironmentVariables is environment variables configs
type EnvironmentVariables struct {
	DB DBConfig `mapstructure:"db"`
}

type DBConfig struct {
	User             string `mapstructure:"user"`
	Pass             string `mapstructure:"pass"`
	Name             string `mapstructure:"name"`
	Host             string `mapstructure:"host"`
	Port             int    `mapstructure:"port"`
	CryptoKey        string `mapstructure:"crypto-key"`
	MaxLifeInMinutes int    `mapstructure:"max-life-minutes"`
	MaxIdleConns     int    `mapstructure:"max-idle-conns"`
	MaxOpenConns     int    `mapstructure:"max-open-conns"`
}

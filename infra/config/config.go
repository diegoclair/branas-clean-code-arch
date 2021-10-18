package config

import (
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

		viper.SetConfigFile("config.toml")
		viper.AutomaticEnv()

		err := viper.ReadInConfig()
		if err != nil {
			logger.Error("Error to read configs: ", err)
			panic(err)
		}

		config = &EnvironmentVariables{}
		config.DB.User = viper.GetString("DB_USER")
		config.DB.Pass = viper.GetString("DB_PASSWORD")
		config.DB.Host = viper.GetString("DB_HOST")
		config.DB.Port = viper.GetInt("DB_PORT")
		config.DB.Name = viper.GetString("DB_NAME")
		config.DB.MaxLifeInMinutes = viper.GetInt("MAX_LIFE_IN_MINUTES")
		config.DB.MaxIdleConns = viper.GetInt("MAX_IDLE_CONNS")
		config.DB.MaxOpenConns = viper.GetInt("MAX_OPEN_CONNS")
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

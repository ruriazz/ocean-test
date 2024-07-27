package configs

import (
	"github.com/spf13/viper"
)

type Config struct {
	Debug     bool   `mapstructure:"DEBUG"`
	RedisHost string `mapstructure:"REDIS_HOST"`
	RedisPort string `mapstructure:"REDIS_PORT"`
	RedisUser string `mapstructure:"REDIS_USER"`
	RedisPass string `mapstructure:"REDIS_PASS"`

	MetaBusinessID  string `mapstructure:"META_BUSINESS_ID"`
	MetaAccessToekn string `mapstructure:"META_ACCESS_TOKEN"`
}

func LoadConfig() (*Config, error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("local")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

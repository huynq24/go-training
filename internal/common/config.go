package common

import "github.com/spf13/viper"

type DbConfig struct {
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Ip       string `mapstructure:"ip"`
	Port     string `mapstructure:"port"`
	DbName   string `mapstructure:"dbName"`
}

type Config struct {
	Mysql DbConfig `mapstructure:"mysql"`
}

func LoadConfig() (config Config, err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

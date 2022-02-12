package configs

import (
	"github.com/spf13/viper"
)

//Config 配置文件
type Config struct {
	Default *Default
}

//Default 默认配置
type Default struct {
	Enable bool
}

//Init init config
func Init(path string) *Config {
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic("Loading config file failed, config file not found!")
		} else {
			panic("Loading config file failed, " + err.Error())
		}
	}

	var cfg Config
	err := viper.Unmarshal(&cfg)
	if err != nil {
		panic("Unmarshal config file failed, " + err.Error())
	}
	return &cfg
}

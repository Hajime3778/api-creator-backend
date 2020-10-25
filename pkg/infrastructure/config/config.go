package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config 設定
type Config struct {
	Server struct {
		Port    string
		Timeout int
	}
	DataBase struct {
		Host     string
		Port     string
		User     string
		Password string
		Database string
	}
}

// NewConfig 設定ファイルを読み込みCondigを作成します
func NewConfig() *Config {

	viper.SetConfigFile("config.json")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}

	c := new(Config)

	// conf読み取り
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}

	// UnmarshalしてConfigにマッピング
	if err := viper.Unmarshal(&c); err != nil {
		panic(fmt.Errorf("unable to decode into struct, %v", err))
	}

	return c
}

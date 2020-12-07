package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config 設定
type Config struct {
	Server struct {
		Host    string
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
func NewConfig(filePath string) *Config {

	viper.SetConfigFile(filePath)
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

// ServerBaseURL ConfigからサーバーのBaseURLを作成して返却します。
func (c *Config) ServerBaseURL() string {
	apiServerBaseurl := fmt.Sprintf("http://%s%s/",
		c.Server.Host,
		c.Server.Port,
	)

	return apiServerBaseurl
}

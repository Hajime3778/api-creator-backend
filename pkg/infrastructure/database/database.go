package database

import (
	"fmt"
	"net/url"

	"github.com/Hajime3778/api-creator-backend/pkg/infrastructure/config"

	"github.com/jinzhu/gorm"
)

// DB Database
type DB struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

// NewDB Configから、DBオブジェクトを作成します
func NewDB(c *config.Config) *DB {
	return &DB{
		Host:     c.DataBase.Host,
		Port:     c.DataBase.Port,
		Username: c.DataBase.User,
		Password: c.DataBase.Password,
		DBName:   c.DataBase.Database,
	}
}

// NewMysqlConnection DBオブジェクトからMysqlの接続を作成します
func (d *DB) NewMysqlConnection() *gorm.DB {
	// MySQLの接続情報を作成
	connectionInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		d.Username,
		d.Password,
		d.Host,
		d.Port,
		d.DBName)

	option := url.Values{}
	option.Add("charset", "utf8")
	option.Add("parseTime", "True")
	option.Add("loc", "Local")

	connection := fmt.Sprintf("%s?%s", connectionInfo, option.Encode())

	// MySQLに接続
	db, err := gorm.Open("mysql", connection)
	if err != nil {
		panic(err.Error())
	}

	return db
}

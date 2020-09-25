package main

import (
	"go-mysql-api/cmd/logger"
	"go-mysql-api/cmd/server"
	"go-mysql-api/pkg/infrastructure/config"
	"go-mysql-api/pkg/infrastructure/database"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	logger.LoggingSetting()
	cfg := config.NewConfig()
	db := database.NewDB(cfg)

	server := server.NewServer(cfg, db)
	server.SetUpRouter()
	server.Run()
}

package main

import (
	"github.com/Hajime3778/api-creator-backend/cmd/logger"
	"github.com/Hajime3778/api-creator-backend/cmd/server"
	"github.com/Hajime3778/api-creator-backend/pkg/infrastructure/config"
	"github.com/Hajime3778/api-creator-backend/pkg/infrastructure/database"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	logger.LoggingSetting("./log/")
	cfg := config.NewConfig("./backend.config.json")
	db := database.NewDB(cfg)

	server := server.NewServer(cfg, db)
	server.SetUpRouter()
	server.Run()
}

package main

import (
	_apiRepository "github.com/Hajime3778/api-creator-backend/pkg/admin/api/repository"
	_methodRepository "github.com/Hajime3778/api-creator-backend/pkg/admin/method/repository"
	_modelRepository "github.com/Hajime3778/api-creator-backend/pkg/admin/model/repository"
	"github.com/Hajime3778/api-creator-backend/pkg/apiserver/handler"
	_apiserverRepository "github.com/Hajime3778/api-creator-backend/pkg/apiserver/repository"
	"github.com/Hajime3778/api-creator-backend/pkg/apiserver/usecase"
	"github.com/Hajime3778/api-creator-backend/pkg/infrastructure/config"
	"github.com/Hajime3778/api-creator-backend/pkg/infrastructure/database"
	"github.com/Hajime3778/api-creator-backend/pkg/infrastructure/server"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// API Server側のMongoDB接続
	apiserverCfg := config.NewConfig("./apiserver.config.json")
	apiServerBaseurl := apiserverCfg.ServerBaseURL()
	mongoDB := database.NewDB(apiserverCfg)
	apiServer := server.NewServer(apiserverCfg)

	// 管理画面で設定されたMysql接続
	mysqlCfg := config.NewConfig("./admin.config.json")
	mysqlDB := database.NewDB(mysqlCfg)
	mysqlConn := mysqlDB.NewMysqlConnection()

	apiRepository := _apiRepository.NewAPIRepository(mysqlConn, apiServerBaseurl)
	methodRepository := _methodRepository.NewMethodRepository(mysqlConn)
	modelRepository := _modelRepository.NewModelRepository(mysqlConn)
	apiserverRepository := _apiserverRepository.NewAPIServerRepository(mongoDB)

	router := apiServer.Router

	apiserverUsecase := usecase.NewAPIServerUsecase(apiRepository, methodRepository, modelRepository, apiserverRepository)
	handler.NewAPIServerHandler(router, apiserverUsecase)

	apiServer.Run()
}

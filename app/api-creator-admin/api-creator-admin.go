package main

import (
	_apiHandler "github.com/Hajime3778/api-creator-backend/pkg/admin/api/handler"
	_apiRepository "github.com/Hajime3778/api-creator-backend/pkg/admin/api/repository"
	_apiUsecase "github.com/Hajime3778/api-creator-backend/pkg/admin/api/usecase"
	_methodHandler "github.com/Hajime3778/api-creator-backend/pkg/admin/method/handler"
	_methodRepository "github.com/Hajime3778/api-creator-backend/pkg/admin/method/repository"
	_methodUsecase "github.com/Hajime3778/api-creator-backend/pkg/admin/method/usecase"
	_modelHandler "github.com/Hajime3778/api-creator-backend/pkg/admin/model/handler"
	_modelRepository "github.com/Hajime3778/api-creator-backend/pkg/admin/model/repository"
	_modelUsecase "github.com/Hajime3778/api-creator-backend/pkg/admin/model/usecase"
	"github.com/Hajime3778/api-creator-backend/pkg/infrastructure/config"
	"github.com/Hajime3778/api-creator-backend/pkg/infrastructure/database"
	"github.com/Hajime3778/api-creator-backend/pkg/infrastructure/logger"
	"github.com/Hajime3778/api-creator-backend/pkg/infrastructure/server"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	logger.LoggingSetting("./log/")
	adminCfg := config.NewConfig("./admin.config.json")
	db := database.NewDB(adminCfg)

	adminServer := server.NewServer(adminCfg)
	router := adminServer.Router

	// Group v1
	apiV1 := router.Group("api/v1")
	conn := db.NewMysqlConnection()

	//repositories
	apiRepository := _apiRepository.NewAPIRepository(conn)
	methodRepository := _methodRepository.NewMethodRepository(conn)
	modelRepository := _modelRepository.NewModelRepository(conn)

	// APIs
	apiUsecase := _apiUsecase.NewAPIUsecase(apiRepository, methodRepository, modelRepository)
	_apiHandler.NewAPIHandler(apiV1, apiUsecase)

	// Methods
	methodUsecase := _methodUsecase.NewMethodUsecase(apiRepository, methodRepository, modelRepository)
	_methodHandler.NewMethodHandler(apiV1, methodUsecase)

	// Models
	modelUsecase := _modelUsecase.NewModelUsecase(modelRepository)
	_modelHandler.NewModelHandler(apiV1, modelUsecase)

	adminServer.Run()
}

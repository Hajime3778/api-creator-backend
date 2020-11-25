package server

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

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func newRouter() *gin.Engine {
	router := gin.Default()

	// アクセス許可設定
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Content-Type"}

	router.Use(cors.New(config))
	router.StaticFile("/", "./index.html")

	return router
}

// SetUpRouter ルーティングを設定します。
func (s *Server) SetUpRouter() *gin.Engine {
	// Group v1
	apiV1 := s.router.Group("api/v1")
	conn := s.db.NewMysqlConnection()

	// APIs
	apiRepository := _apiRepository.NewAPIRepository(conn)
	apiUsecase := _apiUsecase.NewAPIUsecase(apiRepository)
	_apiHandler.NewAPIHandler(apiV1, apiUsecase)

	// Methods
	methodRepository := _methodRepository.NewMethodRepository(conn)
	methodUsecase := _methodUsecase.NewMethodUsecase(methodRepository)
	_methodHandler.NewMethodHandler(apiV1, methodUsecase)

	// Models
	modelRepository := _modelRepository.NewModelRepository(conn)
	modelUsecase := _modelUsecase.NewModelUsecase(modelRepository)
	_modelHandler.NewModelHandler(apiV1, modelUsecase)

	return s.router
}
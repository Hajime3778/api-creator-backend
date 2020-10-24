package server

import (
	_apiHandler "github.com/Hajime3778/api-creator-backend/pkg/apis/api/handler"
	_apiRepository "github.com/Hajime3778/api-creator-backend/pkg/apis/api/repository"
	_apiUsecase "github.com/Hajime3778/api-creator-backend/pkg/apis/api/usecase"
	_methodHandler "github.com/Hajime3778/api-creator-backend/pkg/apis/method/handler"
	_methodRepository "github.com/Hajime3778/api-creator-backend/pkg/apis/method/repository"
	_methodUsecase "github.com/Hajime3778/api-creator-backend/pkg/apis/method/usecase"
	_userHandler "github.com/Hajime3778/api-creator-backend/pkg/apis/user/handler"
	_userRepository "github.com/Hajime3778/api-creator-backend/pkg/apis/user/repository"
	_userUsecase "github.com/Hajime3778/api-creator-backend/pkg/apis/user/usecase"

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

// SetUpRouter Setup all api routing
func (s *Server) SetUpRouter() *gin.Engine {
	// Group v1
	apiV1 := s.router.Group("api/v1")

	// Users
	userRepository := _userRepository.NewUserRepository(s.db)
	userUsecase := _userUsecase.NewUserUsecase(userRepository)
	_userHandler.NewUserHandler(apiV1, userUsecase)

	// APIs
	apiRepository := _apiRepository.NewAPIRepository(s.db)
	apiUsecase := _apiUsecase.NewAPIUsecase(apiRepository)
	_apiHandler.NewAPIHandler(apiV1, apiUsecase)

	// Methods
	methodRepository := _methodRepository.NewMethodRepository(s.db)
	methodUsecase := _methodUsecase.NewMethodUsecase(methodRepository)
	_methodHandler.NewMethodHandler(apiV1, methodUsecase)

	return s.router
}

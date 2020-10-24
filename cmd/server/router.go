package server

import (
	_apiHandler "github.com/Hajime3778/api-creator-backend/pkg/apis/api/handler"
	_apiRepository "github.com/Hajime3778/api-creator-backend/pkg/apis/api/repository"
	_apiUsecase "github.com/Hajime3778/api-creator-backend/pkg/apis/api/usecase"
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

	s.apiRoutes(apiV1)
	s.userRoutes(apiV1)

	return s.router
}

func (s *Server) userRoutes(api *gin.RouterGroup) {
	repository := _userRepository.NewUserRepository(s.db)
	usecase := _userUsecase.NewUserUsecase(repository)
	_userHandler.NewUserHandler(api, usecase)
}

func (s *Server) apiRoutes(api *gin.RouterGroup) {
	repository := _apiRepository.NewAPIRepository(s.db)
	usecase := _apiUsecase.NewAPIUsecase(repository)
	_apiHandler.NewAPIHandler(api, usecase)
}

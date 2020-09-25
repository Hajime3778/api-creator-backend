package server

import (
	"go-mysql-api/pkg/user/handler"
	"go-mysql-api/pkg/user/repository"
	"go-mysql-api/pkg/user/usecase"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func newRouter() *gin.Engine {
	router := gin.Default()

	// すべてのアクセス許可
	config := cors.Config{AllowOrigins: []string{"*"}}
	router.Use(cors.New(config))
	router.StaticFile("/", "./index.html")

	return router
}

// SetUpRouter Setup all api routing
func (s *Server) SetUpRouter() *gin.Engine {
	// Group v1
	apiV1 := s.router.Group("api/v1")
	s.userRoutes(apiV1)
	return s.router
}

func (s *Server) userRoutes(api *gin.RouterGroup) {
	repository := repository.NewUserRepository(s.db)
	usecase := usecase.NewUserUsecase(repository)
	handler.NewUserHandler(api, usecase)
}

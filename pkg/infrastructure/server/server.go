package server

import (
	"net/http"
	"time"

	"github.com/Hajime3778/api-creator-backend/pkg/infrastructure/config"
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

// Server サーバーの情報を定義します。
type Server struct {
	Router *gin.Engine
	server *http.Server
}

// NewServer Serverを初期化します
func NewServer(c *config.Config) *Server {
	r := NewRouter()
	s := newServer(c, r)
	return &Server{
		Router: r,
		server: s,
	}
}

func newServer(c *config.Config, router *gin.Engine) *http.Server {
	s := &http.Server{
		Addr:         c.Server.Port,
		Handler:      router,
		ReadTimeout:  time.Duration(c.Server.Timeout) * time.Second,
		WriteTimeout: time.Duration(c.Server.Timeout) * time.Second,
	}
	return s
}

// NewRouter 新規でデフォルト設定のルーターを作成します
func NewRouter() *gin.Engine {
	router := gin.Default()

	// アクセス許可設定
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Content-Type"}

	router.Use(cors.New(config))

	return router
}

// Run サーバーを実行します
func (s *Server) Run() {
	s.server.ListenAndServe()
}

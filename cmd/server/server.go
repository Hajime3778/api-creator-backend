package server

import (
	"go-mysql-api/pkg/infrastructure/config"
	"go-mysql-api/pkg/infrastructure/database"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Server server
type Server struct {
	db     *database.DB
	router *gin.Engine
	server *http.Server
}

// NewServer Server create
func NewServer(c *config.Config, db *database.DB) *Server {
	r := newRouter()
	s := newServer(c, r)
	return &Server{
		db:     db,
		router: r,
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

// Run Run server
func (s *Server) Run() {
	s.server.ListenAndServe()
}

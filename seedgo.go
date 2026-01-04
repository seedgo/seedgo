package seedgo

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Engine *gin.Engine
}

func NewServer() *Server {
	InitCmd()
	Init()

	if !ServerConfig.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	engine := gin.New()
	s := &Server{
		Engine: engine,
	}

	return s
}

func (s *Server) Start() error {
	port := strconv.Itoa(ServerConfig.Port)
	fmt.Printf("start server at port: %s\n", port)
	return s.Engine.Run(":" + port)
}

func (s *Server) GetEngine() *gin.Engine {
	return s.Engine
}

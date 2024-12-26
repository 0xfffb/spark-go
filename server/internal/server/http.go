package server

import (
	"spark-server/internal/handler"
	"spark-server/internal/manager"
	"spark-server/internal/service"

	"github.com/gin-gonic/gin"
)

type HTTPServer struct {
	engine        *gin.Engine
	clientManager *manager.ClientManager
}	

func NewHTTPServer(clientManager *manager.ClientManager) *HTTPServer {
	clientService := service.NewClientService(clientManager)
	clientHandler := handler.NewClientHandler(clientService)

	engine := gin.Default()

	// 设置路由
	engine.POST("/send/:clientID", clientHandler.SendMessage)

	return &HTTPServer{
		engine:        engine,
		clientManager: clientManager,
	}
}

func (s *HTTPServer) Run(addr string) error {
	return s.engine.Run(addr)
}

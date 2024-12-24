package main

import (
	"log"
	"spark-server/internal/manager"
	"spark-server/internal/server"
)

func main() {
	// 创建客户端管理器
	clientManager := manager.NewClientManager()

	// 创建并启动 Yamux 服务器
	yamuxServer, err := server.NewYamuxServer(":1234", clientManager)
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		if err := yamuxServer.Run(); err != nil {
			log.Printf("yamux server error: %v", err)
		}
	}()

	// 创建并启动 HTTP 服务器
	httpServer := server.NewHTTPServer(clientManager)
	if err := httpServer.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

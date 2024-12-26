package service

import (
	"fmt"
	"log"
	"spark-server/internal/manager"
)

type ClientService struct {
	clientManager *manager.ClientManager
}

func NewClientService(clientManager *manager.ClientManager) *ClientService {
	return &ClientService{
		clientManager: clientManager,
	}
}

func (s *ClientService) SendMessage(clientID, command string, data string) (string, error) {
	session, exists := s.clientManager.GetClient(clientID)
	if !exists {
		return "", fmt.Errorf("client %s not found", clientID)
	}

	// 打开新的流进行通信
	stream, err := session.Open()
	if err != nil {
		return "", err
	}
	defer stream.Close()

	// TODO: 实现具体的消息发送逻辑，例如：
	message := fmt.Sprintf("%s:%s", command, data)
	n, err := stream.Write([]byte(message))
	if err != nil {
		return "", fmt.Errorf("failed to write message: %w", err)
	}
	if n < len(message) {
		return "", fmt.Errorf("incomplete write: wrote %d bytes out of %d", n, len(message))
	}

	// 添加读取响应的逻辑
	response := make([]byte, 1024)
	n, err = stream.Read(response)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %w", err)
	}
	log.Default().Printf("response: %s", response)
	return fmt.Sprintf("%s:%s", command, string(response[:n])), nil
}

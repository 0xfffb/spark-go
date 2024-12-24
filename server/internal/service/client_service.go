package service

import (
	"fmt"
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

func (s *ClientService) SendMessage(clientID, command, data string) error {
	session, exists := s.clientManager.GetClient(clientID)
	if !exists {
		return fmt.Errorf("client %s not found", clientID)
	}

	// 打开新的流进行通信
	stream, err := session.Open()
	if err != nil {
		return err
	}
	defer stream.Close()

	// TODO: 实现具体的消息发送逻辑，例如：
	// message := fmt.Sprintf("%s:%s", command, data)
	// _, err = stream.Write([]byte(message))
	return nil
}

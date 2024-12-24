package server

import (
	"fmt"
	"net"
	"spark-server/internal/manager"

	"github.com/hashicorp/yamux"
)

type YamuxServer struct {
	listener      net.Listener
	clientManager *manager.ClientManager
}

func NewYamuxServer(addr string, clientManager *manager.ClientManager) (*YamuxServer, error) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, fmt.Errorf("failed to start yamux listener: %v", err)
	}

	return &YamuxServer{
		listener:      listener,
		clientManager: clientManager,
	}, nil
}

func (s *YamuxServer) Run() error {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			return fmt.Errorf("failed to accept connection: %v", err)
		}

		go s.handleConnection(conn)
	}
}

func (s *YamuxServer) handleConnection(conn net.Conn) {
	// 创建 yamux 服务端会话
	session, err := yamux.Server(conn, nil)
	if err != nil {
		conn.Close()
		return
	}

	// 等待客户端发送第一条消息，包含客户端 ID
	stream, err := session.Accept()
	if err != nil {
		session.Close()
		return
	}

	// 读取客户端 ID
	buf := make([]byte, 1024)
	n, err := stream.Read(buf)
	if err != nil {
		stream.Close()
		session.Close()
		return
	}
	clientID := string(buf[:n])

	// 将会话添加到管理器
	s.clientManager.AddClient(clientID, session)

	// 关闭认证流
	stream.Close()

	// 处理后续的流
	for {
		stream, err := session.Accept()
		if err != nil {
			s.clientManager.RemoveClient(clientID)
			return
		}
		go s.handleStream(stream)
	}
}

func (s *YamuxServer) handleStream(stream net.Conn) {
	defer stream.Close()
	// TODO: 实现具体的流处理逻辑
}

func (s *YamuxServer) Close() error {
	return s.listener.Close()
}

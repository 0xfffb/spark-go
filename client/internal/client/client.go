package client

import (
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/hashicorp/yamux"
)

type Client struct {
	id      string
	session *yamux.Session
}

func NewClient(id string) *Client {
	return &Client{
		id: id,
	}
}

func (c *Client) Connect(addr string) error {
	// 建立TCP连接
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return fmt.Errorf("连接服务器失败: %v", err)
	}

	// 创建客户端会话
	session, err := yamux.Client(conn, nil)
	if err != nil {
		conn.Close()
		return fmt.Errorf("创建yamux会话失败: %v", err)
	}
	c.session = session

	// 打开一个流来发送客户端ID
	stream, err := session.Open()
	if err != nil {
		session.Close()
		return fmt.Errorf("打开认证流失败: %v", err)
	}

	// 发送客户端ID
	if _, err := stream.Write([]byte(c.id)); err != nil {
		stream.Close()
		session.Close()
		return fmt.Errorf("发送客户端ID失败: %v", err)
	}
	stream.Close()

	// 启动消息处理
	go c.handleMessages()

	log.Printf("客户端 %s 已连接到服务器", c.id)
	return nil
}

func (c *Client) handleMessages() {
	for {
		// 接受新的流
		stream, err := c.session.Accept()
		if err != nil {
			log.Printf("接受流失败: %v", err)
			return
		}

		// 处理消息
		go c.handleStream(stream)
	}
}

func (c *Client) handleStream(stream net.Conn) {
	defer stream.Close()

	// 读取消息
	buf := make([]byte, 1024)
	n, err := stream.Read(buf)
	if err != nil {
		log.Printf("读取消息失败: %v", err)
		return
	}

	// 解析消息
	message := string(buf[:n])
	parts := strings.SplitN(message, ":", 2)
	if len(parts) != 2 {
		log.Printf("无效的消息格式: %s", message)
		return
	}

	command, data := parts[0], parts[1]
	log.Printf("收到命令: %s, 数据: %s", command, data)

	// TODO: 根据command和data执行相应的操作
	switch command {
	case "ping":
		_, _ = stream.Write([]byte("pong"))
		log.Printf("收到ping命令，数据：%s", data)
		// 可以在这里实现ping的响应逻辑
	default:
		log.Printf("未知命令：%s", command)
		_, _ = stream.Write([]byte(data))
	}
}

func (c *Client) Close() error {
	if c.session != nil {
		return c.session.Close()
	}
	return nil
}

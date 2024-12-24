package handler

import (
	"net/http"
	"spark-server/internal/service"

	"github.com/gin-gonic/gin"
)

type ClientHandler struct {
	clientService *service.ClientService
}

func NewClientHandler(clientService *service.ClientService) *ClientHandler {
	return &ClientHandler{
		clientService: clientService,
	}
}

// SendMessage 向指定客户端发送消息
func (h *ClientHandler) SendMessage(c *gin.Context) {
	clientID := c.Param("clientID")
	var message struct {
		Command string `json:"command"`
		Data    string `json:"data"`
	}

	if err := c.BindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.clientService.SendMessage(clientID, message.Command, message.Data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "message sent"})
}

package manager

import (
	"sync"

	"github.com/hashicorp/yamux"
)

type ClientManager struct {
	clients map[string]*yamux.Session
	mu      sync.RWMutex
}

func NewClientManager() *ClientManager {
	return &ClientManager{
		clients: make(map[string]*yamux.Session),
	}
}

func (m *ClientManager) AddClient(id string, session *yamux.Session) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.clients[id] = session
}

func (m *ClientManager) RemoveClient(id string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if session, exists := m.clients[id]; exists {
		session.Close()
		delete(m.clients, id)
	}
}

func (m *ClientManager) GetClient(id string) (*yamux.Session, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	session, exists := m.clients[id]
	return session, exists
}

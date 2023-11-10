package manager

import (
	"github.com/AdamShannag/gosockz/types"
	"github.com/gorilla/websocket"
)

type Manager struct {
	clients     types.ClientRepo
	upgrader    *websocket.Upgrader
	eventRouter types.EventRouter
}

func NewManager(clientRepo types.ClientRepo, upgrader *websocket.Upgrader, eventRouter types.EventRouter) *Manager {
	m := &Manager{
		clients:     clientRepo,
		upgrader:    upgrader,
		eventRouter: eventRouter,
	}
	return m
}

func (m *Manager) Clients() types.ClientRepo {
	return m.clients
}

func (m *Manager) EventRouter() types.EventRouter {
	return m.eventRouter
}

func (m *Manager) WebsocketUpgrader() *websocket.Upgrader {
	return m.upgrader
}

func (m *Manager) AddClient(client types.Client) error {
	return m.Clients().AddClient(client)
}

func (m *Manager) RemoveClient(client types.Client) {
	if _, ok := m.clients.ClientExists(client); ok {
		client.Connection().Close()
		_ = m.clients.DeleteClient(client)
	}
}

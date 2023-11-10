package repo

import (
	"errors"
	"sync"

	"github.com/AdamShannag/gosockz/types"
)

type InMemoryClientRepo struct {
	clients map[types.Client]bool
	sync.RWMutex
}

func NewInMemoryClientRepo() *InMemoryClientRepo {
	return &InMemoryClientRepo{clients: make(map[types.Client]bool)}
}

func (r *InMemoryClientRepo) GetAllClients() []types.Client {
	clients := make([]types.Client, len(r.clients))

	i := 0
	for k := range r.clients {
		clients[i] = k
		i++
	}

	return clients
}

func (r *InMemoryClientRepo) AddClient(client types.Client) error {
	r.Lock()
	defer r.Unlock()

	r.clients[client] = true

	return nil
}
func (r *InMemoryClientRepo) ClientExists(client types.Client) (error, bool) {
	if _, ok := r.clients[client]; ok {
		return nil, ok
	}

	return nil, false
}

func (r *InMemoryClientRepo) DeleteClient(client types.Client) error {
	r.Lock()
	defer r.Unlock()

	err, ok := r.ClientExists(client)

	if err != nil {
		return err
	}

	if !ok {
		return errors.New("client does not exist")
	}

	client.Connection().Close()
	delete(r.clients, client)

	return nil
}

func (r *InMemoryClientRepo) GetClient(clientID string) types.Client {
	for client := range r.clients {
		if client.GetClientID() == clientID {
			return client
		}
	}

	return nil
}

package datautils

import (
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
)

type clientRegistry struct {
	mutex       sync.RWMutex
	connections map[string]*websocket.Conn
}

func newRegistry() *clientRegistry {
	return &clientRegistry{connections: map[string]*websocket.Conn{}}
}

func (c *clientRegistry) set(client *UserConnection) {
	// set online

	c.mutex.Lock()
	c.connections[client.UserID] = client.Connection
	c.mutex.Unlock()

	// cache client connect
}

func (c *clientRegistry) get(ID string) (*websocket.Conn, error) {
	// query := map[string]interface{}{
	// 	"query": map[string]interface{}{
	// 		"match": map[string]interface{}{
	// 			"session_id": sID,
	// 		},
	// 	},
	// }
	// var users []User
	// err := FindAll(query, db.IdxUserConnections, &users)
	// if err != nil {
	// 	fmt.Println("error find avail username")
	// }
	c.mutex.RLock()
	conn, ok := c.connections[ID]
	c.mutex.RUnlock()

	if !ok {
		return nil, fmt.Errorf("could not find client with client UserID %s", ID)
	}
	return conn, nil
}

func (c *clientRegistry) delete(client *UserConnection) {
	// set offline

	c.mutex.Lock()
	delete(c.connections, client.UserID)
	c.mutex.Unlock()

	// cache client disconnect
}

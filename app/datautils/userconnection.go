package datautils

import (
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type UserConnection struct {
	UserID     string          `json:"user_id"`
	SessionID  string          `json:"session_id"`
	Connection *websocket.Conn `json:"connection"`
}

const (
	WriteWait      = 30 * time.Second
	PongWait       = 30 * time.Second
	MaxMessageSize = 1024 * 1024
)

var registry = newRegistry()

func CreateUserConnection(s Session, w http.ResponseWriter, r *http.Request) (*UserConnection, error) {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			check := r.Method == http.MethodGet
			return check
		},
	}

	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		return nil, err
	}

	userConn := &UserConnection{
		UserID:     s.UserID,
		SessionID:  s.ID,
		Connection: conn,
	}

	registry.set(userConn)

	return userConn, nil
}

func ConnectionBySessionID(sID string) (*websocket.Conn, error) {
	return registry.get(sID)
}

func (uc *UserConnection) SendCloseConnection() {
	uc.Connection.SetWriteDeadline(time.Now().Add(WriteWait))
	uc.Connection.WriteMessage(websocket.CloseMessage, nil)
}

func (c *UserConnection) Close() error {
	registry.delete(c)
	return c.Connection.Close()
}

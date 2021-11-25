package chat

import "github.com/gorilla/websocket"

type User struct {
	Username string
	Conn     *websocket.Conn
	Global   *Chat
}

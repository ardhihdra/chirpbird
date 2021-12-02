package chat

import (
	"log"

	"github.com/gorilla/websocket"
)

type Chat struct {
	users    map[string]*User
	messages chan *Message
	join     chan *User
	leave    chan *User
}

var upgrader = websocket.Upgrader{}

func Start(port string) {
	c := &Chat{
		users:    make(map[string]*User),
		messages: make(chan *Message),
		join:     make(chan *User),
		leave:    make(chan *User),
	}

	go c.Run()
}

func (c *Chat) Run() {
	for {
		select {
		case user := <-c.join:
			log.Printf(user.Username)
			log.Printf("joined")
		}
	}
}

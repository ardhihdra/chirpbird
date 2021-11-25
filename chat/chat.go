package chat

import (
	"fmt"
	"log"
	"net/http"

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
	log.Printf("listening to %s", port)

	c := &Chat{
		users:    make(map[string]*User),
		messages: make(chan *Message),
		join:     make(chan *User),
		leave:    make(chan *User),
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from go-chat-react!"))
	})

	go c.Run()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
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

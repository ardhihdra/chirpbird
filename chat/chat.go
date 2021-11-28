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

func handleLogin() http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			// q := r.URL.Query().Get("q")
			if r.Method != "POST" {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("not found"))
				return
			}
			username := r.FormValue("username")
			country := r.FormValue("country")
			interests := r.FormValue("interests")
			profile := r.FormValue("profile")

			message := fmt.Sprintf("Hello %s, %s, %s, %s", username, country, interests, profile)
			w.Write([]byte(message))
		})
}

func checkUniqueUsername() http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			// GET users with username
		})
}

func dashboardData() http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			// GET all users, users by interests, users by country,
			// GET all rooms, rooms by interests, rooms by country

		})
}

func roomsData() http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			// GET room detail
			// CREATE a room

		})
}

func searchStuff() http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			// GET search by name, profile,
		})
}

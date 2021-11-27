package chat

import "github.com/gorilla/websocket"

type User struct {
	ID        int64    `json:"id"`
	Username  string   `json:"username"`
	Country   string   `json:"country"`
	Interests []string `json:"interests"`
	Profile   string   `json:"profile"`
	Conn      []*websocket.Conn
	Global    []*Chat
}

func (user *User) getUsers() {

}

func (user *User) getUsersByInterests(interests []string) User {
	return User{}
}

func (user *User) getUsersByCountry(interests []string) User {
	return User{}
}

func (user *User) getUsersByUsername(interests []string) User {
	return User{}
}

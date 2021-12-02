package models

import (
	"fmt"
	"strings"

	"github.com/ardhihdra/chirpbird/db"
)

type User struct {
	ID        string   `json:"id"`
	Username  string   `json:"username"`
	Email     string   `json:"email"`
	Password  string   `json:"password"`
	Profile   string   `json:"profile"`
	Interests []string `json:"interests"`
	Country   string   `json:"country"`
	CreatedAt int64    `json:"created_at"`
	UpdatedAt int64    `json:"updated_at"`
}

func (u *User) UsernameAvailable() bool {
	// res, _ := datastore.DB.Users.Find(bson.M{"username": strings.ToLower(u.Username)}).Count()
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"username": strings.ToLower(u.Username),
			},
		},
	}
	var users []User
	err := FindAll(query, db.IndexList.Users, &users)
	if err != nil {
		fmt.Println("error find avail username")
	}
	return len(users) == 0
}

func (u *User) EmailAvailable() bool {
	query := map[string]interface{}{
		"email": strings.ToLower(u.Email),
	}
	var users []User
	err := FindAll(query, db.IndexList.Users, &users)
	if err != nil {
		fmt.Println("error find avail email")
	}
	return len(users) == 0
}

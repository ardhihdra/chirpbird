package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/ardhihdra/chirpbird/db"
)

type User struct {
	ID        string   `json:"_id"`
	Username  string   `json:"username"`
	Email     string   `json:"email"`
	Password  string   `json:"password"`
	Profile   string   `json:"Profile"`
	Interests []string `json:"interests"`
	Country   string   `json:"country"`
	CreatedAt int64    `json:"created_at"`
	UpdatedAt int64    `json:"updated_at"`
}

func (u *User) UsernameAvailable() bool {
	// res, _ := datastore.DB.Users.Find(bson.M{"username": strings.ToLower(u.Username)}).Count()
	query := map[string]interface{}{
		"username": strings.ToLower(u.Username),
	}
	var users []User
	err := FindAll(query, db.IndexList.Users, &users)
	if err != nil {
		fmt.Println("error find avail username")
	}
	return len(users) == 0
}

func (u *User) EmailAvailable() bool {
	// res, _ := datastore.DB.Users.Find(bson.M{"email": strings.ToLower(u.Email)}).Count()
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

func (u *User) Create() error {
	// return datastore.DB.Users.Insert(u)
	userMarshal, _ := json.Marshal(u)
	res, err := db.Elastic.Index(
		db.IndexList.Users,                     // Index name
		strings.NewReader(string(userMarshal)), // Document body
		db.Elastic.Index.WithDocumentID(u.ID),  // Document ID
		db.Elastic.Index.WithRefresh("true"),   // Refresh
	)
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}
	defer res.Body.Close()
	if res.IsError() {
		db.PrintErrorResponse(res)
	}

	var b bytes.Buffer
	b.ReadFrom(res.Body)
	return nil
}

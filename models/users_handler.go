package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/ardhihdra/chirpbird/db"
	"github.com/ardhihdra/chirpbird/helper"
	"github.com/asaskevich/govalidator"
	"github.com/twinj/uuid"

	"github.com/pkg/errors"
	"github.com/tidwall/gjson"
	"golang.org/x/crypto/bcrypt"
)

type usersHandler struct {
	UsernameRE string
}

func NewUsersHandler() *usersHandler {
	return &usersHandler{
		UsernameRE: "^[A-Za-z0-9_]{1,15}$", // username length copied from Twitter
	}
}

func (h *usersHandler) Register(u *User) (*User, error) {
	u.ID = uuid.NewV4().String()
	u.Username = strings.TrimSpace(u.Username)
	// u.Email = strings.TrimSpace(u.Email)
	// u.Password = strings.TrimSpace(u.Password)
	u.CreatedAt = helper.Timestamp()
	u.UpdatedAt = helper.Timestamp()

	if err := h.UsernameValid(u); err != nil {
		return nil, err
	}

	// if err := h.EmailValid(u); err != nil {
	// 	return nil, err
	// }

	// if err := h.PasswordValid(u); err != nil {
	// 	return nil, err
	// }

	// hpass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	// if err != nil {
	// 	return nil, err
	// }
	// u.Password = string(hpass)

	if err := Create(u); err != nil {
		return nil, err
	}

	return u, nil
}

func (h *usersHandler) UsernameValid(u *User) error {
	if u.Username == "" {
		return errors.New("username required")
	}

	//if matched, _ := regexp.MatchString(h.UsernameRE, u.Username); !matched {
	//	return errors.New("username invalid")
	//}

	if !u.UsernameAvailable() {
		return errors.New("username exists")
	}
	return nil
}

func (h *usersHandler) EmailValid(u *User) error {
	if u.Email == "" {
		return errors.New("email required")
	}

	if !govalidator.IsEmail(u.Email) {
		return errors.New("email invalid")
	}

	if u.EmailAvailable() {
		return errors.New("email exists")
	}
	return nil
}

func (h *usersHandler) PasswordValid(u *User) error {
	if u.Password == "" {
		return errors.New("password required")
	}
	return nil
}

func (h *usersHandler) ByUsername(username string) (*User, error) {
	var u *User
	query := map[string]interface{}{
		"username":        strings.ToLower(username),
		"_source":         true,
		"terminate_after": 1,
	}
	return u, FindOne(query, db.IndexList.Users, u)
}

func (h *usersHandler) ByEmail(email string) (*User, error) {
	var u *User
	query := map[string]interface{}{
		"email": strings.ToLower(email),
	}
	return u, FindOne(query, db.IndexList.Users, u)
}

func (h *usersHandler) ByID(ID string) (User, error) {
	var u User
	query := map[string]interface{}{
		"id": strings.ToLower(ID),
	}
	return u, FindOne(query, db.IndexList.Users, &u)
}

func (h *usersHandler) Auth(userPassword, password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(password)); err != nil {
		return false
	}
	return true
}

func Create(u *User) error {
	userMarshal, _ := json.Marshal(u)
	res, err := db.Elastic.Index(
		db.IndexList.Users,                     // Index name
		strings.NewReader(string(userMarshal)), // Document body
		db.Elastic.Index.WithDocumentID(u.ID),  // Document ID
		db.Elastic.Index.WithRefresh("true"),   // Refresh
	)
	if err != nil {
		log.Fatalf("ERROR: %s", err)
		return err
	}
	defer res.Body.Close()
	if res.IsError() {
		db.PrintErrorResponse(res)
		return err
	}

	return nil
}

func FindOne(query map[string]interface{}, index string, usr *User) error {
	var b bytes.Buffer
	executeQuery(query, index, &b)
	values := gjson.GetManyBytes(b.Bytes(), "hits.hits.0._id", "hits.hits.0._source")
	json.Unmarshal([]byte(values[1].String()), &usr)
	// usr.ID = values[0].String()
	fmt.Println(values)
	return nil
}

func FindAll(query map[string]interface{}, index string, usr *[]User) error {
	var b bytes.Buffer
	executeQuery(query, index, &b)
	values := gjson.GetManyBytes(b.Bytes(), "hits.hits.0._id", "hits.hits")
	json.Unmarshal([]byte(values[1].String()), &usr)
	return nil
}

func executeQuery(query map[string]interface{}, index string, b *bytes.Buffer) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding: %s", err)
		return err
	}
	// return u, db.DB.Users.Find(bson.M{"username": strings.ToLower(username)}).One(&u)
	searchRes, err := db.Elastic.Search(
		db.Elastic.Search.WithIndex(index),
		db.Elastic.Search.WithBody(&buf),
		db.Elastic.Search.WithPretty(),
	)
	if err != nil {
		fmt.Printf("Error searching: %s\n", err)
		// os.Exit(2)
		return err
	}
	defer searchRes.Body.Close()
	if searchRes.IsError() {
		db.PrintErrorResponse(searchRes)
	}

	// parse with gjson
	b.ReadFrom(searchRes.Body)
	return nil
}

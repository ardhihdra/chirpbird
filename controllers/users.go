package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/ardhihdra/chirpbird/datautils"
	"github.com/ardhihdra/chirpbird/helper/jwt"
	"github.com/ardhihdra/chirpbird/models"
)

type UsersController struct {
	BaseController
}

func NewUsersController() *UsersController {
	return &UsersController{}
}

var users = models.NewUsersHandler()

func (usr *UsersController) Login() http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("not found"))
				return
			}
			username := r.FormValue("username")
			country := r.FormValue("country")
			profile := r.FormValue("profile")
			var interests []string
			json.Unmarshal([]byte(r.FormValue("interests")), &interests)

			login, err := saveLogin(username, country, profile, interests)
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
				return
			}

			// message := fmt.Sprintf("Hello %s, %s, %s, %s", username, country, interests, profile)
			// w.Write([]byte(message))

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]string{"token": jwt.Create(login.ID)})
		})
}

func (usr *UsersController) Register() http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("not found"))
				return
			}
			interests := r.FormValue("interests")
			username := r.FormValue("username")
			country := r.FormValue("country")
			profile := r.FormValue("profile")
			user := datautils.User{
				ID:        "1",
				Username:  username,
				Country:   country,
				Profile:   profile,
				Interests: strings.Split(interests, ","),
			}
			_, err := users.Register(&user)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("failed to register"))
				return
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]interface{}{"data": user, "token": jwt.Create(user.ID)})
		})
}

func (usr *UsersController) Logout() http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			// delete User and Session
			// next step set user and session to nonactive and offline
			id := r.FormValue("id")
			user := &datautils.User{ID: id}
			user.DeleteByID()
			session, err := datautils.GetSessionByUserID(user.ID)
			if err != nil {
				fmt.Println("Failed to delete sesion")
			}
			for i := range session {
				/** TO DO: should only update ofline_at */
				session[i].DeleteByID()
			}
		})
}

func (usr *UsersController) CheckUniqueUsername() http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			// GET users with username
			username := r.URL.Query().Get("username")
			user := &datautils.User{Username: username}
			isValid := user.UsernameAvailable()
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]bool{"valid": isValid})
		})
}

func (usr *UsersController) GetUsers() http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			// GET users with username
			if err := usr.Authenticate(r); err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
				return
			}
			id := r.URL.Query().Get("id")
			username := r.URL.Query().Get("username")
			var us []datautils.User
			if id != "" {
				var user datautils.User
				user.ID = id
				user.GetByID()
				us = append(us, user)
			} else if username != "" {
				// find username like
				user, err := users.ByUsername(username, false)
				us = *user
				if err != nil {
					w.WriteHeader(http.StatusBadRequest)
					json.NewEncoder(w).Encode(map[string]string{"error": "err while get users"})
				}
			} else {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]string{"error": "id or username is required"})
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]interface{}{"userinfo": us})
		})
}

func saveLogin(username, country, profile string, interests []string) (*datautils.User, error) {
	var (
		usr *datautils.User
		err error
	)

	if username != "" {

	} else if username != "" {
		usr, err = users.ByEmail(username)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("username or email is required")
	}

	if !users.Auth(usr.Password, username) {
		return nil, errors.New("password wrong")
	}

	return usr, nil
}

package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/ardhihdra/chirpbird/app/datautils"
	"github.com/ardhihdra/chirpbird/app/helper/jwt"
	"github.com/ardhihdra/chirpbird/app/models"
)

type UsersController interface {
	Login() http.HandlerFunc
	Register() http.HandlerFunc
	Logout() http.HandlerFunc
	CheckUniqueUsername() http.HandlerFunc
	GetUsers() http.HandlerFunc
}

type usersController struct {
	BaseController
}

var userModel models.UserModel

func NewUsersController(model models.UserModel) UsersController {
	userModel = model
	return &usersController{}
}

func (usr *usersController) Login() http.HandlerFunc {
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

func (usr *usersController) Register() http.HandlerFunc {
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
			_, err := userModel.Register(&user)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("failed to register"))
				return
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]interface{}{"data": user, "token": jwt.Create(user.ID)})
		},
	)
}

func (usr *usersController) Logout() http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			// delete User and Session
			// next step set user and session to nonactive and offline
			id := r.FormValue("id")
			userModel.DeleteByID(id)
			session, err := datautils.GetSessionByUserID(id)
			if err != nil {
				fmt.Println("Failed to delete sesion")
			}
			for i := range session {
				/** TO DO: should only update ofline_at */
				session[i].DeleteByID()
			}
		})
}

func (usr *usersController) CheckUniqueUsername() http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			// GET users with username
			username := r.URL.Query().Get("username")
			user := &datautils.User{Username: username}
			isValid := userModel.UsernameAvailable(*user)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]bool{"valid": isValid})
		})
}

func (usr *usersController) GetUsers() http.HandlerFunc {
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
				user, err := userModel.ByID(id)
				if err != nil {
					w.WriteHeader(http.StatusBadRequest)
					json.NewEncoder(w).Encode(map[string]string{"error": "err while get users"})
				}
				us = append(us, *user)
			} else if username != "" {
				// find username like
				user, err := userModel.ByUsername(username, false)
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
		},
	)
}

func saveLogin(username, country, profile string, interests []string) (*datautils.User, error) {
	var (
		usr *datautils.User
		err error
	)

	if username != "" {
		usr, err = userModel.ByEmail(username)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("username or email is required")
	}

	if !userModel.Auth(usr.Password, username) {
		return nil, errors.New("password wrong")
	}

	return usr, nil
}

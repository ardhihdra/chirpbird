package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/ardhihdra/chirpbird/helper/jwt"
	"github.com/ardhihdra/chirpbird/models"
)

type UsersController struct {
}

func NewUsersController() *UsersController {
	return &UsersController{}
}

var users = models.NewUsersHandler()

func (usr *UsersController) Login() http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "POST" {
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
			if r.Method != "POST" {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("not found"))
				return
			}
		})

}

func saveLogin(username, country, profile string, interests []string) (*models.User, error) {
	var (
		usr *models.User
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

package controllers

import (
	"encoding/json"
	"net/http"

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
			interests := r.FormValue("interests")
			profile := r.FormValue("profile")

			login, err := saveLogin(username, country, profile, interests)
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				w.WriteJson(map[string]string{"error": err.Error()})
				return
			}

			// message := fmt.Sprintf("Hello %s, %s, %s, %s", username, country, interests, profile)
			// w.Write([]byte(message))

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]string{"token": jwt.Create(u.ID)})
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
	var () (*models.User, error) {
		var (
			usr *models.User
			err error
		)

		if username != "" {

		} else if email != "" {
			u, err = users.ByEmail(email)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, errors.New("username or email is required")
		}

		if !users.Auth(u.Password, password) {
			return nil, errors.New("password wrong")
		}

		return u, nil
}
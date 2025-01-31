package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/ardhihdra/chirpbird/app/models"
)

type SessionsController interface {
	Create() http.HandlerFunc
}
type sessionsController struct {
	BaseController
}

var sessionModel models.SessionModel

func NewSessionsController(model models.SessionModel) SessionsController {
	sessionModel = model
	return &sessionsController{}
}

func (sc *sessionsController) Create() http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("not found"))
				return
			}

			if err := sc.Authenticate(r); err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
				return
			}
			mUserag := r.Header.Get("User-Agent")
			userag := strings.Split(mUserag, " ")
			deviceID := userag[2]
			platform := userag[1]
			build, _ := strconv.Atoi(r.FormValue("build"))
			name := r.FormValue("name")
			s, err := sessionModel.Create(sc.User.ID, deviceID, platform, build, name)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
				return
			}
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"id":           s.ID,
				"access_token": s.AccessToken,
				// "messaging_url": s.MessagingURL(),
				"created_at": s.CreatedAt,
			})
		},
	)
}

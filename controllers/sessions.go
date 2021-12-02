package controllers

import (
	"net/http"
)

type SessionsController struct{}

func NewSessionsController() *SessionsController {
	return &SessionsController{}
}

func (gc *SessionsController) Create() http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "POST" {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("not found"))
				return
			}

		},
	)
}

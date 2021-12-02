package controllers

import (
	"net/http"
)

type MessagingController struct{}

func NewMessagingController() *MessagingController {
	return &MessagingController{}
}

func (gc *MessagingController) Start() http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "GET" {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("not found"))
				return
			}

		},
	)
}

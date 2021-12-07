package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/ardhihdra/chirpbird/models"
)

type GroupsController struct {
	BaseController
}

func NewGroupController() *GroupsController {
	return &GroupsController{}
}

func (gc *GroupsController) Create() http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("not found"))
				return
			}

			if err := gc.Authenticate(r); err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
				return
			}

			/** Init Group */
			groupName := r.FormValue("name")
			groupMember := strings.Split(r.FormValue("user_ids"), ",")
			groupMember = append(groupMember, gc.User.ID)
			g, err := models.Groups.Create(groupName, gc.User.ID, groupMember)
			if err != nil {
				w.WriteHeader(http.StatusBadGateway)
				json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
				return
			}
			fmt.Println(g.ID)
			/** announce new group */
			// if eg := models.Events.NewGroup(g); eg != nil {
			// 	eg.SaveForUsers(g.ID, g.UserIDs)
			// 	eg.SendToUsers(g.UserIDs)
			// }

			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})

		},
	)
}

func (gc *GroupsController) Join() http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPut {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("not found"))
				return
			}

		},
	)
}

func (gc *GroupsController) Left() http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPut {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("not found"))
				return
			}

		},
	)
}

func (gc *GroupsController) DashboardData() http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			// GET all users, users by interests, users by country,
			// GET all rooms, rooms by interests, rooms by country

		})
}

func (gc *GroupsController) RoomsData() http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			// GET room detail
			// CREATE a room

		})
}

func (gc *GroupsController) SearchStuff() http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			// GET search by name, profile,
		})
}

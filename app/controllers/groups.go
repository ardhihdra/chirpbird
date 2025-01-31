package controllers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/ardhihdra/chirpbird/app/datautils"
	"github.com/ardhihdra/chirpbird/app/models"
)

type GroupsController interface {
	Create() http.HandlerFunc
	Join() http.HandlerFunc
	Left() http.HandlerFunc
	DashboardData() http.HandlerFunc
	RoomsData() http.HandlerFunc
	SearchStuff() http.HandlerFunc
}
type groupsController struct {
	BaseController
}

var groupModel models.GroupModel
var eventModel models.EventModel

func NewGroupController(model models.GroupModel, eventM models.EventModel) GroupsController {
	groupModel = model
	eventModel = eventM
	return &groupsController{}
}

func (gc *groupsController) Create() http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if err := gc.Authenticate(r); err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
				return
			}

			/** Init Group */
			groupName := r.FormValue("name")
			groupMember := strings.Split(r.FormValue("user_ids"), ",")
			groupMember = append(groupMember, gc.User.ID)
			g, err := groupModel.Create(groupName, gc.User.ID, groupMember)
			if err != nil {
				w.WriteHeader(http.StatusBadGateway)
				json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
				return
			}
			/** announce new group */
			if eg := datautils.NewGroup(g); eg != nil {
				eventModel.SaveForUsers(g.ID, g.UserIDs, eg)
				eg.SendToUsers(g.UserIDs)
			}
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(map[string]interface{}{"status": true, "group": g})
		},
	)
}

func (gc *groupsController) Join() http.HandlerFunc {
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

func (gc *groupsController) Left() http.HandlerFunc {
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

func (gc *groupsController) DashboardData() http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			// GET all users, users by interests, users by country,
			// GET all rooms, rooms by interests, rooms by country

		})
}

func (gc *groupsController) RoomsData() http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			// GET room detail
			// CREATE a room
			id := r.URL.Query().Get("id")
			name := r.URL.Query().Get("name")
			user_id := r.URL.Query().Get("user_id")
			var us []datautils.Group
			if id != "" {
				group, err := groupModel.GetByID(id)
				if err != nil {
					w.WriteHeader(http.StatusBadGateway)
					json.NewEncoder(w).Encode(map[string]string{"error": "err while get users"})
				}
				us = append(us, *group)
			} else if name != "" {
				// find username like
				group, err := groupModel.ByName(name, false)
				us = *group
				if err != nil {
					w.WriteHeader(http.StatusBadGateway)
					json.NewEncoder(w).Encode(map[string]string{"error": "err while get users"})
				}
			} else if user_id != "" {
				// find username like
				group, err := groupModel.ByUserIDs(user_id)
				us = *group
				if err != nil {
					w.WriteHeader(http.StatusBadGateway)
					json.NewEncoder(w).Encode(map[string]string{"error": "err while get users"})
				}
			} else {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]string{"error": "id or username is required"})
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]interface{}{"groups": us})
		})
}

func (gc *groupsController) SearchStuff() http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			// GET search by name, profile,
		})
}

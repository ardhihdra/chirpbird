package models

import (
	"encoding/json"
	"strings"

	"github.com/ardhihdra/chirpbird/db"
	"github.com/ardhihdra/chirpbird/helper"
	"github.com/twinj/uuid"
)

type groups struct{}

var Groups = new(groups)

type Group struct {
	ID        string   `json:"_id"`
	Name      string   `json:"name"`
	UserID    string   `json:"user_id"`
	UserIDs   []string `json:"user_ids"`
	Deleted   []string `json:"deleted"`
	CreatedAt int64    `json:"created_at"`
	UpdatedAt int64    `json:"updated_at"`
}

func (groups) Create(name, userID string, userIDs []string) (*Group, error) {
	g := &Group{
		ID:        uuid.NewV4().String(),
		Name:      name,
		UserID:    userID,
		UserIDs:   userIDs,
		CreatedAt: helper.Timestamp(),
		UpdatedAt: helper.Timestamp(),
	}
	userMarshal, _ := json.Marshal(g)
	_, err := db.Elastic.Index(
		db.IndexList.Groups,
		strings.NewReader(string(userMarshal)), // Document body
		db.Elastic.Index.WithDocumentID(g.ID),  // Document ID
		db.Elastic.Index.WithRefresh("true"),
	)
	return g, err
}

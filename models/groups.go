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

func (groups) ByID(ID string) (*Group, error) {
	var g *Group
	query := db.MatchCondition(map[string]interface{}{
		"_id": ID,
	})
	values, err := db.FindOne(query, db.IdxGroups)
	if err != nil {
		return nil, err
	}

	return g, json.Unmarshal([]byte(values[1].String()), &g)
}

func (groups) ByIDAndUserID(ID, userID string) (*Group, error) {
	var g *Group
	query := db.MatchCondition(map[string]interface{}{
		"_id":      ID,
		"user_ids": userID,
	})
	values, err := db.FindOne(query, db.IdxGroups)
	if err != nil {
		return nil, err
	}

	return g, json.Unmarshal([]byte(values[1].String()), &g)
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
	marshaled, _ := json.Marshal(g)
	_, err := db.Elastic.Index(
		db.IdxGroups,
		strings.NewReader(string(marshaled)),  // Document body
		db.Elastic.Index.WithDocumentID(g.ID), // Document ID
		db.Elastic.Index.WithRefresh("true"),
	)
	return g, err
}

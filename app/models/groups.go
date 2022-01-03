package models

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/ardhihdra/chirpbird/db"
	"github.com/ardhihdra/chirpbird/helper"
	"github.com/twinj/uuid"
)

type groups struct{}

// handler
var Groups = new(groups)

type Group struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	UserID    string   `json:"user_id"`
	UserIDs   []string `json:"user_ids"`
	Deleted   []string `json:"deleted"`
	CreatedAt int64    `json:"created_at"`
	UpdatedAt int64    `json:"updated_at"`
}

func (groups) GetByID(ID string) (*Group, error) {
	var g *Group
	query := db.MatchCondition(map[string]interface{}{
		"id": ID,
	})
	values, err := db.FindOne(query, db.IdxGroups)
	if err != nil {
		return nil, err
	}

	return g, json.Unmarshal([]byte(values[1].String()), &g)
}

func (groups) ByIDAndUserID(ID, userID string) (*Group, error) {
	var g *Group
	query := db.MustMatch([]map[string]interface{}{
		{"match": map[string]interface{}{"id": ID}},
		{"match": map[string]interface{}{"user_ids": userID}},
	})
	values, err := db.FindOne(query, db.IdxGroups)
	if err != nil {
		return nil, err
	}

	return g, json.Unmarshal([]byte(values[1].String()), &g)
}

func (groups) ByUserIDs(userID string) (*[]Group, error) {
	var g []Group
	query := db.MatchCondition(map[string]interface{}{
		"user_ids": userID,
	})

	values, err := db.FindAll(query, db.IdxGroups)
	arrVal := values[1].Array()
	var gr Group
	for i := range arrVal {
		json.Unmarshal([]byte(arrVal[i].Get("_source").String()), &gr)
		g = append(g, gr)
	}
	if err != nil {
		return &g, err
	}
	return &g, json.Unmarshal([]byte(values[1].String()), &g)
}

func (groups) ByName(name string, exactmatch bool) (*[]Group, error) {
	var u []Group
	var query map[string]interface{}
	if exactmatch {
		query = db.MatchCondition(map[string]interface{}{
			"name": strings.ToLower(name),
		})
	} else {
		query = db.QueryString(map[string]interface{}{
			"fields": []string{"name"},
			"query":  fmt.Sprintf("*%s*", name),
		})
	}
	values, err := db.FindAll(query, db.IdxGroups)
	if err != nil {
		return nil, err
	}
	return &u, json.Unmarshal([]byte(values[1].String()), &u)
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
	res, err := db.Elastic.Index(
		db.IdxGroups,
		strings.NewReader(string(marshaled)),  // Document body
		db.Elastic.Index.WithDocumentID(g.ID), // Document ID
		db.Elastic.Index.WithRefresh("true"),
	)
	if err != nil {
		fmt.Println(err)
		return g, err
	}

	defer res.Body.Close()
	if res.IsError() {
		db.PrintErrorResponse(res)
		return g, err
	}
	return g, err
}

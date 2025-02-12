package repository

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/ardhihdra/chirpbird/app/datautils"
	"github.com/ardhihdra/chirpbird/app/db"
	"github.com/ardhihdra/chirpbird/app/helper"
	"github.com/twinj/uuid"
)

type GroupRepository interface {
	GetByID(ID string) (*datautils.Group, error)
	ByIDAndUserID(ID, userID string) (*datautils.Group, error)
	ByUserIDs(userID string) (*[]datautils.Group, error)
	ByName(name string, exactmatch bool) (*[]datautils.Group, error)
	Create(name, userID string, userIDs []string) (*datautils.Group, error)
}

/** implementation of Elasticsearch for GroupRepository, create another file for other implementation */
type groupRepo struct{}

func NewGroupElasticRepository() GroupRepository {
	return &groupRepo{}
}

func (r *groupRepo) GetByID(ID string) (*datautils.Group, error) {
	var g *datautils.Group
	query := db.MatchCondition(map[string]interface{}{
		"id": ID,
	})
	values, err := db.FindOne(query, db.IdxGroups)
	if err != nil {
		return nil, err
	}

	return g, json.Unmarshal([]byte(values[1].String()), &g)
}

func (r *groupRepo) ByIDAndUserID(ID, userID string) (*datautils.Group, error) {
	var g *datautils.Group
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

func (r *groupRepo) ByUserIDs(userID string) (*[]datautils.Group, error) {
	var g []datautils.Group
	query := db.MatchCondition(map[string]interface{}{
		"user_ids": userID,
	})

	values, err := db.FindAll(query, db.IdxGroups)
	arrVal := values[1].Array()
	var gr datautils.Group
	for i := range arrVal {
		json.Unmarshal([]byte(arrVal[i].Get("_source").String()), &gr)
		g = append(g, gr)
	}
	if err != nil {
		return &g, err
	}
	return &g, json.Unmarshal([]byte(values[1].String()), &g)
}

func (r *groupRepo) ByName(name string, exactmatch bool) (*[]datautils.Group, error) {
	var u []datautils.Group
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

func (r *groupRepo) Create(name, userID string, userIDs []string) (*datautils.Group, error) {
	g := &datautils.Group{
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

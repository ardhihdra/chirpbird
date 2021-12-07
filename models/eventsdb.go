package models

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/ardhihdra/chirpbird/db"
	"github.com/twinj/uuid"
)

type EventDB struct {
	ID        string    `json:"id"`
	Type      EventType `json:"type"`
	ObjectID  string    `json:"object_id"`
	UserIDs   []string  `json:"user_ids"`
	Timestamp int64     `json:"timestamp"`
}

type events struct{}

var Events = new(events)

func (events) GetByUserIDAndTimestamp(ID string, ts int64) ([]*EventDB, error) {
	var ev []*EventDB
	query := db.MatchFilterCondition(
		map[string]interface{}{"user_ids": ID},
		map[string]interface{}{"timestamp": map[string]interface{}{"gt": ts}},
	)
	values, err := db.FindAll(query, db.IdxEvents)
	if err != nil {
		return nil, err
	}

	return ev, json.Unmarshal([]byte(values[1].String()), &ev)
}

func (events) Create(typ EventType, objectID string, clientIDs []string, ts int64) (*EventDB, error) {
	e := &EventDB{
		ID:        uuid.NewV4().String(),
		Type:      typ,
		ObjectID:  objectID,
		UserIDs:   clientIDs,
		Timestamp: ts,
	}
	eMarshal, _ := json.Marshal(e)
	res, err := db.Elastic.Index(
		db.IdxEvents,                          // Index name
		strings.NewReader(string(eMarshal)),   // Document body
		db.Elastic.Index.WithDocumentID(e.ID), // Document ID
		db.Elastic.Index.WithRefresh("true"),  // Refresh
	)
	if err != nil {
		log.Fatalf("ERROR: %s", err)
		return nil, err
	}
	defer res.Body.Close()
	if res.IsError() {
		db.PrintErrorResponse(res)
		return nil, err
	}

	return e, nil
}

func (events) DeleteOldEvents(objectID string, typ EventType, ts int64) {
	query := db.MatchFilterCondition(
		map[string]interface{}{"object_id": objectID, "type": typ},
		map[string]interface{}{"timestamp": map[string]interface{}{"lt": ts}},
	)
	values, err := db.FindAll(query, db.IdxEvents)
	if err != nil {
		return
	}
	var toBeDeleted []*EventDB
	json.Unmarshal([]byte(values[1].String()), &toBeDeleted)
	for idx := range toBeDeleted {
		db.Elastic.Delete(db.IdxEvents, toBeDeleted[idx].ID)
	}
}

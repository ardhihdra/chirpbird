package repository

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/ardhihdra/chirpbird/app/datautils"
	"github.com/ardhihdra/chirpbird/app/db"
	"github.com/twinj/uuid"
)

type EventRepository interface {
	GetByUserIDAndTimestamp(ID string, ts int64) ([]*datautils.EventDB, error)
	CreateEvent(typ datautils.EventType, messageID string, clientIDs []string, ts int64) (*datautils.EventDB, error)
	DeleteOldEvents(messageID string, typ datautils.EventType, ts int64)
	SaveForUser(messageID, userID string, e *datautils.Event)
	SaveForUsers(messageID string, userIDs []string, e *datautils.Event)
}

type eventRepo struct{}

func NewEventElasticRepository() EventRepository {
	return &eventRepo{}
}

func (er *eventRepo) GetByUserIDAndTimestamp(ID string, ts int64) ([]*datautils.EventDB, error) {
	var ev []*datautils.EventDB
	var i_id interface{} = ID
	query := db.MatchFilterCondition(
		map[string]interface{}{"user_ids": i_id},
		map[string]interface{}{"timestamp": map[string]interface{}{"gt": ts}},
	)
	values, err := db.FindAll(query, db.IdxEvents)
	if err != nil {
		return ev, err
	}
	arrVal := values[1].Array()
	for i := range arrVal {
		var evd datautils.EventDB
		if err = json.Unmarshal([]byte(arrVal[i].Get("_source").String()), &evd); err != nil {
			fmt.Println("err", err)
			return ev, err
		}
		ev = append(ev, &evd)
	}
	return ev, nil
}

func (er *eventRepo) CreateEvent(typ datautils.EventType, messageID string, clientIDs []string, ts int64) (*datautils.EventDB, error) {
	e := &datautils.EventDB{
		ID:        uuid.NewV4().String(),
		Type:      typ,
		MessageID: messageID,
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

func (er *eventRepo) DeleteOldEvents(messageID string, typ datautils.EventType, ts int64) {
	var i_messageID interface{} = messageID
	query := db.MatchFilterCondition(
		map[string]interface{}{"message_id": i_messageID, "type": typ},
		map[string]interface{}{"timestamp": map[string]interface{}{"lt": ts}},
	)
	values, err := db.FindAll(query, db.IdxEvents)
	if err != nil {
		return
	}
	var toBeDeleted []*datautils.EventDB
	arrVal := values[1].Array()
	var tbd datautils.EventDB
	for i := range arrVal {
		if err = json.Unmarshal([]byte(arrVal[i].Get("_source").String()), &tbd); err != nil {
			return
		}
		toBeDeleted = append(toBeDeleted, &tbd)
	}
	for idx := range toBeDeleted {
		db.Elastic.Delete(db.IdxEvents, toBeDeleted[idx].ID)
	}
}

func (er *eventRepo) SaveForUser(messageID, userID string, e *datautils.Event) {
	er.SaveForUsers(messageID, []string{userID}, e)
}

func (er *eventRepo) SaveForUsers(messageID string, userIDs []string, e *datautils.Event) {
	er.CreateEvent(e.Type, messageID, userIDs, e.Timestamp)
}

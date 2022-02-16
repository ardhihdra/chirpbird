package models

import (
	"encoding/json"

	"github.com/ardhihdra/chirpbird/app/datautils"
	"github.com/ardhihdra/chirpbird/app/db"
	"github.com/ardhihdra/chirpbird/app/repository"
)

type EventModel interface {
	GetByUserIDAndTimestamp(ID string, ts int64) ([]*datautils.EventDB, error)
	CreateEvent(typ datautils.EventType, messageID string, clientIDs []string, ts int64) (*datautils.EventDB, error)
	DeleteOldEvents(messageID string, typ datautils.EventType, ts int64)
	SaveForUser(messageID, userID string, e *datautils.Event)
	SaveForUsers(messageID string, userIDs []string, e *datautils.Event)
}

type eventModel struct{}

var (
	eventRepo repository.EventRepository
)

func NewEventModel(repos repository.EventRepository) EventModel {
	eventRepo = repos
	return &eventModel{}
}

func (em *eventModel) GetByUserIDAndTimestamp(ID string, ts int64) ([]*datautils.EventDB, error) {
	return eventRepo.GetByUserIDAndTimestamp(ID, ts)
}

func (em *eventModel) CreateEvent(typ datautils.EventType, messageID string, clientIDs []string, ts int64) (*datautils.EventDB, error) {
	return eventRepo.CreateEvent(typ, messageID, clientIDs, ts)
}

func (em *eventModel) DeleteOldEvents(messageID string, typ datautils.EventType, ts int64) {
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

func (em *eventModel) SaveForUser(messageID, userID string, e *datautils.Event) {
	eventRepo.SaveForUser(messageID, userID, e)
}

func (em *eventModel) SaveForUsers(messageID string, userIDs []string, e *datautils.Event) {
	eventRepo.SaveForUsers(messageID, userIDs, e)
}

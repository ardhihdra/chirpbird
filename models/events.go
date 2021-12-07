package models

import (
	"encoding/json"
	"log"

	"github.com/ardhihdra/chirpbird/datautils"
	"github.com/ardhihdra/chirpbird/db"
)

type EventType int

const (
	EVENT_MESSAGE           EventType = 20
	EVENT_MESSAGE_SENT      EventType = 21
	EVENT_MESSAGE_DELIVERED EventType = 22
	EVENT_MESSAGE_READ      EventType = 23
	//EVENT_MESSAGE_UPDATED   EventType = 24
	//EVENT_MESSAGE_DELETED   EventType = 25
	EVENT_TYPING_START EventType = 40
	EVENT_TYPING_END   EventType = 41
	EVENT_GROUP        EventType = 70
	//EVENT_GROUP_UPDATED     EventType = 71
	EVENT_GROUP_JOINED EventType = 72
	EVENT_GROUP_LEFT   EventType = 73
)

type EventMessage struct {
	MessageID string `json:"message_id,omitempty"`
	Data      string `json:"data,omitempty"`
}

type EventMessageSent struct {
	MessageID string `json:"message_id,omitempty"`
}

type EventMessageDelivered struct {
	MessageID string `json:"message_id,omitempty"`
}

type EventTypingStart struct {
	GroupID string `json:"group_id,omitempty"`
	UserID  string `json:"user_id,omitempty"`
}

type EventTypingEnd struct {
	GroupID string `json:"group_id,omitempty"`
	UserID  string `json:"user_id,omitempty"`
}

type EventGroup struct {
	GroupID string   `json:"group_id,omitempty"`
	Name    string   `json:"name,omitempty"`
	UserIDs []string `json:"user_ids,omitempty"`
}

type EventGroupJoined struct {
	GroupID string `json:"group_id,omitempty"`
	UserID  string `json:"user_id,omitempty"`
}

type EventGroupLeft struct {
	GroupID string `json:"group_id,omitempty"`
	UserID  string `json:"user_id,omitempty"`
}

type Event struct {
	Type      EventType   `json:"type"`
	Timestamp int64       `json:"timestamp,omitempty"`
	Body      interface{} `json:"body,omitempty"`
}

/** ws messaging stuff */
func NewEvent(t EventType, ts int64) *Event {
	pe := &Event{
		Type:      t,
		Timestamp: ts,
	}

	return pe
}

func (e *Event) SendToUser(userID string) {
	e.SendToUsersWithoutMe("", []string{userID})
}

func (e *Event) SendToUsers(userIDs []string) {
	e.SendToUsersWithoutMe("", userIDs)
}

func (e *Event) SendToUsersWithoutMe(sessionID string, userIDs []string) {
	for _, uID := range userIDs {
		sessions, err := datautils.GetSessionByUserID(uID)
		if err != nil {
			continue
		}
		for _, s := range sessions {
			if sessionID != s.ID {
				c := &datautils.UserConnection{
					UserID:    uID,
					SessionID: s.ID,
				}
				e.SendEvent(c)
			}
		}
	}
}

func (e *Event) SendEvent(c *datautils.UserConnection) error {
	data, err := json.Marshal(e)
	if err != nil {
		log.Fatalf("ERROR: %s", err)
		return err
	}
	db.Redis.Publish(c.UserID, data)
	return nil
}

func (e *Event) SaveForUser(objectID, userID string) {
	e.SaveForUsers(objectID, []string{userID})
}

func (e *Event) SaveForUsers(objectID string, userIDs []string) {
	Events.Create(e.Type, objectID, userIDs, e.Timestamp)
}

func (e *Event) DeleteOldEvents(objectID string) {
	Events.DeleteOldEvents(objectID, e.Type, e.Timestamp)
}

func NewGroup(g *Group) *Event {
	event := NewEvent(EVENT_GROUP, g.CreatedAt)
	event.Body = EventGroup{
		GroupID: g.ID,
		Name:    g.Name,
		UserIDs: g.UserIDs,
	}
	return event
}

func NewGroupJoined(g *Group) *Event {
	event := NewEvent(EVENT_GROUP_JOINED, g.CreatedAt)
	event.Body = EventGroupJoined{
		GroupID: g.ID,
		UserID:  g.UserID,
	}
	return event
}

func NewGroupLeft(g *Group) *Event {
	event := NewEvent(EVENT_GROUP_LEFT, g.CreatedAt)
	event.Body = EventGroupLeft{
		GroupID: g.ID,
		UserID:  g.UserID,
	}
	return event
}

func NewMessage(m *Message) *Event {
	event := NewEvent(EVENT_MESSAGE, m.CreatedAt)
	event.Body = EventMessage{
		MessageID: m.ID,
		Data:      m.Body.Data,
	}
	return event
}

func NewMessageSent(messageID string, ts int64) *Event {
	event := NewEvent(EVENT_MESSAGE_SENT, ts)
	event.Body = EventMessageSent{
		MessageID: messageID,
	}
	return event
}

func NewMessageDelivered(messageID string, ts int64) *Event {
	event := NewEvent(EVENT_MESSAGE_DELIVERED, ts)
	event.Body = EventMessageDelivered{
		MessageID: messageID,
	}
	return event
}

func NewMessageRead(messageID string, ts int64) *Event {
	event := NewEvent(EVENT_MESSAGE_READ, ts)
	event.Body = EventMessageSent{
		MessageID: messageID,
	}
	return event
}

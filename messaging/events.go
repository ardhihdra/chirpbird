package messaging

import (
	"errors"
	"fmt"

	"github.com/ardhihdra/chirpbird/datautils"
	"github.com/ardhihdra/chirpbird/models"
)

type Events struct{}

func newEvents() *Events {
	return &Events{}
}

func (e *Events) Get(c *datautils.UserConnection, p *datautils.RpcMessageGet) {
	es, err := models.Events.GetByUserIDAndTimestamp(c.UserID, p.Timestamp)
	if err != nil {
		return
	}
	for _, event := range es {
		ev, err := e.createMessagePayload(event)
		if err != nil {
			continue
		}
		if ev == nil {
			fmt.Println("Failed to execute event")
		}
		ev.SendToUser(c.UserID)
	}
}

func (e *Events) createMessagePayload(event *models.EventDB) (*models.Event, error) {
	switch event.Type {
	case models.EVENT_MESSAGE:
		return e.messagePayload(event.ObjectID)
	case models.EVENT_MESSAGE_SENT:
		return e.messagePayloadSent(event.ObjectID, event.Timestamp)
	case models.EVENT_MESSAGE_DELIVERED:
		return e.messagePayloadDelivered(event.ObjectID, event.Timestamp)
	case models.EVENT_MESSAGE_READ:
		return e.messagePayloadRead(event.ObjectID, event.Timestamp)
	case models.EVENT_GROUP:
		return e.messagePayloadGroup(event.ObjectID)
	case models.EVENT_GROUP_JOINED:
		return e.messagePayloadGroupJoined(event.ObjectID)
	case models.EVENT_GROUP_LEFT:
		return e.messagePayloadGroupLeft(event.ObjectID)
	}

	return nil, errors.New("wrong event type")
}

func (e *Events) messagePayload(messageID string) (*models.Event, error) {
	m, err := models.Messages.ByID(messageID)
	if err != nil {
		return nil, err
	}
	return models.NewMessage(m), nil
}

func (e *Events) messagePayloadSent(messageID string, ts int64) (*models.Event, error) {
	m, err := models.Messages.ByID(messageID)
	if err != nil {
		return nil, err
	}
	return models.NewMessageSent(m.ID, ts), nil
}

func (e *Events) messagePayloadDelivered(messageID string, ts int64) (*models.Event, error) {
	m, err := models.Messages.ByID(messageID)
	if err != nil {
		return nil, err
	}
	return models.NewMessageDelivered(m.ID, ts), nil
}

func (e *Events) messagePayloadRead(messageID string, ts int64) (*models.Event, error) {
	m, err := models.Messages.ByID(messageID)
	if err != nil {
		return nil, err
	}
	return models.NewMessageRead(m.ID, ts), nil
}

func (e *Events) messagePayloadGroup(groupID string) (*models.Event, error) {
	g, err := models.Groups.ByID(groupID)
	if err != nil {
		return nil, err
	}
	return models.NewGroup(g), nil
}

func (e *Events) messagePayloadGroupJoined(groupID string) (*models.Event, error) {
	g, err := models.Groups.ByID(groupID)
	if err != nil {
		return nil, err
	}
	return models.NewGroupJoined(g), nil
}

func (e *Events) messagePayloadGroupLeft(groupID string) (*models.Event, error) {
	g, err := models.Groups.ByID(groupID)
	if err != nil {
		return nil, err
	}
	return models.NewGroupLeft(g), nil
}

package messaging

import (
	"errors"
	"fmt"

	"github.com/ardhihdra/chirpbird/app/datautils"
	"github.com/ardhihdra/chirpbird/app/models"
)

type Events struct{}

func newEvents() *Events {
	return &Events{}
}

func (e *Events) Get(c *datautils.UserConnection, p *datautils.RpcMessageGet) {
	es, err := models.Events.GetByUserIDAndTimestamp(c.UserID, p.Timestamp)
	if err != nil {
		fmt.Println(err)
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

func (e *Events) createMessagePayload(event *models.EventDB) (*datautils.Event, error) {
	switch event.Type {
	case datautils.EVENT_MESSAGE:
		return e.messagePayload(event.MessageID)
	case datautils.EVENT_MESSAGE_SENT:
		return e.messagePayloadSent(event.MessageID, event.Timestamp)
	case datautils.EVENT_MESSAGE_DELIVERED:
		return e.messagePayloadDelivered(event.MessageID, event.Timestamp)
	case datautils.EVENT_MESSAGE_READ:
		return e.messagePayloadRead(event.MessageID, event.Timestamp)
	case datautils.EVENT_GROUP:
		return e.messagePayloadGroup(event.MessageID)
	case datautils.EVENT_GROUP_JOINED:
		return e.messagePayloadGroupJoined(event.MessageID)
	case datautils.EVENT_GROUP_LEFT:
		return e.messagePayloadGroupLeft(event.MessageID)
	}

	return nil, errors.New("wrong event type")
}

func (e *Events) messagePayload(messageID string) (*datautils.Event, error) {
	m, err := messageModel.ByID(messageID)
	if err != nil {
		return nil, err
	}
	user, _ := userModel.ByID(m.UserID)
	return datautils.NewMessage(m, user), nil
}

func (e *Events) messagePayloadSent(messageID string, ts int64) (*datautils.Event, error) {
	m, err := messageModel.ByID(messageID)
	if err != nil {
		return nil, err
	}
	return datautils.NewMessageSent(m, ts), nil
}

func (e *Events) messagePayloadDelivered(messageID string, ts int64) (*datautils.Event, error) {
	m, err := messageModel.ByID(messageID)
	if err != nil {
		return nil, err
	}
	return datautils.NewMessageDelivered(m, ts), nil
}

func (e *Events) messagePayloadRead(messageID string, ts int64) (*datautils.Event, error) {
	m, err := messageModel.ByID(messageID)
	if err != nil {
		return nil, err
	}
	return datautils.NewMessageRead(m.ID, ts), nil
}

func (e *Events) messagePayloadGroup(groupID string) (*datautils.Event, error) {
	g, err := groupModel.GetByID(groupID)
	if err != nil {
		return nil, err
	}
	return datautils.NewGroup(g), nil
}

func (e *Events) messagePayloadGroupJoined(groupID string) (*datautils.Event, error) {
	g, err := groupModel.GetByID(groupID)
	if err != nil {
		return nil, err
	}
	return datautils.NewGroupJoined(g), nil
}

func (e *Events) messagePayloadGroupLeft(groupID string) (*datautils.Event, error) {
	g, err := groupModel.GetByID(groupID)
	if err != nil {
		return nil, err
	}
	return datautils.NewGroupLeft(g), nil
}

package messaging

import (
	"github.com/ardhihdra/chirpbird/datautils"
	"github.com/ardhihdra/chirpbird/helper"
	"github.com/ardhihdra/chirpbird/models"
)

type messages struct{}

func newMessages() *messages {
	return &messages{}
}

func (m *messages) Send(c *datautils.UserConnection, p *datautils.RpcMessageSend) {
	group := withGroup(p.GroupID, c.UserID)
	msg, _ := models.Messages.Create(p.GroupID, c.UserID, p.Data, helper.Timestamp())

	e := models.NewMessage(msg)
	e.SaveForUsers(msg.ID, group.UserIDs)

	es := models.NewMessageSent(msg.ID, e.Timestamp)
	es.SaveForUser(msg.ID, group.UserID)
	es.SendToUser(group.UserID)

	e.SendToUsersWithoutMe(c.SessionID, group.UserIDs)
}

func (m *messages) Read(c *datautils.UserConnection, p *datautils.RpcMessageRead) {
	msg, err := models.Messages.ByID(p.MessageID)
	if err != nil {
		return
	}
	group := withGroup(msg.GroupID, c.UserID)
	if group == nil {
		return
	}
	e := models.NewMessageRead(msg.ID, helper.Timestamp())
	e.SaveForUser(msg.ID, msg.UserID)
	e.SendToUsersWithoutMe(c.SessionID, []string{msg.UserID, c.UserID})
	e.DeleteOldEvents(msg.ID)
}

func (m *messages) Delivered(c *datautils.UserConnection, p *datautils.RpcMessageDelivered) {
	msg, err := models.Messages.ByID(p.MessageID)
	if err != nil {
		return
	}
	group := withGroup(msg.GroupID, c.UserID)
	if group == nil {
		return
	}
	e := models.NewMessageDelivered(msg.ID, helper.Timestamp())
	e.SaveForUser(msg.ID, msg.UserID)
	e.SendToUsersWithoutMe(c.SessionID, []string{msg.UserID, c.UserID})
	e.DeleteOldEvents(msg.ID)
}

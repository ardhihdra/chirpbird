package messaging

import (
	"github.com/ardhihdra/chirpbird/app/datautils"
	"github.com/ardhihdra/chirpbird/app/helper"
	"github.com/ardhihdra/chirpbird/app/models"
)

type messages struct{}

func newMessages() *messages {
	return &messages{}
}

func (m *messages) Send(c *datautils.UserConnection, p *datautils.RpcMessageSend) {
	group := withGroup(p.GroupID, c.UserID)
	msg, _ := messageModel.Create(p.GroupID, c.UserID, p.Data, helper.Timestamp())
	user, _ := userModel.ByID(msg.UserID)

	e := datautils.NewMessage(msg, user)
	models.Events.SaveForUsers(msg.ID, group.UserIDs, e)

	es := datautils.NewMessageSent(msg, e.Timestamp)
	models.Events.SaveForUser(msg.ID, group.UserID, es)

	// e.SendToUsersWithoutMe(c.SessionID, group.UserIDs)
	e.SendToUsers(group.UserIDs)
	/** notify sent */
	es.SendToUser(group.UserID)
}

func (m *messages) Read(c *datautils.UserConnection, p *datautils.RpcMessageRead) {
	msg, err := messageModel.ByID(p.MessageID)
	if err != nil {
		return
	}
	group := withGroup(msg.GroupID, c.UserID)
	if group == nil {
		return
	}
	e := datautils.NewMessageRead(msg.ID, helper.Timestamp())
	models.Events.SaveForUser(msg.ID, msg.UserID, e)
	e.SendToUsersWithoutMe(c.SessionID, []string{msg.UserID, c.UserID})
	models.Events.DeleteOldEvents(msg.ID, e.Type, e.Timestamp)
}

func (m *messages) Delivered(c *datautils.UserConnection, p *datautils.RpcMessageDelivered) {
	msg, err := messageModel.ByID(p.MessageID)
	if err != nil {
		return
	}
	group := withGroup(msg.GroupID, c.UserID)
	if group == nil {
		return
	}
	e := datautils.NewMessageDelivered(msg, helper.Timestamp())
	models.Events.SaveForUser(msg.ID, msg.UserID, e)
	e.SendToUsersWithoutMe(c.SessionID, []string{msg.UserID, c.UserID})
	models.Events.DeleteOldEvents(msg.ID, e.Type, e.Timestamp)
}

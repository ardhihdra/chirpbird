package messaging

import (
	"fmt"

	"github.com/ardhihdra/chirpbird/app/datautils"
	"github.com/ardhihdra/chirpbird/app/helper"
	"github.com/ardhihdra/chirpbird/app/models"
)

type messages struct{}

var messageM models.MessageModel

func newMessages(messageModels models.MessageModel) *messages {
	messageM = messageModels
	return &messages{}
}

func (m *messages) Send(c *datautils.UserConnection, p *datautils.RpcMessageSend) {
	group := withGroup(p.GroupID, c.UserID)
	fmt.Println(messageM)
	msg, _ := messageM.Create(p.GroupID, c.UserID, p.Data, helper.Timestamp())
	user, _ := BaseModel.UserModel.ByID(msg.UserID)

	e := datautils.NewMessage(msg, user)
	BaseModel.EventModel.SaveForUsers(msg.ID, group.UserIDs, e)

	es := datautils.NewMessageSent(msg, e.Timestamp)
	BaseModel.EventModel.SaveForUser(msg.ID, group.UserID, es)

	// e.SendToUsersWithoutMe(c.SessionID, group.UserIDs)
	e.SendToUsers(group.UserIDs)
	/** notify sent */
	es.SendToUser(group.UserID)
}

func (m *messages) Read(c *datautils.UserConnection, p *datautils.RpcMessageRead) {
	msg, err := BaseModel.MessageModel.ByID(p.MessageID)
	if err != nil {
		return
	}
	group := withGroup(msg.GroupID, c.UserID)
	if group == nil {
		return
	}
	e := datautils.NewMessageRead(msg.ID, helper.Timestamp())
	BaseModel.EventModel.SaveForUser(msg.ID, msg.UserID, e)
	e.SendToUsersWithoutMe(c.SessionID, []string{msg.UserID, c.UserID})
	BaseModel.EventModel.DeleteOldEvents(msg.ID, e.Type, e.Timestamp)
}

func (m *messages) Delivered(c *datautils.UserConnection, p *datautils.RpcMessageDelivered) {
	msg, err := BaseModel.MessageModel.ByID(p.MessageID)
	if err != nil {
		return
	}
	group := withGroup(msg.GroupID, c.UserID)
	if group == nil {
		return
	}
	e := datautils.NewMessageDelivered(msg, helper.Timestamp())
	BaseModel.EventModel.SaveForUser(msg.ID, msg.UserID, e)
	e.SendToUsersWithoutMe(c.SessionID, []string{msg.UserID, c.UserID})
	BaseModel.EventModel.DeleteOldEvents(msg.ID, e.Type, e.Timestamp)
}

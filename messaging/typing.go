package messaging

import (
	"github.com/ardhihdra/chirpbird/datautils"
	"github.com/ardhihdra/chirpbird/helper"
	"github.com/ardhihdra/chirpbird/models"
)

type typing struct{}

func newTyping() *typing {
	return &typing{}
}

func (t *typing) Start(c *datautils.UserConnection, p *datautils.RpcTypingStart) {
	group := withGroup(p.GroupID, c.UserID)
	e := models.NewEvent(models.EVENT_TYPING_START, helper.Timestamp())
	e.Body = models.EventTypingStart{
		GroupID: group.ID,
		UserID:  c.UserID,
	}
	e.SendToUsersWithoutMe(c.SessionID, group.UserIDs)
}

func (t *typing) End(c *datautils.UserConnection, p *datautils.RpcTypingEnd) {
	group := withGroup(p.GroupID, c.UserID)
	e := models.NewEvent(models.EVENT_TYPING_END, helper.Timestamp())
	e.Body = models.EventTypingEnd{
		GroupID: group.ID,
		UserID:  c.UserID,
	}
	e.SendToUsersWithoutMe(c.SessionID, group.UserIDs)
}
